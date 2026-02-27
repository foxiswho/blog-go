package service

import (
	"context"

	"github.com/foxiswho/blog-go/infrastructure/entityRam"
	"github.com/foxiswho/blog-go/infrastructure/repositoryRam"
	"github.com/foxiswho/blog-go/pkg/consts/constsRam/resourceTypeCategoryPg"
	"github.com/foxiswho/blog-go/pkg/consts/constsRam/typeAttrPg"
	"github.com/foxiswho/blog-go/pkg/enum/state/enumStatePg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/foxiswho/blog-go/pkg/model"
	"github.com/gin-gonic/gin"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"

	"reflect"

	"github.com/pangu-2/go-tools/tools/numberPg"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
)

func init() {
	gs.Provide(new(RamResourceGroupAuthorizationService)).Init(func(s *RamResourceGroupAuthorizationService) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

// RamResourceGroupAuthorizationService 资源组 授权方法动作
// @Description:
type RamResourceGroupAuthorizationService struct {
	sv     *repositoryRam.RamResourceGroupRepository     `autowire:"?"`
	authDb *repositoryRam.RamResourceAuthorityRepository `autowire:"?"`
	ra     *RamResourceAuthorizationService              `autowire:"?"`
	log    *log2.Logger                                  `autowire:"?"`
}

// Enable 启用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamResourceGroupAuthorizationService) Enable(ctx *gin.Context, ct model.BaseIdsCt[string]) (rt rg.Rs[string]) {
	return c.State(ctx, ct.Ids, enumStatePg.ENABLE)
}

// Disable 禁用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamResourceGroupAuthorizationService) Disable(ctx *gin.Context, ct model.BaseIdsCt[string]) (rt rg.Rs[string]) {
	return c.State(ctx, ct.Ids, enumStatePg.GetType(enumStatePg.DISABLE))
}

// State 状态 启用/禁用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamResourceGroupAuthorizationService) State(ctx *gin.Context, ids []string, state enumStatePg.State) (rt rg.Rs[string]) {
	if len(ids) < 1 {
		return rt.ErrorMessage("id错误")
	}
	r := c.sv
	finds, b := r.FindAllByIdStringIn(ids)
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	for _, info := range finds {
		if info.State != state.IndexInt8() {
			r.Update(entityRam.RamResourceGroupEntity{State: state.IndexInt8()}, info.ID)
		}
	}
	//有效
	//if enumStatePg.ENABLE.IsExistInt8(state.IndexInt8()) {
	//	for _, info := range finds {
	//		c.ra.AuthorityValid(ctx, utilsRam.ResourceAuthorityMarkByInt64(iamConstant.GroupResourceTypeCategory, info.ID))
	//	}
	//} else {
	//	//无效
	//	for _, info := range finds {
	//		c.ra.AuthorityInValid(ctx, utilsRam.ResourceAuthorityMarkByInt64(iamConstant.GroupResourceTypeCategory, info.ID))
	//	}
	//}
	return rt.Ok()
}

// StateEnableDisable 状态 设置 有效 停用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamResourceGroupAuthorizationService) StateEnableDisable(ctx *gin.Context, ids []string, state enumStatePg.State) (rt rg.Rs[string]) {
	if !state.IsEnableDisable() {
		return rt.ErrorMessage("状态错误")
	}
	return c.State(ctx, ids, state)
}

// LogicalDeletion 逻辑删除
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamResourceGroupAuthorizationService) LogicalDeletion(ctx *gin.Context, ids []string) (rt rg.Rs[string]) {
	c.log.Infof("ct=%+v", ids)
	if len(ids) < 1 {
		return rt.ErrorMessage("id错误")
	}
	repository := c.sv
	finds, b := repository.FindAllByIdStringIn(ids)
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	//物理删除
	if c.sv.Config().Data.Delete {
		for _, info := range finds {
			c.log.Infof("id=%v,TenantId=%v", info.ID, info.TenantNo)

			//删除授权及权限规则
			//c.ra.Delete(ctx, utilsRam.ResourceAuthorityMarkByInt64(iamConstant.GroupResourceTypeCategory, info.ID))
		}
		repository.DeleteByIdsString(ids)

	} else {
		for _, info := range finds {
			enum := enumStatePg.State(info.State)
			// 有效 停用，反转 为对应的 取消 弃置
			if ok, reverse := enum.ReverseEnableDisable(); ok {
				repository.Update(entityRam.RamResourceGroupEntity{State: reverse.IndexInt8()}, info.ID)
			}

			//设置无效
			//c.ra.AuthorityInValid(ctx, utilsRam.ResourceAuthorityMarkByInt64(iamConstant.GroupResourceTypeCategory, info.ID))
		}
	}

	return rt.Ok()
}

// LogicalRecovery 逻辑删除恢复
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamResourceGroupAuthorizationService) LogicalRecovery(ctx *gin.Context, ids []string) (rt rg.Rs[string]) {
	c.log.Infof("ct=%+v", ids)
	if len(ids) < 1 {
		return rt.ErrorMessage("id错误")
	}
	repository := c.sv
	finds, b := repository.FindAllByIdStringIn(ids)
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	for _, info := range finds {
		enum := enumStatePg.State(info.State)
		//  取消 弃置 批量删除，反转 为对应的 有效 停用 停用
		if ok, reverse := enum.ReverseCancelLayAside(); ok {
			repository.Update(entityRam.RamResourceGroupEntity{State: reverse.IndexInt8()}, info.ID)
		}

		//设置有效
		//c.ra.AuthorityValid(ctx, utilsRam.ResourceAuthorityMarkByInt64(iamConstant.GroupResourceTypeCategory, info.ID))
	}
	return rt.Ok()
}

// PhysicalDeletion 物理删除
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamResourceGroupAuthorizationService) PhysicalDeletion(ctx *gin.Context, ids []string) (rt rg.Rs[string]) {
	c.log.Infof("ct=%+v", ids)
	if len(ids) < 1 {
		return rt.ErrorMessage("id错误")
	}
	r := c.sv
	authDb := c.authDb
	finds, b := r.FindAllByIdStringIn(ids)
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	idsNew := make([]int64, 0)
	idsCategory := make([]string, 0)
	for _, info := range finds {
		c.log.Infof("id=%v,TenantId=%v", info.ID, info.TenantNo)
		//判断 是否是 分类
		if typeAttrPg.CategoryLast.IsEqual(info.TypeAttr) || typeAttrPg.Category.IsEqual(info.TypeAttr) {
			link, b := r.FindAllByIdLink(numberPg.Int64ToString(info.ID))
			if b {
				for _, entity := range link {
					idsCategory = append(idsCategory, numberPg.Int64ToString(entity.ID))
				}
				if len(idsCategory) > 0 {
					_, result := authDb.FindAllByTypeCategoryAndGroupIdStringIn(resourceTypeCategoryPg.Group.String(), idsCategory)
					if result {
						return rt.ErrorMessage("该分类下存在子数据，请先删除子数据")
					}
				}
				idsNew = append(idsNew, info.ID)
			}
		}
		//删除授权及权限规则
		//c.ra.Delete(ctx, utilsRam.ResourceAuthorityMarkByInt64(iamConstant.GroupResourceTypeCategory, info.ID))
	}
	if len(idsNew) > 0 {
		r.DeleteByIds(idsNew)
	}
	return rt.Ok()
}

package service

import (
	"context"

	"github.com/foxiswho/blog-go/app/manage/domainRam/model/modRamResourceRelation"
	"github.com/foxiswho/blog-go/infrastructure/entityRam"
	"github.com/foxiswho/blog-go/infrastructure/repositoryRam"
	"github.com/foxiswho/blog-go/pkg/consts/constsRam/typeAttrPg"
	"github.com/foxiswho/blog-go/pkg/enum/state/enumStatePg"
	"github.com/foxiswho/blog-go/pkg/holderPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/gin-gonic/gin"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"

	"reflect"
	"time"

	"github.com/jinzhu/copier"
	"github.com/pangu-2/go-tools/tools/numberPg"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
)

func init() {
	gs.Provide(new(RamResourceAuthorizationService)).Init(func(s *RamResourceAuthorizationService) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

// RamResourceAuthorizationService 资源授权动作 公共调用
// @Description:
type RamResourceAuthorizationService struct {
	ra            *RamResourceAuthorityService                  `autowire:"?"`
	resourceAuth  *repositoryRam.RamResourceAuthorityRepository `autowire:"?"`
	resourceGroup *repositoryRam.RamResourceGroupRepository     `autowire:"?"`
	resourceRel   *repositoryRam.RamResourceRelationRepository  `autowire:"?"`
	resource      *repositoryRam.RamResourceRepository          `autowire:"?"`
	casbin        *RamResourceCasbinService                     `autowire:"?"`
	log           *log2.Logger                                  `autowire:"?"`
}

// UpdateByResourceGroup 更新权限
//
//	@Description:
//	@receiver c
//	@param ct
//	@return rt
func (c *RamResourceAuthorizationService) UpdateByResourceGroup(ctx *gin.Context, ct modRamResourceRelation.UpdateByResourceGroupCt) (rt rg.Rs[string]) {
	if ct.SourceId <= 0 {
		return rt.ErrorMessage("资源组id不能为空")
	}
	source, b := c.resourceGroup.FindByIdString(ct.SourceId.ToString())
	if !b {
		return rt.ErrorMessage("资源组 不存在")
	}
	auth, b := c.ra.CreateByResourceGroup(ctx, source)
	if !b {
		return rt.ErrorMessage("创建资源权限失败")
	}
	c.log.Infof("authPg=%+v", auth)
	holder := holderPg.GetContextAccount(ctx)
	c.resourceRel.DeleteByAuthorityId(auth.ID)
	if len(ct.Ids) > 0 {
		c.resourceRel.DeleteByAuthorityId(auth.ID)
		now := time.Now()
		//增加权限
		infos, b := c.resource.FindAllByIdStringIn(ct.Ids)
		if !b {
			slice := make([]*entityRam.RamResourceRelationEntity, 0)
			//保存数据库
			for _, item := range infos {
				if !typeAttrPg.Resource.IsEqual(item.TypeAttr) {
					continue
				}
				if "" == item.Path {
					continue
				}
				if "" == item.Method {
					continue
				}
				var info entityRam.RamResourceRelationEntity
				copier.Copy(&info, &item)
				info.ID = 0
				info.CreateAt = &now
				info.TypeCategory = auth.TypeCategory
				info.Mark = auth.Mark
				info.AuthorityId = auth.ID
				info.TypeValue = ct.SourceId.ToString()
				info.ResourceId = item.ID
				c.log.Infof("info%+v", info)
				info.TenantNo = holder.GetTenantNo()
				c.resourceRel.Create(&info)
				c.log.Infof("save=%+v", info)

				slice = append(slice, &info)
			}
			//保存到 casbin
			c.casbin.UpdateCasbin(numberPg.Int64ToString(auth.ID), slice)
		} else {
			return rt.ErrorMessage("权限 不存在")
		}
	}

	return rt.Ok()
}

// Delete 删除
func (c *RamResourceAuthorizationService) Delete(ctx *gin.Context, code string) {
	//删除授权id
	info, result := c.resourceAuth.FindByMark(code)
	if result {
		c.resourceAuth.DeleteByMark(code)
		//删除 资源与授权id关系
		c.resourceRel.DeleteByAuthorityId(info.ID)
		//删除 casbin 内权限规则
		c.casbin.ClearCasbin(int(info.ID))
	}
}

// AuthorityValid code string 标记
//
// AuthorityValid 设置授权id 有效
func (c *RamResourceAuthorizationService) AuthorityValid(ctx *gin.Context, code string) {
	//删除授权id
	info, result := c.resourceAuth.FindByMark(code)
	if result {
		//设置有效
		c.resourceAuth.Update(entityRam.RamResourceAuthorityEntity{State: enumStatePg.ENABLE.Index()}, info.ID)
		//删除 资源与授权id关系
		infos := c.resourceRel.FindAll(entityRam.RamResourceRelationEntity{Mark: code})
		//保存到 casbin
		c.casbin.UpdateCasbin(numberPg.Int64ToString(info.ID), infos)
	}
}

// AuthorityInValid code string 标记
//
// AuthorityInValid 设置授权id 无效
func (c *RamResourceAuthorizationService) AuthorityInValid(ctx *gin.Context, code string) {
	//删除授权id
	info, result := c.resourceAuth.FindByMark(code)
	if result {
		//设置无效
		c.resourceAuth.Update(entityRam.RamResourceAuthorityEntity{State: enumStatePg.DISABLE.Index()}, info.ID)
		//保存到 casbin ，直接清空规则
		c.casbin.ClearCasbin(int(info.ID))
	}
}

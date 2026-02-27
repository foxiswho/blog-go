package service

import (
	"context"

	"github.com/foxiswho/blog-go/app/manage/domainRam/model/modRamResourceRelation"
	"github.com/foxiswho/blog-go/infrastructure/entityRam"
	"github.com/foxiswho/blog-go/infrastructure/repositoryRam"
	"github.com/foxiswho/blog-go/pkg/enum/state/enumStatePg"
	"github.com/foxiswho/blog-go/pkg/holderPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/foxiswho/blog-go/pkg/model"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/gin-gonic/gin"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"

	"reflect"

	"github.com/jinzhu/copier"
	"github.com/pangu-2/go-tools/tools/dbPg/pagePg"
	"github.com/pangu-2/go-tools/tools/numberPg"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
)

func init() {
	gs.Provide(new(RamResourceRelationService)).Init(func(s *RamResourceRelationService) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

// RamResourceRelationService 资源关系
// @Description:
type RamResourceRelationService struct {
	sv  *repositoryRam.RamResourceRelationRepository `autowire:"?"`
	log *log2.Logger                                 `autowire:"?"`
}

// Create 新增
//
//	@Description:
//	@receiver c
//	@param ct
//	@return rt
func (c *RamResourceRelationService) Create(ctx *gin.Context, ct modRamResourceRelation.CreateCt) (rt rg.Rs[string]) {
	if "" == ct.Name {
		return rt.ErrorMessage("名称不能为空")
	}
	holder := holderPg.GetContextAccount(ctx)

	var info entityRam.RamResourceRelationEntity
	copier.Copy(&info, &ct)
	c.log.Infof("info%+v", info)
	info.TenantNo = holder.GetTenantNo()
	c.sv.Create(&info)
	c.log.Infof("save=%+v", info)
	return rg.OkData(numberPg.Int64ToString(info.ID))
}

// Update 更新
//
//	@Description:
//	@receiver c
//	@param ct
//	@return rt
func (c *RamResourceRelationService) Update(ctx *gin.Context, ct modRamResourceRelation.UpdateCt) (rt rg.Rs[string]) {
	if ct.ID < 1 {
		return rt.ErrorMessage("id错误")
	}
	if "" == ct.Name {
		return rt.ErrorMessage("名称不能为空")
	}
	r := c.sv
	_, b := r.FindById(ct.ID.ToInt64())
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	var info entityRam.RamResourceRelationEntity
	copier.Copy(&info, &ct)
	r.Update(info, info.ID)
	return rt.Ok()
}

// Detail 详情
//
//	@Description:
//	@receiver c
//	@param id
func (c *RamResourceRelationService) Detail(ctx *gin.Context, id int64) (rt rg.Rs[modRamResourceRelation.Vo]) {
	if id < 1 {
		return rt.ErrorMessage("id错误")
	}
	find, b := c.sv.FindById(id, repositoryPg.GetOption(ctx))
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	var info modRamResourceRelation.Vo
	copier.Copy(&info, find)
	return rt.OkData(info)
}

// Delete 删除
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamResourceRelationService) Delete(ctx *gin.Context, ct model.BaseIdsCt[string]) (rt rg.Rs[string]) {
	if len(ct.Ids) < 1 {
		return rt.ErrorMessage("id错误")
	}
	r := c.sv
	finds, b := r.FindAllByIdStringIn(ct.Ids)
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	for _, info := range finds {
		c.log.Infof("id=%v,TenantId=%v", info.ID, info.TenantNo)
		r.DeleteById(info.ID)
	}
	return rt.Ok()
}

// Enable 启用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamResourceRelationService) Enable(ctx *gin.Context, ct model.BaseIdsCt[string]) (rt rg.Rs[string]) {
	return c.State(ctx, ct.Ids, enumStatePg.ENABLE)
}

// Disable 禁用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamResourceRelationService) Disable(ctx *gin.Context, ct model.BaseIdsCt[string]) (rt rg.Rs[string]) {
	return c.State(ctx, ct.Ids, enumStatePg.GetType(enumStatePg.DISABLE))
}

// State 状态 启用/禁用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamResourceRelationService) State(ctx *gin.Context, ids []string, state enumStatePg.State) (rt rg.Rs[string]) {
	if len(ids) < 1 {
		return rt.ErrorMessage("id错误")
	}
	r := c.sv
	_, b := r.FindAllByIdStringIn(ids)
	if b {
		return rt.ErrorMessage("数据不存在")
	}
	//for _, info := range finds {
	//	if info.StateOrder != state.IndexInt8() {
	//		r.Update(entityRam.RamResourceRelationEntity{StateOrder: state.IndexInt8()}, info.ID)
	//	}
	//}
	return rt.Ok()
}

// StateEnableDisable 状态 设置 有效 停用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamResourceRelationService) StateEnableDisable(ctx *gin.Context, ids []string, state enumStatePg.State) (rt rg.Rs[string]) {
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
func (c *RamResourceRelationService) LogicalDeletion(ctx *gin.Context, ids []string) (rt rg.Rs[string]) {
	c.log.Infof("ct=%+v", ids)
	if len(ids) < 1 {
		return rt.ErrorMessage("id错误")
	}
	repository := c.sv
	finds, b := repository.FindAllByIdStringIn(ids, repositoryPg.GetOption(ctx))
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	if c.sv.Config().Data.Delete {
		for _, info := range finds {
			c.log.Infof("id=%v,TenantId=%v", info.ID, info.TenantNo)
		}
		repository.DeleteByIdsString(ids, repositoryPg.GetOption(ctx))
	} else {
		//for _, info := range finds {
		//	enum := enumStatePg.State(info.StateOrder)
		//	// 有效 停用，反转 为对应的 取消 弃置
		//	if ok, reverse := enum.ReverseEnableDisable(); ok {
		//		repository.Update(entityRam.RamResourceRelationEntity{StateOrder: reverse.IndexInt8()}, info.ID)
		//	}
		//}
	}

	return rt.Ok()
}

// LogicalRecovery 逻辑删除恢复
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamResourceRelationService) LogicalRecovery(ctx *gin.Context, ids []string) (rt rg.Rs[string]) {
	c.log.Infof("ct=%+v", ids)
	if len(ids) < 1 {
		return rt.ErrorMessage("id错误")
	}
	repository := c.sv
	_, b := repository.FindAllByIdStringIn(ids)
	if b {
		return rt.ErrorMessage("数据不存在")
	}
	//for _, info := range finds {
	//	enum := enumStatePg.State(info.StateOrder)
	//	//  取消 弃置 批量删除，反转 为对应的 有效 停用 停用
	//	if ok, reverse := enum.ReverseCancelLayAside(); ok {
	//		repository.Update(entityRam.RamResourceRelationEntity{StateOrder: reverse.IndexInt8()}, info.ID)
	//	}
	//}
	return rt.Ok()
}

// PhysicalDeletion 物理删除
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamResourceRelationService) PhysicalDeletion(ctx *gin.Context, ids []string) (rt rg.Rs[string]) {
	c.log.Infof("ct=%+v", ids)
	if len(ids) < 1 {
		return rt.ErrorMessage("id错误")
	}
	cn := c.sv
	finds, b := cn.FindAllByIdStringIn(ids, repositoryPg.GetOption(ctx))
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	idsNew := make([]int64, 0)
	for _, info := range finds {
		c.log.Infof("id=%v,TenantId=%v", info.ID, info.TenantNo)
		idsNew = append(idsNew, info.ID)
	}
	if len(idsNew) > 0 {
		cn.DeleteByIds(idsNew, repositoryPg.GetOption(ctx))
	}
	return rt.Ok()
}

// Query 查询
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamResourceRelationService) Query(ctx *gin.Context, ct modRamResourceRelation.QueryCt) (rt rg.Rs[pagePg.PaginatorPg[modRamResourceRelation.Vo]]) {
	c.log.Infof("ct=%+v", ct)
	var query entityRam.RamResourceRelationEntity
	copier.Copy(&query, &ct)
	slice := make([]modRamResourceRelation.Vo, 0)
	rt.Data.Data = slice
	page, err := c.sv.FindAllPage(query, func(c *pagePg.PaginatorPg[*entityRam.RamResourceRelationEntity]) {
		c.PageNum = ct.PageNum
		c.PageSize = ct.PageSize
	})
	if nil != err {
		return rt.Ok()
	}

	if page.Total > 0 && page.Data != nil && len(page.Data) > 0 {

		pg := pagePg.NewPaginatorPg(func(c *pagePg.PaginatorPg[modRamResourceRelation.Vo]) {
			c.TotalPage = page.TotalPage
			c.Total = page.Total
			c.PageSize = page.PageSize
			c.PageNum = page.PageNum
		})
		//字段赋值
		for _, item := range page.Data {
			var vo modRamResourceRelation.Vo
			copier.Copy(&vo, &item)
			slice = append(slice, vo)
		}
		pg.Data = slice
		pg.Pageable = page.Pageable
		rt.Data = pg
		return rt.Ok()
	}
	return rt.Ok()
}

// SelectNodePublic 查询
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamResourceRelationService) SelectNodePublic(ctx *gin.Context, ct modRamResourceRelation.QueryCt) (rt rg.Rs[[]model.BaseNode]) {
	c.log.Infof("ct=%+v", ct)
	var query entityRam.RamResourceRelationEntity
	copier.Copy(&query, &ct)
	slice := make([]model.BaseNode, 0)
	rt.Data = slice
	infos := c.sv.FindAll(query, repositoryPg.GetOption(ctx))
	if len(infos) > 0 {
		//for _, item := range infos {
		//	slice = append(slice, model.BaseNode{Key: numberPg.Int64ToString(item.ID),
		//		Label: item.Name,
		//	})
		//}
		rt.Data = slice
	}
	return rt.Ok()
}

// SelectNodeAllPublic 查询
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamResourceRelationService) SelectNodeAllPublic(ctx *gin.Context, ct modRamResourceRelation.QueryCt) (rt rg.Rs[[]model.BaseNode]) {
	c.log.Infof("ct=%+v", ct)
	var query entityRam.RamResourceRelationEntity
	copier.Copy(&query, &ct)
	slice := make([]model.BaseNode, 0)
	rt.Data = slice
	infos := c.sv.FindAll(query, repositoryPg.GetOption(ctx))
	if len(infos) > 0 {
		for _, item := range infos {
			var vo modRamResourceRelation.Vo
			copier.Copy(&vo, &item)
			slice = append(slice, model.BaseNode{Key: numberPg.Int64ToString(item.ID),
				//Label:  item.Name,
				Extend: vo,
			})
		}
		rt.Data = slice
	}
	return rt.Ok()
}

// SelectPublic 查询
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamResourceRelationService) SelectPublic(ctx *gin.Context, ct modRamResourceRelation.QueryCt) (rt rg.Rs[[]modRamResourceRelation.Vo]) {
	c.log.Infof("ct=%+v", ct)
	var query entityRam.RamResourceRelationEntity
	copier.Copy(&query, &ct)
	rt.Data = []modRamResourceRelation.Vo{}
	infos := c.sv.FindAll(query, repositoryPg.GetOption(ctx))
	if len(infos) > 0 {
		slice := make([]modRamResourceRelation.Vo, 0)
		for _, item := range infos {
			var vo modRamResourceRelation.Vo
			copier.Copy(&vo, &item)
			slice = append(slice, vo)
		}
		rt.Data = slice
	}
	return rt.Ok()
}

// Selected 查询已选中的
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamResourceRelationService) Selected(ctx *gin.Context, code string) (rt rg.Rs[[]string]) {
	var query entityRam.RamResourceRelationEntity
	query.Mark = code
	slice := make([]string, 0)
	rt.Data = slice
	infos := c.sv.FindAll(query, repositoryPg.GetOption(ctx))
	if len(infos) > 0 {
		for _, item := range infos {
			slice = append(slice, numberPg.Int64ToString(item.ResourceId))
		}
		rt.Data = slice
	}
	return rt.Ok()
}

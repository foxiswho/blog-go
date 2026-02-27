package service

import (
	"context"
	"reflect"

	"github.com/foxiswho/blog-go/app/system/ram/model/modRamMenuRelation"
	"github.com/foxiswho/blog-go/infrastructure/entityRam"
	"github.com/foxiswho/blog-go/infrastructure/repositoryRam"
	"github.com/foxiswho/blog-go/pkg/consts/automatedPg"
	"github.com/foxiswho/blog-go/pkg/enum/state/enumStatePg"
	"github.com/foxiswho/blog-go/pkg/holderPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/foxiswho/blog-go/pkg/model"
	"github.com/foxiswho/blog-go/pkg/tools/noPg"
	"github.com/gin-gonic/gin"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
	"github.com/jinzhu/copier"
	"github.com/pangu-2/go-tools/tools/dbPg/pagePg"
	"github.com/pangu-2/go-tools/tools/numberPg"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
)

func init() {
	gs.Provide(new(RamMenuRelationService)).Init(func(s *RamMenuRelationService) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

// RamMenuRelationService 菜单关系
// @Description:
type RamMenuRelationService struct {
	sv  *repositoryRam.RamMenuRelationRepository `autowire:"?"`
	log *log2.Logger                             `autowire:"?"`
}

// Create 新增
//
//	@Description:
//	@receiver c
//	@param ct
//	@return rt
func (c *RamMenuRelationService) Create(ctx *gin.Context, ct modRamMenuRelation.CreateCt) (rt rg.Rs[string]) {
	c.log.Infof("ct=%+v", ct)
	if "" == ct.Name {
		return rt.ErrorMessage("名称不能为空")
	}
	if "" == ct.Code {
		return rt.ErrorMessage("编号不能为空")
	}
	r := c.sv
	//判断是否是自动,不是自动
	if !automatedPg.IsCreateCode(ct.Code) {
		//判断格式是否满足要求
		if !automatedPg.FormatVerify(ct.Code) {
			return rt.ErrorMessage("编号格式不能为空")
		}
		//不是自动
		_, result := r.FindByCode(ct.Code)
		if result {
			return rt.ErrorMessage("编号已存在")
		}
	}
	holder := holderPg.GetContextAccount(ctx)
	var info entityRam.RamMenuRelationEntity
	copier.Copy(&info, &ct)
	c.log.Infof("info%+v", info)
	info.TenantNo = holder.GetTenantNo()
	info.No = noPg.No()
	r.Create(&info)
	//自动设置编号
	if automatedPg.IsCreateCode(info.Code) {
		info.Code = numberPg.Int64ToString(info.ID)
		r.Update(entityRam.RamMenuRelationEntity{Code: info.Code}, info.ID)
	}
	c.log.Infof("save=%+v", info)
	return rg.OkData(numberPg.Int64ToString(info.ID))
}

// Update 更新
//
//	@Description:
//	@receiver c
//	@param ct
//	@return rt
func (c *RamMenuRelationService) Update(ctx *gin.Context, ct modRamMenuRelation.UpdateCt) (rt rg.Rs[string]) {
	c.log.Infof("ct=%+v", ct)
	if ct.ID < 1 {
		return rt.ErrorMessage("id错误")
	}
	if "" == ct.Name {
		return rt.ErrorMessage("名称不能为空")
	}
	if "" == ct.Code {
		return rt.ErrorMessage("编号不能为空")
	}
	r := c.sv
	_, result := r.FindByCodeAndIdNot(ct.Code, ct.ID.ToString())
	if result {
		return rt.ErrorMessage("编号已存在")
	}
	_, b := r.FindById(ct.ID.ToInt64())
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	var info entityRam.RamMenuRelationEntity
	copier.Copy(&info, &ct)
	r.Update(info, info.ID)
	return rt.Ok()
}

// Detail 详情
//
//	@Description:
//	@receiver c
//	@param id
func (c *RamMenuRelationService) Detail(ctx *gin.Context, id int64) (rt rg.Rs[modRamMenuRelation.Vo]) {
	if id < 1 {
		return rt.ErrorMessage("id错误")
	}
	find, b := c.sv.FindById(id)
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	var info modRamMenuRelation.Vo
	copier.Copy(&info, &find)
	return rt.OkData(info)
}

// Enable 启用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamMenuRelationService) Enable(ctx *gin.Context, ct model.BaseIdsCt[string]) (rt rg.Rs[string]) {
	c.log.Infof("ct=%+v", ct)
	return c.State(ctx, ct.Ids, enumStatePg.ENABLE)
}

// Disable 禁用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamMenuRelationService) Disable(ctx *gin.Context, ct model.BaseIdsCt[string]) (rt rg.Rs[string]) {
	c.log.Infof("ct=%+v", ct)
	return c.State(ctx, ct.Ids, enumStatePg.GetType(enumStatePg.DISABLE))
}

// State 状态 启用/禁用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamMenuRelationService) State(ctx *gin.Context, ids []string, state enumStatePg.State) (rt rg.Rs[string]) {
	if len(ids) < 1 {
		return rt.ErrorMessage("id错误")
	}
	r := c.sv
	_, b := r.FindAllByIdStringIn(ids)
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	//for _, info := range finds {
	//	if info.StateOrder != state.IndexInt8() {
	//		r.Update(entityRam.RamMenuRelationEntity{StateOrder: state.IndexInt8()}, info.ID)
	//	}
	//}
	return rt.Ok()
}

// StateEnableDisable 状态 设置 有效 停用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamMenuRelationService) StateEnableDisable(ctx *gin.Context, ids []string, state enumStatePg.State) (rt rg.Rs[string]) {
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
func (c *RamMenuRelationService) LogicalDeletion(ctx *gin.Context, ids []string) (rt rg.Rs[string]) {
	c.log.Infof("ct=%+v", ids)
	if len(ids) < 1 {
		return rt.ErrorMessage("id错误")
	}
	repository := c.sv
	finds, b := repository.FindAllByIdStringIn(ids)
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	if c.sv.Config().Data.Delete {
		for _, info := range finds {
			c.log.Infof("id=%v,TenantId=%v", info.ID, info.TenantNo)
		}
		repository.DeleteByIdsString(ids)
	} else {
		//for _, info := range finds {
		//	enum := enumStatePg.State(info.StateOrder)
		//	// 有效 停用，反转 为对应的 取消 弃置
		//	if ok, reverse := enum.ReverseEnableDisable(); ok {
		//		repository.Update(entityRam.RamMenuRelationEntity{StateOrder: reverse.IndexInt8()}, info.ID)
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
func (c *RamMenuRelationService) LogicalRecovery(ctx *gin.Context, ids []string) (rt rg.Rs[string]) {
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
	//		repository.Update(entityRam.RamMenuRelationEntity{StateOrder: reverse.IndexInt8()}, info.ID)
	//	}
	//}
	return rt.Ok()
}

// PhysicalDeletion 物理删除
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamMenuRelationService) PhysicalDeletion(ctx *gin.Context, ids []string) (rt rg.Rs[string]) {
	c.log.Infof("ct=%+v", ids)
	if len(ids) < 1 {
		return rt.ErrorMessage("id错误")
	}
	cn := c.sv
	finds, b := cn.FindAllByIdStringIn(ids)
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	idsNew := make([]int64, 0)
	for _, info := range finds {
		c.log.Infof("id=%v,TenantId=%v", info.ID, info.TenantNo)
		idsNew = append(idsNew, info.ID)
	}
	if len(idsNew) > 0 {
		cn.DeleteByIds(idsNew)
	}
	return rt.Ok()
}

// Query 查询
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamMenuRelationService) Query(ctx *gin.Context, ct modRamMenuRelation.QueryCt) (rt rg.Rs[pagePg.PaginatorPg[modRamMenuRelation.Vo]]) {
	c.log.Infof("ct=%+v", ct)
	var query entityRam.RamMenuRelationEntity
	copier.Copy(&query, &ct)
	slice := make([]modRamMenuRelation.Vo, 0)
	rt.Data.Data = slice
	r := c.sv
	page, err := r.FindAllPageQuery(query, func(p *pagePg.PageCondition[*entityRam.RamMenuRelationEntity]) {
		p.PageOption = func(c *pagePg.PaginatorPg[*entityRam.RamMenuRelationEntity]) {
			c.PageNum = ct.PageNum
			c.PageSize = ct.PageSize
		}
		p.Condition = r.DbModel().Order("create_at asc")
		//自定义查询
		if "" != ct.Wd {
			p.Condition.Where("name like ?", "%"+ct.Wd+"%")
		}
	})
	if nil != err {
		return rt.Ok()
	}

	if page.Total > 0 && page.Data != nil && len(page.Data) > 0 {

		pg := pagePg.NewPaginatorPg(func(c *pagePg.PaginatorPg[modRamMenuRelation.Vo]) {
			c.TotalPage = page.TotalPage
			c.Total = page.Total
			c.PageSize = page.PageSize
			c.PageNum = page.PageNum
		})
		//字段赋值
		for _, item := range page.Data {
			var vo modRamMenuRelation.Vo
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
func (c *RamMenuRelationService) SelectNodePublic(ctx *gin.Context, ct modRamMenuRelation.QueryCt) (rt rg.Rs[[]model.BaseNode]) {
	c.log.Infof("ct=%+v", ct)
	var query entityRam.RamMenuRelationEntity
	copier.Copy(&query, &ct)
	slice := make([]model.BaseNode, 0)
	rt.Data = slice
	infos := c.sv.FindAll(query)
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
func (c *RamMenuRelationService) SelectNodeAllPublic(ctx *gin.Context, ct modRamMenuRelation.QueryCt) (rt rg.Rs[[]model.BaseNode]) {
	c.log.Infof("ct=%+v", ct)
	var query entityRam.RamMenuRelationEntity
	copier.Copy(&query, &ct)
	slice := make([]model.BaseNode, 0)
	rt.Data = slice
	infos := c.sv.FindAll(query)
	if len(infos) > 0 {
		for _, item := range infos {
			var vo modRamMenuRelation.Vo
			copier.Copy(&vo, &item)
			slice = append(slice, model.BaseNode{Key: numberPg.Int64ToString(item.ID),
				//Label:  item.Name,
				Extend: vo,
				No:     item.No,
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
func (c *RamMenuRelationService) SelectPublic(ctx *gin.Context, ct modRamMenuRelation.QueryCt) (rt rg.Rs[[]modRamMenuRelation.Vo]) {
	c.log.Infof("ct=%+v", ct)
	var query entityRam.RamMenuRelationEntity
	copier.Copy(&query, &ct)
	rt.Data = []modRamMenuRelation.Vo{}
	infos := c.sv.FindAll(query)
	if len(infos) > 0 {
		slice := make([]modRamMenuRelation.Vo, 0)
		for _, item := range infos {
			var vo modRamMenuRelation.Vo
			copier.Copy(&vo, &item)
			slice = append(slice, vo)
		}
		rt.Data = slice
	}
	return rt.Ok()
}

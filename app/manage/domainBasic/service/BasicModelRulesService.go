package service

import (
	"context"
	"reflect"
	"strings"

	"github.com/foxiswho/blog-go/app/manage/domainBasic/model/modBasicModelRules"
	"github.com/foxiswho/blog-go/app/manage/domainBasic/service/modelRules"
	"github.com/foxiswho/blog-go/infrastructure/entityBasic"
	"github.com/foxiswho/blog-go/infrastructure/repositoryBasic"
	"github.com/foxiswho/blog-go/pkg/enum/request/enumParameterPg"
	"github.com/foxiswho/blog-go/pkg/enum/state/enumStatePg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/foxiswho/blog-go/pkg/model"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/gin-gonic/gin"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
	"github.com/jinzhu/copier"
	"github.com/pangu-2/go-tools/tools/dbPg/pagePg"
	"github.com/pangu-2/go-tools/tools/numberPg"
	"github.com/pangu-2/go-tools/tools/strPg"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
)

func init() {
	gs.Provide(new(BasicModelRulesService)).Init(func(s *BasicModelRulesService) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type BasicModelRulesService struct {
	log *log2.Logger                               `autowire:"?"`
	sv  *repositoryBasic.BasicModelRulesRepository `autowire:"?"`
	sp  *modelRules.Sp                             `autowire:"?"`
}

func (c *BasicModelRulesService) CreateUpdateData(ctx *gin.Context, ct modBasicModelRules.CreateUpdateDataCt) (rt rg.Rs[string]) {
	return modelRules.NewCreateUpdateData(c.sp, ct, true).Process(ctx)
}
func (c *BasicModelRulesService) CreateUpdate(ctx *gin.Context, ct modBasicModelRules.CreateUpdateCt) (rt rg.Rs[string]) {
	return modelRules.NewCreateUpdate(c.sp, ct, true).Process(ctx)
}

// Detail 详情
//func (c *BasicModelRulesService) Detail(ctx *gin.Context, id string) (rt rg.Rs[modBasicModelRules.CreateUpdateCt]) {
//	return configEvent.NewDetail(c.sp).Process(ctx, id)
//}

// Enable 启用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *BasicModelRulesService) Enable(ctx *gin.Context, ct model.BaseIdsCt[string]) (rt rg.Rs[string]) {
	return c.State(ctx, ct.Ids, enumStatePg.ENABLE)
}

// Disable 禁用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *BasicModelRulesService) Disable(ctx *gin.Context, ct model.BaseIdsCt[string]) (rt rg.Rs[string]) {
	return c.State(ctx, ct.Ids, enumStatePg.GetType(enumStatePg.DISABLE))
}

// State 状态 启用/禁用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *BasicModelRulesService) State(ctx *gin.Context, ids []string, state enumStatePg.State) (rt rg.Rs[string]) {
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
			r.Update(entityBasic.BasicModelRulesEntity{State: state.IndexInt8()}, info.ID)
		}
	}
	return rt.Ok()
}

// StateEnableDisable 状态 设置 有效 停用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *BasicModelRulesService) StateEnableDisable(ctx *gin.Context, ids []string, state enumStatePg.State) (rt rg.Rs[string]) {
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
func (c *BasicModelRulesService) LogicalDeletion(ctx *gin.Context, ct model.BaseIdsCt[string]) (rt rg.Rs[string]) {
	if nil == ct.Ids || len(ct.Ids) < 1 {
		return rt.ErrorMessage("id错误")
	}
	if ct.Extend == nil {
		return rt.ErrorMessage("参数错误")
	}
	fieldNo := ""
	if obj, ok := ct.Extend["fieldNo"]; ok {
		fieldNo = obj.(string)
	}
	if strPg.IsBlank(fieldNo) {
		return rt.ErrorMessage("字段错误")
	}
	ids := make([]string, 0)
	for _, id := range ct.Ids {
		if strPg.IsNotBlank(id) {
			ids = append(ids, strings.TrimSpace(id))
		}
	}
	if len(ids) < 1 {
		return rt.ErrorMessage("请选择数据")
	}
	repository := c.sv
	finds, b := repository.FindAllByIdStringIn(ids)
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	if c.sv.Config().Data.Delete {
		idsInt := make([]int64, 0)
		for _, info := range finds {
			c.log.Infof("id=%v", info.ID)
			if info.FieldNo == fieldNo {
				idsInt = append(idsInt, info.ID)
			}
		}
		if len(idsInt) > 0 {
			repository.DeleteByIds(idsInt)
		}
	} else {
		for _, info := range finds {
			if info.FieldNo != fieldNo {
				continue
			}
			enum := enumStatePg.State(info.State)
			// 有效 停用，反转 为对应的 取消 弃置
			if ok, reverse := enum.ReverseEnableDisable(); ok {
				repository.Update(entityBasic.BasicModelRulesEntity{State: reverse.IndexInt8()}, info.ID)
			}
		}
	}

	return rt.Ok()
}

// LogicalRecovery 逻辑删除恢复
//
//	@Description:
//	@receiver c
//	@param ct
func (c *BasicModelRulesService) LogicalRecovery(ctx *gin.Context, ids []string) (rt rg.Rs[string]) {
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
			repository.Update(entityBasic.BasicModelRulesEntity{State: reverse.IndexInt8()}, info.ID)
		}
	}
	return rt.Ok()
}

// PhysicalDeletion 物理删除
//
//	@Description:
//	@receiver c
//	@param ct
func (c *BasicModelRulesService) PhysicalDeletion(ctx *gin.Context, ids []string) (rt rg.Rs[string]) {
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
		c.log.Infof("id=%v", info.ID)
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
func (c *BasicModelRulesService) Query(ctx *gin.Context, ct modBasicModelRules.QueryCt) (rt rg.Rs[pagePg.Paginator[modBasicModelRules.Vo]]) {
	c.log.Infof("ct=%+v", ct)
	var query entityBasic.BasicModelRulesEntity
	copier.Copy(&query, &ct)
	slice := make([]modBasicModelRules.Vo, 0)
	rt.Data.Data = slice
	r := c.sv
	page, err := r.FindAllPageQuery(ctx, query, repositoryPg.WithOptionPg(func(arg *repositoryPg.OptionParams) {
		if ct.PageSize < 1 {
			ct.PageSize = 20
		}
		arg.Pageable = new(pagePg.PageablePageSize(0, ct.PageNum, ct.PageSize))
		//自定义查询
		arg.Db.Order("create_at asc")
		if strPg.IsNotBlank(ct.Wd) {
			arg.Db.Where("name like ?", "%"+ct.Wd+"%").
				Or("description like ?", "%"+ct.Wd+"%")
		}
	}), repositoryPg.WithCtx(ctx))

	if nil != err {
		return rt.Ok()
	}

	if page.Total > 0 && page.Data != nil && len(page.Data) > 0 {
		pg := pagePg.NewPaginatorByPageable[modBasicModelRules.Vo](page.Pageable)
		//字段赋值
		for _, item := range page.Data {
			var vo modBasicModelRules.Vo
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

// SelectNodeAllPublic 查询
//
//	@Description:
//	@receiver c
//	@param ct
func (c *BasicModelRulesService) SelectNodeAllPublic(ctx *gin.Context, ct modBasicModelRules.QueryPublicCt) (rt rg.Rs[[]model.BaseNodeNo]) {
	var query entityBasic.BasicModelRulesEntity
	copier.Copy(&query, &ct)
	slice := make([]model.BaseNodeNo, 0)
	rt.Data = slice
	infos := c.sv.FindAll(query)
	if len(infos) > 0 {
		for _, item := range infos {
			var vo modBasicModelRules.Vo
			copier.Copy(&vo, &item)
			//
			code := model.BaseNodeNo{
				Key:      item.No,
				Id:       item.No,
				No:       item.No,
				Label:    item.Name,
				ParentNo: "",
				ParentId: "",
				Extend:   vo,
			}
			//编码
			if !enumParameterPg.NodeQueryByNo.IsEqual(ct.By) {
				code.Key = numberPg.Int64ToString(item.ID)
				code.Id = code.Key
				code.ParentId = ""
			}
			slice = append(slice, code)
		}
		rt.Data = slice
	}
	return rt.Ok()
}

// AllByValueNo 查询
//
//	@Description:
//	@receiver c
//	@param ct
func (c *BasicModelRulesService) AllByValueNo(ctx *gin.Context, ct modBasicModelRules.QueryPublicCt) (rt rg.Rs[[]modBasicModelRules.Vo]) {
	var query entityBasic.BasicModelRulesEntity
	copier.Copy(&query, &ct)
	//
	slice := make([]modBasicModelRules.Vo, 0)
	rt.Data = slice
	//
	if strPg.IsBlank(ct.ValueNo) {
		return rt.ErrorMessage("模型/事件编号错误")
	}
	infos := c.sv.FindAll(query)
	if len(infos) > 0 {
		for _, item := range infos {
			var vo modBasicModelRules.Vo
			copier.Copy(&vo, &item)
			slice = append(slice, vo)
		}
		rt.Data = slice
	}
	return rt.Ok()
}

// ExistName 查重
//
//	@Description:
//	@receiver c
//	@param ct
func (c *BasicModelRulesService) ExistName(ctx *gin.Context, ct model.BaseExistWdCt[string]) (rt rg.Rs[string]) {
	if "" == ct.Wd {
		return rt.ErrorMessage("关键词不能为空")
	}
	id := "0"
	if strPg.IsNotBlank(ct.Id) {
		id = ct.Id
	}
	_, result := c.sv.FindByNameAndIdNot(ct.Wd, id)
	if result {
		return rt.ErrorMessage("重复，已存在")
	}
	return rt.OkMessage("可以使用")
}

// ExistCode 查重
//
//	@Description:
//	@receiver c
//	@param ct
func (c *BasicModelRulesService) ExistCode(ctx *gin.Context, ct model.BaseExistWdCt[string]) (rt rg.Rs[string]) {
	if "" == ct.Wd {
		return rt.ErrorMessage("关键词不能为空")
	}
	id := "0"
	if strPg.IsNotBlank(ct.Id) {
		id = ct.Id
	}
	_, result := c.sv.FindByCodeAndIdNot(ct.Wd, id)
	if result {
		return rt.ErrorMessage("重复，已存在")
	}
	return rt.OkMessage("可以使用")
}

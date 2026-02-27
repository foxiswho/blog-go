package service

import (
	"context"
	"reflect"

	"github.com/foxiswho/blog-go/app/manage/domainBasic/model/modBasicAccountApplyDenyList"
	"github.com/foxiswho/blog-go/infrastructure/entityBasic"
	"github.com/foxiswho/blog-go/infrastructure/repositoryBasic"
	"github.com/foxiswho/blog-go/pkg/consts/constsRam/typeDomainPg"
	"github.com/foxiswho/blog-go/pkg/enum/enumCommonPg/typeExprPg"
	"github.com/foxiswho/blog-go/pkg/enum/enumCommonPg/typeFieldPg"
	"github.com/foxiswho/blog-go/pkg/enum/enumCommonPg/typeSysPg"
	"github.com/foxiswho/blog-go/pkg/enum/request/enumParameterPg"
	"github.com/foxiswho/blog-go/pkg/enum/state/enumStatePg"
	"github.com/foxiswho/blog-go/pkg/holderPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/foxiswho/blog-go/pkg/model"
	"github.com/foxiswho/blog-go/pkg/tools/excelPg"
	"github.com/foxiswho/blog-go/pkg/tools/noPg"
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
	gs.Provide(new(BasicAccountApplyDenyListService)).Init(func(s *BasicAccountApplyDenyListService) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

// BasicAccountApplyDenyListService 账号申请禁用列表
// @Description:
type BasicAccountApplyDenyListService struct {
	sv  *repositoryBasic.BasicAccountApplyDenyListEntityRepository `autowire:"?"`
	log *log2.Logger                                               `autowire:"?"`
}

// Create 新增
//
//	@Description:
//	@receiver c
//	@param ct
//	@return rt
func (c *BasicAccountApplyDenyListService) Create(ctx *gin.Context, ct modBasicAccountApplyDenyList.CreateCt) (rt rg.Rs[string]) {
	c.log.Infof("ct=%#v", ct)
	var info entityBasic.BasicAccountApplyDenyListEntity
	err := copier.Copy(&info, &ct)
	if err != nil {
		c.log.Infof("copier.Copy error: %+v", err)
	}
	if "" == ct.Name {
		return rt.ErrorMessage("名称不能为空")
	}
	if strPg.IsBlank(ct.Expr) {
		return rt.ErrorMessage("表达式不能为空")
	}
	if strPg.IsBlank(ct.TypeDomain) {
		return rt.ErrorMessage("域类型不能为空")
	}
	if strPg.IsBlank(ct.TypeExpr) {
		return rt.ErrorMessage("表达式类型不能为空")
	}
	if strPg.IsBlank(ct.TypeField) {
		return rt.ErrorMessage("字段类型不能为空")
	}
	if strPg.IsBlank(ct.TypeSys) {
		return rt.ErrorMessage("系统类型不能为空")
	}
	if _, ok := typeSysPg.IsExistTypeSys(ct.TypeSys); !ok {
		return rt.ErrorMessage("类型错误")
	}
	if _, ok := typeDomainPg.IsExistTypeDomain(ct.TypeDomain); !ok {
		return rt.ErrorMessage("域类型错误")
	}
	if _, ok := typeExprPg.IsExistTypeExpr(ct.TypeExpr); !ok {
		return rt.ErrorMessage("表达式类型错误")
	}
	if _, ok := typeFieldPg.IsExistTypeField(ct.TypeField); !ok {
		return rt.ErrorMessage("字段类型错误")
	}
	r := c.sv
	if strPg.IsNotBlank(ct.Expr) {
		_, result := r.FindByExprAndIdNot(ct.Expr, "0")
		if result {
			return rt.ErrorMessage("标志已存在")
		}
	}

	holder := holderPg.GetContextAccount(ctx)
	info.No = noPg.No()
	info.TenantNo = holder.GetTenantNo()
	c.log.Infof("info%+v", info)
	err, _ = r.Create(&info)
	if err != nil {
		return rt.ErrorMessage("保存失败 " + err.Error())
	}
	return rg.OkData(numberPg.Int64ToString(info.ID))
}

// Update 更新
//
//	@Description:
//	@receiver c
//	@param ct
//	@return rt
func (c *BasicAccountApplyDenyListService) Update(ctx *gin.Context, ct modBasicAccountApplyDenyList.UpdateCt) (rt rg.Rs[string]) {
	c.log.Infof("ct=%#v", ct)
	var info entityBasic.BasicAccountApplyDenyListEntity
	copier.Copy(&info, &ct)
	if ct.ID < 1 {
		return rt.ErrorMessage("id错误")
	}
	if "" == ct.Name {
		return rt.ErrorMessage("名称不能为空")
	}
	if strPg.IsBlank(ct.Expr) {
		return rt.ErrorMessage("表达式不能为空")
	}
	if strPg.IsBlank(ct.TypeDomain) {
		return rt.ErrorMessage("域类型不能为空")
	}
	if strPg.IsBlank(ct.TypeExpr) {
		return rt.ErrorMessage("表达式类型不能为空")
	}
	if strPg.IsBlank(ct.TypeField) {
		return rt.ErrorMessage("字段类型不能为空")
	}
	if strPg.IsBlank(ct.TypeSys) {
		return rt.ErrorMessage("系统类型不能为空")
	}
	if _, ok := typeSysPg.IsExistTypeSys(ct.TypeSys); !ok {
		return rt.ErrorMessage("类型错误")
	}
	if _, ok := typeDomainPg.IsExistTypeDomain(ct.TypeDomain); !ok {
		return rt.ErrorMessage("域类型错误")
	}
	if _, ok := typeExprPg.IsExistTypeExpr(ct.TypeExpr); !ok {
		return rt.ErrorMessage("表达式类型错误")
	}
	if _, ok := typeFieldPg.IsExistTypeField(ct.TypeField); !ok {
		return rt.ErrorMessage("字段类型错误")
	}
	r := c.sv
	if strPg.IsNotBlank(ct.Expr) {
		_, result := r.FindByExprAndIdNot(ct.Expr, ct.ID.ToString())
		if result {
			return rt.ErrorMessage("标志已存在")
		}
	}
	find, b := r.FindById(ct.ID.ToInt64())
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	info.No = ""
	err := r.Update(info, find.ID)
	if err != nil {
		c.log.Errorf("update error=%+v", err)
		return rt.ErrorMessage(err.Error())
	}
	c.log.Infof("save.info=%+v", info)
	return rt.Ok()
}

// CacheOverride 缓存重载
//
//	@Description:
//	@receiver c
func (c *BasicAccountApplyDenyListService) CacheOverride(ctx *gin.Context) {
}

// Detail 详情
//
//	@Description:
//	@receiver c
//	@param id
func (c *BasicAccountApplyDenyListService) Detail(ctx *gin.Context, id int64) (rt rg.Rs[modBasicAccountApplyDenyList.Vo]) {
	if id < 1 {
		return rt.ErrorMessage("id错误")
	}
	find, b := c.sv.FindById(id)
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	var info modBasicAccountApplyDenyList.Vo
	copier.Copy(&info, &find)
	return rt.OkData(info)
}

// Enable 启用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *BasicAccountApplyDenyListService) Enable(ctx *gin.Context, ct model.BaseIdsCt[string]) (rt rg.Rs[string]) {
	return c.State(ctx, ct.Ids, enumStatePg.ENABLE)
}

// Disable 禁用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *BasicAccountApplyDenyListService) Disable(ctx *gin.Context, ct model.BaseIdsCt[string]) (rt rg.Rs[string]) {
	return c.State(ctx, ct.Ids, enumStatePg.GetType(enumStatePg.DISABLE))
}

// State 状态 启用/禁用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *BasicAccountApplyDenyListService) State(ctx *gin.Context, ids []string, state enumStatePg.State) (rt rg.Rs[string]) {
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
			r.Update(entityBasic.BasicAccountApplyDenyListEntity{State: state.IndexInt8()}, info.ID)
		}
	}
	return rt.Ok()
}

// StateEnableDisable 状态 设置 有效 停用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *BasicAccountApplyDenyListService) StateEnableDisable(ctx *gin.Context, ids []string, state enumStatePg.State) (rt rg.Rs[string]) {
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
func (c *BasicAccountApplyDenyListService) LogicalDeletion(ctx *gin.Context, ids []string) (rt rg.Rs[string]) {
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
			c.log.Infof("id=%v", info.ID)
		}
		repository.DeleteByIdsString(ids)
	} else {
		for _, info := range finds {
			enum := enumStatePg.State(info.State)
			// 有效 停用，反转 为对应的 取消 弃置
			if ok, reverse := enum.ReverseEnableDisable(); ok {
				repository.Update(entityBasic.BasicAccountApplyDenyListEntity{State: reverse.IndexInt8()}, info.ID)
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
func (c *BasicAccountApplyDenyListService) LogicalRecovery(ctx *gin.Context, ids []string) (rt rg.Rs[string]) {
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
			repository.Update(entityBasic.BasicAccountApplyDenyListEntity{State: reverse.IndexInt8()}, info.ID)
		}
	}
	return rt.Ok()
}

// PhysicalDeletion 物理删除
//
//	@Description:
//	@receiver c
//	@param ct
func (c *BasicAccountApplyDenyListService) PhysicalDeletion(ctx *gin.Context, ids []string) (rt rg.Rs[string]) {
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
func (c *BasicAccountApplyDenyListService) Query(ctx *gin.Context, ct modBasicAccountApplyDenyList.QueryCt) (rt rg.Rs[pagePg.PaginatorPg[modBasicAccountApplyDenyList.Vo]]) {
	c.log.Infof("ct=%+v", ct)
	var query entityBasic.BasicAccountApplyDenyListEntity
	copier.Copy(&query, &ct)
	slice := make([]modBasicAccountApplyDenyList.Vo, 0)
	rt.Data.Data = slice
	r := c.sv
	page, err := r.FindAllPageQuery(query, func(p *pagePg.PageCondition[*entityBasic.BasicAccountApplyDenyListEntity]) {
		p.PageOption = func(c *pagePg.PaginatorPg[*entityBasic.BasicAccountApplyDenyListEntity]) {
			c.PageNum = ct.PageNum
			c.PageSize = ct.PageSize
		}
		//自定义查询
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
		pg := pagePg.NewPaginatorPg(func(c *pagePg.PaginatorPg[modBasicAccountApplyDenyList.Vo]) {
			c.TotalPage = page.TotalPage
			c.Total = page.Total
			c.PageSize = page.PageSize
			c.PageNum = page.PageNum
		})
		//字段赋值
		for _, item := range page.Data {
			var vo modBasicAccountApplyDenyList.Vo
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
func (c *BasicAccountApplyDenyListService) SelectNodePublic(ctx *gin.Context, ct modBasicAccountApplyDenyList.QueryPublicCt) (rt rg.Rs[[]model.BaseNodeNo]) {
	var query entityBasic.BasicAccountApplyDenyListEntity
	copier.Copy(&query, &ct)
	slice := make([]model.BaseNodeNo, 0)
	rt.Data = slice
	infos := c.sv.FindAll(query)
	if len(infos) > 0 {
		for _, item := range infos {
			var vo modBasicAccountApplyDenyList.Vo
			copier.Copy(&vo, &item)
			code := model.BaseNodeNo{
				Key:    item.No,
				Id:     item.No,
				No:     item.No,
				Label:  item.Name,
				Extend: vo,
			}
			//编码
			if !enumParameterPg.NodeQueryByNo.IsEqual(ct.By) {
				code.Key = numberPg.Int64ToString(item.ID)
				code.Id = code.Key
			}
			slice = append(slice, code)
		}
		rt.Data = slice
	}
	return rt.Ok()
}

// SelectNodeAllPublic 查询
//
//	@Description:
//	@receiver c
//	@param ct
func (c *BasicAccountApplyDenyListService) SelectNodeAllPublic(ctx *gin.Context, ct modBasicAccountApplyDenyList.QueryPublicCt) (rt rg.Rs[[]model.BaseNodeNo]) {
	var query entityBasic.BasicAccountApplyDenyListEntity
	copier.Copy(&query, &ct)
	slice := make([]model.BaseNodeNo, 0)
	rt.Data = slice
	infos := c.sv.FindAll(query)
	if len(infos) > 0 {
		for _, item := range infos {
			var vo modBasicAccountApplyDenyList.Vo
			copier.Copy(&vo, &item)
			code := model.BaseNodeNo{
				Key:    item.No,
				Id:     item.No,
				No:     item.No,
				Label:  item.Name,
				Extend: vo,
			}
			//编码
			if !enumParameterPg.NodeQueryByNo.IsEqual(ct.By) {
				code.Key = numberPg.Int64ToString(item.ID)
				code.Id = code.Key
			}
			slice = append(slice, code)
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
func (c *BasicAccountApplyDenyListService) SelectPublic(ctx *gin.Context, ct modBasicAccountApplyDenyList.QueryPublicCt) (rt rg.Rs[[]modBasicAccountApplyDenyList.Vo]) {
	c.log.Infof("ct=%+v", ct)
	var query entityBasic.BasicAccountApplyDenyListEntity
	copier.Copy(&query, &ct)
	slice := make([]modBasicAccountApplyDenyList.Vo, 0)
	rt.Data = slice
	infos := c.sv.FindAll(query)
	if len(infos) > 0 {
		for _, item := range infos {
			var vo modBasicAccountApplyDenyList.Vo
			copier.Copy(&vo, &item)
			slice = append(slice, vo)
		}
		rt.Data = slice
	}
	return rt.Ok()
}

// ExportExcel 导出
//
//	@Description:
//	@receiver c
//	@param ct
func (c *BasicAccountApplyDenyListService) ExportExcel(ctx *gin.Context, ct modBasicAccountApplyDenyList.QueryCt) {
	c.log.Infof("ct=%+v", ct)
	var query entityBasic.BasicAccountApplyDenyListEntity
	copier.Copy(&query, &ct)
	infos := c.sv.FindAll(query)
	if len(infos) > 0 {
		slice := make([]interface{}, 0)
		for _, item := range infos {
			var vo modBasicAccountApplyDenyList.Vo
			copier.Copy(&vo, &item)
			slice = append(slice, vo)
		}
		c.log.Infof("导出数据 %+v", slice)
		strings := []string{"ID", "名称", "名称外文",
			"编号代号",
			"全称",
			"状态:1启用;2禁用",
			"删除:1是;2否",
			"描述",
			"创建时间",
			"更新时间",
			"创建人",
			"更新人", "组织id"}
		err := excelPg.ExportExcelByStruct(ctx, strings, slice, "area", "Sheet1")
		if nil != err {
			r := rg.Rs[string]{}
			ctx.JSON(200, r.ErrorMessage(err.Error()))
		}
	} else {
		r := rg.Rs[string]{}
		ctx.JSON(200, r.ErrorMessage("没有任何数据"))
	}

}

// ExistName 查重
//
//	@Description:
//	@receiver c
//	@param ct
func (c *BasicAccountApplyDenyListService) ExistName(ctx *gin.Context, ct model.BaseExistWdCt[string]) (rt rg.Rs[string]) {
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

// ExistExpr 查重
//
//	@Description:
//	@receiver c
//	@param ct
func (c *BasicAccountApplyDenyListService) ExistExpr(ctx *gin.Context, ct model.BaseExistWdCt[string]) (rt rg.Rs[string]) {
	if "" == ct.Wd {
		return rt.ErrorMessage("关键词不能为空")
	}
	id := "0"
	if strPg.IsNotBlank(ct.Id) {
		id = ct.Id
	}
	_, result := c.sv.FindByExprAndIdNot(ct.Wd, id)
	if result {
		return rt.ErrorMessage("重复，已存在")
	}
	return rt.OkMessage("可以使用")
}

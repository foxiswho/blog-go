package service

import (
	"context"
	"reflect"

	"github.com/foxiswho/blog-go/app/manage/domainBasic/model/modBasicArea"
	"github.com/foxiswho/blog-go/infrastructure/entityBasic"
	"github.com/foxiswho/blog-go/infrastructure/repositoryBasic"
	"github.com/foxiswho/blog-go/pkg/consts/automatedPg"
	"github.com/foxiswho/blog-go/pkg/consts/constNodePg"
	"github.com/foxiswho/blog-go/pkg/enum/request/enumParameterPg"
	"github.com/foxiswho/blog-go/pkg/enum/state/enumStatePg"
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
	"github.com/pangu-2/go-tools/tools/slicePg"
	"github.com/pangu-2/go-tools/tools/strPg"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
)

func init() {
	gs.Provide(new(BasicAreaService)).Init(func(s *BasicAreaService) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

// BasicAreaService 省市区
// @Description:
type BasicAreaService struct {
	sv  *repositoryBasic.BasicAreaRepository `autowire:"?"`
	log *log2.Logger                         `autowire:"?"`
}

// Create 新增
//
//	@Description:
//	@receiver c
//	@param ct
//	@return rt
func (c *BasicAreaService) Create(ctx *gin.Context, ct modBasicArea.CreateCt) (rt rg.Rs[string]) {
	c.log.Infof("ct=%#v", ct)
	var info entityBasic.BasicAreaEntity
	err := copier.Copy(&info, &ct)
	if err != nil {
		c.log.Infof("copier.Copy error: %+v", err)
	}
	if "" == ct.Name {
		return rt.ErrorMessage("名称不能为空")
	}
	if strPg.IsBlank(ct.Code) {
		info.Code = automatedPg.CREATE_CODE
	}
	parent := &entityBasic.BasicAreaEntity{}
	r := c.sv
	//判断是否是自动,不是自动
	if !automatedPg.IsCreateCode(info.Code) {
		//判断格式是否满足要求
		if !automatedPg.FormatVerify(info.Code) {
			return rt.ErrorMessage("标志格式不能为空")
		}
		//不是自动
		_, result := r.FindByCode(info.Code)
		if result {
			return rt.ErrorMessage("标志已存在")
		}
	}
	result := false
	if strPg.IsNotBlank(ct.ParentNo) {
		parent, result = r.FindByNo(ct.ParentNo)
		if !result {
			return rt.ErrorMessage("上级不存在")
		}
	}
	info.No = noPg.No()
	//自动设置编号
	if automatedPg.IsCreateCode(info.Code) {
		info.Code = strPg.GenerateNumberId22()
	}
	c.log.Infof("info%+v", info)
	err, _ = r.Create(&info)
	if err != nil {
		return rt.ErrorMessage("保存失败 " + err.Error())
	}
	//设置上级 link
	if strPg.IsNotBlank(ct.ParentNo) {
		info.IdLink = constNodePg.NoLinkAssemble(parent.IdLink, numberPg.Int64ToString(info.ID))
		info.NoLink = constNodePg.NoLinkAssemble(parent.NoLink, info.No)
		info.ParentNo = parent.No
		info.ParentId = numberPg.Int64ToString(parent.ID)
	} else {
		info.IdLink = constNodePg.NoLinkDefault(numberPg.Int64ToString(info.ID))
		info.NoLink = constNodePg.NoLinkDefault(info.No)
		info.ParentId = ""
		info.ParentNo = ""
	}
	err = r.Update(info, info.ID)
	if err != nil {
		return rt.ErrorMessage(err.Error())
	}
	return rg.OkData(numberPg.Int64ToString(info.ID))
}

// Update 更新
//
//	@Description:
//	@receiver c
//	@param ct
//	@return rt
func (c *BasicAreaService) Update(ctx *gin.Context, ct modBasicArea.UpdateCt) (rt rg.Rs[string]) {
	c.log.Infof("ct=%#v", ct)
	var info entityBasic.BasicAreaEntity
	copier.Copy(&info, &ct)
	if ct.ID < 1 {
		return rt.ErrorMessage("id错误")
	}
	if "" == ct.Name {
		return rt.ErrorMessage("名称不能为空")
	}
	if strPg.IsBlank(ct.Code) {
		info.Code = ""
	}
	r := c.sv
	if strPg.IsBlank(ct.Code) {
		info.Code = ""
	} else {
		_, result := r.FindByCodeAndIdNot(ct.Code, ct.ID.ToString())
		if result {
			return rt.ErrorMessage("标志已存在")
		}
	}
	find, b := r.FindById(ct.ID.ToInt64())
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	//上级
	parent := &entityBasic.BasicAreaEntity{}
	var childData []*entityBasic.BasicAreaEntity
	if strPg.IsNotBlank(ct.ParentNo) {
		result := false
		parent, result = r.FindByNo(ct.ParentNo)
		if !result {
			return rt.ErrorMessage("上级不存在")
		}
		if parent.ID == ct.ID.ToInt64() {
			return rt.ErrorMessage("上级不能等于自己")
		}
		//新的ID 不等于 旧的上级时,检测是否已经 在新的子集已存在
		if parent.No != find.ParentNo {
			result2 := false
			childData, result2 = r.FindAllByNoLink(find.IdLink)
			if result2 {
				//c.log.Infof("data=%+v \n", childData)
				for _, item := range childData {
					if item.No == parent.No {
						return rt.ErrorMessage("无法保存，不能设置为自己的子集")
					}
				}
			}
		}
	}

	//设置上级 link
	if strPg.IsNotBlank(ct.ParentNo) {
		info.IdLink = constNodePg.NoLinkAssemble(parent.IdLink, numberPg.Int64ToString(find.ID))
		info.NoLink = constNodePg.NoLinkAssemble(parent.NoLink, find.No)
		info.ParentNo = parent.No
		info.ParentId = numberPg.Int64ToString(parent.ID)
	} else {
		info.IdLink = constNodePg.NoLinkDefault(numberPg.Int64ToString(find.ID))
		info.NoLink = constNodePg.NoLinkDefault(find.No)
		info.ParentNo = ""
		info.ParentId = ""
	}
	info.No = ""
	c.log.Infof("info.IdLink=%+v", info.IdLink)
	err := r.Update(info, info.ID)
	if err != nil {
		c.log.Errorf("update error=%+v", err)
		return rt.ErrorMessage(err.Error())
	}
	c.log.Infof("save.info=%+v", info)
	//更改上级后，相关子集修改
	if strPg.IsNotBlank(ct.ParentNo) && nil != childData {
		maps := slicePg.ToMapArray(childData, func(t *entityBasic.BasicAreaEntity) (string, *entityBasic.BasicAreaEntity) {
			if strPg.IsBlank(t.ParentNo) {
				return constNodePg.ROOT, t
			}
			return t.ParentNo, t
		})
		if strPg.IsBlank(info.ParentNo) {
			info.ParentNo = constNodePg.ROOT
		}
		for _, item := range maps[info.ParentNo] {
			item.IdLink = constNodePg.NoLinkAssemble(info.IdLink, numberPg.Int64ToString(find.ID))
			item.NoLink = constNodePg.NoLinkAssemble(info.NoLink, item.No)
			c.childParentIdLink(maps, item)
		}
		c.log.Infof("maps=%+v", maps)
		for _, val := range maps {
			for _, item := range val {
				if item.ID == find.ID {
					continue
				}
				r.Update(entityBasic.BasicAreaEntity{IdLink: item.IdLink, NoLink: item.NoLink}, item.ID)
			}
		}
		maps = nil
	}
	return rt.Ok()
}

// ChildParentIdLink 子集 上级 link更新
//
//	@Description:
//	@receiver c
//	@param id
func (c *BasicAreaService) childParentIdLink(maps map[string][]*entityBasic.BasicAreaEntity, parent *entityBasic.BasicAreaEntity) {
	key := parent.ParentNo
	if strPg.IsBlank(parent.ParentNo) {
		key = constNodePg.ROOT
	}
	entities := maps[key]
	for _, item := range entities {
		item.IdLink = constNodePg.NoLinkAssemble(parent.IdLink, numberPg.Int64ToString(item.ID))
		item.NoLink = constNodePg.NoLinkAssemble(parent.NoLink, item.No)
	}
}

// CacheOverride 缓存重载
//
//	@Description:
//	@receiver c
func (c *BasicAreaService) CacheOverride(ctx *gin.Context) {
	r := c.sv
	infos, b := r.FindAllData()
	if !b {
		return
	}
	maps := slicePg.ToMapArray(infos, func(t *entityBasic.BasicAreaEntity) (string, *entityBasic.BasicAreaEntity) {
		if strPg.IsBlank(t.ParentNo) {
			return constNodePg.ROOT, t
		}
		return t.ParentNo, t
	})
	for _, item := range maps[constNodePg.ROOT] {
		item.IdLink = constNodePg.NoLinkDefault(numberPg.Int64ToString(item.ID))
		item.NoLink = constNodePg.NoLinkDefault(item.No)
		c.childParentIdLink(maps, item)
	}
	c.log.Infof("maps=%+v", maps)
	for _, val := range maps {
		for _, item := range val {
			r.Update(entityBasic.BasicAreaEntity{
				IdLink: item.IdLink,
				NoLink: item.NoLink},
				item.ID)
		}
	}
	maps = nil
}

// Detail 详情
//
//	@Description:
//	@receiver c
//	@param id
func (c *BasicAreaService) Detail(ctx *gin.Context, id int64) (rt rg.Rs[modBasicArea.Vo]) {
	if id < 1 {
		return rt.ErrorMessage("id错误")
	}
	find, b := c.sv.FindById(id)
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	var info modBasicArea.Vo
	copier.Copy(&info, &find)
	return rt.OkData(info)
}

// Enable 启用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *BasicAreaService) Enable(ctx *gin.Context, ct model.BaseIdsCt[string]) (rt rg.Rs[string]) {
	return c.State(ctx, ct.Ids, enumStatePg.ENABLE)
}

// Disable 禁用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *BasicAreaService) Disable(ctx *gin.Context, ct model.BaseIdsCt[string]) (rt rg.Rs[string]) {
	return c.State(ctx, ct.Ids, enumStatePg.GetType(enumStatePg.DISABLE))
}

// State 状态 启用/禁用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *BasicAreaService) State(ctx *gin.Context, ids []string, state enumStatePg.State) (rt rg.Rs[string]) {
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
			r.Update(entityBasic.BasicAreaEntity{State: state.IndexInt8()}, info.ID)
		}
	}
	return rt.Ok()
}

// StateEnableDisable 状态 设置 有效 停用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *BasicAreaService) StateEnableDisable(ctx *gin.Context, ids []string, state enumStatePg.State) (rt rg.Rs[string]) {
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
func (c *BasicAreaService) LogicalDeletion(ctx *gin.Context, ids []string) (rt rg.Rs[string]) {
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
				repository.Update(entityBasic.BasicAreaEntity{State: reverse.IndexInt8()}, info.ID)
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
func (c *BasicAreaService) LogicalRecovery(ctx *gin.Context, ids []string) (rt rg.Rs[string]) {
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
			repository.Update(entityBasic.BasicAreaEntity{State: reverse.IndexInt8()}, info.ID)
		}
	}
	return rt.Ok()
}

// PhysicalDeletion 物理删除
//
//	@Description:
//	@receiver c
//	@param ct
func (c *BasicAreaService) PhysicalDeletion(ctx *gin.Context, ids []string) (rt rg.Rs[string]) {
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
func (c *BasicAreaService) Query(ctx *gin.Context, ct modBasicArea.QueryCt) (rt rg.Rs[pagePg.PaginatorPg[modBasicArea.Vo]]) {
	c.log.Infof("ct=%+v", ct)
	var query entityBasic.BasicAreaEntity
	copier.Copy(&query, &ct)
	slice := make([]modBasicArea.Vo, 0)
	rt.Data.Data = slice
	r := c.sv
	page, err := r.FindAllPageQuery(query, func(p *pagePg.PageCondition[*entityBasic.BasicAreaEntity]) {
		p.PageOption = func(c *pagePg.PaginatorPg[*entityBasic.BasicAreaEntity]) {
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
		pg := pagePg.NewPaginatorPg(func(c *pagePg.PaginatorPg[modBasicArea.Vo]) {
			c.TotalPage = page.TotalPage
			c.Total = page.Total
			c.PageSize = page.PageSize
			c.PageNum = page.PageNum
		})
		//字段赋值
		for _, item := range page.Data {
			var vo modBasicArea.Vo
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
func (c *BasicAreaService) SelectNodePublic(ctx *gin.Context, ct modBasicArea.QueryPublicCt) (rt rg.Rs[[]model.BaseNodeNo]) {
	var query entityBasic.BasicAreaEntity
	copier.Copy(&query, &ct)
	slice := make([]model.BaseNodeNo, 0)
	rt.Data = slice
	infos := c.sv.FindAll(query)
	if len(infos) > 0 {
		for _, item := range infos {
			var vo modBasicArea.Vo
			copier.Copy(&vo, &item)
			code := model.BaseNodeNo{
				Key:      item.No,
				Id:       item.No,
				No:       item.No,
				Label:    item.Name,
				ParentNo: item.ParentNo,
				ParentId: item.ParentNo,
				Extend:   vo,
			}
			//编码
			if !enumParameterPg.NodeQueryByNo.IsEqual(ct.By) {
				code.Key = numberPg.Int64ToString(item.ID)
				code.Id = code.Key
				code.ParentId = item.ParentId
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
func (c *BasicAreaService) SelectNodeAllPublic(ctx *gin.Context, ct modBasicArea.QueryPublicCt) (rt rg.Rs[[]model.BaseNodeNo]) {
	var query entityBasic.BasicAreaEntity
	copier.Copy(&query, &ct)
	slice := make([]model.BaseNodeNo, 0)
	rt.Data = slice
	infos := c.sv.FindAll(query)
	if len(infos) > 0 {
		for _, item := range infos {
			var vo modBasicArea.Vo
			copier.Copy(&vo, &item)
			code := model.BaseNodeNo{
				Key:      item.No,
				Id:       item.No,
				No:       item.No,
				Label:    item.Name,
				ParentNo: item.ParentNo,
				ParentId: item.ParentNo,
				Extend:   vo,
			}
			//编码
			if !enumParameterPg.NodeQueryByNo.IsEqual(ct.By) {
				code.Key = numberPg.Int64ToString(item.ID)
				code.Id = code.Key
				code.ParentId = item.ParentId
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
func (c *BasicAreaService) SelectPublic(ctx *gin.Context, ct modBasicArea.QueryPublicCt) (rt rg.Rs[[]modBasicArea.Vo]) {
	c.log.Infof("ct=%+v", ct)
	var query entityBasic.BasicAreaEntity
	copier.Copy(&query, &ct)
	slice := make([]modBasicArea.Vo, 0)
	rt.Data = slice
	infos := c.sv.FindAll(query)
	if len(infos) > 0 {
		for _, item := range infos {
			var vo modBasicArea.Vo
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
func (c *BasicAreaService) ExportExcel(ctx *gin.Context, ct modBasicArea.QueryCt) {
	c.log.Infof("ct=%+v", ct)
	var query entityBasic.BasicAreaEntity
	copier.Copy(&query, &ct)
	infos := c.sv.FindAll(query)
	if len(infos) > 0 {
		slice := make([]interface{}, 0)
		for _, item := range infos {
			var vo modBasicArea.Vo
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
func (c *BasicAreaService) ExistName(ctx *gin.Context, ct model.BaseExistWdCt[string]) (rt rg.Rs[string]) {
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
func (c *BasicAreaService) ExistCode(ctx *gin.Context, ct model.BaseExistWdCt[string]) (rt rg.Rs[string]) {
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

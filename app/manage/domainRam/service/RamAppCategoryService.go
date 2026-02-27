package service

import (
	"context"
	"reflect"

	"github.com/foxiswho/blog-go/app/manage/domainRam/model/modRamAppCategory"
	"github.com/foxiswho/blog-go/infrastructure/entityRam"
	"github.com/foxiswho/blog-go/infrastructure/repositoryRam"
	"github.com/foxiswho/blog-go/pkg/consts/automatedPg"
	"github.com/foxiswho/blog-go/pkg/consts/constNodePg"
	"github.com/foxiswho/blog-go/pkg/enum/enumCommonPg/typeSysPg"
	"github.com/foxiswho/blog-go/pkg/enum/request/enumParameterPg"
	"github.com/foxiswho/blog-go/pkg/enum/state/enumStatePg"
	"github.com/foxiswho/blog-go/pkg/holderPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/foxiswho/blog-go/pkg/model"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
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
	gs.Provide(new(RamAppCategoryService)).Init(func(s *RamAppCategoryService) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

// RamAppCategoryService 分类
// @Description:
type RamAppCategoryService struct {
	log *log2.Logger                            `autowire:"?"`
	sv  *repositoryRam.RamAppCategoryRepository `autowire:"?"`
}

// Create 新增
//
//	@Description:
//	@receiver c
//	@param ct
//	@return rt
func (c *RamAppCategoryService) Create(ctx *gin.Context, ct modRamAppCategory.CreateCt) (rt rg.Rs[string]) {
	c.log.Infof("ct=%#v", ct)
	var info entityRam.RamAppCategoryEntity
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
	holder := holderPg.GetContextAccount(ctx)
	parent := &entityRam.RamAppCategoryEntity{}
	r := c.sv
	//判断是否是自动,不是自动
	if !automatedPg.IsCreateCode(info.Code) {
		//判断格式是否满足要求
		if !automatedPg.FormatVerify(info.Code) {
			return rt.ErrorMessage("标志格式不能为空")
		}
		//不是自动
		_, result := r.FindByCode(info.Code, repositoryPg.GetOption(ctx))
		if result {
			return rt.ErrorMessage("标志已存在")
		}
	}
	result := false
	if strPg.IsNotBlank(info.ParentNo) {
		parent, result = r.FindByNo(info.ParentNo, repositoryPg.GetOption(ctx))
		if !result {
			return rt.ErrorMessage("上级不存在")
		}
	} else {
		info.ParentNo = ""
	}
	info.No = noPg.No()
	//自动设置编号
	if automatedPg.IsCreateCode(info.Code) {
		info.Code = strPg.GenerateNumberId22()
	}
	info.TenantNo = holder.GetTenantNo()
	c.log.Infof("info=%+v", info)
	err, _ = r.Create(&info)
	if err != nil {
		return rt.ErrorMessage("保存失败 " + err.Error())
	}
	//设置上级 link
	if strPg.IsNotBlank(info.ParentNo) {
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
func (c *RamAppCategoryService) Update(ctx *gin.Context, ct modRamAppCategory.UpdateCt) (rt rg.Rs[string]) {
	c.log.Infof("ct=%#v", ct)
	var info entityRam.RamAppCategoryEntity
	copier.Copy(&info, &ct)
	if ct.ID < 1 {
		return rt.ErrorMessage("id错误")
	}
	if "" == ct.Name {
		return rt.ErrorMessage("名称不能为空")
	}
	r := c.sv
	if strPg.IsBlank(ct.Code) {
		info.Code = ""
	} else {
		_, result := r.FindByCodeAndIdNot(ct.Code, ct.ID.ToString(), repositoryPg.GetOption(ctx))
		if result {
			return rt.ErrorMessage("标志已存在")
		}
	}
	find, b := r.FindById(ct.ID.ToInt64(), repositoryPg.GetOption(ctx))
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	//上级
	parent := &entityRam.RamAppCategoryEntity{}
	var childData []*entityRam.RamAppCategoryEntity
	if strPg.IsNotBlank(ct.ParentNo) {
		result := false
		parent, result = r.FindByNo(ct.ParentNo, repositoryPg.GetOption(ctx))
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

	info.TypeSys = typeSysPg.General.Index()
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
		maps := slicePg.ToMapArray(childData, func(t *entityRam.RamAppCategoryEntity) (string, *entityRam.RamAppCategoryEntity) {
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
				r.Update(entityRam.RamAppCategoryEntity{IdLink: item.IdLink,
					NoLink: item.NoLink},
					item.ID)
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
func (c *RamAppCategoryService) childParentIdLink(maps map[string][]*entityRam.RamAppCategoryEntity, parent *entityRam.RamAppCategoryEntity) {
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
func (c *RamAppCategoryService) CacheOverride(ctx *gin.Context) {
	r := c.sv
	infos, b := r.FindAllData(repositoryPg.GetOption(ctx))
	if !b {
		return
	}
	maps := slicePg.ToMapArray(infos, func(t *entityRam.RamAppCategoryEntity) (string, *entityRam.RamAppCategoryEntity) {
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
			r.Update(entityRam.RamAppCategoryEntity{
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
func (c *RamAppCategoryService) Detail(ctx *gin.Context, id int64) (rt rg.Rs[modRamAppCategory.Vo]) {
	if id < 1 {
		return rt.ErrorMessage("id错误")
	}
	find, b := c.sv.FindById(id, repositoryPg.GetOption(ctx))
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	var info modRamAppCategory.Vo
	copier.Copy(&info, &find)
	return rt.OkData(info)
}

// Enable 启用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamAppCategoryService) Enable(ctx *gin.Context, ct model.BaseIdsCt[string]) (rt rg.Rs[string]) {
	return c.State(ctx, ct.Ids, enumStatePg.ENABLE)
}

// Disable 禁用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamAppCategoryService) Disable(ctx *gin.Context, ct model.BaseIdsCt[string]) (rt rg.Rs[string]) {
	return c.State(ctx, ct.Ids, enumStatePg.GetType(enumStatePg.DISABLE))
}

// State 状态 启用/禁用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamAppCategoryService) State(ctx *gin.Context, ids []string, state enumStatePg.State) (rt rg.Rs[string]) {
	if len(ids) < 1 {
		return rt.ErrorMessage("id错误")
	}
	r := c.sv
	finds, b := r.FindAllByIdStringIn(ids, repositoryPg.GetOption(ctx))
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	for _, info := range finds {
		if info.State != state.IndexInt8() {
			r.Update(entityRam.RamAppCategoryEntity{State: state.IndexInt8()}, info.ID)
		}
	}
	return rt.Ok()
}

// StateEnableDisable 状态 设置 有效 停用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamAppCategoryService) StateEnableDisable(ctx *gin.Context, ids []string, state enumStatePg.State) (rt rg.Rs[string]) {
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
func (c *RamAppCategoryService) LogicalDeletion(ctx *gin.Context, ids []string) (rt rg.Rs[string]) {
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
			c.log.Infof("id=%v,TenantId=%v", info.ID, 0)
		}
		repository.DeleteByIdsString(ids, repositoryPg.GetOption(ctx))
	} else {
		for _, info := range finds {
			enum := enumStatePg.State(info.State)
			// 有效 停用，反转 为对应的 取消 弃置
			if ok, reverse := enum.ReverseEnableDisable(); ok {
				repository.Update(entityRam.RamAppCategoryEntity{State: reverse.IndexInt8()}, info.ID)
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
func (c *RamAppCategoryService) LogicalRecovery(ctx *gin.Context, ids []string) (rt rg.Rs[string]) {
	c.log.Infof("ct=%+v", ids)
	if len(ids) < 1 {
		return rt.ErrorMessage("id错误")
	}
	repository := c.sv
	finds, b := repository.FindAllByIdStringIn(ids, repositoryPg.GetOption(ctx))
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	for _, info := range finds {
		enum := enumStatePg.State(info.State)
		//  取消 弃置 批量删除，反转 为对应的 有效 停用 停用
		if ok, reverse := enum.ReverseCancelLayAside(); ok {
			repository.Update(entityRam.RamAppCategoryEntity{State: reverse.IndexInt8()}, info.ID)
		}
	}
	return rt.Ok()
}

// PhysicalDeletion 物理删除
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamAppCategoryService) PhysicalDeletion(ctx *gin.Context, ids []string) (rt rg.Rs[string]) {
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
		c.log.Infof("id=%v,TenantId=%v", info.ID, 0)
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
func (c *RamAppCategoryService) Query(ctx *gin.Context, ct modRamAppCategory.QueryCt) (rt rg.Rs[pagePg.PaginatorPg[modRamAppCategory.Vo]]) {
	c.log.Infof("ct=%+v", ct)
	var query entityRam.RamAppCategoryEntity
	copier.Copy(&query, &ct)
	slice := make([]modRamAppCategory.Vo, 0)
	rt.Data.Data = slice
	r := c.sv
	page, err := r.FindAllPageQuery(query, func(p *pagePg.PageCondition[*entityRam.RamAppCategoryEntity]) {
		p.PageOption = func(c *pagePg.PaginatorPg[*entityRam.RamAppCategoryEntity]) {
			c.PageNum = ct.PageNum
			c.PageSize = ct.PageSize
		}
		//自定义查询
		p.Condition = r.DbModel().Order("create_at asc")
		if "" != ct.Wd {
			p.Condition.Where("name like ?", "%"+ct.Wd+"%")
		}
	}, repositoryPg.GetOption(ctx))
	if nil != err {
		return rt.Ok()
	}

	if page.Total > 0 && page.Data != nil && len(page.Data) > 0 {
		pg := pagePg.NewPaginatorPg(func(c *pagePg.PaginatorPg[modRamAppCategory.Vo]) {
			c.TotalPage = page.TotalPage
			c.Total = page.Total
			c.PageSize = page.PageSize
			c.PageNum = page.PageNum
		})
		//字段赋值
		for _, item := range page.Data {
			var vo modRamAppCategory.Vo
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

// QueryPublic 查询
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamAppCategoryService) QueryPublic(ctx *gin.Context, ct modRamAppCategory.QueryCt) (rt rg.Rs[pagePg.PaginatorPg[modRamAppCategory.Vo]]) {
	var query entityRam.RamAppCategoryEntity
	copier.Copy(&query, &ct)
	slice := make([]modRamAppCategory.Vo, 0)
	rt.Data.Data = slice
	r := c.sv
	page, err := r.FindAllPageQuery(query, func(p *pagePg.PageCondition[*entityRam.RamAppCategoryEntity]) {
		p.PageOption = func(c *pagePg.PaginatorPg[*entityRam.RamAppCategoryEntity]) {
			c.PageNum = ct.PageNum
			c.PageSize = ct.PageSize
		}
		//自定义查询
		p.Condition = r.DbModel().Order("create_at asc")
		if "" != ct.Wd {
			p.Condition.Where("name like ?", "%"+ct.Wd+"%")
		}
	}, repositoryPg.GetOption(ctx))
	if nil != err {
		return rt.Ok()
	}

	if page.Total > 0 && page.Data != nil && len(page.Data) > 0 {

		pg := pagePg.NewPaginatorPg(func(c *pagePg.PaginatorPg[modRamAppCategory.Vo]) {
			c.TotalPage = page.TotalPage
			c.Total = page.Total
			c.PageSize = page.PageSize
			c.PageNum = page.PageNum
		})
		//字段赋值
		for _, item := range page.Data {
			var vo modRamAppCategory.Vo
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
func (c *RamAppCategoryService) SelectNodePublic(ctx *gin.Context, ct modRamAppCategory.QueryPublicCt) (rt rg.Rs[[]model.BaseNodeNo]) {
	var query entityRam.RamAppCategoryEntity
	copier.Copy(&query, &ct)
	slice := make([]model.BaseNodeNo, 0)
	rt.Data = slice
	infos := c.sv.FindAll(query, repositoryPg.GetOption(ctx))
	if len(infos) > 0 {
		for _, item := range infos {
			var vo modRamAppCategory.Vo
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
func (c *RamAppCategoryService) SelectNodeAllPublic(ctx *gin.Context, ct modRamAppCategory.QueryPublicCt) (rt rg.Rs[[]model.BaseNodeNo]) {
	var query entityRam.RamAppCategoryEntity
	copier.Copy(&query, &ct)
	slice := make([]model.BaseNodeNo, 0)
	rt.Data = slice
	infos := c.sv.FindAll(query, repositoryPg.GetOption(ctx))
	if len(infos) > 0 {
		for _, item := range infos {
			var vo modRamAppCategory.Vo
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
func (c *RamAppCategoryService) SelectPublic(ctx *gin.Context, ct modRamAppCategory.QueryPublicCt) (rt rg.Rs[[]modRamAppCategory.Vo]) {
	var query entityRam.RamAppCategoryEntity
	copier.Copy(&query, &ct)
	slice := make([]modRamAppCategory.Vo, 0)
	rt.Data = slice
	infos := c.sv.FindAll(query, repositoryPg.GetOption(ctx))
	if len(infos) > 0 {
		for _, item := range infos {
			var vo modRamAppCategory.Vo
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
func (c *RamAppCategoryService) ExistName(ctx *gin.Context, ct model.BaseExistWdCt[string]) (rt rg.Rs[string]) {
	if "" == ct.Wd {
		return rt.ErrorMessage("关键词不能为空")
	}
	id := "0"
	if strPg.IsNotBlank(ct.Id) {
		id = ct.Id
	}
	_, result := c.sv.FindByNameAndIdNot(ct.Wd, id, repositoryPg.GetOption(ctx))
	if result {
		return rt.ErrorMessage("重复，已存在")
	}
	return rt.OkMessage("可以使用")
}

// ExistNo 查重
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamAppCategoryService) ExistNo(ctx *gin.Context, ct model.BaseExistWdCt[string]) (rt rg.Rs[string]) {
	if "" == ct.Wd {
		return rt.ErrorMessage("关键词不能为空")
	}
	id := "0"
	if strPg.IsNotBlank(ct.Id) {
		id = ct.Id
	}
	_, result := c.sv.FindByCodeAndIdNot(ct.Wd, id, repositoryPg.GetOption(ctx))
	if result {
		return rt.ErrorMessage("重复，已存在")
	}
	return rt.OkMessage("可以使用")
}

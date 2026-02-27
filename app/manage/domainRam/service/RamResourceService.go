package service

import (
	"context"

	"github.com/foxiswho/blog-go/app/manage/domainRam/model/modRamResource"
	"github.com/foxiswho/blog-go/infrastructure/entityRam"
	"github.com/foxiswho/blog-go/infrastructure/repositoryRam"
	"github.com/foxiswho/blog-go/pkg/consts/constNodePg"
	"github.com/foxiswho/blog-go/pkg/consts/constsRam/typeAttrPg"
	"github.com/foxiswho/blog-go/pkg/enum/request/enumParameterPg"
	"github.com/foxiswho/blog-go/pkg/enum/state/enumStatePg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/foxiswho/blog-go/pkg/model"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/foxiswho/blog-go/pkg/tools/noPg"
	"github.com/gin-gonic/gin"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"

	"reflect"

	"github.com/jinzhu/copier"
	"github.com/pangu-2/go-tools/tools/dbPg/pagePg"
	"github.com/pangu-2/go-tools/tools/numberPg"
	"github.com/pangu-2/go-tools/tools/slicePg"
	"github.com/pangu-2/go-tools/tools/strPg"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
)

func init() {
	gs.Provide(new(RamResourceService)).Init(func(s *RamResourceService) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

// RamResourceService 资源
// @Description:
type RamResourceService struct {
	sv  *repositoryRam.RamResourceRepository `autowire:"?"`
	log *log2.Logger                         `autowire:"?"`
}

// Create 新增
//
//	@Description:
//	@receiver c
//	@param ct
//	@return rt
func (c *RamResourceService) Create(ctx *gin.Context, ct modRamResource.CreateCt) (rt rg.Rs[string]) {
	c.log.Infof("ct=%+v", ct)
	var info entityRam.RamResourceEntity
	err2 := copier.Copy(&info, &ct)
	if err2 != nil {
		c.log.Error("copier.Copy=%+v", err2)
		return rt.ErrorMessage(err2.Error())
	}
	if "" == ct.Name {
		return rt.ErrorMessage("名称不能为空")
	}
	//holder := holderPg.GetContextAccount(ctx)
	parent := &entityRam.RamResourceEntity{}
	r := c.sv
	if strPg.IsBlank(ct.TypeAttr) {
		return rt.ErrorMessage("属性不能为空")
	}
	if strPg.IsNotBlank(ct.ParentNo) {
		result := false
		parent, result = r.FindByNo(ct.ParentNo, repositoryPg.GetOption(ctx))
		if !result {
			return rt.ErrorMessage("上级不存在")
		}
		if strPg.IsNotBlank(parent.ParentNo) {
			return rt.ErrorMessage("只允许2级存在")
		}
		if !typeAttrPg.CategoryLast.IsEqual(parent.TypeAttr) {
			return rt.ErrorMessage("上级属性为分类属性时,才能有下级")
		}
		// 当前属性
		if !typeAttrPg.Resource.IsEqual(ct.TypeAttr) {
			return rt.ErrorMessage("属性必须为资源属性")
		}
		info.ParentId = numberPg.Int64ToString(parent.ID)
		info.ParentNo = parent.No
	} else {
		if !typeAttrPg.CategoryLast.IsEqual(ct.TypeAttr) {
			return rt.ErrorMessage("顶级属性为分类属性")
		}
	}
	info.No = noPg.No()
	//
	c.log.Infof("info=%+v", info)
	err, _ := r.Create(&info)
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
		info.ParentNo = ""
		info.ParentId = ""
	}
	err = r.Update(info, info.ID)
	if err != nil {
		return rt.ErrorMessage(err.Error())
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
func (c *RamResourceService) Update(ctx *gin.Context, ct modRamResource.UpdateCt) (rt rg.Rs[string]) {
	c.log.Infof("ct=%+v", ct)
	var info entityRam.RamResourceEntity
	err := copier.Copy(&info, &ct)
	if err != nil {
		return rt.ErrorMessage(err.Error())
	}
	if ct.ID < 1 {
		return rt.ErrorMessage("id错误")
	}
	if "" == ct.Name {
		return rt.ErrorMessage("名称不能为空")
	}
	if strPg.IsBlank(ct.TypeAttr) {
		return rt.ErrorMessage("属性不能为空")
	}
	r := c.sv
	find, b := r.FindById(ct.ID.ToInt64(), repositoryPg.GetOption(ctx))
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	//上级
	parent := &entityRam.RamResourceEntity{}
	var childData []*entityRam.RamResourceEntity
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
					if item.ID == parent.ID {
						return rt.ErrorMessage("无法保存，不能设置为自己的子集")
					}
				}
			}
		}
	}
	// 如果是分类，上级应为空
	if typeAttrPg.CategoryLast.IsEqual(ct.TypeAttr) {
		if strPg.IsNotBlank(info.ParentNo) {
			return rt.ErrorMessage("上级应为空")
		}
	} else if typeAttrPg.Resource.IsEqual(ct.TypeAttr) {
		//如果是资源,上级不能为空
		if strPg.IsBlank(info.ParentNo) {
			return rt.ErrorMessage("上级不能为空")
		} else {
			//且上级必须是分类
			if !typeAttrPg.CategoryLast.IsEqual(parent.TypeAttr) {
				return rt.ErrorMessage("上级属性为分类属性")
			}
		}
	}
	//
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
	err = r.Update(info, info.ID)
	if err != nil {
		c.log.Errorf("update error=%+v", err)
		return rt.ErrorMessage(err.Error())
	}
	c.log.Infof("save.info=%+v", info)
	//更改上级后，相关子集修改
	if strPg.IsNotBlank(ct.ParentNo) && nil != childData {
		maps := slicePg.ToMapArray(childData, func(t *entityRam.RamResourceEntity) (string, *entityRam.RamResourceEntity) {
			if len(t.ParentId) == 0 {
				return constNodePg.ROOT, t
			}
			return t.ParentId, t
		})
		for _, item := range maps[numberPg.Int64ToString(find.ID)] {
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
				err = r.Update(entityRam.RamResourceEntity{IdLink: item.IdLink, NoLink: item.NoLink}, item.ID)
				if err != nil {
					return rt.ErrorMessage(err.Error())
				}
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
func (c *RamResourceService) childParentIdLink(maps map[string][]*entityRam.RamResourceEntity, parent *entityRam.RamResourceEntity) {
	entities := maps[numberPg.Int64ToString(parent.ID)]
	for _, item := range entities {
		item.IdLink = constNodePg.NoLinkAssemble(parent.IdLink, numberPg.Int64ToString(item.ID))
		item.NoLink = constNodePg.NoLinkAssemble(parent.NoLink, item.No)
	}
}

// CacheOverride 缓存重载
//
//	@Description:
//	@receiver c
func (c *RamResourceService) CacheOverride(ctx *gin.Context) {
	r := c.sv
	infos, b := r.FindAllData(repositoryPg.GetOption(ctx))
	if !b {
		return
	}
	maps := slicePg.ToMapArray(infos, func(t *entityRam.RamResourceEntity) (string, *entityRam.RamResourceEntity) {
		if len(t.ParentId) == 0 {
			return constNodePg.ROOT, t
		}
		return t.ParentId, t
	})
	for _, item := range maps[constNodePg.ROOT] {
		item.IdLink = constNodePg.NoLinkDefault(numberPg.Int64ToString(item.ID))
		item.NoLink = constNodePg.NoLinkDefault(item.No)
		c.childParentIdLink(maps, item)
	}
	c.log.Infof("maps=%+v", maps)
	for _, val := range maps {
		for _, item := range val {
			r.Update(entityRam.RamResourceEntity{IdLink: item.IdLink, NoLink: item.NoLink}, item.ID)
		}
	}
	maps = nil
}

// Detail 详情
//
//	@Description:
//	@receiver c
//	@param id
func (c *RamResourceService) Detail(ctx *gin.Context, id int64) (rt rg.Rs[modRamResource.Vo]) {
	if id < 1 {
		return rt.ErrorMessage("id错误")
	}
	find, b := c.sv.FindById(id, repositoryPg.GetOption(ctx))
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	var info modRamResource.Vo
	copier.Copy(&info, &find)
	return rt.OkData(info)
}

// Enable 启用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamResourceService) Enable(ctx *gin.Context, ct model.BaseIdsCt[string]) (rt rg.Rs[string]) {
	return c.State(ctx, ct.Ids, enumStatePg.ENABLE)
}

// Disable 禁用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamResourceService) Disable(ctx *gin.Context, ct model.BaseIdsCt[string]) (rt rg.Rs[string]) {
	return c.State(ctx, ct.Ids, enumStatePg.GetType(enumStatePg.DISABLE))
}

// State 状态 启用/禁用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamResourceService) State(ctx *gin.Context, ids []string, state enumStatePg.State) (rt rg.Rs[string]) {
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
			r.Update(entityRam.RamResourceEntity{State: state.IndexInt8()}, info.ID)
		}
	}
	return rt.Ok()
}

// StateEnableDisable 状态 设置 有效 停用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamResourceService) StateEnableDisable(ctx *gin.Context, ids []string, state enumStatePg.State) (rt rg.Rs[string]) {
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
func (c *RamResourceService) LogicalDeletion(ctx *gin.Context, ids []string) (rt rg.Rs[string]) {
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
			c.log.Infof("id=%v,TenantId=%v", info.ID, " info.TenantNo")
		}
		repository.DeleteByIdsString(ids, repositoryPg.GetOption(ctx))
	} else {
		for _, info := range finds {
			enum := enumStatePg.State(info.State)
			// 有效 停用，反转 为对应的 取消 弃置
			if ok, reverse := enum.ReverseEnableDisable(); ok {
				repository.Update(entityRam.RamResourceEntity{State: reverse.IndexInt8()}, info.ID)
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
func (c *RamResourceService) LogicalRecovery(ctx *gin.Context, ids []string) (rt rg.Rs[string]) {
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
			repository.Update(entityRam.RamResourceEntity{State: reverse.IndexInt8()}, info.ID)
		}
	}
	return rt.Ok()
}

// PhysicalDeletion 物理删除
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamResourceService) PhysicalDeletion(ctx *gin.Context, ids []string) (rt rg.Rs[string]) {
	c.log.Infof("ct=%+v", ids)
	if len(ids) < 1 {
		return rt.ErrorMessage("id错误")
	}
	r := c.sv
	finds, b := r.FindAllByIdStringIn(ids, repositoryPg.GetOption(ctx))
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	idsNew := make([]int64, 0)
	for _, info := range finds {
		c.log.Infof("id=%v,TenantId=%v", info.ID, 0)
		//判断 是否是 分类
		if typeAttrPg.CategoryLast.IsEqual(info.TypeAttr) || typeAttrPg.Category.IsEqual(info.TypeAttr) {
			idString, result := r.CountByParentIdString(numberPg.Int64ToString(info.ID))
			if result && idString > 0 {
				return rt.ErrorMessage("该分类下存在子数据，请先删除子数据")
			}
		}
		idsNew = append(idsNew, info.ID)
	}
	if len(idsNew) > 0 {
		r.DeleteByIds(idsNew)
	}
	return rt.Ok()
}

// Query 查询
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamResourceService) Query(ctx *gin.Context, ct modRamResource.QueryCt) (rt rg.Rs[pagePg.PaginatorPg[modRamResource.Vo]]) {
	c.log.Infof("ct=%+v", ct)
	var query entityRam.RamResourceEntity
	copier.Copy(&query, &ct)
	slice := make([]modRamResource.Vo, 0)
	rt.Data.Data = slice
	r := c.sv
	page, err := r.FindAllPageQuery(query, func(p *pagePg.PageCondition[*entityRam.RamResourceEntity]) {
		p.PageOption = func(c *pagePg.PaginatorPg[*entityRam.RamResourceEntity]) {
			c.PageNum = ct.PageNum
			c.PageSize = ct.PageSize
			if c.PageSize < 1 {
				c.PageSize = 20
			}
		}
		p.Condition = r.DbModel().Order("name,create_at desc")
		//自定义查询
		if "" != ct.Wd {
			p.Condition.Where("name like ?", "%"+ct.Wd+"%")
		}
		//是否读取全部
		if !ct.ALL {
			p.Condition.Where("parent_no !='' ")
		}
	})
	if nil != err {
		return rt.Ok()
	}

	if page.Total > 0 && page.Data != nil && len(page.Data) > 0 {

		pg := pagePg.NewPaginatorPg(func(c *pagePg.PaginatorPg[modRamResource.Vo]) {
			c.TotalPage = page.TotalPage
			c.Total = page.Total
			c.PageSize = page.PageSize
			c.PageNum = page.PageNum
		})
		//字段赋值
		for _, item := range page.Data {
			var vo modRamResource.Vo
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
func (c *RamResourceService) QueryPublic(ctx *gin.Context, ct modRamResource.QueryCt) (rt rg.Rs[pagePg.PaginatorPg[modRamResource.Vo]]) {
	c.log.Infof("ct=%+v", ct)
	var query entityRam.RamResourceEntity
	copier.Copy(&query, &ct)
	slice := make([]modRamResource.Vo, 0)
	rt.Data.Data = slice
	r := c.sv
	page, err := r.FindAllPageQuery(query, func(p *pagePg.PageCondition[*entityRam.RamResourceEntity]) {
		p.PageOption = func(c *pagePg.PaginatorPg[*entityRam.RamResourceEntity]) {
			c.PageNum = ct.PageNum
			c.PageSize = ct.PageSize
			if c.PageSize < 1 {
				c.PageSize = 20
			}
		}
		p.Condition = r.DbModel().Order("name,create_at desc")
		//自定义查询
		if "" != ct.Wd {
			p.Condition.Where("name like ?", "%"+ct.Wd+"%")
		}
		//是否读取全部
		if !ct.ALL {
			p.Condition.Where("parent_no !='' ")
		}
	})
	if nil != err {
		return rt.Ok()
	}

	if page.Total > 0 && page.Data != nil && len(page.Data) > 0 {

		pg := pagePg.NewPaginatorPg(func(c *pagePg.PaginatorPg[modRamResource.Vo]) {
			c.TotalPage = page.TotalPage
			c.Total = page.Total
			c.PageSize = page.PageSize
			c.PageNum = page.PageNum
		})
		//字段赋值
		for _, item := range page.Data {
			var vo modRamResource.Vo
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
func (c *RamResourceService) SelectNodePublic(ctx *gin.Context, ct modRamResource.QueryPublicCt) (rt rg.Rs[[]model.BaseNodeNo]) {
	c.log.Infof("ct=%+v", ct)
	var query entityRam.RamResourceEntity
	copier.Copy(&query, &ct)
	slice := make([]model.BaseNodeNo, 0)
	rt.Data = slice
	infos := c.sv.FindAll(query, repositoryPg.GetOption(ctx))
	if len(infos) > 0 {
		for _, item := range infos {
			var vo modRamResource.Vo
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
func (c *RamResourceService) SelectNodeAllPublic(ctx *gin.Context, ct modRamResource.QueryPublicCt) (rt rg.Rs[[]model.BaseNodeNo]) {
	c.log.Infof("ct=%+v", ct)
	var query entityRam.RamResourceEntity
	copier.Copy(&query, &ct)
	slice := make([]model.BaseNodeNo, 0)
	rt.Data = slice
	infos := c.sv.FindAll(query, repositoryPg.GetOption(ctx))
	if len(infos) > 0 {
		for _, item := range infos {
			var vo modRamResource.Vo
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
func (c *RamResourceService) SelectPublic(ctx *gin.Context, ct modRamResource.QueryCt) (rt rg.Rs[[]modRamResource.Vo]) {
	c.log.Infof("ct=%+v", ct)
	var query entityRam.RamResourceEntity
	copier.Copy(&query, &ct)
	slice := make([]modRamResource.Vo, 0)
	rt.Data = slice
	infos := c.sv.FindAll(query, repositoryPg.GetOption(ctx))
	if len(infos) > 0 {
		for _, item := range infos {
			var vo modRamResource.Vo
			copier.Copy(&vo, &item)
			slice = append(slice, vo)
		}
		rt.Data = slice
	}
	return rt.Ok()
}

// SelectCategoryPublic 分类
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamResourceService) SelectCategoryPublic(ctx *gin.Context, ct modRamResource.QueryCt) (rt rg.Rs[[]modRamResource.Vo]) {
	c.log.Infof("ct=%+v", ct)
	var query entityRam.RamResourceEntity
	copier.Copy(&query, &ct)
	slice := make([]modRamResource.Vo, 0)
	rt.Data = slice
	infos, result := c.sv.FindByParentIdRoot()
	if result {
		for _, item := range infos {
			var vo modRamResource.Vo
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
func (c *RamResourceService) ExistName(ctx *gin.Context, ct model.BaseExistWdCt[string]) (rt rg.Rs[string]) {
	c.log.Infof("ct=%+v", ct)
	if "" == ct.Wd {
		return rt.ErrorMessage("查询内容不能为空")
	}
	_, result := c.sv.FindByNameAndIdNot(ct.Wd, numberPg.StrToInt64(ct.Id))
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
func (c *RamResourceService) ExistCode(ctx *gin.Context, ct model.BaseExistWdCt[string]) (rt rg.Rs[string]) {
	c.log.Infof("ct=%+v", ct)
	if "" == ct.Wd {
		return rt.ErrorMessage("查询内容不能为空")
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

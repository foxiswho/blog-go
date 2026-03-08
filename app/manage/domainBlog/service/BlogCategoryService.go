package service

import (
	"context"
	"reflect"

	"github.com/foxiswho/blog-go/app/manage/domainBlog/model/modBlogCategory"
	"github.com/foxiswho/blog-go/infrastructure/entityBlog"
	"github.com/foxiswho/blog-go/infrastructure/repositoryBlog"
	"github.com/foxiswho/blog-go/pkg/consts/automatedPg"
	"github.com/foxiswho/blog-go/pkg/consts/constNodePg"
	"github.com/foxiswho/blog-go/pkg/enum/enumCommonPg/typeSysPg"
	"github.com/foxiswho/blog-go/pkg/enum/request/enumParameterPg"
	"github.com/foxiswho/blog-go/pkg/enum/state/enumStatePg"
	"github.com/foxiswho/blog-go/pkg/holderPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/foxiswho/blog-go/pkg/model"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
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
	gs.Provide(new(BlogCategoryService)).Init(func(s *BlogCategoryService) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

// BlogCategoryService 分类
// @Description:
type BlogCategoryService struct {
	log *log2.Logger                           `autowire:"?"`
	rep *repositoryBlog.BlogCategoryRepository `autowire:"?"`
}

// Create 新增
//
//	@Description:
//	@receiver c
//	@param ct
//	@return rt
func (c *BlogCategoryService) Create(ctx *gin.Context, ct modBlogCategory.CreateCt) (rt rg.Rs[string]) {
	if "" == ct.Name {
		return rt.ErrorMessage("名称不能为空")
	}
	if "" == ct.No {
		return rt.ErrorMessage("编号不能为空")
	}
	r := c.rep
	//判断是否是自动,不是自动
	if !automatedPg.IsCreateCode(ct.No) {
		//判断格式是否满足要求
		if !automatedPg.FormatVerify(ct.No) {
			return rt.ErrorMessage("编号格式不能为空")
		}
		//不是自动
		_, result := r.FindByNo(ct.No)
		if result {
			return rt.ErrorMessage("编号已存在")
		}
	}
	holder := holderPg.GetContextAccount(ctx)
	parent := &entityBlog.BlogCategoryEntity{}
	var info entityBlog.BlogCategoryEntity
	err := copier.Copy(&info, &ct)
	if err != nil {
		c.log.Infof("copier.Copy error: %+v", err)
	}
	info.TypeSys = typeSysPg.General.Index()
	//
	if len(ct.ParentId) > 0 {
		result := false
		parent, result = r.FindByIdString(ct.ParentId)
		if !result {
			return rt.ErrorMessage("上级不存在")
		}
	}
	_, result := r.FindByNoAndIdNot(ct.No, "0", repositoryPg.GetOption(ctx))
	if result {
		return rt.ErrorMessage("字段已存在")
	}
	c.log.Infof("info%+v", info)
	info.TenantNo = holder.GetTenantNo()
	//自动设置编号
	if automatedPg.IsCreateCode(ct.No) {
		info.No = strPg.GenerateNumberId22()
	}
	err, _ = r.Create(&info)
	if err != nil {
		return rt.ErrorMessage("保存失败 " + err.Error())
	}
	//设置上级 link
	if len(ct.ParentId) > 0 {
		info.IdLink = parent.IdLink + numberPg.Int64ToString(info.ID) + "|"
		info.NoLink = parent.NoLink + info.No + "|"
		info.ParentNo = parent.No
	} else {
		info.IdLink = numberPg.Int64ToString(info.ID) + "|"
		info.NoLink = info.No + "|"
	}
	r.Update(info, info.ID)

	c.log.Infof("save=%+v", info)
	return rg.OkData(numberPg.Int64ToString(info.ID))
}

// Update 更新
//
//	@Description:
//	@receiver c
//	@param ct
//	@return rt
func (c *BlogCategoryService) Update(ctx *gin.Context, ct modBlogCategory.UpdateCt) (rt rg.Rs[string]) {
	if ct.ID < 1 {
		return rt.ErrorMessage("id错误")
	}
	if "" == ct.Name {
		return rt.ErrorMessage("名称不能为空")
	}
	if "" == ct.No {
		return rt.ErrorMessage("编号不能为空")
	}
	r := c.rep
	_, result := r.FindByNoAndIdNot(ct.No, ct.ID.ToString(), repositoryPg.GetOption(ctx))
	if result {
		return rt.ErrorMessage("编号已存在")
	}
	find, b := r.FindById(ct.ID.ToInt64(), repositoryPg.GetOption(ctx))
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	//上级
	parent := &entityBlog.BlogCategoryEntity{}
	var childData []*entityBlog.BlogCategoryEntity
	if len(ct.ParentId) > 0 {
		result := false
		parent, result = r.FindByIdString(ct.ParentId)
		if !result {
			return rt.ErrorMessage("上级不存在")
		}
		if parent.ID == ct.ID.ToInt64() {
			return rt.ErrorMessage("上级不能等于自己")
		}
		//新的ID 不等于 旧的上级时,检测是否已经 在新的子集已存在
		if numberPg.Int64ToString(parent.ID) != find.ParentId {
			result2 := false
			childData, result2 = r.FindAllByParentIdLink(numberPg.Int64ToString(find.ID))
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

	var info entityBlog.BlogCategoryEntity
	copier.Copy(&info, &ct)
	info.TypeSys = typeSysPg.General.Index()
	//设置上级 link
	if len(ct.ParentId) > 0 {
		info.IdLink = parent.IdLink + numberPg.Int64ToString(info.ID) + "|"
		info.NoLink = parent.NoLink + info.No + "|"
		info.ParentNo = parent.No
	} else {
		info.IdLink = numberPg.Int64ToString(info.ID)
		info.NoLink = info.No + "|"
		info.ParentNo = ""
		info.ParentId = ""
	}

	err := r.Update(info, info.ID)
	if err != nil {
		c.log.Errorf("update error=%+v", err)
		return rt.ErrorMessage(err.Error())
	}
	c.log.Infof("save.info=%+v", info)
	r.UpdateMap(map[string]interface{}{
		"parent_id": info.ParentId,
		"parent_no": info.ParentNo,
	}, info.ID)
	//更改上级后，相关子集修改
	if len(ct.ParentId) > 0 && nil != childData {
		maps := slicePg.ToMapArray(childData, func(t *entityBlog.BlogCategoryEntity) (string, *entityBlog.BlogCategoryEntity) {
			if len(t.ParentId) == 0 {
				return constNodePg.ROOT, t
			}
			return t.ParentId, t
		})
		for _, item := range maps[numberPg.Int64ToString(info.ID)] {
			item.IdLink = info.IdLink + numberPg.Int64ToString(item.ID) + "|"
			item.NoLink = info.NoLink + item.No + "|"
			c.childParentIdLink(maps, item)
		}
		c.log.Infof("maps=%+v", maps)
		for _, val := range maps {
			for _, item := range val {
				if item.ID == info.ID {
					continue
				}
				r.Update(entityBlog.BlogCategoryEntity{IdLink: item.IdLink, NoLink: item.NoLink}, item.ID)
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
func (c *BlogCategoryService) childParentIdLink(maps map[string][]*entityBlog.BlogCategoryEntity, parent *entityBlog.BlogCategoryEntity) {
	entities := maps[numberPg.Int64ToString(parent.ID)]
	for _, item := range entities {
		item.IdLink = parent.IdLink + numberPg.Int64ToString(item.ID) + "|"
	}
}

// CacheOverride 缓存重载
//
//	@Description:
//	@receiver c
func (c *BlogCategoryService) CacheOverride(ctx *gin.Context) {
	r := c.rep
	infos, b := r.FindAllData(repositoryPg.GetOption(ctx))
	if !b {
		return
	}
	maps := slicePg.ToMapArray(infos, func(t *entityBlog.BlogCategoryEntity) (string, *entityBlog.BlogCategoryEntity) {
		if len(t.ParentId) == 0 {
			return constNodePg.ROOT, t
		}
		return t.ParentId, t
	})
	for _, item := range maps[constNodePg.ROOT] {
		item.IdLink = numberPg.Int64ToString(item.ID)
		c.childParentIdLink(maps, item)
	}
	c.log.Infof("maps=%+v", maps)
	for _, val := range maps {
		for _, item := range val {
			r.Update(entityBlog.BlogCategoryEntity{IdLink: item.IdLink}, item.ID)
		}
	}
	maps = nil
}

// Detail 详情
//
//	@Description:
//	@receiver c
//	@param id
func (c *BlogCategoryService) Detail(ctx *gin.Context, id int64) (rt rg.Rs[modBlogCategory.Vo]) {
	if id < 1 {
		return rt.ErrorMessage("id错误")
	}
	find, b := c.rep.FindById(id, repositoryPg.GetOption(ctx))
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	var info modBlogCategory.Vo
	copier.Copy(&info, &find)
	return rt.OkData(info)
}

// Enable 启用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *BlogCategoryService) Enable(ctx *gin.Context, ct model.BaseIdsCt[string]) (rt rg.Rs[string]) {
	return c.State(ctx, ct.Ids, enumStatePg.ENABLE)
}

// Disable 禁用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *BlogCategoryService) Disable(ctx *gin.Context, ct model.BaseIdsCt[string]) (rt rg.Rs[string]) {
	return c.State(ctx, ct.Ids, enumStatePg.GetType(enumStatePg.DISABLE))
}

// State 状态 启用/禁用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *BlogCategoryService) State(ctx *gin.Context, ids []string, state enumStatePg.State) (rt rg.Rs[string]) {
	if len(ids) < 1 {
		return rt.ErrorMessage("id错误")
	}
	r := c.rep
	finds, b := r.FindAllByIdStringIn(ids, repositoryPg.GetOption(ctx))
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	for _, info := range finds {
		if info.State != state.IndexInt8() {
			r.Update(entityBlog.BlogCategoryEntity{State: state.IndexInt8()}, info.ID)
		}
	}
	return rt.Ok()
}

// StateEnableDisable 状态 设置 有效 停用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *BlogCategoryService) StateEnableDisable(ctx *gin.Context, ids []string, state enumStatePg.State) (rt rg.Rs[string]) {
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
func (c *BlogCategoryService) LogicalDeletion(ctx *gin.Context, ids []string) (rt rg.Rs[string]) {
	if len(ids) < 1 {
		return rt.ErrorMessage("id错误")
	}
	repository := c.rep
	finds, b := repository.FindAllByIdStringIn(ids, repositoryPg.GetOption(ctx))
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	if c.rep.Config().Data.Delete {
		for _, info := range finds {
			c.log.Infof("id=%v,TenantId=%v", info.ID, 0)
		}
		repository.DeleteByIdsString(ids, repositoryPg.GetOption(ctx))
	} else {
		for _, info := range finds {
			enum := enumStatePg.State(info.State)
			// 有效 停用，反转 为对应的 取消 弃置
			if ok, reverse := enum.ReverseEnableDisable(); ok {
				repository.Update(entityBlog.BlogCategoryEntity{State: reverse.IndexInt8()}, info.ID)
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
func (c *BlogCategoryService) LogicalRecovery(ctx *gin.Context, ids []string) (rt rg.Rs[string]) {
	if len(ids) < 1 {
		return rt.ErrorMessage("id错误")
	}
	repository := c.rep
	finds, b := repository.FindAllByIdStringIn(ids, repositoryPg.GetOption(ctx))
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	for _, info := range finds {
		enum := enumStatePg.State(info.State)
		//  取消 弃置 批量删除，反转 为对应的 有效 停用 停用
		if ok, reverse := enum.ReverseCancelLayAside(); ok {
			repository.Update(entityBlog.BlogCategoryEntity{State: reverse.IndexInt8()}, info.ID)
		}
	}
	return rt.Ok()
}

// PhysicalDeletion 物理删除
//
//	@Description:
//	@receiver c
//	@param ct
func (c *BlogCategoryService) PhysicalDeletion(ctx *gin.Context, ids []string) (rt rg.Rs[string]) {
	if len(ids) < 1 {
		return rt.ErrorMessage("id错误")
	}
	cn := c.rep
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
func (c *BlogCategoryService) Query(ctx *gin.Context, ct modBlogCategory.QueryCt) (rt rg.Rs[pagePg.PaginatorPg[modBlogCategory.Vo]]) {
	var query entityBlog.BlogCategoryEntity
	copier.Copy(&query, &ct)
	slice := make([]modBlogCategory.Vo, 0)
	rt.Data.Data = slice
	r := c.rep
	page, err := r.FindAllPageQuery(query, func(p *pagePg.PageCondition[*entityBlog.BlogCategoryEntity]) {
		p.PageOption = func(c *pagePg.PaginatorPg[*entityBlog.BlogCategoryEntity]) {
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

		pg := pagePg.NewPaginatorPg(func(c *pagePg.PaginatorPg[modBlogCategory.Vo]) {
			c.TotalPage = page.TotalPage
			c.Total = page.Total
			c.PageSize = page.PageSize
			c.PageNum = page.PageNum
		})
		//字段赋值
		for _, item := range page.Data {
			var vo modBlogCategory.Vo
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
func (c *BlogCategoryService) QueryPublic(ctx *gin.Context, ct modBlogCategory.QueryCt) (rt rg.Rs[pagePg.PaginatorPg[modBlogCategory.Vo]]) {
	var query entityBlog.BlogCategoryEntity
	copier.Copy(&query, &ct)
	slice := make([]modBlogCategory.Vo, 0)
	rt.Data.Data = slice
	r := c.rep
	page, err := r.FindAllPageQuery(query, func(p *pagePg.PageCondition[*entityBlog.BlogCategoryEntity]) {
		p.PageOption = func(c *pagePg.PaginatorPg[*entityBlog.BlogCategoryEntity]) {
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

		pg := pagePg.NewPaginatorPg(func(c *pagePg.PaginatorPg[modBlogCategory.Vo]) {
			c.TotalPage = page.TotalPage
			c.Total = page.Total
			c.PageSize = page.PageSize
			c.PageNum = page.PageNum
		})
		//字段赋值
		for _, item := range page.Data {
			var vo modBlogCategory.Vo
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
func (c *BlogCategoryService) SelectNodePublic(ctx *gin.Context, ct modBlogCategory.QueryPublicCt) (rt rg.Rs[[]model.BaseNodeCode]) {
	var query entityBlog.BlogCategoryEntity
	copier.Copy(&query, &ct)
	slice := make([]model.BaseNodeCode, 0)
	rt.Data = slice
	infos := c.rep.FindAll(query, repositoryPg.GetOption(ctx))
	if len(infos) > 0 {

		for _, item := range infos {
			code := model.BaseNodeCode{Key: item.No, Id: item.No, Label: item.Name, ParentId: item.ParentId, Extend: item, Code: item.No}
			//编码
			if !enumParameterPg.NodeQueryByCode.IsEqual(ct.By) {
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
func (c *BlogCategoryService) SelectNodeAllPublic(ctx *gin.Context, ct modBlogCategory.QueryPublicCt) (rt rg.Rs[[]model.BaseNodeCode]) {
	var query entityBlog.BlogCategoryEntity
	copier.Copy(&query, &ct)
	slice := make([]model.BaseNodeCode, 0)
	rt.Data = slice
	infos := c.rep.FindAll(query, repositoryPg.GetOption(ctx))
	if len(infos) > 0 {

		for _, item := range infos {
			var vo modBlogCategory.Vo
			copier.Copy(&vo, &item)
			code := model.BaseNodeCode{Key: item.No, Id: item.No, Label: item.Name, ParentId: item.ParentNo, Extend: vo, Code: item.No}
			//编码
			if !enumParameterPg.NodeQueryByCode.IsEqual(ct.By) {
				code.Key = numberPg.Int64ToString(item.ID)
				code.Id = code.Key
				code.ParentId = vo.ParentId
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
func (c *BlogCategoryService) SelectPublic(ctx *gin.Context, ct modBlogCategory.QueryPublicCt) (rt rg.Rs[[]modBlogCategory.Vo]) {
	var query entityBlog.BlogCategoryEntity
	copier.Copy(&query, &ct)
	slice := make([]modBlogCategory.Vo, 0)
	rt.Data = slice
	infos := c.rep.FindAll(query, repositoryPg.GetOption(ctx))
	if len(infos) > 0 {
		for _, item := range infos {
			var vo modBlogCategory.Vo
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
func (c *BlogCategoryService) ExistName(ctx *gin.Context, ct model.BaseExistWdCt[string]) (rt rg.Rs[string]) {
	if "" == ct.Wd {
		return rt.ErrorMessage("关键词不能为空")
	}
	id := "0"
	if strPg.IsNotBlank(ct.Id) {
		id = ct.Id
	}
	_, result := c.rep.FindByNameAndIdNot(ct.Wd, id, repositoryPg.GetOption(ctx))
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
func (c *BlogCategoryService) ExistCode(ctx *gin.Context, ct model.BaseExistWdCt[string]) (rt rg.Rs[string]) {
	if "" == ct.Wd {
		return rt.ErrorMessage("关键词不能为空")
	}
	if "" == ct.Id {
		return rt.ErrorMessage("id不能为空")
	}
	id := "0"
	if strPg.IsNotBlank(ct.Id) {
		id = ct.Id
	}
	_, result := c.rep.FindByNoAndIdNot(ct.Wd, id, repositoryPg.GetOption(ctx))
	if result {
		return rt.ErrorMessage("重复，已存在")
	}
	return rt.OkMessage("可以使用")
}

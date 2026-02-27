package service

import (
	"context"

	"github.com/foxiswho/blog-go/app/manage/domainRam/utilsRam"
	"github.com/foxiswho/blog-go/app/system/ram/model/modRamResourceAuthority"
	"github.com/foxiswho/blog-go/app/system/ram/service/ramResource"
	"github.com/foxiswho/blog-go/infrastructure/entityRam"
	"github.com/foxiswho/blog-go/infrastructure/repositoryRam"
	"github.com/foxiswho/blog-go/pkg/consts/common/typeSysPg"
	"github.com/foxiswho/blog-go/pkg/consts/constsRam/resourceTypeCategoryPg"
	"github.com/foxiswho/blog-go/pkg/consts/constsRam/typeAttrPg"
	"github.com/foxiswho/blog-go/pkg/enum/state/enumStatePg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/foxiswho/blog-go/pkg/model"
	"github.com/gin-gonic/gin"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"

	"reflect"
	"strings"
	"time"

	"github.com/jinzhu/copier"
	"github.com/pangu-2/go-tools/tools/dbPg/pagePg"
	"github.com/pangu-2/go-tools/tools/numberPg"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
)

func init() {
	gs.Provide(new(RamResourceAuthorityService)).Init(func(s *RamResourceAuthorityService) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

// RamResourceAuthorityService 资源授权
// @Description:
type RamResourceAuthorityService struct {
	sv         *repositoryRam.RamResourceAuthorityRepository     `autowire:"?"`
	resDb      *repositoryRam.RamResourceRepository              `autowire:"?"`
	groupDb    *repositoryRam.RamResourceGroupRepository         `autowire:"?"`
	roleDb     *repositoryRam.RamRoleRepository                  `autowire:"?"`
	grDb       *repositoryRam.RamResourceGroupRelationRepository `autowire:"?"`
	relationDb *repositoryRam.RamResourceRelationRepository      `autowire:"?"`
	//
	log *log2.Logger `autowire:"?"`
}

// Create 新增
//
//	@Description:
//	@receiver c
//	@param ct
//	@return rt
func (c *RamResourceAuthorityService) Create(ctx *gin.Context, ct modRamResourceAuthority.CreateCt) (rt rg.Rs[string]) {
	if "" == ct.Name {
		return rt.ErrorMessage("名称不能为空")
	}
	var info entityRam.RamResourceAuthorityEntity
	copier.Copy(&info, &ct)
	c.log.Infof("info%+v", info)
	c.sv.Create(&info)
	c.log.Infof("save=%+v", info)
	return rg.OkData(numberPg.Int64ToString(info.ID))
}

// CreatByGroup
//
//	@Description: 批量添加 资源组权限
//	@receiver c
//	@param ctx
//	@param ct
//	@return rt
func (c *RamResourceAuthorityService) CreatByGroup(ctx *gin.Context, ct modRamResourceAuthority.CreatByGroupCt) (rt rg.Rs[string]) {
	c.log.Infof("ct=%+v", ct)
	if ct.GroupId.ToInt64() < 1 {
		return rt.ErrorMessage("请选择资源组")
	}
	if ct.Ids == nil || len(ct.Ids) < 1 {
		return rt.ErrorMessage("请选择数据")
	}
	ids := make([]string, 0)
	for _, id := range ct.Ids {
		id = strings.TrimSpace(id)
		ids = append(ids, id)
	}
	if len(ids) < 1 {
		return rt.ErrorMessage("请选择数据")
	}
	group, result := c.groupDb.FindById(ct.GroupId.ToInt64())
	if !result {
		return rt.ErrorMessage("资源组不存在")
	}
	resourceData, r := c.resDb.FindAllByIdStringIn(ids)
	if !r {
		return rt.ErrorMessage("资源数据不存在")
	}
	auth := c.sv
	for _, item := range resourceData {
		//不是资源属性，跳过
		if !typeAttrPg.Resource.IsEqual(item.TypeAttr) {
			continue
		}
		var info entityRam.RamResourceAuthorityEntity
		info.Name = item.Name
		info.NameFl = item.NameFl
		info.NameFull = item.NameFull
		info.Code = item.Code
		info.Description = item.Description
		info.State = enumStatePg.ENABLE.Index()
		info.TypeCategory = resourceTypeCategoryPg.Group.String()
		info.TypeSys = typeSysPg.General.String()
		info.TypeAttr = typeAttrPg.Resource.String()
		info.TypeDomain = group.TypeDomain
		info.GroupId = group.ID
		info.TypeValue = numberPg.Int64ToString(group.ID)
		info.ResourceId = item.ID
		info.Path = item.Path
		info.Method = item.Method
		info.TypeValueSource = numberPg.Int64ToString(item.ID)
		info.Mark = utilsRam.ResourceAuthorityMarkByUint64(resourceTypeCategoryPg.Group, group.ID, item.ID)
		//已存在 则跳过
		_, b := auth.FindByMark(info.Mark)
		if b {
			continue
		}
		//
		c.log.Infof("info%+v", info)
		err, _ := auth.Create(&info)
		if err != nil {
			return rt.ErrorMessage("保存失败")
		}
		c.log.Infof("save=%+v", info)
	}
	return rg.OkData("操作成功")
}

// Update 更新
//
//	@Description:
//	@receiver c
//	@param ct
//	@return rt
func (c *RamResourceAuthorityService) Update(ctx *gin.Context, ct modRamResourceAuthority.UpdateCt) (rt rg.Rs[string]) {
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
	var info entityRam.RamResourceAuthorityEntity
	copier.Copy(&info, &ct)
	r.Update(info, info.ID)
	return rt.Ok()
}

// Detail 详情
//
//	@Description:
//	@receiver c
//	@param id
func (c *RamResourceAuthorityService) Detail(ctx *gin.Context, id int64) (rt rg.Rs[modRamResourceAuthority.Vo]) {
	if id < 1 {
		return rt.ErrorMessage("id错误")
	}
	find, b := c.sv.FindById(id)
	if !b {
		return rt.ErrorMessage("数据不存在")
	}
	var info modRamResourceAuthority.Vo
	copier.Copy(&info, &find)
	return rt.OkData(info)
}

// Enable 启用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamResourceAuthorityService) Enable(ctx *gin.Context, ct model.BaseIdsCt[string]) (rt rg.Rs[string]) {
	return c.State(ctx, ct.Ids, enumStatePg.ENABLE)
}

// Disable 禁用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamResourceAuthorityService) Disable(ctx *gin.Context, ct model.BaseIdsCt[string]) (rt rg.Rs[string]) {
	return c.State(ctx, ct.Ids, enumStatePg.GetType(enumStatePg.DISABLE))
}

// State 状态 启用/禁用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamResourceAuthorityService) State(ctx *gin.Context, ids []string, state enumStatePg.State) (rt rg.Rs[string]) {
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
			r.Update(entityRam.RamResourceAuthorityEntity{State: state.IndexInt8()}, info.ID)
		}
	}
	return rt.Ok()
}

// StateEnableDisable 状态 设置 有效 停用
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamResourceAuthorityService) StateEnableDisable(ctx *gin.Context, ids []string, state enumStatePg.State) (rt rg.Rs[string]) {
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
func (c *RamResourceAuthorityService) LogicalDeletion(ctx *gin.Context, ids []string) (rt rg.Rs[string]) {
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
			c.log.Infof("id=%v,TenantId=%v", info.ID, " info.TenantNo")
		}
		repository.DeleteByIdsString(ids)
	} else {
		for _, info := range finds {
			enum := enumStatePg.State(info.State)
			// 有效 停用，反转 为对应的 取消 弃置
			if ok, reverse := enum.ReverseEnableDisable(); ok {
				repository.Update(entityRam.RamResourceAuthorityEntity{State: reverse.IndexInt8()}, info.ID)
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
func (c *RamResourceAuthorityService) LogicalRecovery(ctx *gin.Context, ids []string) (rt rg.Rs[string]) {
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
			repository.Update(entityRam.RamResourceAuthorityEntity{State: reverse.IndexInt8()}, info.ID)
		}
	}
	return rt.Ok()
}

// PhysicalDeletion 物理删除
//
//	@Description:
//	@receiver c
//	@param ct
func (c *RamResourceAuthorityService) PhysicalDeletion(ctx *gin.Context, ids []string) (rt rg.Rs[string]) {
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
		c.log.Infof("id=%v,TenantId=%v", info.ID, 0)
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
func (c *RamResourceAuthorityService) Query(ctx *gin.Context, ct modRamResourceAuthority.QueryCt) (rt rg.Rs[pagePg.PaginatorPg[modRamResourceAuthority.Vo]]) {
	c.log.Infof("ct=%+v", ct)
	var query entityRam.RamResourceAuthorityEntity
	copier.Copy(&query, &ct)
	slice := make([]modRamResourceAuthority.Vo, 0)
	rt.Data.Data = slice
	r := c.sv
	page, err := r.FindAllPageQuery(query, func(p *pagePg.PageCondition[*entityRam.RamResourceAuthorityEntity]) {
		p.PageOption = func(c *pagePg.PaginatorPg[*entityRam.RamResourceAuthorityEntity]) {
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

		pg := pagePg.NewPaginatorPg(func(c *pagePg.PaginatorPg[modRamResourceAuthority.Vo]) {
			c.TotalPage = page.TotalPage
			c.Total = page.Total
			c.PageSize = page.PageSize
			c.PageNum = page.PageNum
		})
		//字段赋值
		for _, item := range page.Data {
			var vo modRamResourceAuthority.Vo
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
func (c *RamResourceAuthorityService) SelectNodePublic(ctx *gin.Context, ct modRamResourceAuthority.QueryCt) (rt rg.Rs[[]model.BaseNode]) {
	c.log.Infof("ct=%+v", ct)
	var query entityRam.RamResourceAuthorityEntity
	copier.Copy(&query, &ct)
	slice := make([]model.BaseNode, 0)
	rt.Data = slice
	infos := c.sv.FindAll(query)
	if len(infos) > 0 {
		for _, item := range infos {
			slice = append(slice, model.BaseNode{Key: numberPg.Int64ToString(item.ID),
				Label: item.Name,
				Code:  numberPg.Int64ToString(item.ID),
				Id:    numberPg.Int64ToString(item.ID),
			})
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
func (c *RamResourceAuthorityService) SelectNodeAllPublic(ctx *gin.Context, ct modRamResourceAuthority.QueryCt) (rt rg.Rs[[]model.BaseNode]) {
	c.log.Infof("ct=%+v", ct)
	var query entityRam.RamResourceAuthorityEntity
	copier.Copy(&query, &ct)
	slice := make([]model.BaseNode, 0)
	rt.Data = slice
	infos := c.sv.FindAll(query)
	if len(infos) > 0 {
		for _, item := range infos {
			var vo modRamResourceAuthority.Vo
			copier.Copy(&vo, &item)
			slice = append(slice, model.BaseNode{Key: numberPg.Int64ToString(item.ID),
				Label:  item.Name,
				Code:   numberPg.Int64ToString(item.ID),
				Id:     numberPg.Int64ToString(item.ID),
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
func (c *RamResourceAuthorityService) SelectPublic(ctx *gin.Context, ct modRamResourceAuthority.QueryCt) (rt rg.Rs[[]modRamResourceAuthority.Vo]) {
	c.log.Infof("ct=%+v", ct)
	var query entityRam.RamResourceAuthorityEntity
	copier.Copy(&query, &ct)
	rt.Data = []modRamResourceAuthority.Vo{}
	infos := c.sv.FindAll(query)
	if len(infos) > 0 {
		slice := make([]modRamResourceAuthority.Vo, 0)
		for _, item := range infos {
			var vo modRamResourceAuthority.Vo
			copier.Copy(&vo, &item)
			slice = append(slice, vo)
		}
		rt.Data = slice
	}
	return rt.Ok()
}

// CreateByResourceGroup 新增
//
//	@Description:
//	@receiver c
//	@param ct
//	@return rt
func (c *RamResourceAuthorityService) CreateByResourceGroup(ctx *gin.Context, value *entityRam.RamResourceGroupEntity) (*entityRam.RamResourceAuthorityEntity, bool) {
	var info entityRam.RamResourceAuthorityEntity
	copier.Copy(&info, value)
	now := time.Now()
	info.ID = 0
	info.CreateAt = &now
	info.UpdateAt = &now
	info.TypeCategory = resourceTypeCategoryPg.Group.String()
	info.TypeSys = typeSysPg.General.String()
	info.TypeAttr = typeAttrPg.Resource.String()
	info.TypeValue = numberPg.Int64ToString(value.ID)
	info.Mark = info.TypeCategory + ":" + info.TypeValue
	r := c.sv
	mark, b := r.FindByMark(info.Mark)
	if b {
		return mark, true
	}
	r.Create(&info)
	return &info, true
}

// UpdateByRole
//
//	@Description: 授权 角色 资源权限
//	@receiver c
//	@param ctx
//	@param ct
//	@return rt
func (c *RamResourceAuthorityService) UpdateByRole(ctx *gin.Context, ct modRamResourceAuthority.UpdateByTypeValueCt) (rt rg.Rs[string]) {
	c.log.Infof("ct=%+v", ct)
	return ramResource.NewUpdateByTypeValue(c.log, c.roleDb, c.sv, c.resDb, c.groupDb, c.grDb, c.relationDb, resourceTypeCategoryPg.Role, ct, ctx).Process()
}

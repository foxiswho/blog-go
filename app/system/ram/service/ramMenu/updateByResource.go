package ramMenu

import (
	"github.com/foxiswho/blog-go/app/manage/domainRam/utilsRam"
	"github.com/foxiswho/blog-go/app/system/ram/model/modRamResourceMenu"
	"github.com/foxiswho/blog-go/infrastructure/entityRam"
	"github.com/foxiswho/blog-go/infrastructure/repositoryRam"
	"github.com/foxiswho/blog-go/pkg/consts/common/typeSysPg"
	iamConstant2 "github.com/foxiswho/blog-go/pkg/consts/constsRam/menuTypePg"
	"github.com/foxiswho/blog-go/pkg/consts/constsRam/typeAttrPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/gin-gonic/gin"

	"github.com/jinzhu/copier"
	"github.com/pangu-2/go-tools/tools/numberPg"
	"github.com/pangu-2/go-tools/tools/strPg"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
)

// UpdateByResource
// @Description: 更新菜单 资源/资源组关系
type UpdateByResource struct {
	log            *log2.Logger
	menuDb         *repositoryRam.RamMenuRepository              `autowire:"?"`
	menuRelationDb *repositoryRam.RamMenuRelationRepository      `autowire:"?"`
	authDb         *repositoryRam.RamResourceAuthorityRepository `autowire:"?"`
	resMenuDb      *repositoryRam.RamResourceMenuRepository      `autowire:"?"`
	groupDb        *repositoryRam.RamResourceGroupRepository     `autowire:"?"`
	resDb          *repositoryRam.RamResourceRepository          `autowire:"?"`
	ct             modRamResourceMenu.UpdateByMenuCt
	ctx            *gin.Context
	//
	menuData       *entityRam.RamMenuEntity
	authorityData  *entityRam.RamResourceAuthorityEntity
	idsGroup       []string
	idsResource    []string
	idsGroupNew    []string
	idsResourceNew []string
}

// NewUpdateByResource
//
//	@Description: 更新菜单 资源/资源组关系
//	@param log
//	@param menuDb
//	@param menuRelationDb
//	@param authDb
//	@param resMenuDb
//	@param groupDb
//	@param resDb
//	@param ct
//	@param ctx
//	@return *UpdateByResource
func NewUpdateByResource(
	log *log2.Logger,
	menuDb *repositoryRam.RamMenuRepository,
	menuRelationDb *repositoryRam.RamMenuRelationRepository,
	authDb *repositoryRam.RamResourceAuthorityRepository,
	resMenuDb *repositoryRam.RamResourceMenuRepository,
	groupDb *repositoryRam.RamResourceGroupRepository,
	resDb *repositoryRam.RamResourceRepository,
	ct modRamResourceMenu.UpdateByMenuCt,
	ctx *gin.Context) *UpdateByResource {
	return &UpdateByResource{
		log:            log,
		menuDb:         menuDb,
		menuRelationDb: menuRelationDb,
		authDb:         authDb,
		resMenuDb:      resMenuDb,
		groupDb:        groupDb,
		resDb:          resDb,
		ctx:            ctx,
		ct:             ct,
	}
}

// Process
//
//	@Description:  处理
//	@receiver c
//	@return rt
func (c *UpdateByResource) Process() (rt rg.Rs[string]) {
	if c.ct.Data == nil || 0 == len(c.ct.Data) {
		return rt.ErrorMessage("数据不能为空")
	}
	if c.ct.MenuId <= 0 {
		return rt.ErrorMessage("菜单id不能为空")
	}
	result := false
	c.menuData, result = c.menuDb.FindByIdString(c.ct.MenuId.ToString())
	if !result {
		return rt.ErrorMessage("菜单不存在")
	}
	c.idsGroup = make([]string, 0)
	c.idsResource = make([]string, 0)
	c.idsGroupNew = make([]string, 0)
	c.idsResourceNew = make([]string, 0)
	//获取数据
	c.getData()
	//保存 菜单资源/资源组关系
	c.updateDataRelation(c.idsGroup, c.idsResource)
	// 更新实际资源菜单数据
	c.updateResourceMenu()
	return rt.Ok()
}

// 获取数据
func (c *UpdateByResource) getData() {
	for _, v := range c.ct.Data {
		//资源组
		if iamConstant2.Group.IsEqual(v.Type) && strPg.IsNotBlank(v.Id) {
			c.idsGroup = append(c.idsGroup, v.Id)
		}
		//资源
		if iamConstant2.Resource.IsEqual(v.Type) && strPg.IsNotBlank(v.Id) {
			c.idsResource = append(c.idsResource, v.Id)
		}
	}
	//

}

// updateDataRelation
//
//	@Description:  保存 菜单资源/资源组关系
//	@receiver c
//	@param idsGroup
//	@param idsResource
func (c *UpdateByResource) updateDataRelation(idsGroup, idsResource []string) {
	saveData := make([]entityRam.RamMenuRelationEntity, 0)
	mapData := make(map[string]struct{})
	mapDataOld := make(map[string]struct{})
	//资源组
	if len(idsGroup) > 0 {
		list, result := c.groupDb.FindAllByIdStringIn(idsGroup)
		if result {
			for _, item := range list {
				var info entityRam.RamMenuRelationEntity
				info.Name = item.Name
				info.Code = item.Code
				info.Type = iamConstant2.Group.String()
				info.TypeValue = numberPg.Int64ToString(item.ID)
				info.TypeSys = typeSysPg.General.String()
				info.MenuId = c.menuData.ID
				saveData = append(saveData, info)
				//
				mapData[utilsRam.MenuTypeByRelation(info.TypeValue, info.Type)] = struct{}{}
			}
		}
	}
	//资源
	if len(idsResource) > 0 {
		list, result := c.resDb.FindAllByIdStringIn(idsResource)
		if result {
			for _, item := range list {
				var info entityRam.RamMenuRelationEntity
				info.Name = item.Name
				info.Code = item.Code
				info.Type = iamConstant2.Resource.String()
				info.TypeValue = numberPg.Int64ToString(item.ID)
				info.TypeSys = typeSysPg.General.String()
				info.MenuId = c.menuData.ID
				saveData = append(saveData, info)
				//
				mapData[utilsRam.MenuTypeByRelation(info.TypeValue, info.Type)] = struct{}{}
			}
		}
	}

	//获取已经保存过的数据
	oldData, result := c.menuRelationDb.FindAllByMenuIdIn([]int64{c.menuData.ID})
	if result {
		for _, item := range oldData {
			mapDataOld[utilsRam.MenuTypeByRelation(item.TypeValue, item.Type)] = struct{}{}
		}
		c.log.Infof("mapDataOld=%+v", mapDataOld)
	}
	//保存数据
	if len(saveData) > 0 {
		for _, item := range saveData {
			c.log.Infof("find.%+v,menuId=%+v", utilsRam.MenuTypeByRelation(item.TypeValue, item.Type), item.MenuId)
			//已经存在的跳过
			if _, ok := mapDataOld[utilsRam.MenuTypeByRelation(item.TypeValue, item.Type)]; ok {
				continue
			}
			c.menuRelationDb.Create(&item)
		}
	}
}

// updateResourceMenu
//
//	@Description: 更新实际资源菜单数据
//	@receiver c
func (c *UpdateByResource) updateResourceMenu() {
	//删除 原始数据
	c.resMenuDb.DeleteByMenuId(c.menuData.ID)
	//
	//获取已经保存过的数据
	{
		oldData, result := c.menuRelationDb.FindAllByMenuIdIn([]int64{c.menuData.ID})
		if result {
			for _, item := range oldData {
				//资源组
				if iamConstant2.Group.IsEqual(item.Type) {
					c.idsGroupNew = append(c.idsGroupNew, item.TypeValue)
				} else if iamConstant2.Resource.IsEqual(item.Type) {
					//资源
					c.idsResourceNew = append(c.idsResourceNew, item.TypeValue)
				}
			}
		}
	}
	c.log.Debugf("idsGroupNew=[%+v],idsResourceNew=[%+v]", c.idsGroupNew, c.idsResourceNew)
	saveData := make([]entityRam.RamResourceMenuEntity, 0)
	//资源组
	{
		if len(c.idsGroupNew) > 0 {
			list, result := c.authDb.FindAllByGroupIdStringIn(c.idsGroupNew)
			if result {
				for _, item := range list {
					if !typeAttrPg.Resource.IsEqual(item.TypeAttr) {
						continue
					}
					var info entityRam.RamResourceMenuEntity
					copier.Copy(&info, &item)
					info.ID = 0
					info.ResourceId = item.ID
					info.GroupId = item.GroupId
					info.ParentId = c.menuData.ParentId
					info.IdLink = c.menuData.IdLink
					info.MenuId = c.menuData.ID
					saveData = append(saveData, info)
				}
			}
		}
	}
	//资源
	{
		if len(c.idsResourceNew) > 0 {
			list, result := c.resDb.FindAllByIdStringIn(c.idsResourceNew)
			if result {
				for _, item := range list {
					//已经存在的跳过
					if !typeAttrPg.Resource.IsEqual(item.TypeAttr) {
						continue
					}
					var info entityRam.RamResourceMenuEntity
					copier.Copy(&info, &item)
					info.ID = 0
					info.ResourceId = item.ID
					info.ParentId = c.menuData.ParentId
					info.IdLink = c.menuData.IdLink
					info.MenuId = c.menuData.ID
					saveData = append(saveData, info)
				}
			}
		}
	}
	c.log.Debugf("saveData=%+v", saveData)
	//保存数据
	if len(saveData) > 0 {
		for _, item := range saveData {
			c.resMenuDb.Create(&item)
		}
	}
}

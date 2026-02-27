package controller

import (
	"github.com/foxiswho/blog-go/app/system/ram/model/modRamMenuRelation"
	"github.com/foxiswho/blog-go/app/system/ram/model/modRamResourceMenu"
	service2 "github.com/foxiswho/blog-go/app/system/ram/service"
	"github.com/foxiswho/blog-go/middleware/validatorPg"
	"github.com/foxiswho/blog-go/pkg/common/controllerPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/gin-gonic/gin"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"

	"github.com/foxiswho/blog-go/pkg/model"
)

// MenuRelationController
// @Description: 资源菜单关系
type MenuRelationController struct {
	controllerPg.SpSystemAuth
	sv  *service2.RamResourceMenuService `autowire:"?"`
	rel *service2.RamMenuRelationService `autowire:"?"`
	log *log2.Logger                     `autowire:"?"`
}

// UpdateByMenu 更新资源菜单关系
//
//	@Description:
//	@receiver c
//	@param ctx
func (c *MenuRelationController) UpdateByMenu(ctx *gin.Context) {
	var ct modRamResourceMenu.UpdateByMenuCt
	if err := ctx.ShouldBind(&ct); err != nil {
		//对 返回 错误进行转义 成中文
		translate := validatorPg.Translate(err, &ct)
		if len(translate) > 0 {
			ctx.JSON(200, rg.ErrorMessageData[string](translate))
			return
		}
		ctx.JSON(200, rg.ErrorDefault[string]())
		return
	}
	ctx.JSON(200, c.sv.UpdateByMenu(ctx, ct))
}

// Query 列表
//
//	@Description:
//	@receiver c
//	@param ctx
func (c *MenuRelationController) Query(ctx *gin.Context) {
	var ct modRamMenuRelation.QueryCt
	if err := ctx.ShouldBind(&ct); err != nil {
		//对 返回 错误进行转义 成中文
		translate := validatorPg.Translate(err, &ct)
		if len(translate) > 0 {
			ctx.JSON(200, rg.ErrorMessageData[string](translate))
			return
		}
		ctx.JSON(200, rg.ErrorDefault[string]())
		return
	}
	ctx.JSON(200, c.rel.Query(ctx, ct))
}

// PhysicalDeletion 物理删除
//
//	@Description:
//	@receiver c
//	@param ctx
func (c *MenuRelationController) PhysicalDeletion(ctx *gin.Context) {
	var ct model.BaseIdsCt[string]
	if err := ctx.ShouldBind(&ct); err != nil {
		//对 返回 错误进行转义 成中文
		translate := validatorPg.Translate(err, &ct)
		if len(translate) > 0 {
			ctx.JSON(200, rg.ErrorMessageData[string](translate))
			return
		}
		ctx.JSON(200, rg.ErrorDefault[string]())
		return
	}
	ctx.JSON(200, c.rel.PhysicalDeletion(ctx, ct.Ids))
}

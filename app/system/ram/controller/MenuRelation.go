package controller

import (
	"github.com/foxiswho/blog-go/app/system/ram/model/modRamMenuRelation"
	"github.com/foxiswho/blog-go/app/system/ram/model/modRamResourceMenu"
	service2 "github.com/foxiswho/blog-go/app/system/ram/service"
	"github.com/foxiswho/blog-go/middleware/authPg"
	"github.com/foxiswho/blog-go/middleware/serverPg/ginServer"
	"github.com/foxiswho/blog-go/pkg/common/controllerPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/foxiswho/blog-go/pkg/model"
	"github.com/foxiswho/blog-go/pkg/routerPg"
	"github.com/gin-gonic/gin"
	"go-spring.org/spring/gs"
)

func init() {
	gs.Provide(new(MenuRelationController)).Name("SystemMenuRelationController").Export(gs.As[routerPg.RouteRegistrar]())
}

type MenuRelationController struct {
	routerPg.RouteRegistrar
	controllerPg.SpSystemAuth
	sv  *service2.RamResourceMenuService `autowire:"?"`
	rel *service2.RamMenuRelationService `autowire:"?"`
	log *log2.Logger                     `autowire:"?"`
}

func (c *MenuRelationController) RegisterRoutes(e *gin.Engine) {
	r := ginServer.GinServerDefault
	group := r.Group("/pg2lq/sys/ram/menu-relation", authPg.GroupSystemMiddleware(c.Sp))
	group.POST("/updateByMenu", c.UpdateByMenu)
	group.POST("/query", c.Query)
	group.POST("/physicalDeletion", c.PhysicalDeletion)
}

func (c *MenuRelationController) UpdateByMenu(ctx *gin.Context) {
	var ct modRamResourceMenu.UpdateByMenuCt
	if !routerPg.BindJson(ctx, &ct) {
		return
	}
	ctx.JSON(200, c.sv.UpdateByMenu(ctx, ct))
}

func (c *MenuRelationController) Query(ctx *gin.Context) {
	var ct modRamMenuRelation.QueryCt
	if !routerPg.BindJson(ctx, &ct) {
		return
	}
	ctx.JSON(200, c.rel.Query(ctx, ct))
}

func (c *MenuRelationController) PhysicalDeletion(ctx *gin.Context) {
	var ct model.BaseIdsCt[string]
	if !routerPg.BindJson(ctx, &ct) {
		return
	}
	ctx.JSON(200, c.rel.PhysicalDeletion(ctx, ct.Ids))
}

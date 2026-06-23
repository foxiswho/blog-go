package controller

import (
	"github.com/foxiswho/blog-go/app/system/ram/model/modRamResourceRelation"
	"github.com/foxiswho/blog-go/app/system/ram/service"
	"github.com/foxiswho/blog-go/middleware/authPg"
	"github.com/foxiswho/blog-go/middleware/serverPg/ginServer"
	"github.com/foxiswho/blog-go/pkg/common/controllerPg"
	"github.com/foxiswho/blog-go/pkg/enum/state/enumStatePg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/foxiswho/blog-go/pkg/routerPg"
	"github.com/gin-gonic/gin"
	"go-spring.org/spring/gs"
)

func init() {
	gs.Provide(new(ResourceRelationController)).Name("SystemResourceRelationController").Export(gs.As[routerPg.RouteRegistrar]())
}

type ResourceRelationController struct {
	routerPg.RouteRegistrar
	controllerPg.SpSystemAuth
	sv  *service.RamResourceRelationService `autowire:"?"`
	log *log2.Logger                        `autowire:"?"`
}

func (c *ResourceRelationController) RegisterRoutes(e *gin.Engine) {
	r := ginServer.GinServerDefault
	group := r.Group("/pg2lq/sys/ram/resource-relation", authPg.GroupSystemMiddleware(c.Sp))
	group.POST("/selected", c.Selected)
}

func (c *ResourceRelationController) Query(ctx *gin.Context) {
	var ct modRamResourceRelation.QueryCt
	if !routerPg.BindJson(ctx, &ct) {
		return
	}
	ctx.JSON(200, c.sv.Query(ctx, ct))
}

func (c *ResourceRelationController) SelectNodePublic(ctx *gin.Context) {
	ct := modRamResourceRelation.QueryCt{State: enumStatePg.ENABLE.IndexPg()}
	ctx.JSON(200, c.sv.SelectNodePublic(ctx, ct))
}

func (c *ResourceRelationController) SelectNodeAllPublic(ctx *gin.Context) {
	ct := modRamResourceRelation.QueryCt{State: enumStatePg.ENABLE.IndexPg()}
	ctx.JSON(200, c.sv.SelectNodeAllPublic(ctx, ct))
}

func (c *ResourceRelationController) SelectPublic(ctx *gin.Context) {
	ct := modRamResourceRelation.QueryCt{State: enumStatePg.ENABLE.IndexPg()}
	ctx.JSON(200, c.sv.SelectPublic(ctx, ct))
}

func (c *ResourceRelationController) Selected(ctx *gin.Context) {
	var ct modRamResourceRelation.QuerySelectedCt
	if !routerPg.BindJson(ctx, &ct) {
		return
	}
	ctx.JSON(200, c.sv.Selected(ctx, ct.Code))
}

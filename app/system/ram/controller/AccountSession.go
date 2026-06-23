package controller

import (
	"github.com/foxiswho/blog-go/app/system/ram/model/modRamAccountSession"
	"github.com/foxiswho/blog-go/app/system/ram/service"
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
	gs.Provide(new(AccountSessionController)).Name("SystemAccountSessionController").Export(gs.As[routerPg.RouteRegistrar]())
}

type AccountSessionController struct {
	routerPg.RouteRegistrar
	controllerPg.SpSystemAuth
	sv  *service.RamAccountSessionService `autowire:"?"`
	log *log2.Logger                      `autowire:"?"`
}

func (c *AccountSessionController) RegisterRoutes(e *gin.Engine) {
	r := ginServer.GinServerDefault
	group := r.Group("/pg2lq/sys/ram/account-session", authPg.GroupSystemMiddleware(c.Sp))
	group.POST("/physicalDeletion", c.PhysicalDeletion)
	group.POST("/query", c.Query)
}

func (c *AccountSessionController) PhysicalDeletion(ctx *gin.Context) {
	var ct model.BaseIdsCt[string]
	if !routerPg.BindJson(ctx, &ct) {
		return
	}
	ctx.JSON(200, c.sv.PhysicalDeletion(ctx, ct.Ids))
}

func (c *AccountSessionController) Query(ctx *gin.Context) {
	var ct modRamAccountSession.QueryCt
	if !routerPg.BindJson(ctx, &ct) {
		return
	}
	ctx.JSON(200, c.sv.Query(ctx, ct))
}

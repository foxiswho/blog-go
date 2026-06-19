package controller

import (
	"github.com/foxiswho/blog-go/app/system/ram/service"
	"github.com/foxiswho/blog-go/middleware/serverPg/ginServer"
	"github.com/foxiswho/blog-go/pkg/holderPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/foxiswho/blog-go/pkg/routerPg"
	"github.com/gin-gonic/gin"
	"go-spring.org/spring/gs"
)

func init() {
	gs.Provide(new(LogoutController)).Name("SystemLogoutController").Export(gs.As[routerPg.RouteRegistrar]())
}

type LogoutController struct {
	routerPg.RouteRegistrar
	sv  *service.AccountLogoutService `autowire:"?"`
	log *log2.Logger                  `autowire:"?"`
}

func (c *LogoutController) RegisterRoutes(e *gin.Engine) {
	r := ginServer.GinServerDefault
	group := r.Group("/pg2lq/auth/sys")
	group.Any("/logout", c.Logout)
}

func (c *LogoutController) Logout(ctx *gin.Context) {
	ctx.JSON(200, c.sv.Logout(holderPg.GetContextAccount(ctx)))
}

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hongmengzhu/xianfu-blog-go/app/system/ram/service"
	"github.com/hongmengzhu/xianfu-blog-go/middleware/serverPg/ginServer"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/holderPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/log2"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/routerPg"
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
	group := r.Group("/xianfu/auth/sys")
	group.Any("/logout", c.Logout)
}

func (c *LogoutController) Logout(ctx *gin.Context) {
	ctx.JSON(200, c.sv.Logout(holderPg.GetContextAccount(ctx)))
}

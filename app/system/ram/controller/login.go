package controller

import (
	"github.com/gin-gonic/gin"
	modRamLogin2 "github.com/hongmengzhu/xianfu-blog-go/app/system/ram/model/modRamLogin"
	"github.com/hongmengzhu/xianfu-blog-go/app/system/ram/service"
	"github.com/hongmengzhu/xianfu-blog-go/middleware/serverPg/ginServer"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/enum/enumCommonPg/appModulePg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/log2"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/routerPg"
	"go-spring.org/spring/gs"
)

func init() {
	gs.Provide(new(LoginController)).Name("SystemLoginController").Export(gs.As[routerPg.RouteRegistrar]())
}

type LoginController struct {
	routerPg.RouteRegistrar
	sv  *service.AccountLoginService `autowire:"?"`
	log *log2.Logger                 `autowire:"?"`
}

func (c *LoginController) RegisterRoutes(e *gin.Engine) {
	r := ginServer.GinServerDefault
	group := r.Group("/xianfu/auth/sys")
	group.POST("/login", c.Login)
	group.POST("/refresh", c.RefreshToken)
}

func (c *LoginController) Login(ctx *gin.Context) {
	var ct modRamLogin2.LoginCt
	if !routerPg.BindJson(ctx, &ct) {
		return
	}
	ctx.JSON(200, c.sv.Login(ctx, ct, appModulePg.System))
}

func (c *LoginController) RefreshToken(ctx *gin.Context) {
	var ct modRamLogin2.TokenRefreshCt
	if !routerPg.BindJson(ctx, &ct) {
		return
	}
	ctx.JSON(200, c.sv.RefreshToken(ctx, ct))
}

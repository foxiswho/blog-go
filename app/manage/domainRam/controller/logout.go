package controller

import (
	"github.com/foxiswho/blog-go/app/manage/domainRam/service"
	"github.com/foxiswho/blog-go/pkg/holderPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/foxiswho/blog-go/pkg/routerPg"
	"github.com/gin-gonic/gin"
	"go-spring.org/spring/gs"
)

func init() {
	gs.Provide(new(LogoutController)).Name("ManageLogoutController").Export(gs.As[routerPg.RouteRegistrar]())
}

// LogoutController 退出
// @Description:
type LogoutController struct {
	routerPg.RouteRegistrar
	sv  *service.AccountLogoutService `autowire:"?"`
	log *log2.Logger                  `autowire:"?"`
}

// RegisterRoutes 注册路由
//
//	@Description:
//	@receiver c
//	@param e
func (c *LogoutController) RegisterRoutes(e *gin.Engine) {
	group := e.Group("/pg2lq/auth/manage")
	group.Any("/logout", c.Logout)
}

// Logout 退出
//
//	@Description:
//	@receiver c
//	@param ctx
func (c *LogoutController) Logout(ctx *gin.Context) {
	ctx.JSON(200, c.sv.Logout(holderPg.GetContextAccount(ctx)))
}

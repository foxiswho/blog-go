package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hongmengzhu/xianfu-blog-go/app/manage/domainRam/service"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/holderPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/log2"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/routerPg"
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

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hongmengzhu/xianfu-blog-go/app/system/ram/model/modPublic"
	"github.com/hongmengzhu/xianfu-blog-go/app/system/ram/service"
	"github.com/hongmengzhu/xianfu-blog-go/cmd"
	"github.com/hongmengzhu/xianfu-blog-go/middleware/authPg"
	"github.com/hongmengzhu/xianfu-blog-go/middleware/serverPg/ginServer"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/common/controllerPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/holderPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/log2"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/routerPg"
	"go-spring.org/spring/gs"
)

func init() {
	gs.Provide(new(PublicController)).Name("SystemPublicController").Export(gs.As[routerPg.RouteRegistrar]())
}

type PublicController struct {
	routerPg.RouteRegistrar
	controllerPg.SpSystemAuth
	sv  *service.RamAccountPublicService `autowire:"?"`
	log *log2.Logger                     `autowire:"?"`
}

func (c *PublicController) RegisterRoutes(e *gin.Engine) {
	r := ginServer.GinServerDefault
	group := r.Group("/xianfu/sys/public", authPg.GroupSystemMiddleware(c.Sp))
	group.GET("/info", c.Public)
	group.GET("/infoPublic", c.InfoPublic)
	group.POST("/password", c.UpdatePassword)
	group.GET("/envInfoPublic", cmd.GetVersion)
}

func (c *PublicController) Public(ctx *gin.Context) {
	ctx.JSON(200, c.sv.Public(holderPg.GetContextAccount(ctx)))
}
func (c *PublicController) InfoPublic(ctx *gin.Context) {
	ctx.JSON(200, c.sv.InfoPublic(holderPg.GetContextAccount(ctx)))
}

func (c *PublicController) UpdatePassword(ctx *gin.Context) {
	var ct modPublic.PasswordCt
	if !routerPg.BindJson(ctx, &ct) {
		return
	}
	ctx.JSON(200, c.sv.UpdatePassword(ctx, ct))
}

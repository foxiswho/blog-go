package controller

import (
	"github.com/foxiswho/blog-go/app/system/ram/model/modPublic"
	"github.com/foxiswho/blog-go/app/system/ram/service"
	"github.com/foxiswho/blog-go/cmd"
	"github.com/foxiswho/blog-go/middleware/authPg"
	"github.com/foxiswho/blog-go/middleware/validatorPg"
	"github.com/foxiswho/blog-go/middleware/serverPg/ginServer"
	"github.com/foxiswho/blog-go/pkg/common/controllerPg"
	"github.com/foxiswho/blog-go/pkg/holderPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/foxiswho/blog-go/pkg/routerPg"
	"github.com/gin-gonic/gin"
	"github.com/go-spring/spring-core/gs"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
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
	group := r.Group("/pg2lq/sys/public", authPg.GroupSystemMiddleware(c.Sp))
	group.GET("/info", c.Public)
	group.POST("/password", c.UpdatePassword)
	group.GET("/envInfoPublic", cmd.GetVersion)
}

func (c *PublicController) Public(ctx *gin.Context) {
	ctx.JSON(200, c.sv.Public(holderPg.GetContextAccount(ctx)))
}

func (c *PublicController) UpdatePassword(ctx *gin.Context) {
	var ct modPublic.PasswordCt
	if err := ctx.ShouldBind(&ct); err != nil {
		translate := validatorPg.Translate(err, &ct)
		if len(translate) > 0 {
			ctx.JSON(200, rg.ErrorMessageData[string](translate))
			return
		}
		ctx.JSON(200, rg.ErrorDefault[string]())
		return
	}
	ctx.JSON(200, c.sv.UpdatePassword(ctx, ct))
}

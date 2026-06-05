package controller

import (
	"github.com/foxiswho/blog-go/app/web/api/model/modBlogCollect"
	"github.com/foxiswho/blog-go/app/web/api/service"
	"github.com/foxiswho/blog-go/middleware/authPg"
	"github.com/foxiswho/blog-go/middleware/validatorPg"
	"github.com/foxiswho/blog-go/pkg/routerPg"
	"github.com/gin-gonic/gin"
	"github.com/go-spring/spring-core/gs"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
)

func init() {
	gs.Provide(new(CollectController)).Export(gs.As[routerPg.RouteRegistrar]())
}

type CollectController struct {
	routerPg.RouteRegistrar
	Sp *authPg.GroupApiMiddlewareSp `autowire:""`
	sv *service.CollectService      `autowire:""`
}

// RegisterRoutes
//
//	@Description: 注册路由
//	@receiver c
//	@param e
func (c *CollectController) RegisterRoutes(e *gin.Engine) {
	group := e.Group("/api/collect", authPg.GroupApiMiddleware(c.Sp))
	group.POST("/push", c.Push)
	group.POST("/pushAll", c.PushAll)
}

// Push 推送
//
//	@Description:
//	@receiver c
//	@param ctx
func (c *CollectController) Push(ctx *gin.Context) {
	var ct modBlogCollect.PushCt
	if err := ctx.ShouldBind(&ct); err != nil {
		//对 返回 错误进行转义 成中文
		translate := validatorPg.Translate(err, &ct)
		if len(translate) > 0 {
			ctx.JSON(200, rg.ErrorMessageData[string](translate))
			return
		}
		ctx.JSON(200, rg.ErrorDefault[string]())
		return
	}
	ctx.JSON(200, c.sv.Push(ctx, ct))
}

// PushAll 推送
//
//	@Description:
//	@receiver c
//	@param ctx
func (c *CollectController) PushAll(ctx *gin.Context) {
	var ct modBlogCollect.PushAll
	if err := ctx.ShouldBind(&ct); err != nil {
		//对 返回 错误进行转义 成中文
		translate := validatorPg.Translate(err, &ct)
		if len(translate) > 0 {
			ctx.JSON(200, rg.ErrorMessageData[string](translate))
			return
		}
		ctx.JSON(200, rg.ErrorDefault[string]())
		return
	}
	ctx.JSON(200, c.sv.PushAll(ctx, ct))
}

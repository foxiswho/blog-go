package controller

import (
	"github.com/foxiswho/blog-go/app/web/api/model/modBlogCollect"
	"github.com/foxiswho/blog-go/app/web/api/service"
	"github.com/foxiswho/blog-go/middleware/authPg"
	"github.com/foxiswho/blog-go/middleware/validatorPg"
	"github.com/gin-gonic/gin"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
)

type CollectController struct {
	Sp *authPg.GroupApiMiddlewareSp `autowire:""`
	sv *service.CollectService      `autowire:""`
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

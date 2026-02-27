package controller

import (
	"github.com/foxiswho/blog-go/app/web/api/model/modBlogArticleCategory"
	"github.com/foxiswho/blog-go/app/web/api/service"
	"github.com/foxiswho/blog-go/middleware/authPg"
	"github.com/foxiswho/blog-go/middleware/validatorPg"
	"github.com/gin-gonic/gin"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
)

type ArticleCategoryController struct {
	Sp *authPg.GroupApiMiddlewareSp    `autowire:""`
	sv *service.ArticleCategoryService `autowire:""`
}

func (c *ArticleCategoryController) SelectNodeAllPublic(ctx *gin.Context) {
	var ct modBlogArticleCategory.QueryPublicCt
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
	ctx.JSON(200, c.sv.SelectNodeAllPublic(ctx, ct))
}

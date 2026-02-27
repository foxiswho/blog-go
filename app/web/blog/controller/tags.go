package controller

import (
	"github.com/foxiswho/blog-go/middleware/authPg"
	"github.com/gin-gonic/gin"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
)

type TagsController struct {
	Sp *authPg.GroupWebMiddlewareSp `autowire:""`
}

func (c *TagsController) Detail(ctx *gin.Context) {

	ctx.JSON(200, rg.Ok[string]("ok"))
}

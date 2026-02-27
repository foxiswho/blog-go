package controller

import (
	"github.com/foxiswho/blog-go/app/web/blog/service"
	"github.com/foxiswho/blog-go/middleware/authPg"
	"github.com/foxiswho/blog-go/pkg/templatePg"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ArticleController struct {
	Sp *authPg.GroupWebMiddlewareSp `autowire:""`
	sv *service.ArticleService      `autowire:"?"`
}

func (c *ArticleController) Detail(ctx *gin.Context) {
	param := ctx.Param("id")
	hMap := gin.H{
		"title":  "详情",
		"ctxPg":  templatePg.NewHttpPg(ctx),
		"dataIs": false,
	}
	rt := c.sv.Detail(ctx, param)
	if rt.SuccessIs() {
		hMap["dataIs"] = true
		hMap["info"] = rt.Data
	}
	//syslog.Infof("Data=%+v", rt.Data)
	ctx.HTML(http.StatusOK, "blog/blog/get.tpl", hMap)
}

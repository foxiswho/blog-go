package controller

import (
	"context"
	"net/http"
	"strings"

	"github.com/foxiswho/blog-go/app/manage/domainTc/model/cacheTc"
	"github.com/foxiswho/blog-go/app/web/blog/model/modBlogArticle"
	"github.com/foxiswho/blog-go/app/web/blog/service"
	"github.com/foxiswho/blog-go/middleware/authPg"
	"github.com/foxiswho/blog-go/pkg/templatePg"
	"github.com/gin-gonic/gin"
	syslog "github.com/go-spring/log"
	"github.com/pangu-2/go-tools/tools/strPg"
)

type IndexController struct {
	Sp *authPg.GroupWebMiddlewareSp `autowire:""`
	ca *cacheTc.TenantDomainCache   `autowire:"?"`
	sv *service.ArticleService      `autowire:"?"`
}

func (c *IndexController) Index(ctx *gin.Context) {
	var ct modBlogArticle.QueryCt
	ctx.Bind(&ct)
	//
	hMap := gin.H{
		"title":  "首页",
		"ctxPg":  templatePg.NewHttpPg(ctx),
		"dataIs": false,
	}
	//
	rt := c.sv.Query(ctx, ct)
	if rt.SuccessIs() {
		hMap["dataIs"] = true
		hMap["data"] = rt.Data
	}
	syslog.Infof(context.Background(), syslog.TagBizDef, "Data=%+v", rt.Data)
	ctx.HTML(http.StatusOK, "blog/blog/index.tpl", hMap)
}

func (c *IndexController) Page(ctx *gin.Context) {

	var ct modBlogArticle.QueryCt
	ctx.Bind(&ct)
	//
	hMap := gin.H{
		"title":  "首页",
		"ctxPg":  templatePg.NewHttpPg(ctx),
		"dataIs": false,
	}
	//
	rt := c.sv.Query(ctx, ct)
	if rt.SuccessIs() {
		hMap["dataIs"] = true
		hMap["data"] = rt.Data
	}
	syslog.Infof(context.Background(), syslog.TagBizDef, "Data=%+v", rt.Data)
	ctx.HTML(http.StatusOK, "blog/blog/index.tpl", hMap)
}

func (c *IndexController) Tag(ctx *gin.Context) {
	hMap := gin.H{
		"title":  "标签",
		"ctxPg":  templatePg.NewHttpPg(ctx),
		"dataIs": false,
	}
	param := ctx.Param("tag")
	param = strings.TrimSpace(param)
	if strPg.IsNotBlank(param) {
		var ct modBlogArticle.QueryCt
		ct.TagsQuery = make([]string, 0)
		ct.TagsQuery = append(ct.TagsQuery, param)
		//
		//
		rt := c.sv.Query(ctx, ct)
		if rt.SuccessIs() {
			hMap["dataIs"] = true
			hMap["data"] = rt.Data
		}
		syslog.Infof(context.Background(), syslog.TagBizDef, "Data=%+v", rt.Data)
	}
	ctx.HTML(http.StatusOK, "blog/blog/index.tpl", hMap)
}

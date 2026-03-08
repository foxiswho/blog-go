package controller

import (
	"context"

	"github.com/foxiswho/blog-go/app/core/blog/serviceCore"
	"github.com/foxiswho/blog-go/app/web/blog/model/modBlogArticle"
	"github.com/foxiswho/blog-go/app/web/blog/service"
	"github.com/foxiswho/blog-go/app/web/utils/webPg"
	"github.com/foxiswho/blog-go/middleware/authPg"
	"github.com/foxiswho/blog-go/pkg/templatePg"
	"github.com/gin-gonic/gin"
	syslog "github.com/go-spring/log"
)

type ArticleController struct {
	Sp       *authPg.GroupWebMiddlewareSp     `autowire:"?"`
	sv       *service.ArticleService          `autowire:"?"`
	catCache *serviceCore.CoreArticleCategory `autowire:"?"`
}

func (c *ArticleController) Detail(ctx *gin.Context) {
	dataIs := false
	var data any
	param := ctx.Param("id")
	rt := c.sv.Detail(ctx, param)
	if rt.SuccessIs() {
		dataIs = true
		data = rt.Data
	}
	//
	tenantNo := webPg.GetTenantNo(ctx)
	tree, _ := c.catCache.FormatTree(ctx, tenantNo)
	//fmt.Printf("Data: %+v\n", data)
	// 模版
	templatePg.HTML(ctx, "blog/detail",
		templatePg.WithDataByResult(dataIs, data),
		templatePg.WithHtmlObjSet("categorys", tree),
		templatePg.WithHtmlObjSet("pageUrl", "article"),
		templatePg.WithSitePage(templatePg.SitePage{
			Title:       "详情",
			Description: "博客",
			Keywords:    "博客",
			SiteName:    "博客",
		}))
}

func (c *ArticleController) List(ctx *gin.Context) {
	var ct modBlogArticle.QueryCt
	ctx.Bind(&ct)
	//
	rt := c.sv.Query(ctx, ct)
	syslog.Infof(context.Background(), syslog.TagBizDef, "Data=%+v", rt.Data.Pageable)
	//
	tenantNo := webPg.GetTenantNo(ctx)
	tree, _ := c.catCache.FormatTree(ctx, tenantNo)

	//fmt.Printf("Data: %+v\n", tree)
	// 模版
	templatePg.HTML(ctx, "blog/article_list",
		templatePg.WithDataByResult(rt.SuccessIs(), rt.Data),
		templatePg.WithHtmlObjSet("categorys", tree),
		templatePg.WithHtmlObjSet("pageUrl", "/article/search"),
		templatePg.WithSitePage(templatePg.SitePage{
			Title:       "博客",
			Description: "博客",
			Keywords:    "博客",
			SiteName:    "博客",
		}))
}

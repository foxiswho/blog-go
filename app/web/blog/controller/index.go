package controller

import (
	"context"

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
	rt := c.sv.Query(ctx, ct)
	syslog.Infof(context.Background(), syslog.TagBizDef, "Data=%+v", rt.Data.Pageable)
	//fmt.Printf("Data: %+v\n", rt.Data.Pageable)
	// 模版
	templatePg.HTML(ctx, "blog/index",
		templatePg.WithDataByResult(rt.SuccessIs(), rt.Data),
		templatePg.WithSitePage(templatePg.SitePage{
			Title:       "博客",
			Description: "博客",
			Keywords:    "博客",
			SiteName:    "博客",
		}))
}

func (c *IndexController) Page(ctx *gin.Context) {
	param := ctx.Param("page")
	var ct modBlogArticle.QueryCt
	ct.PageSize = 20
	ct.PageNum = strPg.ToInt64(param)
	if ct.PageNum < 1 {
		ct.PageNum = 1
	}
	//
	rt := c.sv.Query(ctx, ct)
	// 模版
	templatePg.HTML(ctx, "blog/index",
		templatePg.WithDataByResult(rt.SuccessIs(), rt.Data),
		templatePg.WithSitePage(templatePg.SitePage{
			Title:       "博客",
			Description: "博客",
			Keywords:    "博客",
			SiteName:    "博客",
		}))
}

package controller

import (
	"strings"

	"github.com/foxiswho/blog-go/app/manage/domainTc/model/cacheTc"
	"github.com/foxiswho/blog-go/app/web/blog/model/modBlogArticle"
	"github.com/foxiswho/blog-go/app/web/blog/service"
	"github.com/foxiswho/blog-go/middleware/authPg"
	"github.com/foxiswho/blog-go/pkg/templatePg"
	"github.com/gin-gonic/gin"
	"github.com/pangu-2/go-tools/tools/strPg"
)

type TagController struct {
	Sp *authPg.GroupWebMiddlewareSp `autowire:""`
	ca *cacheTc.TenantDomainCache   `autowire:"?"`
	sv *service.ArticleService      `autowire:"?"`
}

func (c *TagController) List(ctx *gin.Context) {
	dataIs := false
	var data any
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
			dataIs = true
			data = rt.Data
		}
	}
	// 模版
	templatePg.HTML(ctx, "blog/tag",
		templatePg.WithDataByResult(dataIs, data),
		templatePg.WithSitePage(templatePg.SitePage{
			Title:       "标签",
			Description: "博客",
			Keywords:    "博客",
			SiteName:    "博客",
		}))
}

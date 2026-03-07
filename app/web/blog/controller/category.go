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

type CategoryController struct {
	Sp *authPg.GroupWebMiddlewareSp `autowire:""`
	ca *cacheTc.TenantDomainCache   `autowire:"?"`
	sv *service.ArticleService      `autowire:"?"`
}

func (c *CategoryController) List(ctx *gin.Context) {
	dataIs := false
	var data any
	param := ctx.Param("cat")
	param = strings.TrimSpace(param)
	if strPg.IsNotBlank(param) {
		var ct modBlogArticle.QueryCt
		ct.CategoryQuery = make([]string, 0)
		ct.CategoryQuery = append(ct.CategoryQuery, param)
		//
		//
		rt := c.sv.Query(ctx, ct)
		if rt.SuccessIs() {
			dataIs = true
			data = rt.Data
		}
	}
	// 模版
	templatePg.HTML(ctx, "blog/category_list",
		templatePg.WithDataByResult(dataIs, data),
		templatePg.WithSitePage(templatePg.SitePage{
			Title:       "标签",
			Description: "博客",
			Keywords:    "博客",
			SiteName:    "博客",
		}))
}

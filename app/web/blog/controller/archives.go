package controller

import (
	"strings"
	"time"

	core "github.com/foxiswho/blog-go/app/core/blog/service"
	"github.com/foxiswho/blog-go/app/manage/domainTc/model/cacheTc"
	"github.com/foxiswho/blog-go/app/web/blog/model/modBlogArticle"
	"github.com/foxiswho/blog-go/app/web/blog/service"
	"github.com/foxiswho/blog-go/app/web/utils/webPg"
	"github.com/foxiswho/blog-go/middleware/authPg"
	"github.com/foxiswho/blog-go/pkg/templatePg"
	"github.com/foxiswho/blog-go/pkg/tools/timePg"
	"github.com/foxiswho/blog-go/pkg/tools/typePg"
	"github.com/gin-gonic/gin"
	"github.com/pangu-2/go-tools/tools/strPg"
)

// ArchivesController 归档
type ArchivesController struct {
	Sp       *authPg.GroupWebMiddlewareSp `autowire:""`
	ca       *cacheTc.TenantDomainCache   `autowire:"?"`
	sv       *service.ArticleService      `autowire:"?"`
	catCache *core.CoreArticleCategory    `autowire:"?"`
}

func (c *ArchivesController) List(ctx *gin.Context) {
	dataIs := false
	var data any
	param := ctx.Param("code")
	param = strings.TrimSpace(param)
	if strPg.IsNotBlank(param) {
		flexible, b := timePg.IsValidYearMonthFlexible(param)
		if b {
			firstDay := time.Date(flexible.Year(), flexible.Month(), 1, 0, 0, 0, 0, flexible.Location())
			monthEnd := firstDay.AddDate(0, 1, 0).Add(-1 * time.Second)
			var ct modBlogArticle.QueryCt
			ct.CreateAtStart = new(typePg.Time(firstDay))
			ct.CreateAtEnd = new(typePg.Time(monthEnd))
			//
			//
			rt := c.sv.Query(ctx, ct)
			if rt.SuccessIs() {
				dataIs = true
				data = rt.Data
			}
		}
	}
	//
	tenantNo := webPg.GetTenantNo(ctx)
	tree, _ := c.catCache.FormatTree(ctx, tenantNo)
	// 模版
	templatePg.HTML(ctx, "blog/archive",
		templatePg.WithDataByResult(dataIs, data),
		templatePg.WithHtmlObjSet("categorys", tree),
		templatePg.WithHtmlObjSet("subTitle", param),
		templatePg.WithSitePage(templatePg.SitePage{
			Title:       "归档",
			Description: "博客",
			Keywords:    "博客",
			SiteName:    "博客",
		}))
}

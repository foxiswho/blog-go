package controller

import (
	"strings"

	"github.com/foxiswho/blog-go/app/core/blog/serviceCore"
	"github.com/foxiswho/blog-go/app/manage/domainTc/model/cacheTc"
	"github.com/foxiswho/blog-go/app/web/blog/model/modBlogArticle"
	"github.com/foxiswho/blog-go/app/web/blog/service"
	"github.com/foxiswho/blog-go/app/web/utils/webPg"
	"github.com/foxiswho/blog-go/middleware/authPg"
	"github.com/foxiswho/blog-go/pkg/routerPg"
	"github.com/foxiswho/blog-go/pkg/templatePg"
	"github.com/gin-gonic/gin"
	"github.com/go-spring/spring-core/gs"
	"github.com/pangu-2/go-tools/tools/strPg"
)

func init() {
	gs.Provide(new(TagController)).Export(gs.As[routerPg.RouteRegistrar]())
}

type TagController struct {
	routerPg.RouteRegistrar
	Sp       *authPg.GroupWebMiddlewareSp     `autowire:""`
	ca       *cacheTc.TenantDomainCache       `autowire:"?"`
	sv       *service.ArticleService          `autowire:"?"`
	catCache *serviceCore.CoreArticleCategory `autowire:"?"`
}

// RegisterRoutes
//
//	@Description: 注册路由
//	@receiver c
//	@param e
func (c *TagController) RegisterRoutes(e *gin.Engine) {
	group := e.Group("", authPg.GroupWebMiddleware(c.Sp))
	group.GET("/tag/:tag", c.List)
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
	//
	tenantNo := webPg.GetTenantNo(ctx)
	tree, _ := c.catCache.FormatTree(ctx, tenantNo)
	// 模版
	templatePg.HTML(ctx, "blog/article_list",
		templatePg.WithDataByResult(dataIs, data),
		templatePg.WithHtmlObjSet("categorys", tree),
		templatePg.WithHtmlObjSet("pageUrl", "tag"),
		templatePg.WithSitePage(templatePg.SitePage{
			Title:       "标签",
			Description: "博客",
			Keywords:    "博客",
			SiteName:    "博客",
		}))
}

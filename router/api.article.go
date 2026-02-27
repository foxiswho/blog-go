package router

import (
	"github.com/foxiswho/blog-go/app/web/api/controller"
	"github.com/foxiswho/blog-go/middleware/authPg"
	"github.com/foxiswho/blog-go/middleware/serverPg/ginServer"
	"github.com/go-spring/spring-core/gs"
)

func init() {
	//// 文章
	//gs.Root(gs.Object(new(controller.ArticleController)).Init(func(c *controller.ArticleController) {
	//	r := ginServer.GinServerDefault
	//	group := r.Group("/api/article", authPg.GroupApiMiddleware(c.Sp))
	//	group.POST("/push", c.Push)
	//}))
	////文章分类
	//gs.Root(gs.Object(new(controller.ArticleCategoryController)).Init(func(c *controller.ArticleCategoryController) {
	//	r := ginServer.GinServerDefault
	//	group := r.Group("/api/article-category", authPg.GroupApiMiddleware(c.Sp))
	//	group.POST("/selectNodeAllPublic", c.SelectNodeAllPublic)
	//}))

	// 收集
	gs.Root(gs.Object(new(controller.CollectController)).Init(func(c *controller.CollectController) {
		r := ginServer.GinServerDefault
		group := r.Group("/api/collect", authPg.GroupApiMiddleware(c.Sp))
		group.POST("/push", c.Push)
	}))
}

package router

import (
	"github.com/foxiswho/blog-go/app/web/blog/controller"
	"github.com/foxiswho/blog-go/middleware/authPg"
	"github.com/foxiswho/blog-go/middleware/serverPg/ginServer"
	"github.com/go-spring/spring-core/gs"
)

func init() {
	// 首页
	gs.Root(gs.Object(new(controller.IndexController)).Init(func(c *controller.IndexController) {
		r := ginServer.GinServerDefault
		group := r.Group("", authPg.GroupWebMiddleware(c.Sp))
		group.GET("/", c.Index)
		group.GET("/page/:page", c.Page)
	}))
	//文章
	gs.Root(gs.Object(new(controller.ArticleController)).Init(func(c *controller.ArticleController) {
		r := ginServer.GinServerDefault
		group := r.Group("", authPg.GroupWebMiddleware(c.Sp))
		group.GET("/article/:id", c.Detail)
	}))
	//文章分类
	gs.Root(gs.Object(new(controller.CategoryController)).Init(func(c *controller.CategoryController) {
		r := ginServer.GinServerDefault
		group := r.Group("", authPg.GroupWebMiddleware(c.Sp))
		group.GET("/category/:cat", c.List)
	}))

	// 标签
	gs.Root(gs.Object(new(controller.TagController)).Init(func(c *controller.TagController) {
		r := ginServer.GinServerDefault
		group := r.Group("", authPg.GroupWebMiddleware(c.Sp))
		group.GET("/tag/:tag", c.List)
	}))
	// 归档
	gs.Root(gs.Object(new(controller.ArchivesController)).Init(func(c *controller.ArchivesController) {
		r := ginServer.GinServerDefault
		group := r.Group("", authPg.GroupWebMiddleware(c.Sp))
		group.GET("/archives/date/:code", c.List)
	}))
}

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
		group.GET("/page", c.Page)
		group.GET("/tag/:tag", c.Page)
	}))
	//文章
	gs.Root(gs.Object(new(controller.ArticleController)).Init(func(c *controller.ArticleController) {
		r := ginServer.GinServerDefault
		group := r.Group("", authPg.GroupWebMiddleware(c.Sp))
		group.GET("/article/:id", c.Detail)
	}))
}

package router

import (
	"github.com/foxiswho/blog-go/app/manage/domainBlog/controller"
	"github.com/foxiswho/blog-go/middleware/authPg"
	"github.com/foxiswho/blog-go/middleware/serverPg/ginServer"
	"github.com/go-spring/spring-core/gs"
)

func init() {
	//
	gs.Root(gs.Object(new(controller.ArticleController)).Init(func(c *controller.ArticleController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/blog/article", authPg.GroupManageMiddleware(c.Sp))
		group.POST("/create", c.Create)
		group.POST("/update", c.Update)
		group.GET("/detail/:id", c.Detail)
		group.POST("/enable", c.Enable)
		group.POST("/disable", c.Disable)
		group.POST("/state", c.State)
		group.POST("/delete", c.Delete)
		group.POST("/recovery", c.Recovery)
		group.POST("/physicalDeletion", c.PhysicalDeletion)
		group.POST("/query", c.Query)
		group.POST("/selectPublic", c.SelectPublic)
		group.POST("/selectNodePublic", c.SelectNodePublic)
		group.POST("/selectNodeAllPublic", c.SelectNodeAllPublic)
		group.POST("/existName", c.ExistName)
		group.POST("/existNo", c.ExistNo)
	}))
	//
	gs.Root(gs.Object(new(controller.ArticleCategoryController)).Init(func(c *controller.ArticleCategoryController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/blog/article-category", authPg.GroupManageMiddleware(c.Sp))
		group.POST("/create", c.Create)
		group.POST("/update", c.Update)
		group.GET("/detail/:id", c.Detail)
		group.POST("/enable", c.Enable)
		group.POST("/disable", c.Disable)
		group.POST("/state", c.State)
		group.POST("/delete", c.Delete)
		group.POST("/recovery", c.Recovery)
		group.POST("/physicalDeletion", c.PhysicalDeletion)
		group.POST("/query", c.Query)
		group.POST("/selectPublic", c.SelectPublic)
		group.POST("/selectNodePublic", c.SelectNodePublic)
		group.POST("/selectNodeAllPublic", c.SelectNodeAllPublic)
		group.POST("/existName", c.ExistName)
		group.POST("/existNo", c.ExistNo)
	}))
	gs.Root(gs.Object(new(controller.CollectController)).Init(func(c *controller.CollectController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/blog/collect", authPg.GroupManageMiddleware(c.Sp))
		group.POST("/create", c.Create)
		group.POST("/update", c.Update)
		group.GET("/detail/:id", c.Detail)
		group.POST("/enable", c.Enable)
		group.POST("/disable", c.Disable)
		group.POST("/state", c.State)
		group.POST("/delete", c.Delete)
		group.POST("/recovery", c.Recovery)
		group.POST("/physicalDeletion", c.PhysicalDeletion)
		group.POST("/query", c.Query)
		group.POST("/selectPublic", c.SelectPublic)
		group.POST("/selectNodePublic", c.SelectNodePublic)
		group.POST("/selectNodeAllPublic", c.SelectNodeAllPublic)
		group.POST("/existName", c.ExistName)
		group.POST("/existNo", c.ExistNo)
	}))
	//
	gs.Root(gs.Object(new(controller.CollectCategoryController)).Init(func(c *controller.CollectCategoryController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/blog/collect-category", authPg.GroupManageMiddleware(c.Sp))
		group.POST("/create", c.Create)
		group.POST("/update", c.Update)
		group.GET("/detail/:id", c.Detail)
		group.POST("/enable", c.Enable)
		group.POST("/disable", c.Disable)
		group.POST("/state", c.State)
		group.POST("/delete", c.Delete)
		group.POST("/recovery", c.Recovery)
		group.POST("/physicalDeletion", c.PhysicalDeletion)
		group.POST("/query", c.Query)
		group.POST("/selectPublic", c.SelectPublic)
		group.POST("/selectNodePublic", c.SelectNodePublic)
		group.POST("/selectNodeAllPublic", c.SelectNodeAllPublic)
		group.POST("/existName", c.ExistName)
		group.POST("/existNo", c.ExistNo)
	}))
	//
	gs.Root(gs.Object(new(controller.TopicController)).Init(func(c *controller.TopicController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/blog/topic", authPg.GroupManageMiddleware(c.Sp))
		group.POST("/create", c.Create)
		group.POST("/update", c.Update)
		group.GET("/detail/:id", c.Detail)
		group.POST("/enable", c.Enable)
		group.POST("/disable", c.Disable)
		group.POST("/state", c.State)
		group.POST("/delete", c.Delete)
		group.POST("/recovery", c.Recovery)
		group.POST("/physicalDeletion", c.PhysicalDeletion)
		group.POST("/query", c.Query)
		group.POST("/selectPublic", c.SelectPublic)
		group.POST("/selectNodePublic", c.SelectNodePublic)
		group.POST("/selectNodeAllPublic", c.SelectNodeAllPublic)
		group.POST("/existName", c.ExistName)
		group.POST("/existNo", c.ExistNo)
	}))
	//
	gs.Root(gs.Object(new(controller.TopicCategoryController)).Init(func(c *controller.TopicCategoryController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/blog/topic-category", authPg.GroupManageMiddleware(c.Sp))
		group.POST("/create", c.Create)
		group.POST("/update", c.Update)
		group.GET("/detail/:id", c.Detail)
		group.POST("/enable", c.Enable)
		group.POST("/disable", c.Disable)
		group.POST("/state", c.State)
		group.POST("/delete", c.Delete)
		group.POST("/recovery", c.Recovery)
		group.POST("/physicalDeletion", c.PhysicalDeletion)
		group.POST("/query", c.Query)
		group.POST("/selectPublic", c.SelectPublic)
		group.POST("/selectNodePublic", c.SelectNodePublic)
		group.POST("/selectNodeAllPublic", c.SelectNodeAllPublic)
		group.POST("/existName", c.ExistName)
		group.POST("/existNo", c.ExistNo)
	}))
	//
	gs.Root(gs.Object(new(controller.TopicRelationController)).Init(func(c *controller.TopicRelationController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/blog/topic-relation", authPg.GroupManageMiddleware(c.Sp))
		group.POST("/addByTopic", c.AddByTopic)
		group.POST("/physicalDeletion", c.PhysicalDeletion)
		group.POST("/delete", c.PhysicalDeletion)
		group.POST("/query", c.Query)
	}))
}

package router

import (
	"github.com/foxiswho/blog-go/app/manage/domainApi/controller"
	"github.com/foxiswho/blog-go/middleware/authPg"
	"github.com/foxiswho/blog-go/middleware/serverPg/ginServer"
	"github.com/go-spring/spring-core/gs"
)

func init() {
	//
	gs.Root(gs.Object(new(controller.DiplController)).Init(func(c *controller.DiplController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/api/dipl", authPg.GroupManageMiddleware(c.Sp))
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
		group.POST("/existCode", c.ExistCode)
	}))
	//
	gs.Root(gs.Object(new(controller.DiplAccessKeyController)).Init(func(c *controller.DiplAccessKeyController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/api/dipl-access-key", authPg.GroupManageMiddleware(c.Sp))
		group.POST("/enable", c.Enable)
		group.POST("/disable", c.Disable)
		group.POST("/state", c.State)
		group.POST("/delete", c.Delete)
		group.POST("/recovery", c.Recovery)
		group.POST("/physicalDeletion", c.PhysicalDeletion)
		group.POST("/query", c.Query)
		group.POST("/selectPublic", c.SelectPublic)
		group.POST("/makeNew", c.MakeNewRecord)
	}))
	//
	gs.Root(gs.Object(new(controller.DiplCategoryController)).Init(func(c *controller.DiplCategoryController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/api/dipl-category", authPg.GroupManageMiddleware(c.Sp))
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
	}))
}

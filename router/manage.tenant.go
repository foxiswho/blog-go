package router

import (
	"github.com/foxiswho/blog-go/app/manage/domainTc/controller"
	"github.com/foxiswho/blog-go/middleware/authPg"
	"github.com/foxiswho/blog-go/middleware/serverPg/ginServer"
	"github.com/go-spring/spring-core/gs"
)

func init() {
	//
	gs.Root(gs.Object(new(controller.LevelController)).Init(func(c *controller.LevelController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/tc/level", authPg.GroupManageMiddleware(c.Sp))
		//group.POST("/create", c.Create)
		//group.POST("/update", c.Update)
		group.GET("/detail/:id", c.Detail)
		//group.POST("/enable", c.Enable)
		//group.POST("/disable", c.Disable)
		//group.POST("/state", c.State)
		//group.POST("/delete", c.Delete)
		//group.POST("/recovery", c.Recovery)
		//group.POST("/physicalDeletion", c.PhysicalDeletion)
		group.POST("/query", c.Query)
		group.POST("/selectPublic", c.SelectPublic)
		group.POST("/selectNodePublic", c.SelectNodePublic)
		group.POST("/selectNodeAllPublic", c.SelectNodeAllPublic)
		group.POST("/existName", c.ExistName)
		group.POST("/existNo", c.ExistNo)
	}))
	//
	gs.Root(gs.Object(new(controller.TenantController)).Init(func(c *controller.TenantController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/tc/tenant", authPg.GroupManageMiddleware(c.Sp))
		//group.POST("/create", c.Create)
		//group.POST("/update", c.Update)
		group.GET("/detail/:id", c.Detail)
		//group.POST("/enable", c.Enable)
		//group.POST("/disable", c.Disable)
		//group.POST("/state", c.State)
		//group.POST("/delete", c.Delete)
		//group.POST("/recovery", c.Recovery)
		//group.POST("/physicalDeletion", c.PhysicalDeletion)
		group.POST("/query", c.Query)
		group.POST("/selectPublic", c.SelectPublic)
		group.POST("/selectNodePublic", c.SelectNodePublic)
		group.POST("/selectNodeAllPublic", c.SelectNodeAllPublic)
		group.POST("/existName", c.ExistName)
		group.POST("/existCode", c.ExistCode)
		group.POST("/existNo", c.ExistNo)
	}))
	//
	gs.Root(gs.Object(new(controller.TenantDomainController)).Init(func(c *controller.TenantDomainController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/tc/tenant-domain", authPg.GroupManageMiddleware(c.Sp))
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
		group.POST("/existNo", c.ExistNo)
		group.POST("/setDefaulted", c.SetDefaulted)
	}))
}

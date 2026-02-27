package router

import (
	"github.com/foxiswho/blog-go/app/manage/domainRam/controller"
	"github.com/foxiswho/blog-go/middleware/authPg"
	"github.com/foxiswho/blog-go/middleware/serverPg/ginServer"
	"github.com/foxiswho/blog-go/pkg/enum/enumCommonPg/appModulePg"
	"github.com/go-spring/spring-core/gs"
)

func init() {
	//  账号
	gs.Root(gs.Object(new(controller.AccountController).SetAppModule(appModulePg.Manage)).Init(func(c *controller.AccountController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/ram/account", authPg.GroupManageMiddleware(c.Sp))
		group.POST("/enable", c.Enable)
		group.POST("/disable", c.Disable)
		group.POST("/state", c.State)
		group.GET("/detail/:id", c.Detail)
		group.POST("/delete", c.Delete)
		group.POST("/recovery", c.Recovery)
		group.POST("/physicalDeletion", c.PhysicalDeletion)
		group.POST("/query", c.Query)
		group.POST("/updatePassword", c.UpdatePassword)
		group.POST("/create", c.Create)
		group.POST("/createAccount", c.CreateAccount)
		group.POST("/update", c.Update)
		group.POST("/updateAccount", c.Update)
		group.POST("/existAccount", c.ExistAccount)
		group.POST("/existPhone", c.ExistPhone)
		group.POST("/existMail", c.ExistMail)
		group.POST("/existIdentityCode", c.ExistIdentityCode)
		group.POST("/existRealName", c.ExistRealName)
	}))
	//  账号设备
	gs.Root(gs.Object(new(controller.AccountDeviceController)).Init(func(c *controller.AccountDeviceController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/ram/account-device", authPg.GroupManageMiddleware(c.Sp))
		group.POST("/physicalDeletion", c.PhysicalDeletion)
		group.POST("/query", c.Query)
	}))
	//  账号登录日志
	gs.Root(gs.Object(new(controller.AccountLoginLogController)).Init(func(c *controller.AccountLoginLogController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/ram/account-login-log", authPg.GroupManageMiddleware(c.Sp))
		group.POST("/physicalDeletion", c.PhysicalDeletion)
		group.POST("/query", c.Query)
	}))
	//   账号会话
	gs.Root(gs.Object(new(controller.AccountSessionController)).Init(func(c *controller.AccountSessionController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/ram/account-session", authPg.GroupManageMiddleware(c.Sp))
		group.POST("/physicalDeletion", c.PhysicalDeletion)
		group.POST("/query", c.Query)
	}))
	//   应用
	gs.Root(gs.Object(new(controller.AppController)).Init(func(c *controller.AppController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/ram/app", authPg.GroupManageMiddleware(c.Sp))
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
	gs.Root(gs.Object(new(controller.AppAccessKeyController)).Init(func(c *controller.AppAccessKeyController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/ram/app-access-key", authPg.GroupManageMiddleware(c.Sp))
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
	gs.Root(gs.Object(new(controller.AppCategoryController)).Init(func(c *controller.AppCategoryController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/ram/app-category", authPg.GroupManageMiddleware(c.Sp))
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
	//
	gs.Root(gs.Object(new(controller.DepartmentController)).Init(func(c *controller.DepartmentController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/ram/department", authPg.GroupManageMiddleware(c.Sp))
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
		group.POST("/exportExcel", c.ExportExcel)
		group.POST("/existName", c.ExistName)
	}))

	gs.Root(gs.Object(new(controller.GroupController)).Init(func(c *controller.GroupController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/ram/group", authPg.GroupManageMiddleware(c.Sp))
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

	gs.Root(gs.Object(new(controller.LevelController)).Init(func(c *controller.LevelController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/ram/level", authPg.GroupManageMiddleware(c.Sp))
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

	gs.Root(gs.Object(new(controller.PositionController)).Init(func(c *controller.PositionController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/ram/position", authPg.GroupManageMiddleware(c.Sp))
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

	gs.Root(gs.Object(new(controller.PostController)).Init(func(c *controller.PostController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/ram/post", authPg.GroupManageMiddleware(c.Sp))
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

	gs.Root(gs.Object(new(controller.RoleController)).Init(func(c *controller.RoleController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/ram/role", authPg.GroupManageMiddleware(c.Sp))
		group.POST("/create", c.Create)
		group.POST("/update", c.Update)
		group.GET("/detail/:id", c.Detail)
		group.POST("/enable", c.Enable)
		group.POST("/disable", c.Disable)
		group.POST("/delete", c.Delete)
		group.POST("/state", c.State)
		group.POST("/recovery", c.Recovery)
		group.POST("/physicalDeletion", c.PhysicalDeletion)
		group.POST("/query", c.Query)
		group.POST("/selectNodePublic", c.SelectNodePublic)
		group.POST("/selectNodeAllPublic", c.SelectNodeAllPublic)
		group.POST("/selectPublic", c.SelectPublic)
		group.POST("/existName", c.ExistName)
	}))

	gs.Root(gs.Object(new(controller.MenuController)).Init(func(c *controller.MenuController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/ram/menu", authPg.GroupManageMiddleware(c.Sp))
		group.POST("/create", c.Create)
		group.POST("/update", c.Update)
		group.GET("/detail/:id", c.Detail)
		group.POST("/enable", c.Enable)
		group.POST("/disable", c.Disable)
		group.POST("/delete", c.Delete)
		group.POST("/state", c.State)
		group.POST("/recovery", c.Recovery)
		group.POST("/physicalDeletion", c.PhysicalDeletion)
		group.POST("/query", c.Query)
		group.POST("/selectNodePublic", c.SelectNodePublic)
		group.POST("/selectNodeAllPublic", c.SelectNodeAllPublic)
		group.POST("/selectPublic", c.SelectPublic)
		group.POST("/existName", c.ExistName)
	}))

	gs.Root(gs.Object(new(controller.MenuRelationController)).Init(func(c *controller.MenuRelationController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/ram/menu-relation", authPg.GroupManageMiddleware(c.Sp))
		group.POST("/updateByMenu", c.UpdateByMenu)
		group.POST("/query", c.Query)
		group.POST("/physicalDeletion", c.PhysicalDeletion)
	}))

	gs.Root(gs.Object(new(controller.ResourceController)).Init(func(c *controller.ResourceController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/ram/resource", authPg.GroupManageMiddleware(c.Sp))
		group.POST("/create", c.Create)
		group.POST("/update", c.Update)
		group.GET("/detail/:id", c.Detail)
		group.POST("/enable", c.Enable)
		group.POST("/disable", c.Disable)
		group.POST("/delete", c.Delete)
		group.POST("/state", c.State)
		group.POST("/recovery", c.Recovery)
		group.POST("/physicalDeletion", c.PhysicalDeletion)
		group.POST("/query", c.Query)
		group.POST("/selectNodePublic", c.SelectNodePublic)
		group.POST("/selectNodeAllPublic", c.SelectNodeAllPublic)
		group.POST("/selectPublic", c.SelectPublic)
		group.POST("/selectCategoryPublic", c.SelectCategoryPublic)
		group.POST("/existName", c.ExistName)
	}))

	gs.Root(gs.Object(new(controller.ResourceGroupController)).Init(func(c *controller.ResourceGroupController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/ram/resource-group", authPg.GroupManageMiddleware(c.Sp))
		group.POST("/create", c.Create)
		group.POST("/update", c.Update)
		group.GET("/detail/:id", c.Detail)
		group.POST("/enable", c.Enable)
		group.POST("/disable", c.Disable)
		group.POST("/delete", c.Delete)
		group.POST("/state", c.State)
		group.POST("/recovery", c.Recovery)
		group.POST("/physicalDeletion", c.PhysicalDeletion)
		group.POST("/query", c.Query)
		group.POST("/selectNodePublic", c.SelectNodePublic)
		group.POST("/selectNodeAllPublic", c.SelectNodeAllPublic)
		group.POST("/selectPublic", c.SelectPublic)
		group.POST("/updateByResourceGroup", c.UpdateByResourceGroup)
		group.POST("/resourceSelected", c.Selected)
		group.POST("/existName", c.ExistName)
	}))

	gs.Root(gs.Object(new(controller.ResourceAuthorityController)).Init(func(c *controller.ResourceAuthorityController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/ram/resource-authority", authPg.GroupManageMiddleware(c.Sp))
		group.POST("/create", c.Create)
		group.POST("/createByGroup", c.CreatByGroup)
		group.POST("/update", c.Update)
		group.POST("/updateByRole", c.UpdateByRole)
		group.GET("/detail/:id", c.Detail)
		group.POST("/enable", c.Enable)
		group.POST("/disable", c.Disable)
		group.POST("/delete", c.Delete)
		group.POST("/state", c.State)
		group.POST("/recovery", c.Recovery)
		group.POST("/physicalDeletion", c.PhysicalDeletion)
		group.POST("/query", c.Query)
		group.POST("/queryByGroup", c.QueryByGroup)
		group.POST("/selectNodePublic", c.SelectNodePublic)
		group.POST("/selectNodeAllPublic", c.SelectNodeAllPublic)
		group.POST("/selectPublic", c.SelectPublic)
	}))

	gs.Root(gs.Object(new(controller.ResourceRelationController)).Init(func(c *controller.ResourceRelationController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/ram/resource-relation", authPg.GroupManageMiddleware(c.Sp))
		group.POST("/selected", c.Selected)
	}))

	gs.Root(gs.Object(new(controller.ResourceGroupRelationController)).Init(func(c *controller.ResourceGroupRelationController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/ram/resource-group-relation", authPg.GroupManageMiddleware(c.Sp))
		group.POST("/selectedByRole", c.SelectedByRole)
	}))

	gs.Root(gs.Object(new(controller.TeamController)).Init(func(c *controller.TeamController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/ram/team", authPg.GroupManageMiddleware(c.Sp))
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

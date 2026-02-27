package router

import (
	"github.com/foxiswho/blog-go/app/manage/domainBasic/controller"
	"github.com/foxiswho/blog-go/middleware/authPg"
	"github.com/foxiswho/blog-go/middleware/serverPg/ginServer"
	"github.com/go-spring/spring-core/gs"
)

func init() {
	//
	gs.Root(gs.Object(new(controller.AttachmentController)).Init(func(c *controller.AttachmentController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/basic/attachment", authPg.GroupManageMiddleware(c.Sp))
		group.POST("/upload", c.Upload)
		group.POST("/upload-more", c.UploadMore)
		group.POST("/upload-link", c.UploadLink)
		group.POST("/upload-list", c.Query)
		group.POST("/query", c.Query)
		group.POST("/makeFileOwnerPublic", c.MakeFileOwner)
		group.POST("/makeFileOwnerAllPublic", c.MakeFileOwnerAll)
		group.POST("/upload-listByOwnerPublic", c.ListByOwner)
		group.POST("/upload-delByOwnerPublic", c.DelByOwner)
		group.POST("/upload-makeFileOwnerAllPublic", c.MakeFileOwnerAll)
		group.POST("/upload-updateByFileOwner", c.UpdateByFileOwner)
	}))
	//
	gs.Root(gs.Object(new(controller.AccountApplyDenyListController)).Init(func(c *controller.AccountApplyDenyListController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/basic/accountApplyDenyList", authPg.GroupManageMiddleware(c.Sp))
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
		group.POST("/exportExcel", c.ExportExcel)
		group.POST("/existName", c.ExistName)
		group.POST("/existExpr", c.ExistExpr)
	}))
	//
	gs.Root(gs.Object(new(controller.AreaController)).Init(func(c *controller.AreaController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/basic/area", authPg.GroupManageMiddleware(c.Sp))
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
		group.POST("/exportExcel", c.ExportExcel)
		group.POST("/existName", c.ExistName)
		group.POST("/existCode", c.ExistCode)
	}))
	//
	gs.Root(gs.Object(new(controller.CountryController)).Init(func(c *controller.CountryController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/basic/country", authPg.GroupManageMiddleware(c.Sp))
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
		group.POST("/selectPublicCountryCode", c.SelectPublicCountryCode)
		group.POST("/existName", c.ExistName)
		group.POST("/existCode", c.ExistCode)
		group.POST("/existCountryCode", c.ExistCountryCode)
	}))
	//
	gs.Root(gs.Object(new(controller.DataDictionaryController)).Init(func(c *controller.DataDictionaryController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/basic/dataDictionary", authPg.GroupManageMiddleware(c.Sp))
		group.POST("/createUpdate", c.CreateUpdate)
		group.GET("/detail/:id", c.Detail)
		group.POST("/enable", c.Enable)
		group.POST("/disable", c.Disable)
		group.POST("/state", c.State)
		group.POST("/delete", c.Delete)
		group.POST("/recovery", c.Recovery)
		group.POST("/physicalDeletion", c.PhysicalDeletion)
		group.POST("/query", c.Query)
		group.POST("/selectNodeAllPublic", c.SelectNodeAllPublic)
		group.POST("/existName", c.ExistName)
		group.POST("/existCode", c.ExistCode)
	}))
	//
	gs.Root(gs.Object(new(controller.DataDictionarySubController)).Init(func(c *controller.DataDictionarySubController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/basic/dataDictionarySub", authPg.GroupManageMiddleware(c.Sp))
		group.POST("/createUpdate", c.CreateUpdate)
		group.GET("/detail/:id", c.Detail)
		group.POST("/enable", c.Enable)
		group.POST("/disable", c.Disable)
		group.POST("/state", c.State)
		group.POST("/delete", c.Delete)
		group.POST("/recovery", c.Recovery)
		group.POST("/physicalDeletion", c.PhysicalDeletion)
		group.POST("/query", c.Query)
		group.POST("/selectNodeAllPublic", c.SelectNodeAllPublic)
		group.POST("/existName", c.ExistName)
		group.POST("/existCode", c.ExistCode)
	}))
	//
	gs.Root(gs.Object(new(controller.TagsController)).Init(func(c *controller.TagsController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/basic/tags", authPg.GroupManageMiddleware(c.Sp))
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
	gs.Root(gs.Object(new(controller.TagsCategoryController)).Init(func(c *controller.TagsCategoryController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/basic/tags-category", authPg.GroupManageMiddleware(c.Sp))
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
	gs.Root(gs.Object(new(controller.TagsRelationController)).Init(func(c *controller.TagsRelationController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/basic/tags-relation", authPg.GroupManageMiddleware(c.Sp))
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
		group.POST("/all", c.All)
		group.POST("/selectPublic", c.SelectPublic)
		group.POST("/selectNodePublic", c.SelectNodePublic)
		group.POST("/selectNodeAllPublic", c.SelectNodeAllPublic)
		group.POST("/existName", c.ExistName)
		group.POST("/existCode", c.ExistCode)
		group.GET("/getCategoryRoot/:category", c.GetCategory)
		group.POST("/getCategoryTagsAll/:category", c.GetCategoryTagsAll)
		group.POST("/getCategoryTags/:category", c.GetCategoryTags)
	}))
	//
	gs.Root(gs.Object(new(controller.BasicConfigModelController)).Init(func(c *controller.BasicConfigModelController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/basic/config-model", authPg.GroupManageMiddleware(c.Sp))
		group.POST("/create", c.Create)
	}))
	//
	gs.Root(gs.Object(new(controller.BasicConfigEventController)).Init(func(c *controller.BasicConfigEventController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/basic/config-event", authPg.GroupManageMiddleware(c.Sp))
		group.POST("/create", c.Create)
	}))
}

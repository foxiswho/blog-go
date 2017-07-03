// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"blog/controllers/admin"
	"blog/controllers"
	"blog/controllers/api"
)

func init() {

	//beego.Router("/admin/login", &admin.LoginController{})
	//beego.Router("/admin/index", &admin.IndexController{})
	ns := beego.NewNamespace("/admin",
		beego.NSRouter("/login", &admin.Login{}),
		beego.NSRouter("/logout", &admin.Logout{}),
		beego.NSRouter("/my_password", &admin.MyPassword{}),
		//blog
		//使用注解路由
		//beego.NSRouter("/blog", &admin.BlogController{}, "get:List"),
		//beego.NSRouter("/blog", &admin.BlogController{}, ),
		//beego.NSRouter("/blog/detail/:id", &admin.BlogController{}, "get:Detail"),
		//beego.NSRouter("/blog/:id", &admin.BlogController{}, "put:Put"), //, "put:Put"
		//beego.NSRouter("/blog/:id", &admin.BlogController{}, "get:Get"),
		//beego.NSRouter("/blog/add", &admin.BlogController{}, "get:Add"),
		//type
		//使用注解路由
		//beego.NSRouter("/types", &admin.TypeController{},"get:List"),
		//beego.NSRouter("/types/add", &admin.TypeController{},"get:Add"),
		//beego.NSRouter("/types/add/:id", &admin.TypeController{},"get:Add"),
		//beego.NSRouter("/types/:id", &admin.TypeController{},"get:ListChild"),
		//beego.NSRouter("/type/:id", &admin.TypeController{},"get:Get"),// "get:Get"
		//beego.NSRouter("/type", &admin.TypeController{}),
		beego.NSInclude(
			&admin.Index{},
			&admin.Area{},
			&admin.Attachment{},
			&admin.Type{},
			&admin.BlogTag{},
			&admin.BlogCat{},
			&admin.Blog{},
			&admin.BlogSync{},
			&admin.Select{},
			&admin.Upload{},
			&admin.AdminUser{},
			&admin.AdminRole{},
			&admin.Member{},
			&admin.MemberGroup{},
			&admin.Site{},
			&admin.Oauth{},
		),
	)
	beego.AddNamespace(ns)
	//首页
	beego.Router("/", &controllers.Blog{}, "get:GetAll")
	beego.Router("/search/", &controllers.Blog{}, "get:GetAll")
	beego.Router("/page/:page", &controllers.Blog{}, "get:GetAll")
	beego.Router("/article/:id", &controllers.Blog{}, "get:Get")
	beego.Router("/tag/:tag", &controllers.Tag{}, "get:GetAll")
	//API
	beego.Router("/api/blog/create", &api.Blog{}, "post:Create")
	beego.Router("/api/blog/cat", &api.BlogCat{}, "get:GetAll")
	beego.Router("/api/blog/tag", &api.BlogTag{}, "get:GetAll")
}

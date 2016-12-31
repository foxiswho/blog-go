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
	"fox/controllers/admin"
	"fox/controllers"
)

func init() {

	//beego.Router("/admin/login", &admin.LoginController{})
	//beego.Router("/admin/index", &admin.IndexController{})
	ns := beego.NewNamespace("/admin",
		beego.NSRouter("/login", &admin.LoginController{}),
		beego.NSRouter("/logout", &admin.LogoutController{}),
		beego.NSRouter("/my_password", &admin.MyPasswordController{}),
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
			&admin.TypeController{},
			&admin.BlogTag{},
			&admin.BlogCat{},
			&admin.BlogController{},
			&admin.Select{},
			&admin.Upload{},
			&admin.AdminUser{},
			&admin.AdminRole{},
			&admin.Member{},
			&admin.MemberGroup{},
		),
	)
	beego.AddNamespace(ns)
	//首页
	beego.Router("/", &controllers.BlogController{}, "get:GetAll")
	beego.Router("/page/:page", &controllers.BlogController{}, "get:GetAll")
	beego.Router("/article/:id", &controllers.BlogController{}, "get:Get")
	beego.Router("/tag/:tag", &controllers.TagController{}, "get:GetAll")
}

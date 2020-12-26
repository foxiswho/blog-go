// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/foxiswho/blog-go/controllers"
	"github.com/foxiswho/blog-go/controllers/admin"
	"github.com/foxiswho/blog-go/controllers/api"
)

func init() {

	//beego.Router("/admin/login", &admin.LoginController{})
	//beego.Router("/admin/index", &admin.IndexController{})
	ns := web.NewNamespace("/admin",
		web.NSRouter("/login", &admin.Login{}),
		web.NSRouter("/logout", &admin.Logout{}),
		web.NSRouter("/my_password", &admin.MyPassword{}),
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
		web.NSInclude(
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
	web.AddNamespace(ns)
	//首页
	web.Router("/", &controllers.Blog{}, "get:GetAll")
	web.Router("/search/", &controllers.Blog{}, "get:GetAll")
	web.Router("/page/:page", &controllers.Blog{}, "get:GetAll")
	web.Router("/article/:id", &controllers.Blog{}, "get:Get")
	web.Router("/tag/:tag", &controllers.Tag{}, "get:GetAll")
	//API
	web.Router("/api/blog/create", &api.Blog{}, "post:Create")
	web.Router("/api/blog/cat", &api.BlogCat{}, "get:GetAll")
	web.Router("/api/blog/tag", &api.BlogTag{}, "get:GetAll")
}

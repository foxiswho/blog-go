package routers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/deepzz0/goblog/RS"
	"github.com/deepzz0/goblog/controllers"
	"github.com/deepzz0/goblog/controllers/background"
	"github.com/deepzz0/goblog/controllers/feed"
	"github.com/deepzz0/goblog/controllers/plugin"
	"github.com/deepzz0/goblog/controllers/proxy"
	"github.com/deepzz0/goblog/helper"
)

const (
	ONE_DAYS = 24 * 3600
)

func init() {
	// config
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "SESSIONID"
	beego.BConfig.WebConfig.Session.SessionCookieLifeTime = ONE_DAYS
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3600

	beego.Router("/", &controllers.HomeController{})
	beego.Router("/p/:page([0-9]+)", &controllers.HomeController{})
	beego.Router("/cat/:cat([a-zA-Z]+)", &controllers.CategoryController{})
	beego.Router("/cat/:cat([a-zA-Z]+)/p/:page([0-9]+)", &controllers.CategoryController{})
	beego.Router("/tag/:tag([a-zA-Z0-9\u4e00-\u9fa5]+)", &controllers.TagController{})
	beego.Router("/tag/:tag([a-zA-Z0-9\u4e00-\u9fa5]+)/p/:page([0-9]+)", &controllers.TagController{})
	beego.Router("/:year([0-9]{4})/:month([0-9]{2})/:day([0-9]{2})/:id([0-9]+).html", &controllers.TopicController{})
	beego.Router("/archives/:year([0-9]{4})/:month([0-9]{2})", &controllers.ArchivesController{})
	beego.Router("/message", &controllers.MessageController{})
	beego.Router("/about", &controllers.AboutController{})
	beego.Router("/search", &controllers.SearchController{})
	// admin
	beego.Router("/login", &background.AuthController{})
	beego.InsertFilter("/admin/*", beego.BeforeRouter, FilterUser)
	beego.Router("/admin/user", &background.UserController{})
	beego.Router("/admin/data", &background.DataController{})
	beego.Router("/admin/topics", &background.TopicsController{})
	beego.Router("/admin/category", &background.CategoryController{})
	beego.Router("/admin/message", &background.MessageController{})
	beego.Router("/admin/social", &background.SocialController{})
	beego.Router("/admin/blogroll", &background.BlogrollController{})
	beego.Router("/admin/ad", &background.ADController{})
	beego.Router("/admin/sysconfig", &background.SysconfigController{})
	beego.Router("/admin/databackup", &background.DataBackupRecover{})
	beego.Router("/admin/datarecover", &background.DataBackupRecover{})
	beego.Router("/admin/syslog", &background.SyslogController{})
	beego.Router("/admin/trash", &background.TrashController{})
	// proxy
	beego.Router("/proxy/:url(.*)", &proxy.ProxyController{})
	// rss
	beego.Get("/feed", feed.Feed)
	beego.Get("/sitemap", feed.SiteMap)
	beego.Get("/robots.txt", feed.Robots)
	// 404
	beego.ErrorHandler("404", HTTPNotFound)
	// redirect https
	beego.InsertFilter("*", beego.BeforeRouter, RedirectHttps)
	// plugin
	beego.Router("/plugin/useragent.html", &plugin.UserAgent{})
}

// 404
func HTTPNotFound(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/404.html")
	if err != nil {
		panic(err)
	}
	err = t.Execute(w, "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// 过滤登录
var FilterUser = func(ctx *context.Context) {
	val, ok := ctx.Input.Session(background.SESSIONNAME).(string)
	if !ok || val == "" {
		if ctx.Request.Method == "GET" {
			ctx.Redirect(302, "/login")
		} else if ctx.Request.Method == "POST" {
			resp := helper.NewResponse()
			resp.Status = RS.RS_user_not_login
			resp.Data = "/login"
			resp.WriteJson(ctx.ResponseWriter)
		}
	}
}

// 重定向到https
var RedirectHttps = func(ctx *context.Context) {
	if beego.BConfig.Listen.EnableHTTPS && ctx.Input.Scheme() == "http" {
		ctx.Redirect(301, fmt.Sprintf("%s%s", controllers.Domain, ctx.Input.URL()))
	}
}

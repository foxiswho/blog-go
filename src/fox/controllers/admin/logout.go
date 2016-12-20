package admin

import "github.com/astaxie/beego"

type LogoutController struct {
	beego.Controller
}

func (this *LogoutController)Get() {
	SESSION_NAME := beego.AppConfig.String("session_name")
	this.DelSession(SESSION_NAME)
	this.TplName = "admin/logout/get.html"
	this.Data["blog_name"] = beego.AppConfig.String("blog_name")
	this.Data["URL"] = "/"
}
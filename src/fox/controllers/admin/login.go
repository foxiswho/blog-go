package admin

import "github.com/astaxie/beego"

type LoginController struct {
	beego.Controller
}

func (this *LoginController)Get() {
	this.TplName = "admin/login/get.html"
	this.Data["blog_name"] = beego.AppConfig.String("blog_name")
	this.Data["URL"] = "/"
}
func (this *LoginController)Post() {

}
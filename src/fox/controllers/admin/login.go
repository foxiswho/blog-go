package admin

import (
	"github.com/astaxie/beego"
	"fox/service"
	"fox/util/Response"
)

type LoginController struct {
	BaseNoLoginController
}

func (this *LoginController)Get() {
	this.TplName = "admin/login/get.html"
	this.Data["blog_name"] = beego.AppConfig.String("blog_name")
	this.Data["URL"] = "/"
}
func (this *LoginController)Post() {
	username := this.GetString("username")
	password := this.GetString("password")
	rsp := Response.NewResponse()
	var u *service.AdminUser

	admin, err := u.Auth(username, password)
	if err != nil {
		rsp.Tips(err.Error(), "INFO")
		rsp.WriteJson(this.Ctx.ResponseWriter)
		return
	} else {
		SESSION_NAME := beego.AppConfig.String("session_name")
		this.SetSession(SESSION_NAME, admin.Username)
	}

}
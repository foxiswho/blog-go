package admin

import (
	"github.com/astaxie/beego"
	"fox/service"
	"fox/util/Response"
	"fmt"
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
	fmt.Println("username:",username)
	rsp := Response.NewResponse()
	defer rsp.WriteJson(this.Ctx.ResponseWriter)
	var adminUser *service.AdminUser
	admin, err := adminUser.Auth(username, password)
	if err != nil {
		rsp.Error(err.Error())
		return
	} else {
		SESSION_NAME := beego.AppConfig.String("session_name")
		this.SetSession(SESSION_NAME, admin.Username)
		rsp.Success("")
		return
	}

}
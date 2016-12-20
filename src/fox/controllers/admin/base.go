package admin

import (
	"github.com/astaxie/beego"
	"fox/models"
)

type BaseController struct {
	beego.Controller
	adminUser *models.Admin //当前登录用户信息
}

func (this *BaseController) Prepare() {
	//session 判断
	SESSION_NAME := beego.AppConfig.String("session_name")
	if val, ok := this.GetSession(SESSION_NAME).(string); ok && val == "" {
		this.Redirect("/admin/login", 302)
	}
}
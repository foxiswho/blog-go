package admin

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	BaseNoLoginController
}

func (this *BaseController) Prepare() {
	//session 判断
	SESSION_NAME := beego.AppConfig.String("session_name")
	if val, ok := this.GetSession(SESSION_NAME).(string); ok && val == "" {
		this.Redirect("/admin/login", 302)
	}
}
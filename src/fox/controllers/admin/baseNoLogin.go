package admin

import (
	"github.com/astaxie/beego"
	"fox/models"
)

type BaseNoLoginController struct {
	beego.Controller
	adminUser *models.Admin //当前登录用户信息
}

func (this *BaseNoLoginController) Prepare() {
	this.Data["__public__"] = "/"
	this.Data["__theme__"] = "/static/Hplus-v.4.1.0/"
}
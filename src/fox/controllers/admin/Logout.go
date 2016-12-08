package admin

import "github.com/astaxie/beego"


type LogoutController struct {
	BaseController
}
/**
退出登陆
*/
func (this *LogoutController) Loginout() {
	this.Ctx.SetCookie("token", "", 0)
	this.redirect(beego.URLFor("LoginController.Tologin"))
}
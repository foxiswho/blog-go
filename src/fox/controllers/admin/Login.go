package admin

import (
	"strconv"
	"fox/service"
	"fox/util/crypt"
)

type LoginController struct {
	BaseController
}

/**
进入登录页面
*/
func (this *LoginController) Tologin() {
	this.show("common/loginPage.html")
}
/**
登陆
*/
func (this *LoginController) Login() {
	accout := this.GetString("accout")
	password := this.GetString("password")
	encodePwd := crypt.EnCodeMD5(password)

	if admUser, err := service.adminService.Authentication(accout, encodePwd); err != nil {
		this.jsonResult(err.Error())
	} else {
		token := strconv.FormatInt(admUser.Id, 10) + "|" + accout + "|" + this.getClientIp()
		token = crypt.EnCodeMD5(token)
		this.Ctx.SetCookie("token", token, 0)
		this.jsonResult(SUCCESS)
	}
}
package admin

import (
	"fox/service"
	"fox/util/Response"
	"fmt"
)

type LoginController struct {
	BaseNoLoginController
}

func (this *LoginController)Get() {
	this.TplName = "admin/login/get.html"
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
		fmt.Println("登录成功：",admin)
		//设置Session
		this.SessionSet(admin)
		rsp.Success("")
		return
	}

}
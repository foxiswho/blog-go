package admin

import (
	"fox/util/Response"
	"fmt"
	"fox/service/admin"
)

type Login struct {
	BaseNoLogin
}

func (this *Login)Get() {
	this.TplName = "admin/login/get.html"
}
func (this *Login)Post() {
	username := this.GetString("username")
	password := this.GetString("password")
	fmt.Println("username:",username)
	rsp := Response.NewResponse()
	defer rsp.WriteJson(this.Ctx.ResponseWriter)
	adminUser :=admin.NewAdminUserService()
	adm, err := adminUser.Auth(username, password)
	if err != nil {
		rsp.Error(err.Error())
		return
	} else {
		fmt.Println("登录成功：",adm)
		//设置Session
		this.SessionSet(adm)
		rsp.Success("")
		return
	}

}
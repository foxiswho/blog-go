package admin

import (
	"fmt"
	"blog/util/Response"
	"blog/service/admin"
)

type MyPassword struct {
	Base
}

func (this *MyPassword)Get() {
	this.Data["username"] = this.Session.Username
	this.Data["true_name"] = this.Session.TrueName
	this.TplName = "admin/my/password.html"
}
func (this *MyPassword)Post() {
	password := this.GetString("password")
	fmt.Println("password:",password)
	rsp := Response.NewResponse()
	defer rsp.WriteJson(this.Ctx.ResponseWriter)
	adminUser :=admin.NewAdminUserService()
	ok, err := adminUser.UpdatePassword(password,this.Session.Aid)
	if !ok {
		rsp.Error(err.Error())
		return
	} else {
		rsp.Success("")
		return
	}

}
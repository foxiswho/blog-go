package admin

import (
	"fmt"
	"fox/util/Response"
	"fox/service/admin"
)

type MyPasswordController struct {
	BaseController
}

func (this *MyPasswordController)Get() {
	this.Data["username"] = this.Session.Username
	this.Data["true_name"] = this.Session.TrueName
	this.TplName = "admin/my/password.html"
}
func (this *MyPasswordController)Post() {
	password := this.GetString("password")
	fmt.Println("password:",password)
	rsp := Response.NewResponse()
	defer rsp.WriteJson(this.Ctx.ResponseWriter)
	var adminUser *admin.AdminUser
	ok, err := adminUser.UpdatePassword(password,this.Session.Id)
	if !ok {
		rsp.Error(err.Error())
		return
	} else {
		rsp.Success("")
		return
	}

}
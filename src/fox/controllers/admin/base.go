package admin

import (
	"fmt"
	"github.com/astaxie/beego"
	"fox/service/admin"
)

type BaseController struct {
	BaseNoLoginController
}
//  框架中的扩展函数
func (this *BaseController) Prepare() {
	//this.BaseNoLoginController.Prepare()
	this.Initialization()
	////session 判断
	ok,_ :=beego.AppConfig.Bool("admin_load")
	if ok {
		var AdminAuth admin.AdminAuth
		sess:=AdminAuth.Validate("admin")
		this.SessionSet(sess)
		this.Session=sess
		return
	}
	session, err := this.SessionGet()
	fmt.Println("session:", session)
	fmt.Println("err", err)
	if err != nil {
		this.Redirect("/admin/login", 302)
	}
	if session == nil {
		this.Redirect("/admin/login", 302)
	}
	//获取用户信息
	//var auth *service.AdminAuth
	//Session := auth.Validate(session)
	//if Session == nil {
	//	//验证不通过，删除session
	//	//adminSession.Del(this)
	//	this.Redirect("/admin/login", 302)
	//}
	this.Session=session
	//
	//session2, ok := this.GetSession(SESSION_NAME+"_JSON").(string)
	//fmt.Println("session:? => ?", session2, ok)
}
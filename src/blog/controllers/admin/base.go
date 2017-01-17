package admin

import (
	"fmt"
	"blog/service/admin"
	"blog/fox/config"
)

type Base struct {
	BaseNoLogin
}
//  框架中的扩展函数
func (c *Base) Prepare() {
	//c.BaseNoLoginController.Prepare()
	c.Initialization()
	////session 判断
	ok, _ := config.Bool("admin_load")
	fmt.Println("admin_load", ok)
	if ok {
		AdminAuth := admin.NewAdminAuthService()
		sess := AdminAuth.Validate("admin")
		c.SessionSet(sess)
		c.Session = sess
		return
	}
	session, err := c.SessionGet()
	if err != nil {
		//清空Session
		c.SessionDel()
		fmt.Println("session err", err)
		c.Redirect("/admin/login", 302)
	}
	if session == nil || session.Aid == 0 {
		//清空Session
		c.Redirect("/admin/login", 302)
	}
	fmt.Println("session 值:", session)
	//获取用户信息
	//var auth *service.AdminAuth
	//Session := auth.Validate(session)
	//if Session == nil {
	//	//验证不通过，删除session
	//	//adminSession.Del(c)
	//	c.Redirect("/admin/login", 302)
	//}
	c.Session = session
	//
	//session2, ok := c.GetSession(SESSION_NAME+"_JSON").(string)
	//fmt.Println("session:? => ?", session2, ok)
}
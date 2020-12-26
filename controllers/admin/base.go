package admin

import (
	"fmt"
	"github.com/foxiswho/blog-go/fox/config"
	"github.com/foxiswho/blog-go/service/admin"
)
//必须登录 基础控制器
type Base struct {
	BaseNoLogin
}
//  框架中的扩展函数
func (c *Base) Prepare() {
	//初始化加载
	c.Initialization()
	//开发session 判断,如果为true直接赋值 admin用户的session ，方便开发
	ok, _ := config.Bool("admin_load")
	fmt.Println("admin_load", ok)
	if ok {
		//初始化
		AdminAuth := admin.NewAdminAuthService()
		//admin session
		sess := AdminAuth.Validate("admin")
		//存入session
		c.SessionSet(sess)
		//赋值
		c.Session = sess
		return
	}
	//获取session 和验证
	session, err := c.SessionGet()
	//错误验证
	if err != nil {
		//错误
		//清空Session
		c.SessionDel()
		fmt.Println("session err", err)
		//跳转到登录
		c.Redirect("/admin/login", 302)
		return
	}
	//判断
	if session == nil || session.Aid == 0 {
		//清空Session
		c.Redirect("/admin/login", 302)
		return
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
	//赋值
	c.Session = session
	//
	//session2, ok := c.GetSession(SESSION_NAME+"_JSON").(string)
	//fmt.Println("session:? => ?", session2, ok)
}

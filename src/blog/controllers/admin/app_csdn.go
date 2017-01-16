package admin

import (
	"fmt"
	"blog/app/csdn"
)

type AppCsdn struct {
	Base
}

func (c *AppCsdn) URLMapping() {
	c.Mapping("List", c.List)
}
//列表
// @router /auth_csdn [get]
func (c *AppCsdn)List() {

	web:=csdn.NewAuthorizeWeb()
	ok,err:=web.SetConfig()
	if err !=nil{
		fmt.Println(err)
	}
	fmt.Println("status:",ok);
	web.SetRedirectUri("http://www.foxwho.com:8080/admin/auth_token")

	c.Data["url"] = web.GetAuthorizeUrl()
	c.TplName = "admin/auth/list.html"
}
// @router /auth_token [get]
func (c *AppCsdn)GetToken() {

	token:=c.GetString("code")
	web:=csdn.NewAuthorizeWeb()
	ok,err:=web.SetConfig()
	if err !=nil{
		fmt.Println(err)
	}
	fmt.Println("status:",ok);
	web.SetRedirectUri("http://www.foxwho.com:8080/admin/auth_token")
	ACCESS,err1:=web.GetAccessToken(token)
	fmt.Println(ACCESS)
	fmt.Println(err1)
	c.Data["token"] = token
	c.TplName = "admin/auth/get_token.html"
}

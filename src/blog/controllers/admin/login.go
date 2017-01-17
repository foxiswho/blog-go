package admin

import (
	"fmt"
	"blog/service/admin"
)
//登录控制器
type Login struct {
	BaseNoLogin
}
//显示
func (c *Login)Get() {
	c.TplName = "admin/login/get.html"
}
//验证
func (c *Login)Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	fmt.Println("username:",username)
	//初始化
	adminUser :=admin.NewAdminUserService()
	//验证
	adm, err := adminUser.Auth(username, password)
	//错误检测
	if err != nil {
		c.Error(err.Error())
	} else {
		fmt.Println("登录成功：",adm)
		//设置Session
		c.SessionSet(adm)
		//返回
		c.Success("操作成功")
	}
}
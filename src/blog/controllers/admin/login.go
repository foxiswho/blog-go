package admin

import (
	"fmt"
	"blog/service/admin"
)

type Login struct {
	BaseNoLogin
}

func (c *Login)Get() {
	c.TplName = "admin/login/get.html"
}
func (c *Login)Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	fmt.Println("username:",username)
	adminUser :=admin.NewAdminUserService()
	adm, err := adminUser.Auth(username, password)
	if err != nil {
		c.Error(err.Error())
		return
	} else {
		fmt.Println("登录成功：",adm)
		//设置Session
		c.SessionSet(adm)
		c.Success("操作成功")
		return
	}

}
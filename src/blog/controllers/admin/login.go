package admin

import (
	"blog/fox/response"
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
	rsp := response.NewResponse()
	defer rsp.WriteJson(c.Ctx.ResponseWriter)
	adminUser :=admin.NewAdminUserService()
	adm, err := adminUser.Auth(username, password)
	if err != nil {
		rsp.Error(err.Error())
		return
	} else {
		fmt.Println("登录成功：",adm)
		//设置Session
		c.SessionSet(adm)
		rsp.Success("")
		return
	}

}
package admin

import (
	"fmt"
	"blog/service/admin"
)
//当前管理员密码修改
type MyPassword struct {
	Base
}

func (c *MyPassword)Get() {
	c.Data["username"] = c.Session.Username
	c.Data["true_name"] = c.Session.TrueName
	c.TplName = "admin/my/password.html"
}
func (c *MyPassword)Post() {
	password := c.GetString("password")
	fmt.Println("password:",password)
	adminUser :=admin.NewAdminUserService()
	ok, err := adminUser.UpdatePassword(password,c.Session.Aid)
	if !ok {
		c.Error(err.Error())
		return
	} else {
		c.Success("操作成功")
		return
	}

}
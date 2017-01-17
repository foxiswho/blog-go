package admin
//退出控制器
type Logout struct {
	BaseNoLogin
}
//
func (c *Logout)Get() {
	//直接删除session
	c.SessionDel()
	//this.TplName = "admin/logout/get.html"
	c.Redirect("/admin/login", 302)
}
package admin

type Logout struct {
	BaseNoLogin
}

func (c *Logout)Get() {
	c.SessionDel()
	//this.TplName = "admin/logout/get.html"
	c.Redirect("/admin/login", 302)
}
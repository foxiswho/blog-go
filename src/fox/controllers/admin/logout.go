package admin

type Logout struct {
	BaseNoLogin
}

func (this *Logout)Get() {
	this.SessionDel()
	//this.TplName = "admin/logout/get.html"
	this.Redirect("/admin/login", 302)
}
package admin

type LogoutController struct {
	BaseNoLoginController
}

func (this *LogoutController)Get() {
	this.SessionDel()
	//this.TplName = "admin/logout/get.html"
	this.Redirect("/admin/login", 302)
}
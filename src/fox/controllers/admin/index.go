package admin


type IndexController struct {
	BaseController
}

func (this *IndexController)Get() {
	this.Data["username"] = this.Session.Username
	this.Data["true_name"] = this.Session.TrueName
	this.TplName = "admin/index/get.html"
}
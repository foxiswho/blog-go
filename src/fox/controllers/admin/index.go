package admin


type Index struct {
	BaseController
}
func (c *Index) URLMapping() {
	c.Mapping("Get", c.Get)
	c.Mapping("V2", c.V2)
}
// @router /index [get]
func (this *Index)Get() {
	this.Data["username"] = this.Session.Username
	this.Data["true_name"] = this.Session.TrueName
	this.TplName = "admin/index/get.html"
}
// @router /index/v2 [get]
func (this *Index)V2() {
	this.TplName = "admin/index/v2.html"
}
package admin

//管理首页控制器
type Index struct {
	Base
}
func (c *Index) URLMapping() {
	c.Mapping("Get", c.Get)
	c.Mapping("V2", c.V2)
}
//首页
// @router /index [get]
func (c *Index)Get() {
	c.Data["username"] = c.Session.Username
	c.Data["true_name"] = c.Session.TrueName
	c.TplName = "admin/index/get.html"
}
//首页第一个加载的页面
// @router /index/v2 [get]
func (c *Index)V2() {
	c.TplName = "admin/index/v2.html"
}
//默认页
// @router / [get]
func (c *Index)Default() {
	c.Redirect("admin/index",302)
}
package admin

type TypeController struct {
	BaseController
}

func (c *TypeController)List() {
	c.Data["username"] = c.Session.Username
	c.Data["true_name"] = c.Session.TrueName
	c.TplName = "admin/type/index.html"
}
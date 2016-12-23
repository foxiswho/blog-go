package admin

type TypeController struct {
	BaseController
}

func (c *TypeController)List() {
	c.TplName = "admin/type/list.html"
}
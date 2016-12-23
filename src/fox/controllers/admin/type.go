package admin

import (
	"fox/controllers"
)

type TypeController struct {
	BaseController
}

func (c *TypeController)List() {
	c.TplName = "admin/type/list.html"
	c.Data["HtmlHead"] = controllers.ExecuteTemplateHtml("admin/type/head.html",c.Data)
}
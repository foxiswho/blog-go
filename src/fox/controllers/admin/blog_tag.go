package admin

import (
	"fox/service/blog"
)

type BlogTag struct {
	BaseController
}
func (c *BlogTag) URLMapping() {
	c.Mapping("List", c.List)
}
//列表
// @router /blog/tag [get]
func (c *BlogTag)List() {
	var blogSer *blog.BlogTag
	data, err := blogSer.Query("")
	//println(data)
	println(err)
	c.Data["data"] = data
	c.Data["title"] = "博客-TAG-列表"
	c.TplName = "admin/blog/tag/list.html"
}
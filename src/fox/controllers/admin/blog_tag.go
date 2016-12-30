package admin

import (
	"fox/model"
	"fmt"
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
	query := make(map[string]interface{})
	fields := []string{}
	str := c.GetString("wd")
	if str != "" {
		query["name"] = str
	}
	mode := model.NewBlogTag()
	data, err := mode.GetAll(query, fields, "tag_id desc", 0, 20)
	if err!=nil{
		fmt.Println(err)
	}
	//println(data)
	c.Data["data"] = data
	c.Data["title"] = "博客-TAG-列表"
	c.TplName = "admin/blog/tag/list.html"
}
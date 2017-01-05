package controllers

import (
	"fox/service/blog"
	"fmt"
)

type Tag struct {
	BaseNoLogin
}

// GetAll ...
// @Title Get All
// @Description get Blog
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Blog
// @Failure 403
// @router / [get]
func (c *Tag) GetAll() {
	idStr := c.Ctx.Input.Param(":tag")
	fields := []string{}
	orderBy := "blog_id desc"
	query := make(map[string]interface{})
	query["name"] = idStr
	mode := blog.NewBlogTagService()
	//分页
	page, _ := c.GetInt("page")
	data, err := mode.GetAll(query, fields, orderBy, page, 20)
	fmt.Println("err", err)
	//fmt.Println("data", data)
	if err != nil {
		//c.Data["data"] = err.Error()
		fmt.Println(err.Error())
	} else {
		c.Data["data"] = data
	}
	c.TplName = "blog/index.html"

}

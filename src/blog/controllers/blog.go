package controllers

import (
	"strconv"

	"regexp"
	"blog/service/blog"
	"fmt"
	"blog/service"
)

type Blog struct {
	BaseNoLogin
}


// GetOne ...
// @Title Get One
// @Description get Blog by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Blog
// @Failure 403 :id is empty
// @router /article/:id [get]
func (c *Blog) Get() {
	idStr := c.Ctx.Input.Param(":id")
	ser := blog.NewBlogService()
	var err error
	var read map[string]interface{}
	if ok, _ := regexp.Match(`^\d+$`, []byte(idStr)); ok {
		id, _ := strconv.Atoi(idStr)
		read, err = ser.Read(id)
	} else {
		read, err = ser.ReadByUrlRewrite(idStr)
	}
	if err != nil {
		c.Error(err.Error())
		return
	} else {
		tmp:=read["info"]
		fmt.Println(tmp)
		B:=tmp.(*blog.Blog)
		fmt.Println(B)
		_,err:=ser.UpdateRead(B.Blog.BlogId)
		fmt.Println(err)
		c.Data["title"] =read["title"]
		c.Data["info"] = read["info"]
		c.Data["TimeAdd"] = read["TimeAdd"]
		c.Data["Content"] = read["Content"]
	}
	c.TplName = "blog/get.html"
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
func (c *Blog) GetAll() {
	fields := []string{}
	orderBy := "blog_id desc"
	query := make(map[string]interface{})
	query["type=?"] = service.TYPE_ARTICLE
	mode := blog.NewBlogService()
	//分页
	id := c.Ctx.Input.Param(":page")
	page, _ := strconv.Atoi(id)
	data, err := mode.GetAll(query, fields, orderBy, page, 10)
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

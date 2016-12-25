package admin

import (
	"fox/util/Response"
	"strconv"
	"fox/models"
	"fmt"
	"fox/service/blog"
)

type BlogCat struct {
	BaseController
}
func (c *BlogCat) URLMapping() {
	c.Mapping("List", c.List)
	c.Mapping("Add", c.Add)
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
	c.Mapping("Put", c.Put)
}
//列表
// @router /blog/cat [get]
func (c *BlogCat)List() {
	var blogSer *blog.Blog
	data, err := blogSer.Query(10001)
	//println(data)
	println(err)
	c.Data["data"] = data
	c.Data["title"] = "博客-分类-列表"
	c.TplName = "admin/blog/cat/list.html"
}
//编辑
// @router /blog/cat/:id [get]
func (c *BlogCat)Get() {
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	var blogSer *blog.Blog
	data, err := blogSer.Read(int_id)
	//println("Detail :", err.Error())
	if err != nil {
		rsp := Response.NewResponse()
		defer rsp.WriteJson(c.Ctx.ResponseWriter)
		rsp.Error(err.Error())
	} else {
		c.Data["info"] = data["Blog"]
		c.Data["statistics"] = data["Statistics"]
		c.Data["TimeAdd"] = data["TimeAdd"]
		c.Data["title"] = "博客-分类-编辑"
		c.Data["_method"] = "put"
		c.Data["is_put"] = true
		c.Data["cat_id"] = blog.CAT_ID
		c.TplName = "admin/blog/cat/get.html"
	}
}
//添加
// @router /blog/cat/add [get]
func (c *BlogCat)Add() {
	c.Data["cat_id"] = blog.CAT_ID
	c.Data["_method"] = "post"
	c.Data["title"] = "博客-分类-添加"
	c.TplName = "admin/blog/cat/get.html"
}
//保存
// @router /blog/cat [post]
func (c *BlogCat)Post() {
	rsp := Response.NewResponse()
	defer rsp.WriteJson(c.Ctx.ResponseWriter)
	blogModel := models.Blog{}

	//参数传递
	if err := c.ParseForm(&blogModel); err != nil {
		rsp.Error(err.Error())
		c.StopRun()
	}
	blogModel.CatId=blog.CAT_ID
	//日期
	date, ok := c.GetDateTime("time_add")
	if ok {
		blogModel.TimeAdd = date
	}
	//创建
	var serv blog.BlogCat
	id, err := serv.Create(&blogModel)
	if err != nil {
		rsp.Error(err.Error())
	} else {
		fmt.Println("创建成功！:", id)
		rsp.Success("")
	}
}
//更新
// @router /blog/cat/:id [put]
func (c *BlogCat)Put() {
	rsp := Response.NewResponse()
	defer rsp.WriteJson(c.Ctx.ResponseWriter)
	//ID 获取 格式化
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	//参数传递
	blogModel := models.Blog{}
	if err := c.ParseForm(&blogModel); err != nil {
		rsp.Error(err.Error())
	}
	blogModel.CatId=blog.CAT_ID
	//日期
	date, ok := c.GetDateTime("time_add")
	if ok {
		blogModel.TimeAdd = date
	}
	//更新
	var ser *blog.BlogCat
	_, err := ser.Update(int_id, &blogModel)
	if err != nil {
		rsp.Error(err.Error())
	} else {
		rsp.Success("")
	}
}
//删除
// @router /blog/cat/:id [delete]
func (c *BlogCat)Delete() {
	rsp := Response.NewResponse()
	defer rsp.WriteJson(c.Ctx.ResponseWriter)
	//ID 获取 格式化
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	//更新
	var ser *blog.BlogCat
	_, err := ser.Delete(int_id)
	if err != nil {
		rsp.Error(err.Error())
	} else {
		rsp.Success("")
	}
}
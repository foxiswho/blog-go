package admin

import (
	"fox/util/Response"
	"strconv"
	"fox/models"
	"fox/service/blog"
	"fmt"
)

type BlogController struct {
	BaseController
}
func (c *BlogController) URLMapping() {
	c.Mapping("CheckTitle", c.CheckTitle)
	c.Mapping("List", c.List)
	c.Mapping("Add", c.Add)
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
	c.Mapping("Put", c.Put)
	c.Mapping("Detail", c.Detail)
}
//检测名称重复
// @router /blog/check_title [post]
func (c *BlogController)CheckTitle() {
	rsp := Response.NewResponse()
	defer rsp.WriteJson(c.Ctx.ResponseWriter)
	//ID 获取 格式化
	int_id,_  := c.GetInt("cat_id")
	id,_  := c.GetInt("id")
	name := c.GetString("title")
	//创建
	var serv blog.Blog
	ok, err := serv.CheckTitleById(int_id,name,id)
	if err != nil {
		rsp.Error(err.Error())
	} else {
		fmt.Println("成功！:", ok)
		rsp.Success("")
	}
}
//列表
// @router /blog [get]
func (c *BlogController)List() {
	var ser *blog.Blog
	data, err := ser.Query(0)
	//println(data)
	println(err)
	c.Data["data"] = data
	c.Data["title"] = "博客-列表"
	c.TplName = "admin/blog/list.html"
}
//编辑
// @router /blog/:id [get]
func (c *BlogController)Get() {
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	var ser *blog.Blog
	data, err := ser.Read(int_id)
	//println("Detail :", err.Error())
	if err != nil {
		rsp := Response.NewResponse()
		defer rsp.WriteJson(c.Ctx.ResponseWriter)
		rsp.Error(err.Error())
	} else {
		c.Data["info"] = data["Blog"]
		c.Data["statistics"] = data["Statistics"]
		c.Data["TimeAdd"] = data["TimeAdd"]
		c.Data["title"] = "博客-编辑"
		c.Data["_method"] = "put"
		c.Data["is_put"] = true
		c.TplName = "admin/blog/get.html"
	}
}
//添加
// @router /blog/add [get]
func (c *BlogController)Add() {
	c.Data["_method"] = "post"
	c.Data["title"] = "博客-添加"
	c.TplName = "admin/blog/get.html"
}
//保存
// @router /blog [post]
func (c *BlogController)Post() {
	rsp := Response.NewResponse()
	defer rsp.WriteJson(c.Ctx.ResponseWriter)
	blogModel := models.Blog{}

	//参数传递
	blog_statistics := models.BlogStatistics{}
	if err := c.ParseForm(&blogModel); err != nil {
		rsp.Error(err.Error())
		c.StopRun()
	}
	if err := c.ParseForm(&blog_statistics); err != nil {
		rsp.Error(err.Error())
		c.StopRun()
	}
	//日期
	date, ok := c.GetDateTime("time_add")
	if ok {
		blogModel.TimeAdd = date
	}
	//创建
	var serv blog.Blog
	id, err := serv.Create(&blogModel, &blog_statistics)
	if err != nil {
		rsp.Error(err.Error())
	} else {
		fmt.Println("创建成功！:", id)
		rsp.Success("")
	}
}
//查看
// @router /blog/detail/:id [get]
func (c *BlogController)Detail() {
	c.Get()
	c.Data["title"] = "博客-查看"
	c.TplName = "admin/blog/detail.html"
}
//更新
// @router /blog/:id [put]
func (c *BlogController)Put() {
	rsp := Response.NewResponse()
	defer rsp.WriteJson(c.Ctx.ResponseWriter)
	//ID 获取 格式化
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	//参数传递
	blogMoel := models.Blog{}
	blog_statistics := models.BlogStatistics{}
	if err := c.ParseForm(&blogMoel); err != nil {
		rsp.Error(err.Error())
	}
	if err := c.ParseForm(&blog_statistics); err != nil {
		rsp.Error(err.Error())
	}
	//日期
	date, ok := c.GetDateTime("time_add")
	if ok {
		blogMoel.TimeAdd = date
	}
	//更新
	var ser *blog.Blog
	_, err := ser.Update(int_id, &blogMoel, &blog_statistics)
	if err != nil {
		rsp.Error(err.Error())
	} else {
		rsp.Success("")
	}
}
//删除
// @router /blog/:id [delete]
func (c *BlogController)Delete() {
	rsp := Response.NewResponse()
	defer rsp.WriteJson(c.Ctx.ResponseWriter)
	//ID 获取 格式化
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	//更新
	var ser *blog.Blog
	_, err := ser.Delete(int_id)
	if err != nil {
		rsp.Error(err.Error())
	} else {
		rsp.Success("")
	}
}

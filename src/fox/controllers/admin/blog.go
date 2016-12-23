package admin

import (
	"fox/service"
	"fox/util/Response"
	"strconv"
	"fox/models"
	"fmt"
)

type BlogController struct {
	BaseController
}
//列表
func (c *BlogController)List() {
	var blog *service.Blog
	data, err := blog.Query()
	//println(data)
	println(err)
	c.Data["data"] = data
	c.Data["title"]= "博客-列表"
	c.TplName = "admin/blog/list.html"
}
//详情
func (c *BlogController)Get() {
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	var blog *service.Blog
	data, err := blog.Read(int_id)
	//println("Detail :", err.Error())
	if err != nil {
		rsp := Response.NewResponse()
		defer rsp.WriteJson(c.Ctx.ResponseWriter)
		rsp.Error(err.Error())
	} else {
		c.Data["info"] = data["Blog"]
		c.Data["statistics"] = data["Statistics"]
		c.Data["TimeAdd"] = data["TimeAdd"]
		c.Data["title"]= "博客-查看"
		c.TplName = "admin/blog/get.html"
	}
}
//添加
func (c *BlogController)Add() {
	c.Data["_method"] = "post"
	c.Data["title"]= "博客-添加"
	c.TplName = "admin/blog/edit.html"
}
//保存
func (c *BlogController)Post() {

	rsp := Response.NewResponse()
	defer rsp.WriteJson(c.Ctx.ResponseWriter)
	blog := models.Blog{}
	//blog.Title=c.GetString("title")
	//blog.Content=c.GetString("content")
	//tmp,_:=c.GetInt("status")
	//blog.Status=tmp
	//tmp2,_:=c.GetInt8("is_open")
	//blog.IsOpen=tmp2
	//c.GetString("time_add")
	//time_add:
	//author:
	//url_source:
	//url:
	//url:
	//thumb:
	//sort:
	//description:
	//seo_title:
	//seo_keyword:
	//seo_description:

	//参数传递
	blog_statistics := models.BlogStatistics{}
	if err := c.ParseForm(&blog); err != nil {
		rsp.Error(err.Error())
		c.StopRun()
	}
	if err := c.ParseForm(&blog_statistics); err != nil {
		rsp.Error(err.Error())
		c.StopRun()
	}
	//日期
	date,ok:=c.GetDateTime("time_add")
	if ok {
		blog.TimeAdd = date
	}
	//创建
	var serv service.Blog
	id, err := serv.Create(&blog, &blog_statistics)
	if err != nil {
		rsp.Error(err.Error())
	} else {
		fmt.Println("创建成功！:", id)
		rsp.Success("")
	}
}
//编辑
func (c *BlogController)Edit() {
	c.Get()
	c.Data["_method"] = "put"
	c.Data["is_put"] = true
	c.Data["title"]= "博客-修改"
	c.TplName = "admin/blog/edit.html"
}
//更新
func (c *BlogController)Put() {
	rsp := Response.NewResponse()
	defer rsp.WriteJson(c.Ctx.ResponseWriter)
	//ID 获取 格式化
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	//参数传递
	blog := models.Blog{}
	blog_statistics := models.BlogStatistics{}
	if err := c.ParseForm(&blog); err != nil {
		rsp.Error(err.Error())
	}
	if err := c.ParseForm(&blog_statistics); err != nil {
		rsp.Error(err.Error())
	}
	//日期
	date,ok:=c.GetDateTime("time_add")
	if ok {
		blog.TimeAdd = date
	}
	//更新
	var ser *service.Blog
	_, err := ser.Update(int_id, &blog, &blog_statistics)
	if err != nil {
		rsp.Error(err.Error())
	} else {
		rsp.Success("")
	}
}
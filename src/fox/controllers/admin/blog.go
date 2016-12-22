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
func (this *BlogController)List() {
	var blog *service.Blog
	data, err := blog.Query()
	//println(data)
	println(err)
	this.Data["data"] = data
	//this.Data["title"]=
	this.TplName = "admin/blog/get.html"
}
//详情
func (this *BlogController)Get() {
	id := this.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	var blog *service.Blog
	data, err := blog.Read(int_id)
	//println("Detail :", err.Error())
	if err != nil {
		rsp := Response.NewResponse()
		defer rsp.WriteJson(this.Ctx.ResponseWriter)
		rsp.Error(err.Error())
	} else {
		this.Data["info"] = data["Blog"]
		this.Data["statistics"] = data["Statistics"]
		this.Data["TimeAdd"] = data["TimeAdd"]
		//this.Data["title"]=
		this.TplName = "admin/blog/detail.html"
	}
}
//添加
func (this *BlogController)Add() {
	this.Data["_method"] = "post"
	this.TplName = "admin/blog/edit.html"
}
//保存
func (this *BlogController)Post() {

	rsp := Response.NewResponse()
	defer rsp.WriteJson(this.Ctx.ResponseWriter)
	blog := models.Blog{}
	//blog.Title=this.GetString("title")
	//blog.Content=this.GetString("content")
	//tmp,_:=this.GetInt("status")
	//blog.Status=tmp
	//tmp2,_:=this.GetInt8("is_open")
	//blog.IsOpen=tmp2
	//this.GetString("time_add")
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
	if err := this.ParseForm(&blog); err != nil {
		rsp.Error(err.Error())
		this.StopRun()
	}
	if err := this.ParseForm(&blog_statistics); err != nil {
		rsp.Error(err.Error())
		this.StopRun()
	}
	//日期
	date,ok:=this.GetDateTime("time_add")
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
func (this *BlogController)Edit() {
	this.Get()
	this.Data["_method"] = "put"
	this.Data["is_put"] = true
	this.TplName = "admin/blog/edit.html"
}
//更新
func (this *BlogController)Put() {
	rsp := Response.NewResponse()
	defer rsp.WriteJson(this.Ctx.ResponseWriter)
	//ID 获取 格式化
	id := this.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	//参数传递
	blog := models.Blog{}
	blog_statistics := models.BlogStatistics{}
	if err := this.ParseForm(&blog); err != nil {
		rsp.Error(err.Error())
	}
	if err := this.ParseForm(&blog_statistics); err != nil {
		rsp.Error(err.Error())
	}
	//日期
	date,ok:=this.GetDateTime("time_add")
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
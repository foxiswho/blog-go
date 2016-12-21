package admin

import (
	"fox/service"
	"fox/util/Response"
	"strconv"
)

type BlogController struct {
	BaseController
}

func (this *BlogController)Get() {
	var blog *service.Blog
	data, err := blog.Get()
	//println(data)
	println(err)
	this.Data["data"] = data
	//this.Data["title"]=
	this.TplName = "admin/blog/get.html"
}
//详情
func (this *BlogController)Detail() {
	id:=this.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	var blog *service.Blog
	data, err := blog.Detail(int_id)
	//println("Detail :", err.Error())
	if err != nil {
		rsp := Response.NewResponse()
		defer rsp.WriteJson(this.Ctx.ResponseWriter)
		rsp.Error(err.Error())
	} else {
		this.Data["info"] = data
		//this.Data["title"]=
		this.TplName = "admin/blog/detail.html"
	}
}
package controllers

import "github.com/astaxie/beego"

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	c.Data["content"] = "page not found"
	c.TplName = "error/404.html"
}

func (c *ErrorController) Error501() {
	c.Data["content"] = "server error"
	c.TplName = "error/501.html"
}


func (c *ErrorController) ErrorDb() {
	c.Data["content"] = "database is now down"
	c.TplName = "error/dberror.html"
}
//  框架中的扩展函数
func (this *ErrorController) Prepare() {
	this.Initialization()
}
// 初始化数据
func (this *ErrorController) Initialization() {
	this.Data["__public__"] = "/"
	this.Data["__static__"] = "/static/"
	this.Data["__theme__"] = "/static/post/"
	this.Data["blog_name"] = beego.AppConfig.String("blog_name")
	//orm.RunSyncdb("default", false, true)
}
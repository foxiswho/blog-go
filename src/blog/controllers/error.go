package controllers

import (
	"blog/fox"
	"blog/fox/config"
)

type Error struct {
	fox.Controller
}
//404错误 内容页面
func (c *Error) Error404() {
	c.Data["content"] = "page not found"
	c.TplName = "error/404.html"
}
//501错误 内容页面
func (c *Error) Error501() {
	c.Data["content"] = "server error"
	c.TplName = "error/501.html"
}

//数据库错误 内容页面
func (c *Error) ErrorDb() {
	c.Data["content"] = "database is now down"
	c.TplName = "error/dberror.html"
}
//  框架中的扩展函数
func (c *Error) Prepare() {
	c.Initialization()
}
// 初始化数据
func (c *Error) Initialization() {
	//模版参数
	c.Data["__public__"] = "/"
	c.Data["__static__"] = "/static/"
	c.Data["__theme__"] = "/static/post/"
	c.Data["site_name"] = config.String("site_name")
}
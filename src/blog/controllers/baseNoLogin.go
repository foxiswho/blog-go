package controllers

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
	"blog/fox/datetime"
	"blog/fox/db"
	"blog/service/admin"
	"blog/fox"
	"blog/fox/log"
)
//基础控制器，不需要登录
type BaseNoLogin struct {
	fox.Controller
	Site *admin.Site
	//Session *service.AdminSession //当前登录用户信息
}

//  框架中的扩展函数
func (c *BaseNoLogin) Prepare() {
	c.Initialization()
}
// 初始化数据
func (c *BaseNoLogin) Initialization() {
	//模版 中JS，CSS等静态文件夹
	//TODO 后期会更改
	c.Data["__public__"] = "/"
	c.Data["__static__"] = "/static/"
	c.Data["__theme__"] = "/static/post/"
	//博客基本参数读取加载
	//初始化
	c.Site = admin.NewSiteService()
	//赋值
	c.Site.SetSiteConfig()
	//检测变量是否赋值成功
	if c.Site!=nil {
		//如果参数加载为nil,则独有配置文件中博客名称
		c.Data["site_name"] = c.Site.GetString("SITE_NAME")
	}

}
//表单日期时间
func (c *BaseNoLogin) GetDateTime(key string) (time.Time, bool) {
	date := c.GetString(key)
	if len(date) > 0 {
		date, err := datetime.FormatTimeStructLocation(date, datetime.Y_M_D_H_I_S)
		if err == nil {
			return date, true
		}
	}
	return time.Time{}, false
}
//错误信息
func (c *BaseNoLogin) Error(key string, def ...map[string]interface{}) {
	if c.IsAjax() {
		c.ErrorJson(key, def...)
	} else {
		c.Data["content"] = key
		c.TplName = "error/404.html"
	}
}
//成功
func (c *BaseNoLogin) Success(key string, def ...map[string]interface{}) {
	c.Data["content"] = key
	if c.IsAjax() {
		c.SuccessJson(key, def...)
	} else {
		c.TplName = "error/success.html"
	}
}
//错误信息JSON
func (c *BaseNoLogin) ErrorJson(key string, def ...map[string]interface{}) {
	m := make(map[string]interface{})
	if len(def) > 0 {
		m = def[0]
	}
	m["info"] = key
	m["code"] = 0
	c.Data["json"] = m
	c.ServeJSON()
}
//成功JSON
func (c *BaseNoLogin) SuccessJson(key string, def ...map[string]interface{}) {
	m := make(map[string]interface{})
	if len(def) > 0 {
		m = def[0]
	}
	m["info"] = key
	m["code"] = 1
	c.Data["json"] = m
	c.ServeJSON()
}
//初始化数据库
func init() {
	//初始化
	db.Init();
	log.Info("init start...")
}
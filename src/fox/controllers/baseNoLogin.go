package controllers

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"fox/util/datetime"
	"fox/util/db"
	"github.com/astaxie/beego/orm"
	"fox/service/admin"
)

type BaseNoLogin struct {
	beego.Controller
	Config map[string]interface{}
	//Session *service.AdminSession //当前登录用户信息
}

//  框架中的扩展函数
func (c *BaseNoLogin) Prepare() {
	c.Initialization()
}
// 初始化数据
func (c *BaseNoLogin) Initialization() {
	c.Data["__public__"] = "/"
	c.Data["__static__"] = "/static/"
	c.Data["__theme__"] = "/static/post/"
	//博客名字
	c.Config = admin.NewTypeService().SiteConfig()
	if len(c.Config) > 0 {
		c.Data["site_name"] = c.Config["SITE_NAME"]
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
//表单日期时间
func (c *BaseNoLogin) Error(key string) {
	c.Data["content"] = key
	c.TplName = "error/404.html"
}
//初始化数据库
func init() {
	//初始化
	db.Init();
	orm.DefaultTimeLoc = time.UTC
	beego.Info("init orm start...")
}
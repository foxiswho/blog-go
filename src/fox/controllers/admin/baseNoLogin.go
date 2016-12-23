package admin

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"fox/util/datetime"
)

type BaseNoLoginController struct {
	AdminSession
	//beego.Controller
	//Session *service.AdminSession //当前登录用户信息
}

//  框架中的扩展函数
func (c *BaseNoLoginController) Prepare() {
	c.Initialization()
}
// 初始化数据
func (c *BaseNoLoginController) Initialization() {
	c.Data["HtmlHead"] = ""
	c.Data["__public__"] = "/"
	c.Data["__static__"] = "/static/"
	c.Data["__theme__"] = "/static/Hplus-v.4.1.0/"
	c.Data["blog_name"] = beego.AppConfig.String("blog_name")
	//c.Layout="admin/public/layout.html"
	//orm.RunSyncdb("default", false, true)
}
//表单日期时间
func (c *BaseNoLoginController) GetDateTime(key string) (time.Time, bool) {
	date := c.GetString(key)
	if len(date) > 0 {
		date, err := datetime.FormatTimeStructLocation(date, datetime.Y_M_D_H_I_S)
		if err == nil {
			return date, true
		}
	}
	return time.Time{}, false
}

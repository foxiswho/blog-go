package admin

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
	"fox/util/datetime"
	"fox/service/admin"
	"strings"
)

type BaseNoLoginController struct {
	AdminSession
	Config map[string]interface{}
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
	//博客名字
	c.Config = admin.NewTypeService().SiteConfig()
	if len(c.Config) > 0 {
		c.Data["site_name"] = c.Config["SITE_NAME"]
	}
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
// GetString returns the input value by key string or the default value while it's present and input is blank
func (c *BaseNoLoginController) GetString(key string, def ...string) string {
	if v := c.Ctx.Input.Query(key); v != "" {
		return strings.TrimSpace(v)
	}
	if len(def) > 0 {
		return def[0]
	}
	return ""
}
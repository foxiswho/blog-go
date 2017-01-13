package admin

import (
	"time"
	"blog/fox/datetime"
	"blog/service/admin"
	"strings"
)

type BaseNoLogin struct {
	AdminSession
	Site *admin.Site
}

//  框架中的扩展函数
func (c *BaseNoLogin) Prepare() {
	c.Initialization()
}
// 初始化数据
func (c *BaseNoLogin) Initialization() {
	c.Data["HtmlHead"] = ""
	c.Data["__public__"] = "/"
	c.Data["__static__"] = "/static/"
	c.Data["__theme__"] = "/static/Hplus-v.4.1.0/"
	//博客名字
	c.Site = admin.NewSiteService()
	c.Site.SetSiteConfig()
	if c.Site!=nil{
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
// GetString returns the input value by key string or the default value while it's present and input is blank
func (c *BaseNoLogin) GetString(key string, def ...string) string {
	if v := c.Ctx.Input.Query(key); v != "" {
		return strings.TrimSpace(v)
	}
	if len(def) > 0 {
		return def[0]
	}
	return ""
}
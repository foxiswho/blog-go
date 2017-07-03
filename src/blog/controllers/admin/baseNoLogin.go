package admin

import (
	"time"
	"blog/fox/datetime"
	"blog/service/admin"
	"strings"
)
//无须登录 基础控制器
type BaseNoLogin struct {
	AdminSession //继承此控制器
	Site *admin.Site //博客参数
}

//  框架中的扩展函数
func (c *BaseNoLogin) Prepare() {
	//初始化
	c.Initialization()
}
// 初始化数据
func (c *BaseNoLogin) Initialization() {
	//模版参数
	c.Data["HtmlHead"] = ""
	c.Data["__public__"] = "/"
	c.Data["__static__"] = "/static/"
	c.Data["__theme__"] = "/static/Hplus-v.4.1.0/"
	//初始化
	c.Site = admin.NewSiteService()
	//博客配置赋值
	c.Site.SetSiteConfig()
	//检测是否存在
	if c.Site!=nil{
		//不存在，则，使用 配置文件中名称
		c.Data["site_name"] = c.Site.GetString("site_name")
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
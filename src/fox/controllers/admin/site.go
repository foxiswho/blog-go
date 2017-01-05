package admin

import (
	"fox/service/admin"
	"fmt"
)

type Site struct {
	Base
}
func (c *Site) URLMapping() {
	c.Mapping("List", c.List)
	c.Mapping("Put", c.Put)
}
//列表
// @router /site [get]
func (c *Site)List() {
	ser :=admin.NewSiteService()
	data, err := ser.Query()
	fmt.Println(err)
	c.Data["data"] = data
	c.Data["title"] = "站点配置"
	c.TplName = "admin/site/list.html"
}
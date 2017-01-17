package admin

import (
	"blog/service/admin"
	"fmt"
)
//站点配置
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
	ser := admin.NewSiteService()
	data, err := ser.Query()
	fmt.Println(err)
	c.Data["data"] = data
	c.Data["title"] = "站点配置"
	c.Data["_method"] = "put"
	c.Data["is_put"] = true
	c.TplName = "admin/site/list.html"
}
//更新
// @router /site [put]
func (c *Site)Put() {
	//参数传递
	ser := admin.NewSiteService()
	_, err := ser.Update(c.Input())
	if err != nil {
		c.Error(err.Error())
	} else {
		c.Success("操作成功")
	}
}
package admin

import (
	"blog/service/admin"
	"fmt"
)
//地区
type Area struct {
	Base
}

func (c *Area) URLMapping() {
	c.Mapping("List", c.List)
}
//列表
// @router /area [get]
func (c *Area)List() {
	id, _ := c.GetInt("parent_id")
	//查询
	where := make(map[string]interface{})
	where["parent_id"] = id
	//初始化
	mod := admin.NewAreaService()
	//查询
	page, _ := c.GetInt("page")
	data, err := mod.GetAll(where, []string{}, "id ASC", page, 999)
	if err != nil {
		fmt.Println(err.Error())
		c.Error(err.Error())
		return
	}
	c.Data["data"] = data
	//扩展查询
	ext := admin.NewAreaExtService()
	data_ext, err := ext.GetAll(where, []string{}, "id ASC,ext_id ASC", page, 999)
	if err != nil {
		c.Error(err.Error())
		return
	}
	c.Data["data_ext"] = data_ext
	c.Data["title"] = "地区-列表"
	c.TplName = "admin/area/list.html"
}
package admin

import (
	"fox/service/admin"
	"fmt"
)

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
	where := make(map[string]interface{})
	where["parent_id"] = id
	mod := admin.NewAreaService()
	page, _ := c.GetInt("page")
	data, err := mod.GetAll(where, []string{}, "id ASC", page, 999)
	if err != nil {
		fmt.Println(err)
	}
	c.Data["data"] = data
	ext := admin.NewAreaExtService()
	data_ext, err := ext.GetAll(where, []string{}, "id ASC,ext_id ASC", page, 999)
	if err != nil {
		fmt.Println(err)
	}
	c.Data["data_ext"] = data_ext
	c.Data["title"] = "地区-列表"
	c.TplName = "admin/area/list.html"
}
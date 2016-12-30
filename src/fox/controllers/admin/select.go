package admin

import (
	"strconv"
	"fox/service/admin"
	"fmt"
	"fox/model"
)

type Select struct {
	BaseController
}

func (c *Select) URLMapping() {
	c.Mapping("Type", c.Type)
}
//类别
// @router /select/type [get]
// @router /select/type/:id [get]
func (c *Select)Type() {
	id := c.Ctx.Input.Param(":id")
	int_id, _ := strconv.Atoi(id)
	ser :=admin.NewTypeService()
	data, err := ser.Query(int_id)
	fmt.Println(err)
	mod := model.NewType()
	c.Data["info"] = mod
	if int_id > 0 {
		data, err := ser.Read(int_id)
		if err == nil {
			c.Data["info"] = data["info"]
		}
	}
	c.Data["id"] = id
	c.Data["data"] = data
	c.Data["title"] = "类别-列表"
	c.TplName = "admin/select/type.html"
}
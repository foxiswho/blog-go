package admin

import (
	"strconv"
	"fox/service/admin"
	"fmt"
	"fox/models"
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
	var ser admin.Type
	data, err := ser.Query(int_id)
	fmt.Println(err)
	c.Data["info"] = models.Type{}
	if int_id > 0 {
		var model *admin.Type
		data, err := model.Read(int_id)
		if err == nil {
			c.Data["info"] = data["info"]
		}
	}
	c.Data["id"] = id
	c.Data["data"] = data
	c.Data["title"] = "类别-列表"
	c.TplName = "admin/select/type.html"
}
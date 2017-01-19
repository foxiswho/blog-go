package admin

import (
	"strconv"
	"blog/service/admin"
	"fmt"
	"blog/model"
)
//各种选择加载数据
type Select struct {
	Base
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
	//初始化
	ser :=admin.NewTypeService()
	data, err := ser.Query(int_id)
	if err!=nil{
		fmt.Println(err.Error())
		c.Error(err.Error())
		return
	}
	mod := model.NewType()
	c.Data["info"] = mod
	if int_id > 0 {
		//获取该信息数据
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
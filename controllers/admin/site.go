package admin

import (
	"fmt"
	"github.com/foxiswho/blog-go/service/admin"
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
func (c *Site) List() {
	//初始化
	ser := admin.NewSiteService()
	data, err := ser.Query()
	if err != nil {
		fmt.Println(err.Error())
		c.Error(err.Error())
		return
	}
	c.Data["data"] = data
	c.Data["title"] = "站点配置"
	c.Data["_method"] = "put"
	c.Data["is_put"] = true
	c.TplName = "admin/site/list.html"
}

//更新
// @router /site [put]
func (c *Site) Put() {
	//参数传递
	ser := admin.NewSiteService()
	input, err2 := c.Input()
	if err2 != nil {
		c.Error(err2.Error())
	}
	_, err := ser.Update(input)
	if err != nil {
		c.Error(err.Error())
	} else {
		c.Success("操作成功")
	}
}

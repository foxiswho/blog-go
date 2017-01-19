package admin

import (
	"blog/service/admin"
	"fmt"
)
//附件
type Attachment struct {
	Base
}

func (c *Attachment) URLMapping() {
	c.Mapping("List", c.List)
}
//列表
// @router /attachment [get]
func (c *Attachment)List() {
	where := make(map[string]interface{})
	//初始化
	mod := admin.NewAttachmentService()
	//查询
	page, _ := c.GetInt("page")
	data, err := mod.GetAll(where, []string{}, "attachment_id DESC", page, 20)
	if err != nil {
		fmt.Println(err.Error())
		c.Error(err.Error())
		return
	}
	c.Data["data"] = data
	c.Data["title"] = "附件-列表"
	c.TplName = "admin/attachment/list.html"
}
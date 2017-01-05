package admin

import (
	"fox/service/admin"
	"fmt"
)

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
	mod := admin.NewAttachmentService()
	page, _ := c.GetInt("page")
	data, err := mod.GetAll(where, []string{}, "attachment_id DESC", page, 20)
	if err != nil {
		fmt.Println(err)
	}
	c.Data["data"] = data
	c.Data["title"] = "附件-列表"
	c.TplName = "admin/attachment/list.html"
}
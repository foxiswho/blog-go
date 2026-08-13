package modBasicConfigEventFields

import (
	"github.com/hongmengzhu/xianfu-blog-go/app/models/basic/modBasicConfigEvent"
)

type CreateUpdateCt struct {
	EventNo    string                       `json:"eventNo" label:"事件编号"`
	Body       []modBasicConfigEvent.ItemVo `json:"body" label:"表体"`
	BodyDelIds []string                     `json:"bodyDelIds" label:"表体删除的id"`
}

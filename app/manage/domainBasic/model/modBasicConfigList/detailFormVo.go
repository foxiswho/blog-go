package modBasicConfigList

import "github.com/hongmengzhu/xianfu-blog-go/app/manage/domainBasic/model/modBasicConfigEvent"

type DetailFormVo struct {
	Form modBasicConfigEvent.ModelForm `json:"form"`
	Data map[string]interface{}        `json:"data"`
}

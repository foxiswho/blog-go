package modBasicDataDictionary

import "github.com/hongmengzhu/xianfu-blog-go/pkg/tools/typePg"

type ExistValue struct {
	Id      typePg.Int64String `json:"id"`
	Wd      string             `json:"wd"`
	OwnerNo string             `json:"ownerNo" label:"所属/拥有者" `
}

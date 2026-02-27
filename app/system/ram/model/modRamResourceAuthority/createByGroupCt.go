package modRamResourceAuthority

import "github.com/foxiswho/blog-go/pkg/tools/typePg"

type CreatByGroupCt struct {
	GroupId typePg.Int64String `json:"groupId" form:"groupId" validate:"required" label:"资源组id" `
	Ids     []string           `json:"ids" form:"ids" validate:"required" label:"数据" `
}

package modRamResourceAuthority

type CreatByGroupCt struct {
	GroupNo string   `json:"groupNo" form:"groupNo" validate:"required" label:"资源组id" `
	Ids     []string `json:"ids" form:"ids" validate:"required" label:"数据" `
}

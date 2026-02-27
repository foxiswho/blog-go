package modRamResourceAuthority

type UpdateByTypeValueCt struct {
	TypeValue string   `json:"typeValue" form:"typeValue" validate:"required" label:"资源组id" `
	Ids       []string `json:"ids" form:"ids" validate:"required" label:"数据" `
}

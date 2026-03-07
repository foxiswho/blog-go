package modBasicAttachment

type MakeFileOwnerAllCt struct {
	Rule []MakeFileOwnerCt `json:"rule" label:"规则"`
	Num  int32             `json:"num" label:"数量"`
	Mark string            `json:"mark" label:"mark标记" `
}

package modBasicAttachment

type DelFileOwnerCt struct {
	Nos       []string `json:"nos" label:"no"`
	FileOwner string   `json:"fileOwner" label:"根据fileOwner查询"`
}

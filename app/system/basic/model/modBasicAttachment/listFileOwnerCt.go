package modBasicAttachment

type ListFileOwnerCt struct {
	GroupData []ListFileOwnerGroup `json:"groupData" label:"分组查询"`
	FileOwner []string             `json:"fileOwner" label:"根据fileOwner查询"`
}

package modBasicAttachment

type ListFileOwnerGroup struct {
	FileOwner string `json:"fileOwner" label:"文件拥有者"`
	Group     string `json:"group" label:"分组"`
}

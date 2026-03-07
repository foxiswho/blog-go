package modBasicAttachment

type ListFIleOwnerVo struct {
	GroupData map[string][]Vo `json:"groupData"`
	Data      []Vo            `json:"data"`
}

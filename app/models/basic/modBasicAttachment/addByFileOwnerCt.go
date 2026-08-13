package modBasicAttachment

type AddByFileOwnerCt struct {
	FileOwner    string   `json:"fileOwner"`
	FileOwnerSub string   `json:"fileOwnerSub"`
	Nos          []string `json:"nos"`
}

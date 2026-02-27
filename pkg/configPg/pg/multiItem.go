package pg

type MultiItem struct {
	//包含的表
	Contain []string `json:"contain" value:"${contain:=[]}"  toml:"contain"`
	//不包含的表
	Not []string `json:"not" value:"${not:=[]}"  toml:"not"`
}

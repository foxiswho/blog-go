package pg

type Profiles struct {
	Active  string   `json:"active" value:"${active}"  toml:"active"`    //激活项
	Include []string `json:"include" value:"${include}"  toml:"include"` // 包含多项
}

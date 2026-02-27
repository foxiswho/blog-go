package pg

// Directory 目录模块
type Directory struct {
	Data string `value:"${data:=data}" toml:"data" json:"data" label:"数据资源根目录"`
}

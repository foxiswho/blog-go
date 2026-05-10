package pg

// 数据
type Data struct {
	Delete bool `toml:"delete" value:"${delete:=false}"` // 点击删除时是否直接删除
}

package pg

type Attachment struct {
	Dir             string `value:"${dir:=attachment}" toml:"dir" json:"dir" label:"目录,在url 中会显示"`
	DirRoot         string `value:"${dirRoot:=data}" toml:"dirRoot" json:"dirRoot"  label:"存储根目录，默认 data，不会在url中显示"`
	Domain          string `value:"${domain}" toml:"domain" json:"domain" label:"域名" `
	MkdirPermission uint32 `value:"${mkdirPermission:=750}" toml:"mkdirPermission" json:"mkdirPermission"` //目录权限代号
	Minio           Minio  `value:"${minio}" toml:"minio" json:"minio"`
}

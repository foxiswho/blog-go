package pg

type Minio struct {
	Enable bool   `value:"${enable:=false}" toml:"enable" json:"enable"` // 是否启用 HTTP
	Host   string `value:"${host:=127.0.0.1}" toml:"host" json:"host"`   // HTTP host
	Port   int    `value:"${port:=9000}" toml:"port" json:"port"`        // HTTP 端口
	Access string `value:"${access:=}" toml:"access" json:"access"`      // Access
	Secret string `value:"${secret:=}" toml:"secret" json:"secret"`      // Secret
	Secure bool   `value:"${secure:=true}" toml:"secure" json:"secure"`  // Secure
	Bucket string `value:"${bucket:=}" toml:"bucket" json:"bucket"`
	Region string `value:"${region:=}" toml:"region" json:"region"`
}

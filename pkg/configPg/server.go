package configPg

// Server 服务配置
type Server struct {
	Port   int    `value:"${port:=18080}"`
	Domain string `value:"${domain:=http://localhost:8080}" toml:"domain"`
}

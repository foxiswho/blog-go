package pg

type Mail struct {
	Host     string `value:"${host:=}" toml:"host" json:"host"`
	Port     int    `value:"${port:=12345}" toml:"port" json:"port"`
	Username string `value:"${username:=}" toml:"username" json:"username"`
	AuthCode string `value:"${authCode:=}" toml:"authCode" json:"authCode"`  //密码或授权码
	SMTPAuth string `value:"${smtpAuth:=}" toml:"smtpAuth" json:"smtpAuth" ` // 验证类型
	From     string `value:"${from:=}" toml:"from" json:"from"`
}

package modRamLogin

type LoginCt struct {
	Account   string `json:"account" json:"account" form:"account" label:"账号"` //账号
	Password  string `json:"password"`                                         //密码
	Code      string `json:"code"`                                             //验证码
	CodeSms   string `json:"codeSms"`                                          //短信验证码
	TypeLogin string `json:"type"`                                             //登陆类型
}

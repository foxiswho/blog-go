package modRamIdpCredential

type CreateCt struct {
	Description string `json:"description" label:"描述" ` // 描述
	Idp         string `json:"idp" label:"身份提供商" `    // 身份提供商
	SourceNo    string `json:"sourceNo" label:"认证源编号" ` // 认证源编号
	CredType    string `json:"credType" label:"凭证类型" `   // 凭证类型
	Value       string `json:"value" label:"凭证值" `        // 凭证值
}

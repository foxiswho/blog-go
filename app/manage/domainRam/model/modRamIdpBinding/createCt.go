package modRamIdpBinding

type CreateCt struct {
	Description   string `json:"description" label:"描述" `   // 描述
	Idp           string `json:"idp" label:"身份提供商" `      // 身份提供商
	OpenId        string `json:"openId" label:"OpenId" `      // OpenId
	UnionId       string `json:"unionId" label:"UnionId" `    // UnionId
	AppMark       string `json:"appMark" label:"应用标识" `    // 应用标识
	Platform      string `json:"platform" label:"平台" `       // 平台
	Protocol      string `json:"protocol" label:"协议" `       // 协议
	Mail          string `json:"mail" label:"邮箱" `           // 邮箱
	Phone         string `json:"phone" label:"手机" `          // 手机
	NickName      string `json:"nickName" label:"昵称" `       // 昵称
	Avatar        string `json:"avatar" label:"头像" `         // 头像
	ExtraData     string `json:"extraData" label:"扩展数据" `  // 扩展数据
	ExternalSub   string `json:"externalSub" label:"外部身份标识" ` // 外部身份标识
}

package modRamAccount

import (
	"github.com/foxiswho/blog-go/pkg/tools/typePg"
)

// CreateAccountCt
// @Description: 创建账号
type CreateAccountCt struct {
	Account      string       `json:"account"  validate:"required,min=1,max=255" label:"账户" ` // 账户
	Cc           string       `json:"cc" label:"国际区号"`                                        // 国际区号
	Phone        string       `json:"phone" label:"手机号" `                                     // 手机号
	Mail         string       `json:"mail" label:"邮箱" `                                       // 邮箱
	Code         string       `json:"code" label:"编码" `                                       // 编码
	RegisterTime *typePg.Time `json:"registerTime" label:"注册时间" `                             // 注册时间
	Description  string       `json:"description" label:"描述" `                                // 描述
}

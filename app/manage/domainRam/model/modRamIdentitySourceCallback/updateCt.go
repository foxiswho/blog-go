package modRamIdentitySourceCallback

import (
	"github.com/hongmengzhu/xianfu-blog-go/pkg/tools/typePg"
)

type UpdateCt struct {
	ID          typePg.Uint64String `json:"id" form:"id" validate:"required" label:"id" `
	Name        string              `json:"name" form:"name" validate:"required,min=1,max=255" label:"名称" ` // 名称
	Callback    string              `json:"callback" form:"callback" validate:"required,min=1,max=255" label:"回调地址" ` // 回调地址
	Description string              `json:"description" label:"描述" ` // 描述
	Platform    string              `json:"platform" label:"平台" `    // 平台
	Protocol    string              `json:"protocol" label:"协议" `    // 协议
	Idp         string              `json:"idp" label:"身份提供商" `    // 身份提供商
}

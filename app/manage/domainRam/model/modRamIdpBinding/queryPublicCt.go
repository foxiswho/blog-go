package modRamIdpBinding

import (
	"github.com/hongmengzhu/xianfu-blog-go/pkg/model"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/tools/typePg"
)

type QueryPublicCt struct {
	model.BaseQueryNodeCt
	ID          typePg.Uint64String `json:"id" label:"" `
	Description string              `json:"description" label:"描述" `   // 描述
	State       typePg.Int8         `json:"state" label:"状态:1启用;2禁用" ` // 状态:1启用;2禁用
	Idp         string              `json:"idp" label:"身份提供商" `        // 身份提供商
	Platform    string              `json:"platform" label:"平台" `        // 平台
	Protocol    string              `json:"protocol" label:"协议" `        // 协议
	CreateBy    string              `json:"createBy" label:"创建人" `     // 创建人
}

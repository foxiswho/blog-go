package modRamIdentityProvider

import (
	"github.com/hongmengzhu/xianfu-blog-go/pkg/model"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/tools/typePg"
)

type QueryPublicCt struct {
	model.BaseQueryNodeCt
	ID          typePg.Uint64String `json:"id" label:"" `
	Name        string              `json:"name" label:"名称" ` // 名称
	Code        string              `json:"code" label:"码值" `
	Description string              `json:"description" label:"描述" `   // 描述
	State       typePg.Int8         `json:"state" label:"状态:1启用;2禁用" ` // 状态:1启用;2禁用
	Platform    string              `json:"platform" label:"平台" `        // 平台
	Protocol    string              `json:"protocol" label:"协议" `        // 协议
	CreateBy    string              `json:"createBy" label:"创建人" `     // 创建人
}

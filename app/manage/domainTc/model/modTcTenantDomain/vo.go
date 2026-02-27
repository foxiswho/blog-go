package modTcTenantDomain

import (
	"github.com/foxiswho/blog-go/pkg/tools/typePg"
	"time"
)

type Vo struct {
	ID          typePg.Uint64String    `json:"id" label:"id" `
	No          string                 `json:"no" label:"编号代号"`
	Name        string                 `json:"name" label:"名称" `          // 名称
	NameFl      string                 `json:"nameFl" label:"名称外文" `      // 名称外文
	Code        string                 `json:"code" label:"编号代号" `        // 编号代号
	NameFull    string                 `json:"nameFull" label:"全称" `      // 全称
	State       typePg.Int8            `json:"state" label:"状态:1启用;2禁用" ` // 状态:1启用;2禁用
	Description string                 `json:"description" label:"描述" `   // 描述
	CreateAt    *time.Time             `json:"createAt" label:"创建时间" `    // 创建时间
	UpdateAt    *time.Time             `json:"updateAt" label:"更新时间" `    // 更新时间
	CreateBy    string                 `json:"createBy" label:"创建人" `     // 创建人
	UpdateBy    string                 `json:"updateBy" label:"更新人" `     // 更新人
	Ext         map[string]interface{} `json:"ext" label:"扩展字段" `
	TenantNo    string                 `json:"tenantNo" label:"租户编号" `
	TenantName  string                 `json:"tenantName" label:"租户编号" `
	Defaulted   typePg.Int8            `json:"defaulted" label:"是否默认:1是;2否" `
}

package modRamResource

import (
	"github.com/foxiswho/blog-go/pkg/model"
	"github.com/foxiswho/blog-go/pkg/tools/typePg"
	"time"
)

type QueryCt struct {
	model.BaseQueryCt
	ID          typePg.Uint64String `json:"id" label:"" `
	OrgId       string              `json:"orgId" label:"组织id" `  // 组织id
	Name        string              `json:"name" label:"名称" `     // 名称
	NameFl      string              `json:"nameFl" label:"名称外文" ` // 名称外文
	Code        string              `json:"code" label:"标志" `
	NameFull    string              `json:"nameFull" label:"全称" `      // 全称
	Description string              `json:"description" label:"描述" `   // 描述
	State       typePg.Int8         `json:"state" label:"状态:1启用;2禁用" ` // 状态:1启用;2禁用
	CreateBy    string              `json:"createBy" label:"创建人" `     // 创建人
	CreateAt    *time.Time          `json:"createAt" label:"创建时间" `    // 创建时间
	ALL         bool                `json:"all" label:"所有数据" `
	ParentId    string              `json:"parentId" label:"上级" `
	TypeAttr    string              `json:"typeAttr" label:"属性" `
}

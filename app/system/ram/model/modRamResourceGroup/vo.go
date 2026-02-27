package modRamResourceGroup

import (
	"github.com/foxiswho/blog-go/pkg/tools/typePg"
	"time"
)

type Vo struct {
	ID          typePg.Uint64String `json:"id" label:"id" `
	Name        string              `json:"name" label:"名称" `     // 名称
	NameFl      string              `json:"nameFl" label:"名称外文" ` // 名称外文
	Code        string              `json:"code" label:"标志" `
	NameFull    string              `json:"nameFull" label:"全称" `                                 // 全称
	State       typePg.Int8         `json:"state" label:"1有效2停用11取消(对应有效)12弃置(对应停用)13批量删除(无状态)" ` // 1有效2停用11取消(对应有效)12弃置(对应停用)13批量删除(无状态)
	Description string              `json:"description" label:"描述" `                              // 描述
	CreateAt    *time.Time          `json:"createAt" label:"创建时间" `                               // 创建时间
	UpdateAt    *time.Time          `json:"updateAt" label:"更新时间" `                               // 更新时间
	CreateBy    string              `json:"createBy" label:"创建人" `                                // 创建人
	UpdateBy    string              `json:"updateBy" label:"更新人" `                                // 更新人
	OrgId       string              `json:"orgId" label:"组织id" `                                  // 组织id
	ParentId    string              `json:"parentId" label:"上级" `
	TypeSys     string              `json:"typeSys" label:"类型" `
	TypeAttr    string              `json:"typeAttr" label:"属性" `
}

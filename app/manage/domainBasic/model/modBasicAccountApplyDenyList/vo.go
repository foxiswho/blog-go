package modBasicAccountApplyDenyList

import (
	"github.com/foxiswho/blog-go/pkg/tools/typePg"
	"time"
)

type Vo struct {
	ID          typePg.Uint64String `json:"id" label:"id" `
	Name        string              `json:"name" label:"名称" `
	State       typePg.Int8         `json:"state" label:"状态:1启用;2禁用" `
	Description string              `json:"description" label:"描述" `
	CreateAt    *time.Time          `json:"createAt" label:"创建时间" `
	UpdateAt    *time.Time          `json:"updateAt" label:"更新时间" `
	CreateBy    string              `json:"createBy" label:"创建人" `
	UpdateBy    string              `json:"updateBy" label:"更新人" `
	TypeSys     string              `json:"typeSys" label:"系统类型" `
	TypeDomain  string              `json:"typeDomain" label:"域类型" `
	TypeField   string              `json:"typeField" label:"字段类型" `
	TypeExpr    string              `json:"typeExpr" label:"表达式类型" `
	Expr        string              `json:"expr" label:"表达式" `
	Module      string              `json:"module" label:"模块" `
	Message     string              `json:"message" label:"错误时消息" `
}

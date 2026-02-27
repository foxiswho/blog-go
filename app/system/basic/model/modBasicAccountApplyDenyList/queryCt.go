package modBasicAccountApplyDenyList

import (
	"github.com/foxiswho/blog-go/pkg/model"
	"github.com/foxiswho/blog-go/pkg/tools/typePg"
	"time"
)

type QueryCt struct {
	model.BaseQueryCt
	ID          typePg.Uint64String `json:"id" label:"" `
	Name        string              `json:"name" label:"名称" `
	Description string              `json:"description" label:"描述" `
	State       typePg.Int8         `json:"state" label:"状态:1启用;2禁用" `
	CreateBy    string              `json:"createBy" label:"创建人" `
	CreateAt    *time.Time          `json:"createAt" label:"创建时间" `
	TypeSys     string              `json:"typeSys" label:"系统类型" `
	TypeDomain  string              `json:"typeDomain" label:"域类型" `
	TypeField   string              `json:"typeField" label:"字段类型" `
	TypeExpr    string              `json:"typeExpr" label:"表达式类型" `
	Expr        string              `json:"expr" label:"表达式" `
	Message     string              `json:"message" label:"错误时消息" `
}

package modBasicModelRules

import (
	"github.com/foxiswho/blog-go/pkg/tools/typePg"
)

type CreateUpdateCt struct {
	ID          typePg.Int64String `json:"id" label:"" `
	Name        string             `json:"name" binding:"required" label:"中文名称"`
	TypeSys     string             `json:"type_sys" label:"类型"`
	Description string             `json:"description" label:"描述"`
	//
	State        typePg.Int8 `json:"state" label:"状态" `
	RuleMode     string      `json:"ruleMode" label:"验证模式类型" `
	Coding       string      `json:"coding" label:"代码" `
	Condition    string      `json:"condition" label:"条件" `
	ErrorMessage string      `json:"errorMessage" label:"错误提示" `
	Structure    string      `json:"structure" label:"验证结构" `
	RuleTarget   []string    `json:"ruleTarget" label:"目标" `
	SharedRuleNo string      `json:"sharedRuleNo" label:"共享规则编号" `
	ValueNo      string      `json:"valueNo" label:"值编号/模块编号" `
}

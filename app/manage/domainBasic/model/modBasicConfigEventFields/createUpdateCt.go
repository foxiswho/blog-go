package modBasicConfigEventFields

import (
	"time"

	"github.com/foxiswho/blog-go/pkg/tools/typePg"
)

type CreateUpdateCt struct {
	ID       typePg.Int64String `json:"id" label:"id" `
	CreateAt *time.Time         `json:"createAt" label:"创建时间" `
	State    typePg.Int8        `json:"state" label:"状态" `
	Sort     typePg.Int64String ` json:"sort" comment:"排序" `
	//
	Description      string      `json:"description" label:"描述" ` // 描述
	TypeSys          string      `json:"typeSys" comment:"类型|普通|系统|api" `
	Name             string      `json:"name" label:"名称" `
	ParentField      string      `json:"parentField" comment:"父字段" `
	Field            string      `json:"field" label:"字段" `
	FieldPath        string      `json:"fieldPath" label:"字段路径" `
	Show             typePg.Int8 `json:"show" comment:"1显示2隐藏" `
	Independent      typePg.Int8 `json:"independent" comment:"独立字段" `
	Binary           typePg.Int8 `json:"binary" comment:"二进制" `
	Value            string      `json:"value" comment:"值" `
	DefaultValue     string      `json:"defaultValue" comment:"默认值" `
	FormCode         string      `json:"formCode" comment:"表单代码" `
	ValueType        string      `json:"valueType" comment:"值类型" `
	ValueAttr        string      `json:"valueAttr" comment:"值属性" `
	Rules            []string    `json:"rules" comment:"验证规则" `
	ParameterSource  string      `json:"parameterSource" comment:"参数来源" `
	ParameterContent string      `json:"parameterContent" comment:"参数内容" `
}

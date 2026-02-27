package modBasicDataDictionary

import "github.com/foxiswho/blog-go/pkg/tools/typePg"

type SelectNodeCt struct {
	TypeCode  string      `json:"typeCode" label:"码值" `
	By        string      `json:"by" label:"排序:默认排序正序" `
	FieldName string      `json:"fieldName" label:"字段名称" `
	State     typePg.Int8 `json:"state" label:"状态:1启用;2禁用" `
}

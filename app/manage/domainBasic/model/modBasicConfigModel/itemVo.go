package modBasicConfigModel

import "github.com/foxiswho/blog-go/pkg/tools/typePg"

type ItemVo struct {
	Name            string      `json:"name" label:"字段中文名称"`
	Field           string      `json:"field" label:"字段英文名称"`
	Show            typePg.Int8 `json:"show" label:"显示隐藏"`
	Binary          typePg.Int8 `json:"binary" label:"是否二进制"`
	DefaultValue    string      `json:"defaultValue" label:"默认值"`
	ValueType       string      `json:"valueType" label:"字段类型"`
	FormCode        string      `json:"formCode" label:"表单类型"`
	ParameterSource string      `json:"parameterSource" label:"参数源"`
	RuleMode        string      `json:"ruleMode" label:"验证模式类型"`
}

package modBasicConfigEvent

type CreateCt struct {
	Name         string    `json:"name" binding:"required" label:"中文名称"`
	Field        string    `json:"field" binding:"required" label:"英文标识"` // Event Identifier
	Model        string    `json:"model" binding:"required" label:"模型英文标识"`
	TypeSys      string    `json:"type_sys" label:"类型"`
	TypeCategory string    `json:"type_category" label:"类型种类"`
	ModuleSub    string    `json:"module_sub" label:"子模块选择"`
	Description  string    `json:"description" label:"描述"`
	Fields       []FieldCt `json:"fields" label:"字段列表"`
}

type FieldCt struct {
	Name            string `json:"name" label:"字段中文名称"`
	Field           string `json:"field" label:"字段英文名称"`
	Show            int8   `json:"show" label:"显示隐藏"`
	Binary          int8   `json:"binary" label:"是否二进制"`
	DefaultValue    string `json:"default_value" label:"默认值"`
	ValueType       string `json:"value_type" label:"字段类型"`
	FormCode        string `json:"form_code" label:"表单类型"`
	ParameterSource string `json:"parameter_source" label:"参数源"`
}

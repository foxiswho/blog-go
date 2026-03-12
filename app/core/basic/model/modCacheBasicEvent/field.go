package modCacheBasicEvent

type FieldCache struct {
	ID         int64  `json:"id" comment:"" `
	State      int8   `json:"state" comment:"1有效2停用" `
	Sort       int64  `json:"sort" comment:"排序" `
	TenantNo   string `json:"tenant_no" comment:"租户编号" `
	OrgNo      string `json:"org_no" comment:"组织编号" `
	StoreNo    string `json:"store_no" comment:"店铺编号" `
	MerchantNo string `json:"merchant_no" comment:"商户" `
	Owner      string `json:"owner" comment:"所属/拥有者" `
	No         string `json:"no" comment:"编号" `
	Name       string `json:"name" comment:"名称" `
	//
	TypeSys          string      `json:"type_sys" comment:"类型;普通;系统;" `
	SharedScope      []string    `json:"shared_scope" comment:"共享范围" `
	EventNo          string      `json:"event_no" comment:"事件编号" `
	ModelNo          string      `json:"model_no" comment:"模型编号" `
	Model            string      `json:"model" comment:"模型" `
	Module           string      `json:"module" comment:"模块" `
	ModuleSub        string      `json:"module_sub" comment:"子模块" `
	ParentField      string      `json:"parent_field" comment:"上级" `
	Field            string      `json:"field" comment:"字段名称" `
	FieldPath        string      `json:"field_path" comment:"路径字段名称" `
	Description      string      `json:"description" comment:"描述" `
	Show             int8        `json:"show" comment:"1显示2隐藏" `
	Independent      int8        `json:"independent" comment:"独立1是2否" `
	Binary           int8        `json:"binary" comment:"值二进制1是2否" `
	ExtraData        interface{} `json:"extraData" label:"额外数据" `
	Value            string      `json:"value" comment:"值" `
	ValueBinary      string      `json:"value_binary" comment:"值二进制" `
	DefaultValue     string      `json:"default_value" comment:"默认值" `
	FormCode         string      `json:"form_code" comment:"表单" `
	ValueType        string      `json:"value_type" comment:"字段值类型" `
	ValueAttr        string      `json:"value_attr" comment:"字段值属性|单值|对象" `
	FormAttr         interface{} `json:"form_attr" label:"表单属性" `
	Rules            []string    `json:"rules" comment:"表单验证规则" `
	ParameterSource  string      `json:"parameter_source" comment:"参数源" `
	ParameterContent string      `json:"parameter_content" comment:"参数内容" `
	ParameterConfig  string      `json:"parameter_config" comment:"参数配置" `
	KindUnique       string      `json:"kind_unique" comment:"字段种类唯一:model_no+field" `
	SharedFieldNo    string      `json:"shared_field_no" comment:"共享字段编号" `
}

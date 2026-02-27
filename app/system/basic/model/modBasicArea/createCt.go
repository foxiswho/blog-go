package modBasicArea

type CreateCt struct {
	Name        string `json:"name" form:"name" validate:"required,min=1,max=255" label:"名称" ` // 名称
	NameFl      string `json:"nameFl" label:"名称外文" `                                           // 名称外文
	Code        string `json:"code" label:"编号代号" `                                             // 编号代号
	NameFull    string `json:"nameFull" label:"全称" `                                           // 全称
	Description string `json:"description" label:"描述" `                                        // 描述
	OrgId       string `json:"orgId" label:"组织id" `                                            // 组织id
	Type        string `json:"type" label:"类型正常别名合并" `
	Source      string `json:"source" label:"源" `
	ZipCode     string `json:"zipCode" label:"邮编" `
	AreaCode    string `json:"areaCode" label:"区号" `
	ParentId    string `json:"parentId" label:"上级" `
	ParentNo    string `json:"parentNo" label:"上级编号" `
	CountryNo   string `json:"countryNo" label:"国家编号" `
}

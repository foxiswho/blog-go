package modRamApp

type CreateCt struct {
	Name        string `json:"name" form:"name" validate:"required,min=1,max=255" label:"名称" `
	NameFl      string `json:"nameFl" label:"名称外文" `
	No          string `json:"no" label:"编号代号"`
	Code        string `json:"code" label:"标志" `
	Type        string `json:"type" label:"类型" validate:"required,min=1,max=255"`
	Url         string `json:"url" label:"url" validate:"required,min=1,max=255"`
	NameFull    string `json:"nameFull" label:"全称" `
	Description string `json:"description" label:"描述" `
	CategoryNo  string `json:"categoryNo" label:"分类编号" `
}

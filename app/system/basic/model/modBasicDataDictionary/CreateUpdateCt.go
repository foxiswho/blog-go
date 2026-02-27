package modBasicDataDictionary

import "github.com/foxiswho/blog-go/pkg/tools/typePg"

type CreateUpdateCt struct {
	ID          typePg.Int64String `json:"id" form:"id" label:"" `
	Name        string             `json:"name" form:"name" validate:"required,min=1,max=255" label:"名称" ` // 名称
	NameFl      string             `json:"nameFl" label:"名称外文" `                                           // 名称外文
	Code        string             `json:"code"  validate:"required"  label:"代号" `                         // 编号代号
	NameFull    string             `json:"nameFull" label:"全称" `                                           // 全称
	Description string             `json:"description" label:"描述" `                                        // 描述
	Range       []string           `json:"range" label:"范围" `                                              // 范围

}

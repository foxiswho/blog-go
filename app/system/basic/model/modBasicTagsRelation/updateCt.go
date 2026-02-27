package modBasicTagsRelation

import (
	"github.com/foxiswho/blog-go/pkg/tools/typePg"
)

type UpdateCt struct {
	ID           typePg.Int64String     `json:"id" form:"id" validate:"required" label:"id" `
	OrgId        string                 `json:"orgId" label:"组织id" `
	Name         string                 `json:"name" form:"name" validate:"required,min=1,max=255" label:"名称" `
	NameFl       string                 `json:"nameFl" label:"名称外文" `
	No           string                 `json:"no" label:"标签" validate:"required,min=1,max=255"`
	Category     string                 `json:"category" label:"分类" validate:"required,min=1,max=255"`
	NameFull     string                 `json:"nameFull" label:"全称" `    // 全称
	Description  string                 `json:"description" label:"描述" ` // 描述
	AttributeMap map[string]interface{} `json:"attributeMap" label:"属性" `
	TypeSys      string                 `json:"typeSys" label:"类型" validate:"required,min=1,max=255"`
	NameShort    string                 `json:"nameShort" label:"名称简称" `
}

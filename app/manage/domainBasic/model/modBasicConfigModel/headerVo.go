package modBasicConfigModel

import "github.com/foxiswho/blog-go/pkg/tools/typePg"

type HeaderVo struct {
	Id            typePg.Uint64String `json:"id" form:"id" label:"id" `
	Name          string              `json:"name" binding:"required" label:"中文名称"`
	Model         string              `json:"model" binding:"required" label:"英文标识"`
	TypeSys       string              `json:"typeSys" label:"类型"`
	ModelCategory string              `json:"modelCategory" label:"模型种类"`
	ModuleSub     string              `json:"moduleSub" label:"子模块选择"`
	Description   string              `json:"description" label:"描述"`
}

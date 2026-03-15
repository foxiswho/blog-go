package modBasicConfigEvent

import (
	"time"

	"github.com/foxiswho/blog-go/pkg/tools/typePg"
)

type Vo struct {
	ID       typePg.Int64String `json:"id" label:"id" `
	CreateAt *time.Time         `json:"createAt" label:"创建时间" `
	UpdateAt *time.Time         `json:"updateAt" label:"更新时间" `
	CreateBy string             `json:"createBy" label:"创建人" `
	UpdateBy string             `json:"updateBy" label:"更新人" `
	State    typePg.Int8        `json:"state" label:"状态" `
	Sort     typePg.Int64String ` json:"sort" comment:"排序" `
	//
	Description   string      `json:"description" label:"描述" ` // 描述
	Name          string      `json:"name" label:"名称" `
	LangCode      string      `json:"langCode" comment:"语言" `
	TypeSys       string      `json:"typeSys" comment:"类型|普通|系统|api" `
	Module        string      `json:"module" comment:"模块" `
	Show          typePg.Int8 `json:"show" comment:"1显示2隐藏" `
	No            string      `json:"no" comment:"编号" `
	Model         string      `json:"model" binding:"required" label:"模型"`
	ModelNo       string      `json:"modelNo" binding:"required" label:"模型编号"`
	Field         string      `json:"field" label:"字段"`
	FieldSource   string      `json:"fieldSource" label:"字段来源/原始字段名称" `
	ModelCategory string      `json:"modelCategory" label:"模型种类"`
	ModuleSub     string      `json:"moduleSub" label:"子模块选择"`
}

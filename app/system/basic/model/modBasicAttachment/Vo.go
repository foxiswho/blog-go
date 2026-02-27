package modBasicAttachment

import (
	"github.com/foxiswho/blog-go/pkg/tools/typePg"
)

type Vo struct {
	ID          typePg.Uint64String `json:"id" label:"id" `
	Name        string              `json:"name" label:"名称" `
	State       typePg.Int8         `json:"state" label:"状态:1启用;2禁用" `
	CreateAt    *typePg.Time        `json:"createAt" label:"创建时间" `
	UpdateAt    *typePg.Time        `json:"updateAt" label:"更新时间" `
	CreateBy    string              `json:"createBy" label:"创建人" `
	UpdateBy    string              `json:"updateBy" label:"更新人" `
	Description string              `json:"description" label:"描述" `
	OrgId       string              `json:"orgId" label:"组织id" `
	SourceName  string              `json:"sourceName" label:"原始名称" `
	Url         string              `json:"url" label:"url" ` //全路径
	File        string              `json:"file" label:"相对路径" `
	Size        int64               `json:"size" label:"大小" `
	Module      string              `json:"module" label:"模块" `
	Value       string              `json:"value" label:"值id" `
	Tag         string              `json:"tag" label:"标签" `
	Label       string              `json:"label" label:"标记" `
	Domain      string              `json:"domain" label:"域名" `
	Type        string              `json:"type" label:"类型" `
	Ext         string              `json:"ext" label:"类型" `
}

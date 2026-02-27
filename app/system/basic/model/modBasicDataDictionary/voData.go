package modBasicDataDictionary

import (
	"github.com/foxiswho/blog-go/pkg/tools/typePg"
)

type VoData struct {
	ID           typePg.Int64String `json:"id" label:"" `
	Name         string             `json:"name" label:"名称" `        // 名称
	NameFl       string             `json:"nameFl" label:"名称外文" `    // 名称外文
	Code         string             `json:"code" label:"编号代号" `      // 编号代号
	NameFull     string             `json:"nameFull" label:"全称" `    // 全称
	Description  string             `json:"description" label:"描述" ` // 描述
	CreateAt     *typePg.Time       `json:"createAt" label:"创建时间" `  // 创建时间
	Value        string             `json:"value" label:"值内容" `      // 值内容
	TypeCode     string             `json:"typeCode" label:"码值" `
	TypeCodeName string             `json:"typeCodeName" label:"码值名称" `
	Sort         typePg.Int64String `json:"sort" label:"排序" `
	State        typePg.Int8        `json:"state" label:"状态:1启用;2禁用" ` // 状态:1启用;2禁用
}

package modBasicDataDictionary

import (
	"github.com/foxiswho/blog-go/pkg/tools/typePg"
)

type Vo struct {
	ID           typePg.Int64String `json:"id" label:"" `
	Name         string             `json:"name" label:"名称" `          // 名称
	NameFl       string             `json:"nameFl" label:"名称外文" `      // 名称外文
	Code         string             `json:"code" label:"编号代号" `        // 编号代号
	NameFull     string             `json:"nameFull" label:"全称" `      // 全称
	State        typePg.Int8        `json:"state" label:"状态:1启用;2禁用" ` // 状态:1启用;2禁用
	Description  string             `json:"description" label:"描述" `   // 描述
	CreateAt     *typePg.Time       `json:"createAt" label:"创建时间" `    // 创建时间
	CreateBy     typePg.Int64String `json:"CreateBy" label:"创建人" `     // 创建人
	Value        string             `json:"value" label:"值内容" `        // 值内容
	Range        []string           `json:"range" label:"范围" `         // 范围
	TypeCode     string             `json:"typeCode" label:"码值" `
	TypeCodeName string             `json:"typeCodeName" label:"码值名称" `
}

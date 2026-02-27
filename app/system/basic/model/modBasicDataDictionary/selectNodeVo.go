package modBasicDataDictionary

import (
	"github.com/foxiswho/blog-go/pkg/tools/typePg"
)

type SelectNodeVo struct {
	ID           typePg.Int64String `json:"id" label:"" `
	Name         string             `json:"name" label:"名称" `        // 名称
	NameFl       string             `json:"nameFl" label:"名称外文" `    // 名称外文
	Code         string             `json:"code" label:"编号代号" `      // 编号代号
	NameFull     string             `json:"nameFull" label:"全称" `    // 全称
	Description  string             `json:"description" label:"描述" ` // 描述
	CreateAt     *typePg.Time       `json:"createAt" label:"创建时间" `  // 创建时间
	Value        string             `json:"value" label:"值内容" `      // 值内容
	Range        []string           `json:"range" label:"范围" `       // 范围
	TypeCodeName string             `json:"typeCodeName" label:"类别名称" `
	Extend       string             `json:"extend" label:"扩展参数" `
}

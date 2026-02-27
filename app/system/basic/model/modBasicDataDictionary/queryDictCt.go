package modBasicDataDictionary

import (
	"github.com/foxiswho/blog-go/pkg/model"
	"github.com/foxiswho/blog-go/pkg/tools/typePg"
)

type QueryDictCt struct {
	model.BaseQueryCt
	Name     string      `json:"name" label:"名称" `     // 名称
	NameFl   string      `json:"nameFl" label:"名称外文" ` // 名称外文
	Code     string      `json:"code" label:"编号代号" `   // 编号代号
	TypeCode string      `json:"typeCode" label:"父级码值" `
	State    typePg.Int8 `json:"state" label:"状态:1启用;2禁用" ` // 状态:1启用;2禁用
}

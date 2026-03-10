package modBasicConfigEventFields

import (
	"github.com/foxiswho/blog-go/pkg/model"
	"github.com/foxiswho/blog-go/pkg/tools/typePg"
)

type QueryPublicCt struct {
	model.BaseQueryNodeCt
	Name    string      `json:"name" label:"名称" `          // 名称
	State   typePg.Int8 `json:"state" label:"状态:1启用;2禁用" ` // 状态:1启用;2禁用
	EventNo string      `json:"eventNo" label:"事件编号" `
}

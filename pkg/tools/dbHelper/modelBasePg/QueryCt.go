package modelBasePg

import (
	"github.com/foxiswho/blog-go/pkg/model"
	"github.com/foxiswho/blog-go/pkg/tools/typePg"
)

type QueryCt struct {
	model.BaseQueryCt
	State typePg.Int8 `json:"state" label:"状态:1启用;2禁用" `
}

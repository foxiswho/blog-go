package modBasicTags

import (
	"github.com/foxiswho/blog-go/pkg/model"
	"github.com/foxiswho/blog-go/pkg/tools/typePg"
	"time"
)

type QueryPublicCt struct {
	model.BaseQueryNodeCt
	ID          typePg.Uint64String `json:"id" label:"" `
	Name        string              `json:"name" label:"名称" ` // 名称
	No          string              `json:"no" label:"编号代号"`
	Description string              `json:"description" label:"描述" `   // 描述
	State       typePg.Int8         `json:"state" label:"状态:1启用;2禁用" ` // 状态:1启用;2禁用
	CreateAt    *time.Time          `json:"createAt" label:"创建时间" `    // 创建时间
	CategoryNo  string              `json:"categoryNo" label:"分类" `
}

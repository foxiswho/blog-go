package modBasicArea

import (
	"github.com/foxiswho/blog-go/pkg/model"
	"github.com/foxiswho/blog-go/pkg/tools/typePg"
	"time"
)

type QueryPublicCt struct {
	model.BaseQueryNodeCt
	ID          typePg.Uint64String `json:"id" label:"" `
	Name        string              `json:"name" label:"名称" `          // 名称
	Code        string              `json:"code" label:"编号代号" `        // 编号代号
	Description string              `json:"description" label:"描述" `   // 描述
	State       typePg.Int8         `json:"state" label:"状态:1启用;2禁用" ` // 状态:1启用;2禁用
	CreateAt    *time.Time          `json:"createAt" label:"创建时间" `    // 创建时间
	Type        string              `json:"type" label:"类型正常别名合并" `
	Source      string              `json:"source" label:"源" `
	ZipCode     string              `json:"zipCode" label:"邮编" `
	AreaCode    string              `json:"areaCode" label:"区号" `
	ParentNo    string              `json:"parentNo" label:"上级编号" `
	CountryNo   string              `json:"countryNo" label:"国家编号" `
}

package modBasicArea

import (
	"github.com/foxiswho/blog-go/pkg/tools/typePg"
	"time"
)

type Vo struct {
	ID          typePg.Uint64String `json:"id" label:"id" `
	Name        string              `json:"name" label:"名称" `          // 名称
	NameFl      string              `json:"nameFl" label:"名称外文" `      // 名称外文
	Code        string              `json:"code" label:"编号代号" `        // 编号代号
	NameFull    string              `json:"nameFull" label:"全称" `      // 全称
	State       typePg.Int8         `json:"state" label:"状态:1启用;2禁用" ` // 状态:1启用;2禁用
	Deleted     typePg.Int8         `json:"deleted" label:"删除:1是;2否" ` // 删除:1是;2否
	Description string              `json:"description" label:"描述" `   // 描述
	CreateAt    *time.Time          `json:"createAt" label:"创建时间" `    // 创建时间
	UpdateAt    *time.Time          `json:"updateAt" label:"更新时间" `    // 更新时间
	CreateBy    string              `json:"createBy" label:"创建人" `     // 创建人
	UpdateBy    string              `json:"updateBy" label:"更新人" `     // 更新人
	ParentId    string              `json:"parentId" label:"上级" `
	Type        string              `json:"type" label:"类型正常别名合并" `
	Source      string              `json:"source" label:"源" `
	ZipCode     string              `json:"zipCode" label:"邮编" `
	AreaCode    string              `json:"areaCode" label:"区号" `
	ParentNo    string              `json:"parentNo" label:"上级编号" `
	CountryNo   string              `json:"countryNo" label:"国家编号" `
}

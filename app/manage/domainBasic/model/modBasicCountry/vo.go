package modBasicCountry

import (
	"github.com/foxiswho/blog-go/pkg/tools/typePg"
	"time"
)

type Vo struct {
	ID           typePg.Int64String `json:"id" label:"id" `
	Name         string             `json:"name" label:"名称" `          // 名称
	NameFl       string             `json:"nameFl" label:"名称外文" `      // 名称外文
	Code         string             `json:"code" label:"编号代号" `        // 编号代号
	NameFull     string             `json:"nameFull" label:"全称" `      // 全称
	State        typePg.Int8        `json:"state" label:"状态:1启用;2禁用" ` // 状态:1启用;2禁用
	Description  string             `json:"description" label:"描述" `   // 描述
	CreateAt     *time.Time         `json:"createAt" label:"创建时间" `    // 创建时间
	UpdateAt     *time.Time         `json:"updateAt" label:"更新时间" `    // 更新时间
	CreateBy     string             `json:"createBy" label:"创建人" `     // 创建人
	UpdateBy     string             `json:"updateBy" label:"更新人" `     // 更新人
	Continent    string             `json:"continent" label:"所属洲" `
	ParentId     string             `json:"parentId" label:"上级" `
	ParentNo     string             `json:"parentNo" label:"上级编号" `
	Iso3         string             `json:"iso3" label:"ISO三字代码" `
	CountryCode  string             `json:"countryCode" label:"国际区号" `
	PhoneUse     typePg.Int8        `json:"phoneUse" label:"电话使用1是2否" `
	DomainSuffix string             `json:"domainSuffix" comment:"域名后缀" `
}

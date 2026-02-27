package modBasicCountry

import (
	"github.com/foxiswho/blog-go/pkg/tools/typePg"
)

type CreateCt struct {
	Name         string      `json:"name" form:"name" validate:"required,min=1,max=255" label:"名称" ` // 名称
	NameFl       string      `json:"nameFl" label:"名称外文" `                                           // 名称外文
	Code         string      `json:"code" label:"编号代号" `                                             // 编号代号
	NameFull     string      `json:"nameFull" label:"全称" `                                           // 全称
	Description  string      `json:"description" label:"描述" `                                        // 描述
	Continent    string      `json:"continent" label:"所属洲" `
	ParentId     string      `json:"parentId" label:"上级" `
	ParentNo     string      `json:"parentNo" label:"上级编号" `
	Iso3         string      `json:"iso3" label:"ISO三字代码" `
	CountryCode  string      `json:"countryCode" label:"国际区号" `
	PhoneUse     typePg.Int8 `json:"phoneUse" label:"电话使用1是2否" `
	DomainSuffix string      `json:"domainSuffix" comment:"域名后缀" `
}

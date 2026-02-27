package holderApiPg

import (
	_ "github.com/foxiswho/blog-go/pkg/interfaces"
)

type DiplHolder struct {
	No       string `json:"no" comment:"编码"`
	Code     string `json:"code" comment:"编码" `
	Name     string `json:"name" label:"名称" `
	TenantNo string `json:"tenantNo" comment:"租户编码"`
}

package appDomainPg

import (
	"github.com/foxiswho/blog-go/pkg/consts/constsRam/typeDomainPg"
)

// 域模式
type AppDomain string

const (
	Tenant      AppDomain = "tenant"
	Manage      AppDomain = "manage"
	System      AppDomain = "system"
	WEB         AppDomain = "WEB"
	ManageOwner AppDomain = "manageOwner"
)

func (c AppDomain) ToTypeDomain() typeDomainPg.TypeDomain {
	switch c {
	case System:
		return typeDomainPg.System
	case Manage:
		return typeDomainPg.Merchant
	case Tenant:
		return typeDomainPg.Tenant
	case WEB:
		return typeDomainPg.General
	default:
		return typeDomainPg.General
	}
}

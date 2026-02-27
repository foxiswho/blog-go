package domainTc

import (
	"github.com/foxiswho/blog-go/app/manage/domainTc/model/cacheTc"
	"github.com/foxiswho/blog-go/app/manage/domainTc/model/dtoTenantDomain"
	"github.com/foxiswho/blog-go/pkg/tools/mapPg"
	"github.com/go-spring/spring-core/gs"
)

func init() {
	gs.Object(&cacheTc.TenantDomainCache{Domain: &mapPg.SafeMap[string, string]{}, DomainData: &mapPg.SafeMap[string, dtoTenantDomain.Dto]{}})
}

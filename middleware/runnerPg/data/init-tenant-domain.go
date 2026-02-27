package data

import (
	"context"

	"github.com/foxiswho/blog-go/app/manage/domainTc/service"
	"github.com/foxiswho/blog-go/pkg/log2"
	syslog "github.com/go-spring/log"
)

// InitTenantDomain
// @Description: 初始化租户域名
type InitTenantDomain struct {
	log    *log2.Logger                        `autowire:"?"`
	domain *service.TcTenantDomainCacheService `autowire:"?"`
}

func (b *InitTenantDomain) Run() error {
	syslog.Infof(context.Background(), syslog.TagAppDef, "初始化 => 域名与租户的关系")
	b.domain.InitTenantDomain()
	return nil
}

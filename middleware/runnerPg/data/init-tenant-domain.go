package data

import (
	"context"

	"github.com/foxiswho/blog-go/app/manage/domainTc/service"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/go-spring/log"
	_ "github.com/go-spring/spring-core/gs"
)

// InitTenantDomain
// @Description: 初始化租户域名
type InitTenantDomain struct {
	log    *log2.Logger                        `autowire:"?"`
	domain *service.TcTenantDomainCacheService `autowire:"?"`
}

func (b *InitTenantDomain) Run(ctx context.Context) error {
	log.Infof(context.Background(), log.TagAppDef, "初始化 => 域名与租户的关系")
	b.domain.InitTenantDomain(context.Background())
	return nil
}

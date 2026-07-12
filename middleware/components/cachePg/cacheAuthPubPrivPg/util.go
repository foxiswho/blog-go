package cacheAuthPubPrivPg

import (
	"github.com/hongmengzhu/xianfu-blog-go/pkg/consts/constsRam/typeDomainPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/enum/state/clientPg"
)

const (
	// 缓存前缀
	CachePrefix = "cac:auth:pubprive:"
)

func KeySystem() string {
	return typeDomainPg.System.Index()
}

func KeyManage(tenantNo string) string {
	return typeDomainPg.Manage.Index() + ":" + tenantNo
}

func Key(tenantNo string, typeDomain typeDomainPg.TypeDomain, client clientPg.Client) string {
	return CachePrefix + tenantNo + ":" + typeDomain.Index() + ":" + client.Index()
}

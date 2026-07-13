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
	return Key(typeDomainPg.System.Index(), typeDomainPg.System, clientPg.Browser)
}

func KeyManage(tenantNo string) string {
	return Key(tenantNo, typeDomainPg.Manage, clientPg.Browser)
}

func Key(tenantNo string, typeDomain typeDomainPg.TypeDomain, client clientPg.Client) string {
	return CachePrefix + tenantNo + ":" + typeDomain.Index() + ":" + client.Index()
}

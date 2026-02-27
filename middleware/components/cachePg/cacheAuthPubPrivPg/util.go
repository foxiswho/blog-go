package cacheAuthPubPrivPg

import (
	"github.com/foxiswho/blog-go/pkg/consts/constsRam/typeDomainPg"
)

func KeySystem() string {
	return typeDomainPg.System.Index()
}

func KeyManage(tenantNo string) string {
	return typeDomainPg.Manage.Index() + ":" + tenantNo
}

func KeyCustomer(tenantNo string) string {
	return typeDomainPg.Customer.Index() + ":" + tenantNo
}

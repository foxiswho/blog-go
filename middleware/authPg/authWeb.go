package authPg

import (
	"context"

	"github.com/foxiswho/blog-go/app/manage/domainTc/model/cacheTc"
	"github.com/foxiswho/blog-go/pkg/configPg"
	"github.com/foxiswho/blog-go/pkg/consts/constHeaderPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/gin-gonic/gin"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
)

func init() {
	gs.Object(&GroupWebMiddlewareSp{})
}

// 中间件 服务
type GroupWebMiddlewareSp struct {
	log    *log2.Logger               `autowire:"?"`
	domain *cacheTc.TenantDomainCache `autowire:"?"`
	pg     configPg.Pg                `value:"${pg}"`
	server configPg.Server            `value:"${server}"`
}

// 权限验证 中间件
func GroupWebMiddleware(m *GroupWebMiddlewareSp) gin.HandlerFunc {
	return func(c *gin.Context) {
		val := "-1"
		// 获取租户
		load, b := m.domain.Domain.Load(c.Request.Host)
		if b {
			val = load
		}
		//本地域名
		if m.domain.IsLocalHostExist(c.Request.Host) {
			val = ""
		}
		// 指定域名
		if m.domain.IsServerHostExist(c.Request.Host, m.server) {
			val = ""
		}
		syslog.Infof(context.Background(), syslog.TagAppDef, "租户.no => %+v", val)
		c.Set(constHeaderPg.WebTenantNo, val)
		c.Next()
	}
}

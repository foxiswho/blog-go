package authPg

import (
	"strings"

	"github.com/foxiswho/blog-go/app/manage/domainTc/model/cacheTc"
	"github.com/foxiswho/blog-go/middleware/components/cachePg/cacheDiplPg"
	"github.com/foxiswho/blog-go/pkg/configPg"
	"github.com/foxiswho/blog-go/pkg/consts/constContextPg"
	"github.com/foxiswho/blog-go/pkg/consts/constHeaderPg"
	"github.com/foxiswho/blog-go/pkg/holderPg/holderApiPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/gin-gonic/gin"
	"github.com/go-spring/spring-core/gs"
	"github.com/pangu-2/go-tools/tools/strPg"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
)

func init() {
	gs.Object(&GroupApiMiddlewareSp{})
}

// 中间件 服务
type GroupApiMiddlewareSp struct {
	log    *log2.Logger               `autowire:"?"`
	domain *cacheTc.TenantDomainCache `autowire:"?"`
	pg     configPg.Pg                `value:"${pg}"`
	server configPg.Server            `value:"${server}"`
}

// 权限验证 中间件
func GroupApiMiddleware(m *GroupApiMiddlewareSp) gin.HandlerFunc {
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
		//fmt.Printf("租户.no => %+v\n", val)
		//syslog.Infof(context.Background(), syslog.TagAppDef, "租户.no => %+v", val)
		header := c.GetHeader(constHeaderPg.HeaderAuthorization)
		if strPg.IsNotBlank(header) {
			split := strings.Split(header, ":")
			//fmt.Printf("Dipl.no => %+v\n", split[0])
			//fmt.Printf("Dipl.no => %+v\n", len(split))
			if len(split) >= 2 && strPg.IsNotBlank(split[0]) {
				get, b2 := cacheDiplPg.Get(split[0])
				//fmt.Printf("cacheDiplPg.Get.b2 => %+v\n", b2)
				//fmt.Printf("cacheDiplPg.Get.get => %+v\n", get)
				//fmt.Printf("cacheDiplPg.HashShaVerify => %+v\n", cacheDiplPg.HashShaVerify(get.Key, get.Secret, split[1]))
				//验证
				if b2 && cacheDiplPg.HashShaVerify(get.Key, get.Secret, split[1]) {
					val = get.TenantNo
					pg := holderApiPg.HolderPg{}
					pg.HolderData = holderApiPg.DiplHolder{
						Name:     get.Name,
						No:       get.No,
						TenantNo: get.TenantNo,
					}
					c.Set(constContextPg.AUTH_LOGIN_API, pg)
				}
				c.Set(constHeaderPg.WebTenantNo, val)
				c.Next()
				return
			}
		}
		c.JSON(200, rg.Error[string](constHeaderPg.HeaderAuthorization+" 参数不能为空"))
		c.Abort()
		return
	}
}

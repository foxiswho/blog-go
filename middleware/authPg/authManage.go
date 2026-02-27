package authPg

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/foxiswho/blog-go/app/manage/domainRam/service"
	"github.com/foxiswho/blog-go/middleware/components/authTokenPg"
	"github.com/foxiswho/blog-go/middleware/components/cachePg/cacheAuthPubPrivPg"
	"github.com/foxiswho/blog-go/pkg/configPg"
	"github.com/foxiswho/blog-go/pkg/consts/constContextPg"
	"github.com/foxiswho/blog-go/pkg/consts/constHeaderPg"
	"github.com/foxiswho/blog-go/pkg/holderPg/multiTenantPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/gin-gonic/gin"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
	"github.com/pangu-2/go-tools/tools/convPg"
	"github.com/pangu-2/go-tools/tools/strPg"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
)

// 验证 ownerCode 必填
var managerUrlPaths = make([]string, 0)

func init() {
	gs.Object(&GroupManageMiddlewareSp{})
	//商品模块
	managerUrlPaths = append(managerUrlPaths, "manage/gc")
}

// 中间件 服务
type GroupManageMiddlewareSp struct {
	sv  *service.RamAccountMiddlewareService `autowire:"?"`
	pg  configPg.Pg                          `value:"${pg}"`
	log *log2.Logger                         `autowire:"?"`
}

// 权限验证 中间件
func GroupManageMiddleware(m *GroupManageMiddlewareSp) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(constHeaderPg.HeaderAuthorization)
		if !strPg.IsBlank(token) {
			var payload map[string]interface{}
			unverified, b := authTokenPg.ParseUnverified(token)
			if b {
				// 解析为map便于查看
				if err := json.Unmarshal(unverified, &payload); err != nil {
					c.JSON(200, rg.Error[string]("解析载荷JSON失败"))
					c.Abort()
					return
				}
			}
			m.log.Debugf("[中间件].unverified= %+v", string(unverified))
			tenantNo := ""
			{
				//获取租户信息
				if tmp, ok := payload[authTokenPg.TenantNo]; ok {
					tenantNo = convPg.ObjToStr(tmp)
				}
				m.log.Debugf("[中间件].payload= %+v", payload)
			}
			m.log.Debugf("[中间件].tenantNo= %+v", tenantNo)
			get, b := cacheAuthPubPrivPg.Get(cacheAuthPubPrivPg.KeyManage(tenantNo))
			if !b {
				c.JSON(200, rg.Error[string]("密钥不存在"))
				c.Abort()
				return
			}
			token = strings.Replace(token, authTokenPg.AuthScheme+" ", "", -1)
			t, r := authTokenPg.VerifyByPublicKey(get.Public, token)
			if r.ErrorIs() {
				syslog.Debugf(context.Background(), syslog.TagAppDef, "JWT= %+v", r)
				c.JSON(200, r)
				c.Abort()
				return
			}
			//获取登录 编号
			loginNo, err2 := t.GetSubject()
			if nil != err2 {
				c.JSON(200, rg.Error[string]("解析失败"))
				c.Abort()
				return
			}
			if strPg.IsBlank(loginNo) {
				c.JSON(200, rg.Error[string]("解析失败"))
				c.Abort()
				return
			}
			//获取登录信息
			rt := m.sv.FindByLoginNo(c, loginNo, tenantNo)
			if rt.ErrorIs() {
				c.JSON(200, rt)
				c.Abort()
				return
			}
			m.log.Debugf("rt= %+v", rt)
			m.log.Debugf("rt.Rule= %+v", rt.Data.Rule)
			//
			var data = rt.Data
			//如果配置了managerUrlPaths，则必须携带ownerCode
			if len(managerUrlPaths) > 0 && data.GetOwnerNo() == "" {
				path := c.Request.URL.Path
				for _, item := range managerUrlPaths {
					//必须登录 验证，url 包含的路径，那么提示错误
					if "" != item && strings.Index(path, item) != -1 {
						m.log.Tracef("pass %+v", c.Request.URL.Path)
						c.JSON(200, rg.Error[string]("所有者错误"))
						c.Abort()
						return
					}
				}
			}
			c.Set(constContextPg.CTX_MULITI_TENANT, multiTenantPg.GetMultiTableManage())

			//
			c.Set(constContextPg.AUTH_LOGIN, rt.Data)
			c.Set(constContextPg.CTX_RULE, rt.Data.Rule)
			//value, exists := c.Get(constContextPg.CTX_RULE)
			//m.log.Infof("Set.constContextPg.CTX_RULE.get= %+v,", value, exists)
			//m.log.WithContext(c.Context()).Tracef("::AdminManageFilter")

			//详细权限
			//err = baseAuth.CasbinHandler(ctx, rt.Data, m.pg, m.casbin)
			//if nil != err {
			//	c.JSON(r.Error(err.Error()))
			//	return
			//}
			c.Next()
			return
		}
		c.JSON(200, rg.Error[string](constHeaderPg.HeaderAuthorization+" 参数不能为空"))
		c.Abort()
		return
	}
}

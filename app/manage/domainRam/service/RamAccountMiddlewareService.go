package service

import (
	"context"

	"github.com/foxiswho/blog-go/infrastructure/repositoryRam"
	"github.com/foxiswho/blog-go/pkg/consts/constsRam/typeDomainPg"
	"github.com/foxiswho/blog-go/pkg/holderPg"
	"github.com/foxiswho/blog-go/pkg/holderPg/multiTenantPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/gin-gonic/gin"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
	"github.com/pangu-2/go-tools/tools/strPg"

	"reflect"

	"github.com/jinzhu/copier"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
)

func init() {
	gs.Provide(NewRamAccountMiddlewareService).Init(func(s *RamAccountMiddlewareService) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

// RamAccountMiddlewareService 账户公共动作
// @Description:
type RamAccountMiddlewareService struct {
	sv    *repositoryRam.RamAccountRepository              `autowire:"?"`
	aAuth *repositoryRam.RamAccountAuthorizationRepository `autowire:"?"`
	log   *log2.Logger                                     `autowire:"?"`
}

func NewRamAccountMiddlewareService() *RamAccountMiddlewareService {
	return new(RamAccountMiddlewareService)
}

// FindByLoginNo
//
//	@Description: 登录信息
//	@receiver c
//	@param jwt
//	@return rt
func (c *RamAccountMiddlewareService) FindByLoginNo(ctx *gin.Context, no, tenantNo string) (rt rg.Rs[holderPg.HolderPg]) {
	c.log.Debugf("jwt=%+v", no)
	if strPg.IsBlank(no) {
		return rt.ErrorMessage("账号登陆失败")
	}
	info, b := c.sv.FindByNo(no)
	if !b {
		return rt.ErrorMessage("账号不存在")
	}
	if tenantNo != info.TenantNo {
		return rt.ErrorMessage("账号不存在")
	}
	pg := rt.Data
	rule := multiTenantPg.NewMultiRuleDefaultBySystem()
	//
	rule.Tenant = true
	rule.Merchant = false
	rule.Org = false
	rule.Owner = false
	rule.MultiOwner = false
	//
	var accountHolder2 holderPg.AccountHolder
	copier.Copy(&accountHolder2, &info)
	//
	accountHolder2.Os = holderPg.NewAccountHolderOs()
	accountHolder2.Os.Departments = info.Os.Data().Departments
	accountHolder2.Os.Roles = info.Os.Data().Roles
	accountHolder2.Os.Orgs = info.Os.Data().Orgs
	accountHolder2.Os.Tenants = info.Os.Data().Tenants
	accountHolder2.Os.Teams = info.Os.Data().Teams
	accountHolder2.Os.Shops = info.Os.Data().Shops
	accountHolder2.Os.Stores = info.Os.Data().Stores
	//
	pg.MultiTenant = multiTenantPg.MultiTenantPg{
		TenantNo: accountHolder2.Os.Tenants,
	}
	pg.HolderData = accountHolder2
	pg.TypeDomain = typeDomainPg.Manage.Index()
	pg.Rule = &rule
	rt.Data = pg
	return rt.Ok()
}

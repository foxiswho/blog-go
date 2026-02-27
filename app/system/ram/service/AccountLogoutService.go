package service

import (
	"context"
	"reflect"

	"github.com/foxiswho/blog-go/infrastructure/repositoryRam"
	"github.com/foxiswho/blog-go/pkg/configPg"
	"github.com/foxiswho/blog-go/pkg/holderPg"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
)

func init() {
	gs.Provide(NewAccountLogoutService).Init(func(s *AccountLogoutService) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

// AccountLogoutService 退出
// @Description:
type AccountLogoutService struct {
	dao *repositoryRam.RamAccountRepository `autowire:"?"`
	pg  configPg.Pg                         `value:"${pg}"`
}

func NewAccountLogoutService() *AccountLogoutService {
	return new(AccountLogoutService)
}

func (c *AccountLogoutService) Logout(holder holderPg.HolderPg) (rt rg.Rs[string]) {
	return rt.Ok()
}

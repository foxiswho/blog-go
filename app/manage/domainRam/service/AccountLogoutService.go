package service

import (
	"context"
	"reflect"

	"github.com/foxiswho/blog-go/infrastructure/repositoryRam"
	"github.com/foxiswho/blog-go/pkg/configPg"
	"github.com/foxiswho/blog-go/pkg/holderPg"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
	"go-spring.org/log"
	"go-spring.org/spring/gs"
)

func init() {
	gs.Provide(NewAccountLogoutService).Init(func(s *AccountLogoutService) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
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

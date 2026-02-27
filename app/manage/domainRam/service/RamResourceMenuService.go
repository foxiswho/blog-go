package service

import (
	"context"

	"github.com/foxiswho/blog-go/app/manage/domainRam/model/modRamResourceMenu"
	"github.com/foxiswho/blog-go/app/manage/domainRam/service/ramMenu"
	"github.com/foxiswho/blog-go/infrastructure/repositoryRam"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/gin-gonic/gin"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"

	"reflect"

	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
)

func init() {
	gs.Provide(new(RamResourceMenuService)).Init(func(s *RamResourceMenuService) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

// RamResourceMenuService 资源菜单关系
// @Description:
type RamResourceMenuService struct {
	sv             *repositoryRam.RamResourceMenuRepository      `autowire:"?"`
	resDb          *repositoryRam.RamResourceRepository          `autowire:"?"`
	groupDb        *repositoryRam.RamResourceGroupRepository     `autowire:"?"`
	resAuthDb      *repositoryRam.RamResourceAuthorityRepository `autowire:"?"`
	menuDb         *repositoryRam.RamMenuRepository              `autowire:"?"`
	menuRelationDb *repositoryRam.RamMenuRelationRepository      `autowire:"?"`
	log            *log2.Logger                                  `autowire:"?"`
}

// UpdateByMenu 更新资源菜单关系
//
//	@Description:
//	@receiver c
//	@param ct
//	@return rt
func (c *RamResourceMenuService) UpdateByMenu(ctx *gin.Context, ct modRamResourceMenu.UpdateByMenuCt) (rt rg.Rs[string]) {
	c.log.Infof("ct=%+v", ct)
	return ramMenu.NewUpdateByResource(c.log, c.menuDb, c.menuRelationDb, c.resAuthDb, c.sv, c.groupDb, c.resDb, ct, ctx).Process()
}

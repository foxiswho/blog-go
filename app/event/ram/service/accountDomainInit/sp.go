package accountDomainInit

import (
	"github.com/hongmengzhu/xianfu-blog-go/infrastructure/repositoryRam"
	"github.com/hongmengzhu/xianfu-blog-go/infrastructure/repositoryTc"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/configPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/log2"
	"go-spring.org/spring/gs"
)

func init() {
	gs.Provide(new(Sp))
}

type Sp struct {
	log      *log2.Logger                                     `autowire:"?"`
	accDb    *repositoryRam.RamAccountRepository              `autowire:"?"`
	authDb   *repositoryRam.RamAccountAuthorizationRepository `autowire:"?"`
	depDb    *repositoryRam.RamDepartmentRepository           `autowire:"?"`
	roleDb   *repositoryRam.RamRoleRepository                 `autowire:"?"`
	teamDb   *repositoryRam.RamTeamRepository                 `autowire:"?"`
	groupDb  *repositoryRam.RamGroupRepository                `autowire:"?"`
	levelDb  *repositoryRam.RamLevelRepository                `autowire:"?"`
	tenantDb *repositoryTc.TcTenantRepository                 `autowire:"?"`
	pg       configPg.Pg                                      `value:"${pg}"`
}

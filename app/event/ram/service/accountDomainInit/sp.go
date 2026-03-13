package accountDomainInit

import (
	"github.com/foxiswho/blog-go/infrastructure/repositoryRam"
	"github.com/foxiswho/blog-go/infrastructure/repositoryTc"
	"github.com/foxiswho/blog-go/pkg/configPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/go-spring/spring-core/gs"
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

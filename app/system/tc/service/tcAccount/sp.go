package tcAccount

import (
	"github.com/foxiswho/blog-go/infrastructure/repositoryRam"
	"github.com/foxiswho/blog-go/infrastructure/repositoryTc"
	"github.com/go-spring/spring-core/gs"
)

func init() {
	gs.Provide(new(Sp))
}

type Sp struct {
	accDb   *repositoryRam.RamAccountRepository              `autowire:"?"`
	authDb  *repositoryRam.RamAccountAuthorizationRepository `autowire:"?"`
	depDb   *repositoryRam.RamDepartmentRepository           `autowire:"?"`
	roleDb  *repositoryRam.RamRoleRepository                 `autowire:"?"`
	teamDb  *repositoryRam.RamTeamRepository                 `autowire:"?"`
	groupDb *repositoryRam.RamGroupRepository                `autowire:"?"`
	levelDb *repositoryRam.RamLevelRepository                `autowire:"?"`
	tenDb   *repositoryTc.TcTenantRepository                 `autowire:"?"`
}

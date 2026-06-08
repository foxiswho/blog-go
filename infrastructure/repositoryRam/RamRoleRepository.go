package repositoryRam

import (
	"context"

	"github.com/foxiswho/blog-go/infrastructure/entityRam"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/support"
	"github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"

	"reflect"
)

func init() {
	gs.Provide(new(RamRoleRepository)).Init(func(s *RamRoleRepository) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[RamRoleRepository])).Init(func(s *support.BaseService[RamRoleRepository]) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type RamRoleRepository struct {
	repositoryPg.BaseRepository[entityRam.RamRoleEntity, int64]
	//
}

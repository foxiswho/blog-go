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
	gs.Provide(new(RamMenuRepository)).Init(func(s *RamMenuRepository) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[RamMenuRepository])).Init(func(s *support.BaseService[RamMenuRepository]) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type RamMenuRepository struct {
	repositoryPg.BaseRepository[entityRam.RamMenuEntity, int64]
	//
}

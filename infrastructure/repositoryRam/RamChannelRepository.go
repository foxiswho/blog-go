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
	gs.Provide(new(RamChannelRepository)).Init(func(s *RamChannelRepository) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[RamChannelRepository])).Init(func(s *support.BaseService[RamChannelRepository]) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type RamChannelRepository struct {
	repositoryPg.BaseRepository[entityRam.RamChannelEntity, int64]
}

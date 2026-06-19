package repositoryRam

import (
	"context"

	"reflect"

	"github.com/foxiswho/blog-go/infrastructure/entityRam"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/support"
	"go-spring.org/log"
	"go-spring.org/spring/gs"
)

func init() {
	gs.Provide(new(RamFavoritesRepository)).Init(func(s *RamFavoritesRepository) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[RamFavoritesRepository])).Init(func(s *support.BaseService[RamFavoritesRepository]) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type RamFavoritesRepository struct {
	repositoryPg.BaseRepository[entityRam.RamFavoritesEntity, int64]
}

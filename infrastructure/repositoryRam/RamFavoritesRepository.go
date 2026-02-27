package repositoryRam

import (
	"context"

	"github.com/foxiswho/blog-go/infrastructure/entityRam"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/support"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
	"gorm.io/gorm"

	"reflect"
)

func init() {
	gs.Provide(new(RamFavoritesRepository)).Init(func(s *RamFavoritesRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[RamFavoritesRepository])).Init(func(s *support.BaseService[RamFavoritesRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type RamFavoritesRepository struct {
	repositoryPg.BaseRepository[entityRam.RamFavoritesEntity, int64]
	db *gorm.DB `autowire:"?"`
}

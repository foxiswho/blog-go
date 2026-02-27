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
	gs.Provide(new(RamGroupRepository)).Init(func(s *RamGroupRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[RamGroupRepository])).Init(func(s *support.BaseService[RamGroupRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type RamGroupRepository struct {
	repositoryPg.BaseRepository[entityRam.RamGroupEntity, int64]
	db *gorm.DB `autowire:"?"`
}

package repositoryRam

import (
	"context"
	"reflect"

	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/support"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
)

func init() {
	gs.Provide(new(RamCasbinRepository)).Init(func(s *RamCasbinRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[RamCasbinRepository])).Init(func(s *support.BaseService[RamCasbinRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type RamCasbinRepository struct {
	repositoryPg.BaseRepository[gormadapter.CasbinRule, uint]
	//
}

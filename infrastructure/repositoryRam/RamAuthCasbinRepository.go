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
	gs.Provide(new(RamAuthCasbinRepository)).Init(func(s *RamAuthCasbinRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[RamAuthCasbinRepository])).Init(func(s *support.BaseService[RamAuthCasbinRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type RamAuthCasbinRepository struct {
	repositoryPg.BaseRepository[gormadapter.CasbinRule, uint]
	//
}

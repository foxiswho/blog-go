package repositoryRam

import (
	"context"
	"reflect"

	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/support"
	"go-spring.org/log"
	"go-spring.org/spring/gs"
)

func init() {
	gs.Provide(new(RamAuthCasbinRepository)).Init(func(s *RamAuthCasbinRepository) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[RamAuthCasbinRepository])).Init(func(s *support.BaseService[RamAuthCasbinRepository]) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type RamAuthCasbinRepository struct {
	repositoryPg.BaseRepository[gormadapter.CasbinRule, uint]
	//
}

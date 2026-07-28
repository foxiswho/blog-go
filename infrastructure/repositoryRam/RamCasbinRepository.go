package repositoryRam

import (
	"context"
	"reflect"

	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/tools/dbHelper/support"
	"go-spring.org/log"
	"go-spring.org/spring/gs"
)

func init() {
	gs.Provide(new(RamCasbinRepository)).Init(func(s *RamCasbinRepository) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[RamCasbinRepository])).Init(func(s *support.BaseService[RamCasbinRepository]) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type RamCasbinRepository struct {
	repositoryPg.BaseRepository[gormadapter.CasbinRule, uint]
	//
}

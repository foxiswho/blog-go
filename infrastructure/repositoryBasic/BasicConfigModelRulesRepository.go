package repositoryBasic

import (
	"context"
	"reflect"

	"github.com/foxiswho/blog-go/infrastructure/entityBasic"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/support"
	"github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
)

func init() {
	gs.Provide(new(BasicConfigModelRulesRepository)).Init(func(s *BasicConfigModelRulesRepository) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[BasicConfigModelRulesRepository])).Init(func(s *support.BaseService[BasicConfigModelRulesRepository]) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type BasicConfigModelRulesRepository struct {
	repositoryPg.BaseRepository[entityBasic.BasicConfigModelRulesEntity, int64]
}

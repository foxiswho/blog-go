package repositoryRam

import (
	"context"

	"github.com/hongmengzhu/xianfu-blog-go/infrastructure/entityRam"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/tools/dbHelper/support"
	"go-spring.org/log"
	"go-spring.org/spring/gs"

	"reflect"
)

func init() {
	gs.Provide(new(RamIdentitySourceCallbackRepository)).Init(func(s *RamIdentitySourceCallbackRepository) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[RamIdentitySourceCallbackRepository])).Init(func(s *support.BaseService[RamIdentitySourceCallbackRepository]) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type RamIdentitySourceCallbackRepository struct {
	repositoryPg.BaseRepository[entityRam.RamIdentitySourceCallbackEntity, int64]
	//
}

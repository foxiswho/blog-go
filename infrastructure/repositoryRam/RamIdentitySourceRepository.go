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
	gs.Provide(new(RamIdentitySourceRepository)).Init(func(s *RamIdentitySourceRepository) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[RamIdentitySourceRepository])).Init(func(s *support.BaseService[RamIdentitySourceRepository]) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type RamIdentitySourceRepository struct {
	repositoryPg.BaseRepository[entityRam.RamIdentitySourceEntity, int64]
	//
}

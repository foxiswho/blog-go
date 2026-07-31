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
	gs.Provide(new(RamIdentityProviderRepository)).Init(func(s *RamIdentityProviderRepository) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[RamIdentityProviderRepository])).Init(func(s *support.BaseService[RamIdentityProviderRepository]) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type RamIdentityProviderRepository struct {
	repositoryPg.BaseRepository[entityRam.RamIdentityProviderEntity, int64]
	//
}

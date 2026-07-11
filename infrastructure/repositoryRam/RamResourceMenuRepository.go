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
	gs.Provide(new(RamResourceMenuRepository)).Init(func(s *RamResourceMenuRepository) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[RamResourceMenuRepository])).Init(func(s *support.BaseService[RamResourceMenuRepository]) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type RamResourceMenuRepository struct {
	repositoryPg.BaseRepository[entityRam.RamResourceMenuEntity, int64]
	//
}

func (c *RamResourceMenuRepository) DeleteByMenuId(ctx context.Context, code int64) {
	c.DbModel().WithContext(ctx).Where("menu_id=?", code).Delete(&entityRam.RamResourceMenuEntity{})
}

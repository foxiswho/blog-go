package repositoryRam

import (
	"context"

	"github.com/foxiswho/blog-go/infrastructure/entityRam"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/support"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"

	"reflect"
)

func init() {
	gs.Provide(new(RamResourceMenuRepository)).Init(func(s *RamResourceMenuRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[RamResourceMenuRepository])).Init(func(s *support.BaseService[RamResourceMenuRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type RamResourceMenuRepository struct {
	repositoryPg.BaseRepository[entityRam.RamResourceMenuEntity, int64]
	//
}

func (c *RamResourceMenuRepository) DeleteByMenuId(ctx context.Context, code int64) {
	c.DbModel().WithContext(ctx).Where("menu_id=?", code).Delete(&entityRam.RamResourceMenuEntity{})
}

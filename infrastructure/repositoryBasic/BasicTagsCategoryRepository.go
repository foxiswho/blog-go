package repositoryBasic

import (
	"context"
	"reflect"

	"github.com/foxiswho/blog-go/infrastructure/entityBasic"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/support"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
)

func init() {
	gs.Provide(new(BasicTagsCategoryRepository)).Init(func(s *BasicTagsCategoryRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[BasicTagsCategoryRepository])).Init(func(s *support.BaseService[BasicTagsCategoryRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type BasicTagsCategoryRepository struct {
	repositoryPg.BaseRepository[entityBasic.BasicTagsCategoryEntity, int64]
}

func (c *BasicTagsCategoryRepository) FindAllByNoLinkAndTypeSys(ctx context.Context, code string, tpSys string) (info []*entityBasic.BasicTagsCategoryEntity, result bool) {
	tx := c.DbModel().WithContext(ctx).Where("type_sys = ?", tpSys).Where("no_link like ?", "%"+code+"%").Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

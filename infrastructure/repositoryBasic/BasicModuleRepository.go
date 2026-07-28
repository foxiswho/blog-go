package repositoryBasic

import (
	"context"
	"reflect"

	"github.com/hongmengzhu/xianfu-blog-go/infrastructure/entityBasic"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/tools/dbHelper/support"
	"go-spring.org/log"
	"go-spring.org/spring/gs"
)

func init() {
	gs.Provide(new(BasicModuleRepository)).Init(func(s *BasicModuleRepository) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[BasicModuleRepository])).Init(func(s *support.BaseService[BasicModuleRepository]) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type BasicModuleRepository struct {
	repositoryPg.BaseRepository[entityBasic.BasicModuleEntity, int64]
}

func (c *BasicModuleRepository) FindAllByCodeLinkAndTypeSys(ctx context.Context, code string, tpSys string) (info []*entityBasic.BasicModuleEntity, result bool) {
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

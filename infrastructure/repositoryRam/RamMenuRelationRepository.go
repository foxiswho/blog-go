package repositoryRam

import (
	"context"

	"github.com/foxiswho/blog-go/infrastructure/entityRam"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/support"
	"go-spring.org/log"
	"go-spring.org/spring/gs"

	"reflect"
)

func init() {
	gs.Provide(new(RamMenuRelationRepository)).Init(func(s *RamMenuRelationRepository) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[RamMenuRelationRepository])).Init(func(s *support.BaseService[RamMenuRelationRepository]) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type RamMenuRelationRepository struct {
	repositoryPg.BaseRepository[entityRam.RamMenuRelationEntity, int64]
	//
}

func (c *RamMenuRelationRepository) DeleteByMenuId(ctx context.Context, code int64) {
	c.DbModel().WithContext(ctx).Where("menu_id=?", code).Delete(&entityRam.RamMenuRelationEntity{})
}
func (c *RamMenuRelationRepository) FindAllByMenuIdIn(ctx context.Context, code []int64) (info []*entityRam.RamMenuRelationEntity, result bool) {
	tx := c.DbModel().WithContext(ctx).Where("menu_id in ?", code).Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

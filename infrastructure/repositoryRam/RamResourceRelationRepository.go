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
	gs.Provide(new(RamResourceRelationRepository)).Init(func(s *RamResourceRelationRepository) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[RamResourceRelationRepository])).Init(func(s *support.BaseService[RamResourceRelationRepository]) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type RamResourceRelationRepository struct {
	repositoryPg.BaseRepository[entityRam.RamResourceRelationEntity, int64]
	//
}

func (c *RamResourceRelationRepository) DeleteByAuthorityId(ctx context.Context, code int64) {
	c.DbModel().WithContext(ctx).Where("authority_id=?", code).Delete(&entityRam.RamResourceRelationEntity{})
}

func (c *RamResourceRelationRepository) FindByMark(ctx context.Context, code string) (info *entityRam.RamResourceRelationEntity, result bool) {
	tx := c.DbModel().WithContext(ctx).Where("mark=?", code).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

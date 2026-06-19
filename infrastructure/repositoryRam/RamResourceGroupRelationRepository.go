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
	gs.Provide(new(RamResourceGroupRelationRepository)).Init(func(s *RamResourceGroupRelationRepository) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[RamResourceGroupRelationRepository])).Init(func(s *support.BaseService[RamResourceGroupRelationRepository]) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type RamResourceGroupRelationRepository struct {
	repositoryPg.BaseRepository[entityRam.RamResourceGroupRelationEntity, int64]
	//
}

func (c *RamResourceGroupRelationRepository) FindByMark(ctx context.Context, code string) (info *entityRam.RamResourceGroupRelationEntity, result bool) {
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

func (c *RamResourceGroupRelationRepository) DeleteByTypeCategoryAndTypeValue(ctx context.Context, typeCategory, typeValue string) error {
	tx := c.DbModel().WithContext(ctx).Where("type_category = ?", typeCategory).Where("type_value = ?", typeValue).Delete(&entityRam.RamResourceGroupRelationEntity{})
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return tx.Error
	}
	return nil
}

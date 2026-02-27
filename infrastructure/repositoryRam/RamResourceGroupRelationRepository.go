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
	gs.Provide(new(RamResourceGroupRelationRepository)).Init(func(s *RamResourceGroupRelationRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[RamResourceGroupRelationRepository])).Init(func(s *support.BaseService[RamResourceGroupRelationRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type RamResourceGroupRelationRepository struct {
	repositoryPg.BaseRepository[entityRam.RamResourceGroupRelationEntity, int64]
	//
}

func (c *RamResourceGroupRelationRepository) FindAllByTypeCategoryAndTypeValue(typeCategory, typeValue string) (info []*entityRam.RamResourceGroupRelationEntity, result bool) {
	tx := c.Db().Where("type_category = ?", typeCategory).Where("type_value = ?", typeValue).Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}
func (c *RamResourceGroupRelationRepository) FindByMark(code string) (info *entityRam.RamResourceGroupRelationEntity, result bool) {
	tx := c.Db().Where("mark=?", code).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *RamResourceGroupRelationRepository) DeleteByTypeCategoryAndTypeValue(typeCategory, typeValue string) error {
	tx := c.Db().Where("type_category = ?", typeCategory).Where("type_value = ?", typeValue).Delete(&entityRam.RamResourceGroupRelationEntity{})
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return tx.Error
	}
	return nil
}

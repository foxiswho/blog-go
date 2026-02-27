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
	gs.Provide(new(RamResourceAuthorityRepository)).Init(func(s *RamResourceAuthorityRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[RamResourceAuthorityRepository])).Init(func(s *support.BaseService[RamResourceAuthorityRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type RamResourceAuthorityRepository struct {
	repositoryPg.BaseRepository[entityRam.RamResourceAuthorityEntity, int64]
	//
}

func (c *RamResourceAuthorityRepository) FindByMark(code string) (info *entityRam.RamResourceAuthorityEntity, result bool) {
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

func (c *RamResourceAuthorityRepository) FindAllByResourceIdStringIn(code []string) (info []*entityRam.RamResourceAuthorityEntity, result bool) {
	tx := c.Db().Where("resource_id in ?", code).Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *RamResourceAuthorityRepository) FindAllByGroupIdStringIn(code []string) (info []*entityRam.RamResourceAuthorityEntity, result bool) {
	tx := c.Db().Where("group_id in ?", code).Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}
func (c *RamResourceAuthorityRepository) FindAllByTypeCategoryAndGroupIdStringIn(typeCategory string, code []string) (info []*entityRam.RamResourceAuthorityEntity, result bool) {
	tx := c.Db().Where("type_category = ?", typeCategory).Where("group_id in ?", code).Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *RamResourceAuthorityRepository) FindAllByTypeCategoryAndTypeValue(typeCategory, typeValue string) (info []*entityRam.RamResourceAuthorityEntity, result bool) {
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

func (c *RamResourceAuthorityRepository) DeleteByTypeCategoryAndTypeValue(typeCategory, typeValue string) error {
	tx := c.Db().Where("type_category = ?", typeCategory).Where("type_value = ?", typeValue).Delete(&entityRam.RamResourceAuthorityEntity{})
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return tx.Error
	}
	return nil
}

func (c *RamResourceAuthorityRepository) DeleteByMark(code string) error {
	tx := c.Db().Where("mark=?", code).Delete(&entityRam.RamResourceAuthorityEntity{})
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return tx.Error
	}
	return nil
}

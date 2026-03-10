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
	gs.Provide(new(BasicConfigEventFieldsRepository)).Init(func(s *BasicConfigEventFieldsRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[BasicConfigEventFieldsRepository])).Init(func(s *support.BaseService[BasicConfigEventFieldsRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type BasicConfigEventFieldsRepository struct {
	repositoryPg.BaseRepository[entityBasic.BasicConfigEventFieldsEntity, int64]
}

func (c *BasicConfigEventFieldsRepository) FindAllByModelNo(no string) (info []*entityBasic.BasicConfigEventFieldsEntity, result bool) {
	tx := c.Db().Where("model_no=?", no).Order("sort asc,create_at").Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *BasicConfigEventFieldsRepository) DeleteAllByModelNoAndIds(no string, ids []string) (info []*entityBasic.BasicConfigEventFieldsEntity, result bool) {
	tx := c.Db().Where("id in ?", ids).Where("model_no=?", no).Delete(&c.Entity)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}
func (c *BasicConfigEventFieldsRepository) FindAllByEventNo(no string) (info []*entityBasic.BasicConfigEventFieldsEntity, result bool) {
	tx := c.Db().Where("event_no=?", no).Order("sort asc,create_at").Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *BasicConfigEventFieldsRepository) DeleteAllByEventNoAndIds(no string, ids []string) (info []*entityBasic.BasicConfigEventFieldsEntity, result bool) {
	tx := c.Db().Where("id in ?", ids).Where("event_no=?", no).Delete(&c.Entity)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

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
	gs.Provide(new(BasicCountryRepository)).Init(func(s *BasicCountryRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[BasicCountryRepository])).Init(func(s *support.BaseService[BasicCountryRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type BasicCountryRepository struct {
	repositoryPg.BaseRepository[entityBasic.BasicCountryEntity, int64]
}

func (c *BasicCountryRepository) FindByCountryCode(code string) (info *entityBasic.BasicCountryEntity, result bool) {
	tx := c.Db().Where("state=1").Where("country_code=?", code).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}
func (c *BasicCountryRepository) FindByCountryCodeAndIdNot(code, id string) (info *entityBasic.BasicCountryEntity, result bool) {
	tx := c.Db().Where("state=1").Where("country_code=?", code).Where("id!=?", id).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

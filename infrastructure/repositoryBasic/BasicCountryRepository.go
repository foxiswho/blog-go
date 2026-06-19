package repositoryBasic

import (
	"context"
	"reflect"

	"github.com/foxiswho/blog-go/infrastructure/entityBasic"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/support"
	"go-spring.org/log"
	"go-spring.org/spring/gs"
)

func init() {
	gs.Provide(new(BasicCountryRepository)).Init(func(s *BasicCountryRepository) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[BasicCountryRepository])).Init(func(s *support.BaseService[BasicCountryRepository]) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type BasicCountryRepository struct {
	repositoryPg.BaseRepository[entityBasic.BasicCountryEntity, int64]
}

func (c *BasicCountryRepository) FindByCountryCode(ctx context.Context, code string) (info *entityBasic.BasicCountryEntity, result bool) {
	tx := c.DbModel().WithContext(ctx).Where("state=1").Where("country_code=?", code).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}
func (c *BasicCountryRepository) FindByCountryCodeAndIdNot(ctx context.Context, code, id string) (info *entityBasic.BasicCountryEntity, result bool) {
	tx := c.DbModel().WithContext(ctx).Where("state=1").Where("country_code=?", code).Where("id!=?", id).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

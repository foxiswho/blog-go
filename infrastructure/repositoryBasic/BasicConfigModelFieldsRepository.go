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
	gs.Provide(new(BasicConfigModelFieldsRepository)).Init(func(s *BasicConfigModelFieldsRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[BasicConfigModelFieldsRepository])).Init(func(s *support.BaseService[BasicConfigModelFieldsRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type BasicConfigModelFieldsRepository struct {
	repositoryPg.BaseRepository[entityBasic.BasicConfigModelFieldsEntity, int64]
}

func (c *BasicConfigModelFieldsRepository) FindAllByModelNo(ctx context.Context, no string) (info []*entityBasic.BasicConfigModelFieldsEntity, result bool) {
	tx := c.DbModel().WithContext(ctx).Where("model_no=?", no).Order("sort asc,create_at").Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *BasicConfigModelFieldsRepository) DeleteAllByModelNoAndIds(ctx context.Context, no string, ids []string) (info []*entityBasic.BasicConfigModelFieldsEntity, result bool) {
	tx := c.DbModel().WithContext(ctx).Where("id in ?", ids).Where("model_no=?", no).Delete(&c.Entity)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *BasicConfigModelFieldsRepository) FindByEventNoAndFieldIn(ctx context.Context, eventNo string, code []string) (info []*entityBasic.BasicConfigModelFieldsEntity, result bool) {
	tx := c.DbModel().WithContext(ctx).Where("model_no=?", eventNo).Where("field in ?", code).Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}
func (c *BasicConfigModelFieldsRepository) FindByEventNoTenantAndFieldIn(ctx context.Context, eventNo string, tenantNo string, code []string) (info []*entityBasic.BasicConfigModelFieldsEntity, result bool) {
	tx := c.DbModel().WithContext(ctx).Where("model_no=?", eventNo).Where("tenant_no=?", tenantNo).Where("field in ?", code).Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

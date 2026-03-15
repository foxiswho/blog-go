package repositoryBasic

import (
	"context"

	"github.com/foxiswho/blog-go/infrastructure/entityBasic"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/support"
	"github.com/go-spring/spring-core/gs"
)

func init() {
	gs.Provide(new(BasicConfigListRepository))

	gs.Provide(new(support.BaseService[BasicConfigListRepository]))
}

type BasicConfigListRepository struct {
	repositoryPg.BaseRepository[entityBasic.BasicConfigListEntity, int64]
}

func (c *BasicConfigListRepository) FindByEventNo(ctx context.Context, eventNo string) (info *entityBasic.BasicConfigListEntity, result bool) {
	tx := c.DbModel().WithContext(ctx).Where("event_no=?", eventNo).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *BasicConfigListRepository) FindByTenantNoAndEventNoAndIdNot(ctx context.Context, tenantNo, eventNo string, id string) (info *entityBasic.BasicConfigListEntity, result bool) {
	tx := c.DbModel().WithContext(ctx).Where("tenant_no=?", tenantNo).Where("event_no=?", eventNo).Where("id != ?", id).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

package repositoryBasic

import (
	"context"

	"github.com/foxiswho/blog-go/infrastructure/entityBasic"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/support"
	"github.com/go-spring/spring-core/gs"
)

func init() {
	gs.Provide(new(BasicModelRulesRepository))

	gs.Provide(new(support.BaseService[BasicModelRulesRepository]))
}

type BasicModelRulesRepository struct {
	repositoryPg.BaseRepository[entityBasic.BasicModelRulesEntity, int64]
}

func (c *BasicModelRulesRepository) DeleteAllByValueNoAndIds(ctx context.Context, no string, ids []string) (info []*entityBasic.BasicModelRulesEntity, result bool) {
	tx := c.DbModel().WithContext(ctx).Where("id in ?", ids).Where("value_no=?", no).Delete(&c.Entity)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

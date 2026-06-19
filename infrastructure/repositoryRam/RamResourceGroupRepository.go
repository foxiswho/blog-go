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
	gs.Provide(new(RamResourceGroupRepository)).Init(func(s *RamResourceGroupRepository) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[RamResourceGroupRepository])).Init(func(s *support.BaseService[RamResourceGroupRepository]) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type RamResourceGroupRepository struct {
	repositoryPg.BaseRepository[entityRam.RamResourceGroupEntity, int64]
	//
}

func (c *RamResourceGroupRepository) FindAllByIdLink(ctx context.Context, code string) (info []*entityRam.RamResourceGroupEntity, result bool) {
	tx := c.DbModel().WithContext(ctx).Where("id_link like ?", "%|"+code+"|%").Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

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
	gs.Provide(new(RamResourceRepository)).Init(func(s *RamResourceRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[RamResourceRepository])).Init(func(s *support.BaseService[RamResourceRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type RamResourceRepository struct {
	repositoryPg.BaseRepository[entityRam.RamResourceEntity, int64]
	//
}

func (c *RamResourceRepository) FindByParentIdRoot(ctx context.Context) (info []*entityRam.RamResourceEntity, result bool) {
	tx := c.DbModel().WithContext(ctx).Where("parent_id='' ").Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

// CountByParentIdString
//
//	@Description: 统计
//	@receiver c
//	@param pid
//	@return info
//	@return result
func (c *RamResourceRepository) CountByParentIdString(ctx context.Context, pid string) (total int64, result bool) {
	tx := c.DbModel().WithContext(ctx).Where("parent_id= ? ", pid).Count(&total)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return 0, false
	}
	if 0 == tx.RowsAffected {
		return 0, false
	}
	return total, true
}

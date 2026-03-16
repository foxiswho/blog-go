package repositoryRam

import (
	"context"

	"reflect"

	"github.com/foxiswho/blog-go/infrastructure/entityRam"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/support"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
)

func init() {
	gs.Provide(new(RamDepartmentRepository)).Init(func(s *RamDepartmentRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[RamDepartmentRepository])).Init(func(s *support.BaseService[RamDepartmentRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type RamDepartmentRepository struct {
	repositoryPg.BaseRepository[entityRam.RamDepartmentEntity, int64]
}

func (c *RamDepartmentRepository) FindAllByNoLinkArr(ctx context.Context, code []string) (info []*entityRam.RamDepartmentEntity, result bool) {
	db := c.DbModel().WithContext(ctx)
	for index, val := range code {
		if 0 == index {
			db.Where("no_link like ?", "%|"+val+"|%")
		} else {
			db.Or("no_link like ?", "%|"+val+"|%")
		}
	}
	tx := db.Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

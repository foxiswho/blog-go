package repositoryRam

import (
	"context"

	"reflect"

	"github.com/hongmengzhu/xianfu-blog-go/infrastructure/entityRam"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/tools/dbHelper/support"
	"go-spring.org/log"
	"go-spring.org/spring/gs"
)

func init() {
	gs.Provide(new(RamDepartmentRepository)).Init(func(s *RamDepartmentRepository) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[RamDepartmentRepository])).Init(func(s *support.BaseService[RamDepartmentRepository]) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
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

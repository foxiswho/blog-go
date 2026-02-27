package repositoryRam

import (
	"context"

	"github.com/foxiswho/blog-go/infrastructure/entityRam"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/support"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
	"gorm.io/gorm"

	"reflect"
)

func init() {
	gs.Provide(new(RamLevelRepository)).Init(func(s *RamLevelRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[RamLevelRepository])).Init(func(s *support.BaseService[RamLevelRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type RamLevelRepository struct {
	repositoryPg.BaseRepository[entityRam.RamLevelEntity, int64]
	db *gorm.DB `autowire:"?"`
}

func (c *RamLevelRepository) FindAllByCodeIn(code []string) (info []*entityRam.RamLevelEntity, result bool) {
	tx := c.Db().Where("code in ?", code).Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

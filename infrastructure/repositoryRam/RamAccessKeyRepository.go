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
	gs.Provide(new(RamAccessKeyRepository)).Init(func(s *RamAccessKeyRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[RamAccessKeyRepository])).Init(func(s *support.BaseService[RamAccessKeyRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type RamAccessKeyRepository struct {
	repositoryPg.BaseRepository[entityRam.RamAccessKeyEntity, int64]
}

func (c *RamAccessKeyRepository) FindByAno(no string) (info *entityRam.RamAccessKeyEntity, result bool) {
	tx := c.Db().Where("ano=?", no).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *RamAccessKeyRepository) FindByAnoAndAppNo(no, appNo string) (info *entityRam.RamAccessKeyEntity, result bool) {
	tx := c.Db().Where("ano=?", no).Where("app_no=?", no).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

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
	gs.Provide(new(RamAppAccessKeyRepository)).Init(func(s *RamAppAccessKeyRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[RamAppAccessKeyRepository])).Init(func(s *support.BaseService[RamAppAccessKeyRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type RamAppAccessKeyRepository struct {
	repositoryPg.BaseRepository[entityRam.RamAppAccessKeyEntity, int64]
}

func (c *RamAppAccessKeyRepository) FindByTenantNoAndAppNo(no, appNo string) (info *entityRam.RamAppAccessKeyEntity, result bool) {
	tx := c.Db().Where("tenant_no=?", no).Where("app_no=?", appNo).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *RamAppAccessKeyRepository) UpdateAllByAppNoAndNoSetState(appNo, id string, state int8) (sum int64, result bool) {
	tx := c.Db().Where("app_no=?", appNo).Where("id=?", id).Updates(entityRam.RamAppAccessKeyEntity{State: state})
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return 0, false
	}
	if 0 == tx.RowsAffected {
		return 0, false
	}
	return tx.RowsAffected, true
}

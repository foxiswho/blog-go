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

	gs.Provide(new(RamAccountAuthorizationRepository)).Init(func(s *RamAccountAuthorizationRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[RamAccountAuthorizationRepository])).Init(func(s *support.BaseService[RamAccountAuthorizationRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type RamAccountAuthorizationRepository struct {
	repositoryPg.BaseRepository[entityRam.RamAccountAuthorizationEntity, int64]
	//
}

func (c *RamAccountAuthorizationRepository) FindByTypePasswordANo(code string) (info *entityRam.RamAccountAuthorizationEntity, result bool) {
	tx := c.Db().Where("type=?", "password").Where("ano=?", code).Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *RamAccountAuthorizationRepository) DeleteByAno(code string) (result bool) {
	tx := c.Db().Where("ano=?", code).Delete(&entityRam.RamAccountAuthorizationEntity{})
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return false
	}
	if 0 == tx.RowsAffected {
		return false
	}
	return true
}

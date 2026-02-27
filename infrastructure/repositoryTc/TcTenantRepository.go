package repositoryTc

import (
	"context"
	"reflect"

	"github.com/foxiswho/blog-go/infrastructure/entityTc"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/support"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
)

func init() {
	gs.Provide(new(TcTenantRepository)).Init(func(s *TcTenantRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[TcTenantRepository])).Init(func(s *support.BaseService[TcTenantRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type TcTenantRepository struct {
	repositoryPg.BaseRepository[entityTc.TcTenantEntity, int64]
}

func (c *TcTenantRepository) FindByFounder(no string) (info *entityTc.TcTenantEntity, result bool) {
	tx := c.Db().Where("founder=?", no).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *TcTenantRepository) FindByTenantAndFounder(no string) (info *entityTc.TcTenantEntity, result bool) {
	tx := c.Db().Where("founder=?", no).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

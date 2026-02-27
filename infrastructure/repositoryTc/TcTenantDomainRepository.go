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
	gs.Provide(new(TcTenantDomainRepository)).Init(func(s *TcTenantDomainRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[TcTenantDomainRepository])).Init(func(s *support.BaseService[TcTenantDomainRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type TcTenantDomainRepository struct {
	repositoryPg.BaseRepository[entityTc.TcTenantDomainEntity, int64]
}

func (c *TcTenantDomainRepository) FindAllByTenantNo(no string) (infos []*entityTc.TcTenantDomainEntity, query bool) {
	tx := c.Db().Where("tenant_no=?", no).Find(&infos)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return infos, true
}

func (c *TcTenantDomainRepository) SetDefaultedByTenantNo(def int8, no string) (infos []*entityTc.TcTenantDomainEntity, query bool) {
	tx := c.Db().Where("tenant_no=?", no).Update("defaulted", def)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return infos, true
}

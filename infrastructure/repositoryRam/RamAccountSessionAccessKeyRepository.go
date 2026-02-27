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
	gs.Provide(new(RamAccountSessionAccessKeyRepository)).Init(func(s *RamAccountSessionAccessKeyRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[RamAccountSessionAccessKeyRepository])).Init(func(s *support.BaseService[RamAccountSessionAccessKeyRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type RamAccountSessionAccessKeyRepository struct {
	repositoryPg.BaseRepository[entityRam.RamAccountSessionAccessKeyEntity, int64]
}

func (c *RamAccountSessionAccessKeyRepository) FindByAno(no string) (info *entityRam.RamAccountSessionAccessKeyEntity, result bool) {
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

func (c *RamAccountSessionAccessKeyRepository) FindByAnoAndAppNo(no, appNo string) (info *entityRam.RamAccountSessionAccessKeyEntity, result bool) {
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

func (c *RamAccountSessionAccessKeyRepository) FindByNoAndState(no string) (info *entityRam.RamAccountSessionAccessKeyEntity, result bool) {
	tx := c.Db().Where("no=?", no).Where("state=1").First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *RamAccountSessionAccessKeyRepository) FindByTenantNoAndNoAndState(tno, no string) (info *entityRam.RamAccountSessionAccessKeyEntity, result bool) {
	tx := c.Db().Where("tenant_no=?", tno).Where("no=?", no).Where("state=1").First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *RamAccountSessionAccessKeyRepository) FindByTypeDomainAndState(domain []string) (info []*entityRam.RamAccountSessionAccessKeyEntity, result bool) {
	tx := c.Db().Where("type_domain in ?", domain).Where("state=1").Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

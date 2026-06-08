package repositoryRam

import (
	"context"

	"github.com/foxiswho/blog-go/infrastructure/entityRam"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/support"
	"github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"

	"reflect"
)

func init() {
	gs.Provide(new(RamAccountSessionAccessKeyRepository)).Init(func(s *RamAccountSessionAccessKeyRepository) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[RamAccountSessionAccessKeyRepository])).Init(func(s *support.BaseService[RamAccountSessionAccessKeyRepository]) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type RamAccountSessionAccessKeyRepository struct {
	repositoryPg.BaseRepository[entityRam.RamAccountSessionAccessKeyEntity, int64]
}

func (c *RamAccountSessionAccessKeyRepository) FindByAno(ctx context.Context, no string) (info *entityRam.RamAccountSessionAccessKeyEntity, result bool) {
	tx := c.DbModel().WithContext(ctx).Where("ano=?", no).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *RamAccountSessionAccessKeyRepository) FindByAnoAndAppNo(ctx context.Context, no, appNo string) (info *entityRam.RamAccountSessionAccessKeyEntity, result bool) {
	tx := c.DbModel().WithContext(ctx).Where("ano=?", no).Where("app_no=?", no).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *RamAccountSessionAccessKeyRepository) FindByNoAndState(ctx context.Context, no string) (info *entityRam.RamAccountSessionAccessKeyEntity, result bool) {
	tx := c.DbModel().WithContext(ctx).Where("no=?", no).Where("state=1").First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}
func (c *RamAccountSessionAccessKeyRepository) FindByNoAndClientAndState(ctx context.Context, no string, client string) (info *entityRam.RamAccountSessionAccessKeyEntity, result bool) {
	tx := c.DbModel().WithContext(ctx).Where("no=?", no).Where("state=1").Where("client=?", client).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *RamAccountSessionAccessKeyRepository) FindByTenantNoAndNoAndState(ctx context.Context, tno, no string) (info *entityRam.RamAccountSessionAccessKeyEntity, result bool) {
	tx := c.DbModel().WithContext(ctx).Where("tenant_no=?", tno).Where("no=?", no).Where("state=1").First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *RamAccountSessionAccessKeyRepository) FindByTypeDomainAndState(ctx context.Context, domain string) (info *entityRam.RamAccountSessionAccessKeyEntity, result bool) {
	tx := c.DbModel().WithContext(ctx).Where("type_domain = ?", domain).Where("state=1").First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *RamAccountSessionAccessKeyRepository) FindByTypeDomainInAndState(ctx context.Context, domain []string) (info []*entityRam.RamAccountSessionAccessKeyEntity, result bool) {
	tx := c.DbModel().WithContext(ctx).Where("type_domain in ?", domain).Where("state=1").Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *RamAccountSessionAccessKeyRepository) FindByTypeDomainInAndClientAndState(ctx context.Context, domain, client []string) (info []*entityRam.RamAccountSessionAccessKeyEntity, result bool) {
	tx := c.DbModel().WithContext(ctx).Where("type_domain in ?", domain).Where("client in ?", client).Where("state=1").Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

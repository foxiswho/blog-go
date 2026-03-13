package repositoryBasic

import (
	"github.com/foxiswho/blog-go/infrastructure/entityBasic"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/support"
	"github.com/go-spring/spring-core/gs"
)

func init() {
	gs.Provide(new(BasicConfigRepository))

	gs.Provide(new(support.BaseService[BasicConfigRepository]))
}

type BasicConfigRepository struct {
	repositoryPg.BaseRepository[entityBasic.BasicConfigEntity, int64]
}

func (c *BasicConfigRepository) FindByTypeDomainTenantAndField(typeDomain string, tenantNo string, code string) (info *entityBasic.BasicConfigEntity, result bool) {
	tx := c.Db().Where("type_domain=?", typeDomain).Where("tenant_no=?", tenantNo).Where("field=?", code).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}
func (c *BasicConfigRepository) FindByEventNoAndFieldIn(eventNo string, code []string) (info []*entityBasic.BasicConfigEntity, result bool) {
	tx := c.Db().Where("event_no=?", eventNo).Where("field in ?", code).Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}
func (c *BasicConfigRepository) FindByEventNoTenantAndFieldIn(eventNo string, tenantNo string, code []string) (info []*entityBasic.BasicConfigEntity, result bool) {
	tx := c.Db().Where("event_no=?", eventNo).Where("tenant_no=?", tenantNo).Where("field in ?", code).Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

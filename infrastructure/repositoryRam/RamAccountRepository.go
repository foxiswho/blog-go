package repositoryRam

import (
	"context"
	"reflect"

	"github.com/foxiswho/blog-go/infrastructure/entityRam"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/support"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
)

func init() {
	gs.Provide(new(RamAccountRepository)).Init(func(s *RamAccountRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[RamAccountRepository])).Init(func(s *support.BaseService[RamAccountRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type RamAccountRepository struct {
	repositoryPg.BaseRepository[entityRam.RamAccountEntity, int64]
}

func (c *RamAccountRepository) FindByAccount(code string, opts ...repositoryPg.Option) (info *entityRam.RamAccountEntity, query bool, err error) {
	tx := c.SetOptionScopes(c.DbModel(), opts...).Where("account=?", code).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false, tx.Error
	}
	if 0 == tx.RowsAffected {
		return nil, false, nil
	}
	return info, true, nil
}

func (c *RamAccountRepository) FindByAccountMd5(code string, opts ...repositoryPg.Option) (info *entityRam.RamAccountEntity, query bool, err error) {
	tx := c.SetOptionScopes(c.DbModel(), opts...).Where("account_md5=?", code).Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false, tx.Error
	}
	if 0 == tx.RowsAffected {
		return nil, false, nil
	}
	return info, true, nil
}
func (c *RamAccountRepository) FindByAccountAndTypeDomain(code, typeDomain string, opts ...repositoryPg.Option) (info *entityRam.RamAccountEntity, query bool) {
	tx := c.SetOptionScopes(c.DbModel(), opts...).Where("account=?", code).Where("type_domain=?", typeDomain).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *RamAccountRepository) FindByAccountAndTypeDomainAndIdNot(code, typeDomain, id string, opts ...repositoryPg.Option) (info *entityRam.RamAccountEntity, query bool) {
	tx := c.SetOptionScopes(c.DbModel(), opts...).Where("account=?", code).Where("type_domain=?", typeDomain).Where("id <> ?", id).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *RamAccountRepository) FindByPhoneAndTypeDomainAndIdNot(code, typeDomain, id string, opts ...repositoryPg.Option) (info *entityRam.RamAccountEntity, query bool) {
	tx := c.SetOptionScopes(c.DbModel(), opts...).Where("phone=?", code).Where("type_domain=?", typeDomain).Where("id <> ?", id).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}
func (c *RamAccountRepository) FindByPhoneMd5AndTypeDomainAndIdNot(code, typeDomain, id string, opts ...repositoryPg.Option) (info *entityRam.RamAccountEntity, query bool) {
	tx := c.SetOptionScopes(c.DbModel(), opts...).Where("phone_md5=?", code).Where("type_domain=?", typeDomain).Where("id <> ?", id).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}
func (c *RamAccountRepository) FindByMailAndTypeDomainAndIdNot(code, typeDomain, id string, opts ...repositoryPg.Option) (info *entityRam.RamAccountEntity, query bool) {
	tx := c.SetOptionScopes(c.DbModel(), opts...).Where("mail=?", code).Where("type_domain=?", typeDomain).Where("id <> ?", id).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}
func (c *RamAccountRepository) FindByCodeAndTypeDomainAndIdNot(code, typeDomain, id string, opts ...repositoryPg.Option) (info *entityRam.RamAccountEntity, query bool) {
	tx := c.SetOptionScopes(c.DbModel(), opts...).Where("code=?", code).Where("type_domain=?", typeDomain).Where("id <> ?", id).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}
func (c *RamAccountRepository) FindByIdentityCodeAndTypeDomainAndIdNot(code, typeDomain, id string, opts ...repositoryPg.Option) (info *entityRam.RamAccountEntity, query bool) {
	tx := c.SetOptionScopes(c.DbModel(), opts...).Where("identity_code=?", code).Where("type_domain=?", typeDomain).Where("id <> ?", id).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}
func (c *RamAccountRepository) FindByRealNameAndTypeDomainAndIdNot(code, typeDomain, id string, opts ...repositoryPg.Option) (info *entityRam.RamAccountEntity, query bool) {
	tx := c.SetOptionScopes(c.DbModel(), opts...).Where("real_name=?", code).Where("type_domain=?", typeDomain).Where("id <> ?", id).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}
func (c *RamAccountRepository) FindByAccountMd5AndTypeDomain(code, typeDomain string, opts ...repositoryPg.Option) (info *entityRam.RamAccountEntity, query bool, err error) {
	tx := c.SetOptionScopes(c.DbModel(), opts...).Where("account_md5=?", code).Where("type_domain=?", typeDomain).Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false, tx.Error
	}
	if 0 == tx.RowsAffected {
		return nil, false, nil
	}
	return info, true, nil
}

func (c *RamAccountRepository) FindByNoAndTypeDomainAndIdNot(code, typeDomain, id string, opts ...repositoryPg.Option) (info *entityRam.RamAccountEntity, query bool) {
	tx := c.SetOptionScopes(c.DbModel(), opts...).Where("no=?", code).Where("type_domain=?", typeDomain).Where("id <> ?", id).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *RamAccountRepository) FindByPhone(code string, opts ...repositoryPg.Option) (info *entityRam.RamAccountEntity, query bool, err error) {
	tx := c.SetOptionScopes(c.DbModel(), opts...).Where("phone=?", code).Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false, tx.Error
	}
	if 0 == tx.RowsAffected {
		return nil, false, nil
	}
	return info, true, nil
}

func (c *RamAccountRepository) FindByMail(code string, opts ...repositoryPg.Option) (info *entityRam.RamAccountEntity, query bool, err error) {
	tx := c.SetOptionScopes(c.DbModel(), opts...).Where("mail=?", code).Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false, tx.Error
	}
	if 0 == tx.RowsAffected {
		return nil, false, nil
	}
	return info, true, nil
}

func (c *RamAccountRepository) FindByIdAndTypeDomain(code int64, typeDomain string, opts ...repositoryPg.Option) (info *entityRam.RamAccountEntity, query bool) {
	arg := repositoryPg.OptionArg{}
	for _, opt := range opts {
		opt(&arg)
	}
	tx := c.SetOptionScopes(c.DbModel(), opts...).Where("id=?", code).Where("type_domain=?", typeDomain).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *RamAccountRepository) FindByNoAndTypeDomain(code int64, typeDomain string, opts ...repositoryPg.Option) (info *entityRam.RamAccountEntity, query bool) {
	tx := c.SetOptionScopes(c.DbModel(), opts...).Where("no=?", code).Where("type_domain=?", typeDomain).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *RamAccountRepository) FindAllByIdStringInAndTypeDomain(ids []string, typeDomain string, opts ...repositoryPg.Option) (infos []*entityRam.RamAccountEntity, query bool) {
	tx := c.SetOptionScopes(c.DbModel(), opts...).Where("id in (?)", ids).Where("type_domain=?", typeDomain).Find(&infos)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return infos, true
}
func (c *RamAccountRepository) FindAllByNoInAndTypeDomain(ids []string, typeDomain string, opts ...repositoryPg.Option) (infos []*entityRam.RamAccountEntity, query bool) {
	tx := c.SetOptionScopes(c.DbModel(), opts...).Where("no in (?)", ids).Where("type_domain=?", typeDomain).Find(&infos)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return infos, true
}

func (c *RamAccountRepository) FindByTenantNoAccountAndTypeDomainAndIdNot(tenantNo, code, typeDomain, id string) (info *entityRam.RamAccountEntity, query bool) {
	tx := c.DbModel().Where("tenant_no=?", tenantNo).Where("account=?", code).Where("type_domain=?", typeDomain).Where("id <> ?", id).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}
func (c *RamAccountRepository) FindByTenantNoAccountMd5AndTypeDomain(tenantNo, code, typeDomain string) (info *entityRam.RamAccountEntity, query bool, err error) {
	tx := c.DbModel().Where("tenant_no=?", tenantNo).Where("account_md5=?", code).Where("type_domain=?", typeDomain).Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false, tx.Error
	}
	if 0 == tx.RowsAffected {
		return nil, false, nil
	}
	return info, true, nil
}

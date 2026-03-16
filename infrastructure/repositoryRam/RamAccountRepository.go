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

func (c *RamAccountRepository) FindByAccount(ctx context.Context, code string, opts ...repositoryPg.OptionPg) (info *entityRam.RamAccountEntity, query bool, err error) {
	tx := c.SetOptionScopes(c.DbModel().WithContext(ctx), opts...).Where("account=?", code).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false, tx.Error
	}
	if 0 == tx.RowsAffected {
		return nil, false, nil
	}
	return info, true, nil
}

func (c *RamAccountRepository) FindByAccountMd5(ctx context.Context, code string, opts ...repositoryPg.OptionPg) (info *entityRam.RamAccountEntity, query bool, err error) {
	tx := c.SetOptionScopes(c.DbModel().WithContext(ctx), opts...).Where("account_md5=?", code).Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false, tx.Error
	}
	if 0 == tx.RowsAffected {
		return nil, false, nil
	}
	return info, true, nil
}
func (c *RamAccountRepository) FindByAccountAndTypeDomain(ctx context.Context, code, typeDomain string, opts ...repositoryPg.OptionPg) (info *entityRam.RamAccountEntity, query bool) {
	tx := c.SetOptionScopes(c.DbModel().WithContext(ctx), opts...).Where("account=?", code).Where("type_domain=?", typeDomain).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *RamAccountRepository) FindByAccountAndTypeDomainAndIdNot(ctx context.Context, code, typeDomain, id string, opts ...repositoryPg.OptionPg) (info *entityRam.RamAccountEntity, query bool) {
	tx := c.SetOptionScopes(c.DbModel().WithContext(ctx), opts...).Where("account=?", code).Where("type_domain=?", typeDomain).Where("id <> ?", id).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *RamAccountRepository) FindByPhoneAndTypeDomainAndIdNot(ctx context.Context, code, typeDomain, id string, opts ...repositoryPg.OptionPg) (info *entityRam.RamAccountEntity, query bool) {
	tx := c.SetOptionScopes(c.DbModel().WithContext(ctx), opts...).Where("phone=?", code).Where("type_domain=?", typeDomain).Where("id <> ?", id).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}
func (c *RamAccountRepository) FindByPhoneMd5AndTypeDomainAndIdNot(ctx context.Context, code, typeDomain, id string, opts ...repositoryPg.OptionPg) (info *entityRam.RamAccountEntity, query bool) {
	tx := c.SetOptionScopes(c.DbModel().WithContext(ctx), opts...).Where("phone_md5=?", code).Where("type_domain=?", typeDomain).Where("id <> ?", id).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}
func (c *RamAccountRepository) FindByMailAndTypeDomainAndIdNot(ctx context.Context, code, typeDomain, id string, opts ...repositoryPg.OptionPg) (info *entityRam.RamAccountEntity, query bool) {
	tx := c.SetOptionScopes(c.DbModel().WithContext(ctx), opts...).Where("mail=?", code).Where("type_domain=?", typeDomain).Where("id <> ?", id).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}
func (c *RamAccountRepository) FindByCodeAndTypeDomainAndIdNot(ctx context.Context, code, typeDomain, id string, opts ...repositoryPg.OptionPg) (info *entityRam.RamAccountEntity, query bool) {
	tx := c.SetOptionScopes(c.DbModel().WithContext(ctx), opts...).Where("code=?", code).Where("type_domain=?", typeDomain).Where("id <> ?", id).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}
func (c *RamAccountRepository) FindByIdentityCodeAndTypeDomainAndIdNot(ctx context.Context, code, typeDomain, id string, opts ...repositoryPg.OptionPg) (info *entityRam.RamAccountEntity, query bool) {
	tx := c.SetOptionScopes(c.DbModel().WithContext(ctx), opts...).Where("identity_code=?", code).Where("type_domain=?", typeDomain).Where("id <> ?", id).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}
func (c *RamAccountRepository) FindByRealNameAndTypeDomainAndIdNot(ctx context.Context, code, typeDomain, id string, opts ...repositoryPg.OptionPg) (info *entityRam.RamAccountEntity, query bool) {
	tx := c.SetOptionScopes(c.DbModel().WithContext(ctx), opts...).Where("real_name=?", code).Where("type_domain=?", typeDomain).Where("id <> ?", id).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}
func (c *RamAccountRepository) FindByAccountMd5AndTypeDomain(ctx context.Context, code, typeDomain string, opts ...repositoryPg.OptionPg) (info *entityRam.RamAccountEntity, query bool, err error) {
	tx := c.SetOptionScopes(c.DbModel().WithContext(ctx), opts...).Where("account_md5=?", code).Where("type_domain=?", typeDomain).Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false, tx.Error
	}
	if 0 == tx.RowsAffected {
		return nil, false, nil
	}
	return info, true, nil
}

func (c *RamAccountRepository) FindByNoAndTypeDomainAndIdNot(ctx context.Context, code, typeDomain, id string, opts ...repositoryPg.OptionPg) (info *entityRam.RamAccountEntity, query bool) {
	tx := c.SetOptionScopes(c.DbModel().WithContext(ctx), opts...).Where("no=?", code).Where("type_domain=?", typeDomain).Where("id <> ?", id).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *RamAccountRepository) FindByPhone(ctx context.Context, code string, opts ...repositoryPg.OptionPg) (info *entityRam.RamAccountEntity, query bool, err error) {
	tx := c.SetOptionScopes(c.DbModel().WithContext(ctx), opts...).Where("phone=?", code).Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false, tx.Error
	}
	if 0 == tx.RowsAffected {
		return nil, false, nil
	}
	return info, true, nil
}

func (c *RamAccountRepository) FindByMail(ctx context.Context, code string, opts ...repositoryPg.OptionPg) (info *entityRam.RamAccountEntity, query bool, err error) {
	tx := c.SetOptionScopes(c.DbModel().WithContext(ctx), opts...).Where("mail=?", code).Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false, tx.Error
	}
	if 0 == tx.RowsAffected {
		return nil, false, nil
	}
	return info, true, nil
}

func (c *RamAccountRepository) FindByIdAndTypeDomain(ctx context.Context, code int64, typeDomain string, opts ...repositoryPg.OptionPg) (info *entityRam.RamAccountEntity, query bool) {
	tx := c.SetOptionScopes(c.DbModel().WithContext(ctx), opts...).Where("id=?", code).Where("type_domain=?", typeDomain).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *RamAccountRepository) FindByNoAndTypeDomain(ctx context.Context, code int64, typeDomain string, opts ...repositoryPg.OptionPg) (info *entityRam.RamAccountEntity, query bool) {
	tx := c.SetOptionScopes(c.DbModel().WithContext(ctx), opts...).Where("no=?", code).Where("type_domain=?", typeDomain).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *RamAccountRepository) FindAllByIdStringInAndTypeDomain(ctx context.Context, ids []string, typeDomain string, opts ...repositoryPg.OptionPg) (infos []*entityRam.RamAccountEntity, query bool) {
	tx := c.SetOptionScopes(c.DbModel().WithContext(ctx), opts...).Where("id in (?)", ids).Where("type_domain=?", typeDomain).Find(&infos)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return infos, true
}
func (c *RamAccountRepository) FindAllByNoInAndTypeDomain(ctx context.Context, ids []string, typeDomain string, opts ...repositoryPg.OptionPg) (infos []*entityRam.RamAccountEntity, query bool) {
	tx := c.SetOptionScopes(c.DbModel().WithContext(ctx), opts...).Where("no in (?)", ids).Where("type_domain=?", typeDomain).Find(&infos)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return infos, true
}

func (c *RamAccountRepository) FindByTenantNoAccountAndTypeDomainAndIdNot(ctx context.Context, tenantNo, code, typeDomain, id string) (info *entityRam.RamAccountEntity, query bool) {
	tx := c.DbModel().WithContext(ctx).Where("tenant_no=?", tenantNo).Where("account=?", code).Where("type_domain=?", typeDomain).Where("id <> ?", id).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}
func (c *RamAccountRepository) FindByTenantNoAccountMd5AndTypeDomain(ctx context.Context, tenantNo, code, typeDomain string) (info *entityRam.RamAccountEntity, query bool, err error) {
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

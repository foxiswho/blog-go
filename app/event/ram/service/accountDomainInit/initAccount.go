package accountDomainInit

import (
	"context"
	"time"

	"github.com/foxiswho/blog-go/app/manage/domainRam/utilsRam"
	"github.com/foxiswho/blog-go/infrastructure/entityRam"
	"github.com/foxiswho/blog-go/infrastructure/entityTc"
	"github.com/foxiswho/blog-go/infrastructure/repositoryRam"
	"github.com/foxiswho/blog-go/infrastructure/repositoryTc"
	"github.com/foxiswho/blog-go/middleware/components/dataFilePg"
	"github.com/foxiswho/blog-go/pkg/configPg"
	"github.com/foxiswho/blog-go/pkg/consts/constsPg"
	"github.com/foxiswho/blog-go/pkg/consts/constsRam/typeDomainPg"
	"github.com/foxiswho/blog-go/pkg/consts/constsRam/typeIdentityPg"
	"github.com/foxiswho/blog-go/pkg/enum/enumRam/enumAuthorizationTypePg"
	"github.com/foxiswho/blog-go/pkg/enum/state/enumStatePg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/foxiswho/blog-go/pkg/tools/noPg"
	"github.com/pangu-2/go-tools/tools/cryptPg"
	"github.com/pangu-2/go-tools/tools/strPg"
	"github.com/pangu-2/go-tools/tools/userPg"
	"gorm.io/datatypes"
)

type Sp struct {
	log      *log2.Logger                                     `autowire:"?"`
	accDb    *repositoryRam.RamAccountRepository              `autowire:"?"`
	authDb   *repositoryRam.RamAccountAuthorizationRepository `autowire:"?"`
	depDb    *repositoryRam.RamDepartmentRepository           `autowire:"?"`
	roleDb   *repositoryRam.RamRoleRepository                 `autowire:"?"`
	teamDb   *repositoryRam.RamTeamRepository                 `autowire:"?"`
	groupDb  *repositoryRam.RamGroupRepository                `autowire:"?"`
	levelDb  *repositoryRam.RamLevelRepository                `autowire:"?"`
	tenantDb *repositoryTc.TcTenantRepository                 `autowire:"?"`
	pg       configPg.Pg                                      `value:"${pg}"`
}

var domain = []string{
	typeDomainPg.System.Index(),
	//iamConstant.Merchant.Index(),
	//iamConstant.Tenant.Index(),
	typeDomainPg.Manage.Index(),
}

// InitAccount
// @Description: 初始化账号
type InitAccount struct {
	sp  *Sp          `autowire:"?"`
	log *log2.Logger `autowire:"?"`
}

func NewInitAccount(log *log2.Logger, sp *Sp) *InitAccount {
	return &InitAccount{
		log: log,
		sp:  sp,
	}
}

func (t *InitAccount) Processor(ctx context.Context) error {
	mapDomain := make(map[string]*entityRam.RamAccountEntity)
	find := make(map[string]bool)
	// 获取超管账号
	info, result := t.sp.accDb.FindAllByNoIn(domain)
	if result {
		if len(info) > 0 {
			for _, v := range info {
				mapDomain[v.No] = v
				find[v.No] = true
			}
		}
	}
	for _, item := range domain {
		//不存在时创建
		if _, ok := find[item]; !ok {
			//系统账号
			if typeDomainPg.System.IsEqual(item) {
				t.systemAccount(ctx, item)
			} else if typeDomainPg.Manage.IsEqual(item) {
				//管理后台账号
				t.manageAccount(ctx, item)
			}
		} else {
			t.log.Debugf("[init].[账号初始化] %+v 已存在,无需初始化", item)
		}
	}
	return nil
}

func (t *InitAccount) systemAccount(ctx context.Context, domain string) {
	no := noPg.No()
	now := time.Now()
	save := entityRam.RamAccountEntity{
		ID:            1,
		No:            domain,
		Code:          domain,
		TenantNo:      domain,
		TypeDomain:    domain,
		TypeIdentity:  typeIdentityPg.General.Index(),
		Account:       constsPg.ACCOUNT_SYSTEM,
		Cc:            constsPg.ACCOUNT_CC,
		Mail:          no + "@xxxx.com",
		AccountVerify: enumStatePg.ENABLE.Index(),
		State:         enumStatePg.ENABLE.Index(),
		Founder:       enumStatePg.ENABLE.Index(),
		RegisterTime:  &now,
	}
	save.AccountMd5 = cryptPg.Md5(save.Account)
	save.PhoneMd5 = cryptPg.Md5(save.Account)
	save.MailMd5 = cryptPg.Md5(save.Account)
	save.Name = constsPg.ACCOUNT_NAME
	save.RealName = constsPg.ACCOUNT_NAME
	//
	os := entityRam.RamAccountJsonOs{
		Departments: make([]string, 0),
		Groups:      make([]string, 0),
		Levels:      make([]string, 0),
		Merchants:   make([]string, 0),
		Orgs:        make([]string, 0),
		Projects:    make([]string, 0),
		Roles:       make([]string, 0),
		Shops:       make([]string, 0),
		Teams:       make([]string, 0),
		Tenants:     make([]string, 0),
	}
	os.Tenants = append(os.Tenants, save.TenantNo)
	save.Os = datatypes.NewJSONType(os)
	t.sp.accDb.Create(&save)
	//
	t.sp.authDb.DeleteByAno(save.No)
	//
	salt := strPg.GetNanoid(8)
	authorizationEntity := entityRam.RamAccountAuthorizationEntity{
		Ano:       save.No,
		Type:      enumAuthorizationTypePg.PASSWORD.String(),
		TenantNo:  save.TenantNo,
		ExtraData: salt,
	}
	pwd := constsPg.ACCOUNT_PASSWORD + "." + strPg.GetNanoid(8)
	authorizationEntity.Value = userPg.PasswordSalt(pwd, salt)
	//设置唯一值
	authorizationEntity.KindUnique = utilsRam.AuthorizationKindUniquePasswordByEntity(authorizationEntity)
	t.sp.authDb.Create(&authorizationEntity)
	//
	//记录文件日志
	dataFilePg.NewAccountFileRecord(t.sp.pg, dataFilePg.MakeContent(save.Account, pwd, typeDomainPg.System.Index())).Write()
}

func (t *InitAccount) manageAccount(ctx context.Context, domain string) {
	no := noPg.No()
	now := time.Now()
	tenant, result := t.sp.tenantDb.FindByNo(constsPg.ACCOUNT_MANAGE_No)
	if !result {
		tenant = &entityTc.TcTenantEntity{
			ID:       1000,
			No:       constsPg.ACCOUNT_MANAGE_No,
			Name:     "默认",
			NameFl:   "默认",
			NameFull: "默认",
			Code:     constsPg.ACCOUNT_MANAGE_No,
			State:    enumStatePg.ENABLE.Index(),
			Founder:  "1000",
		}
		t.sp.tenantDb.Create(tenant)
	}
	save := entityRam.RamAccountEntity{
		ID:            1000,
		No:            domain,
		Code:          domain,
		TenantNo:      tenant.No,
		TypeDomain:    domain,
		TypeIdentity:  typeIdentityPg.General.Index(),
		Account:       constsPg.ACCOUNT_MANAGE,
		Cc:            constsPg.ACCOUNT_CC,
		Mail:          no + "@xxxx.com",
		AccountVerify: enumStatePg.ENABLE.Index(),
		State:         enumStatePg.ENABLE.Index(),
		Founder:       enumStatePg.ENABLE.Index(),
		RegisterTime:  &now,
	}
	save.AccountMd5 = cryptPg.Md5(save.Account)
	save.PhoneMd5 = cryptPg.Md5(save.Account)
	save.MailMd5 = cryptPg.Md5(save.Account)
	save.Name = constsPg.ACCOUNT_NAME
	save.RealName = constsPg.ACCOUNT_NAME
	//
	os := entityRam.RamAccountJsonOs{
		Departments: make([]string, 0),
		Groups:      make([]string, 0),
		Levels:      make([]string, 0),
		Merchants:   make([]string, 0),
		Orgs:        make([]string, 0),
		Projects:    make([]string, 0),
		Roles:       make([]string, 0),
		Shops:       make([]string, 0),
		Teams:       make([]string, 0),
		Tenants:     make([]string, 0),
	}
	os.Tenants = append(os.Tenants, save.TenantNo)
	save.Os = datatypes.NewJSONType(os)
	t.sp.accDb.Create(&save)
	//
	t.sp.authDb.DeleteByAno(save.No)
	//
	salt := strPg.GetNanoid(8)
	authorizationEntity := entityRam.RamAccountAuthorizationEntity{
		Ano:       save.No,
		Type:      enumAuthorizationTypePg.PASSWORD.String(),
		TenantNo:  save.TenantNo,
		ExtraData: salt,
	}
	authorizationEntity.Value = userPg.PasswordSalt(constsPg.ACCOUNT_PASSWORD, salt)
	//设置唯一值
	authorizationEntity.KindUnique = utilsRam.AuthorizationKindUniquePasswordByEntity(authorizationEntity)
	t.sp.authDb.Create(&authorizationEntity)
}

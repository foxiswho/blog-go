package service

import (
	"context"
	"encoding/json"
	"reflect"
	"time"

	"github.com/farseer-go/eventBus"
	"github.com/foxiswho/blog-go/app/system/ram/model/modRamLogin"
	"github.com/foxiswho/blog-go/infrastructure/entityRam"
	"github.com/foxiswho/blog-go/infrastructure/repositoryRam"
	"github.com/foxiswho/blog-go/middleware/components/authTokenPg"
	"github.com/foxiswho/blog-go/middleware/components/cachePg/cacheAuthPubPrivPg"
	"github.com/foxiswho/blog-go/pkg/configPg"
	"github.com/foxiswho/blog-go/pkg/consts/constEventBusPg"
	"github.com/foxiswho/blog-go/pkg/consts/constHeaderPg"
	"github.com/foxiswho/blog-go/pkg/consts/constsRam/typeDomainPg"
	"github.com/foxiswho/blog-go/pkg/enum/enumCommonPg/appModulePg"
	"github.com/foxiswho/blog-go/pkg/enum/state/enumStatePg"
	"github.com/foxiswho/blog-go/pkg/holderPg"
	"github.com/foxiswho/blog-go/pkg/holderPg/multiTenantPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/foxiswho/blog-go/pkg/sdk/ram/model/modRamAccount"
	"github.com/gin-gonic/gin"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
	"github.com/jinzhu/copier"
	"github.com/pangu-2/go-tools/tools/cryptPg"
	"github.com/pangu-2/go-tools/tools/jsonPg"
	"github.com/pangu-2/go-tools/tools/strPg"
	"github.com/pangu-2/go-tools/tools/userPg"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
)

func init() {
	gs.Provide(NewAccountLoginService).Init(func(s *AccountLoginService) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

// AccountLoginService 登录
// @Description:
type AccountLoginService struct {
	dao       *repositoryRam.RamAccountRepository                 `autowire:"?"`
	daoAuth   *repositoryRam.RamAccountAuthorizationRepository    `autowire:"?"`
	sessionAk *repositoryRam.RamAccountSessionAccessKeyRepository `autowire:"?"`
	pg        configPg.Pg                                         `value:"${pg}"`
	log       *log2.Logger                                        `autowire:"?"`
}

func NewAccountLoginService() *AccountLoginService {
	return new(AccountLoginService)
}

// Login
//
//	@Description:
//	@receiver c
//	@param ct
//	@return rt
func (c *AccountLoginService) Login(ctx *gin.Context, ct modRamLogin.LoginCt, tp appModulePg.AppModule) (rt rg.Rs[modRamLogin.LoginSuccess]) {
	c.log.Infof("tp=%+v,ct=%+v", tp, ct)
	if !c.dao.Config().Domain.System {
		return rt.ErrorMessage("系统管理模块已禁用，不允许操作")
	}
	if strPg.IsBlank(ct.Account) {
		return rt.ErrorMessage("账号不能为空")
	}
	if strPg.IsBlank(ct.Password) {
		return rt.ErrorMessage("密码不能为空")
	}

	md5 := cryptPg.Md5(ct.Account)
	info, b, err := c.dao.FindByAccountMd5AndTypeDomain(md5, tp.ToTypeDomain().String())
	if nil != err {
		return rt.Error()
	}
	if !b {
		return rt.ErrorMessage("账号不存在")
	}
	if !enumStatePg.ENABLE.IsExistInt8(info.State) {
		return rt.ErrorMessage("账户已被禁用，不能登陆")
	}
	pwdInfo, result := c.daoAuth.FindByTypePasswordANo(info.No)
	if !result {
		return rt.ErrorMessage("用户密码未设置")
	}
	if !userPg.PasswordVerify(pwdInfo.Value, ct.Password, pwdInfo.ExtraData) {
		c.log.Debugf("pwd=%+v,[%+v],[value]=%+v,[extra]=%+v", ct.Password, userPg.PasswordSalt(ct.Password, pwdInfo.ExtraData), pwdInfo.Value, pwdInfo.ExtraData)
		return rt.ErrorMessage("账号密码错误")
	}
	now := time.Now()
	//租户默认
	mult := multiTenantPg.MultiTenantPg{
		TenantNo: make([]string, 0),
	}
	mult.TenantNo = append(mult.TenantNo, typeDomainPg.System.Index())
	//是否新生成密钥
	isMakeNewKey := false
	privatePubKey := authTokenPg.Result{}
	// 获取 密钥对
	no, r := c.sessionAk.FindByNoAndState(typeDomainPg.System.Index())
	if !r {
		privatePubKey = authTokenPg.MakePublicPrivateKey()
		isMakeNewKey = true
	} else {
		if strPg.IsBlank(no.Data) {
			privatePubKey = authTokenPg.MakePublicPrivateKey()
			isMakeNewKey = true
		} else {
			var privatePubKeyEnt entityRam.RamAsaJsonPrivatePublicKey
			err := json.Unmarshal([]byte(no.Data), &privatePubKeyEnt)
			if err != nil {
				privatePubKey = authTokenPg.MakePublicPrivateKey()
				isMakeNewKey = true
			} else {
				privatePubKey.PrivateKey = privatePubKeyEnt.Private
				privatePubKey.PublicKey = privatePubKeyEnt.Public
			}
		}
	}
	//生成 令牌
	param := authTokenPg.Param{
		UniqueId:    strPg.GenerateNumberId22(),
		MultiTenant: mult,
		LoginNo:     info.No,
		No:          info.No,
		Name:        info.Name,
		Type:        string(tp),
		Result:      privatePubKey,
		TenantNo:    typeDomainPg.System.Index(),
	}
	ret := authTokenPg.MakePaseToken(param, c.pg.Jwt.System)
	if ret.ErrorIs() {
		return rt.ErrorMessage(ret.Message)
	}
	tokenResult := ret.Data
	//判断密钥，是否需要保存
	{
		key := cacheAuthPubPrivPg.KeySystem()
		dataKey := entityRam.RamAsaJsonPrivatePublicKey{
			Private: tokenResult.PrivateKey,
			Public:  tokenResult.PublicKey,
		}
		if isMakeNewKey {
			toJson, _ := jsonPg.ObjToJson(dataKey)
			save := entityRam.RamAccountSessionAccessKeyEntity{
				Ano:        info.No,
				Data:       toJson,
				No:         typeDomainPg.System.Index(),
				TenantNo:   info.TenantNo,
				Key:        tokenResult.PublicKey,
				Type:       typeDomainPg.System.Index(),
				TypeDomain: typeDomainPg.System.Index(),
			}
			save.KindUnique = userPg.SaltMake(tokenResult.PublicKey, toJson+save.No+save.TenantNo+save.TypeDomain)
			c.sessionAk.Create(&save)
			//缓存
			cacheAuthPubPrivPg.Set(key, dataKey)
		}
		//缓存
		_, b2 := cacheAuthPubPrivPg.Get(key)
		if !b2 {
			c.log.Errorf("密钥不存在，重新加载")
			cacheAuthPubPrivPg.Set(key, dataKey)
		}
	}
	//记录登录日志
	{
		var tmp entityRam.RamAccountEntity
		copier.Copy(&tmp, info)
		tmp.LoginTime = &now
		tmp.TenantNo = typeDomainPg.System.Index()
		obj := modRamAccount.LoginLogDto{
			Account:     tmp,
			Ano:         tmp.No,
			AppNo:       "",
			Client:      "",
			LoginSource: "",
			Ip:          ctx.ClientIP(),
			ExtraData:   make(map[string]any),
		}
		ua := ctx.GetHeader(constHeaderPg.HeaderUserAgent)
		obj.ExtraData[constHeaderPg.HeaderUserAgent] = ua
		//保存到数据库
		eventBus.PublishEventAsync(constEventBusPg.RamAccountLoginLog, obj)
	}

	successInfo := modRamLogin.LoginSuccessInfo{
		Account: info.Account,
		Name:    info.Name,
	}
	success := modRamLogin.LoginSuccess{Info: successInfo, Token: tokenResult.Token, AccessToken: tokenResult.Token, AuthCode: []string{"AC_100100", "AC_100110", "AC_100120", "AC_100010"}}
	rt.Data = success
	return rt.Ok()
}

func (c *AccountLoginService) Logout(holder holderPg.HolderPg) (rt rg.Rs[string]) {
	return rt.Ok()
}

// RefreshToken
//
//	@Description:  刷新
//	@receiver c
func (c *AccountLoginService) RefreshToken(ctx *gin.Context, ct modRamLogin.TokenRefreshCt) (rt rg.Rs[modRamLogin.LoginSuccess]) {
	token := ctx.GetHeader("Authorization")
	rt.Data = modRamLogin.LoginSuccess{Token: token, AccessToken: token}
	return rt.Ok()
}

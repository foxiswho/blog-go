package service

import (
	"context"
	"reflect"
	"time"

	"github.com/farseer-go/eventBus"
	"github.com/gin-gonic/gin"
	"github.com/hongmengzhu/xianfu-blog-go/app/core/cache/cacheRam"
	"github.com/hongmengzhu/xianfu-blog-go/app/system/ram/model/modRamLogin"
	"github.com/hongmengzhu/xianfu-blog-go/infrastructure/entityRam"
	"github.com/hongmengzhu/xianfu-blog-go/infrastructure/repositoryRam"
	"github.com/hongmengzhu/xianfu-blog-go/middleware/components/authTokenPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/configPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/configPg/pg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/consts/constEventBusPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/consts/constHeaderPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/consts/constsRam/typeDomainPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/enum/state/clientPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/enum/state/enumStatePg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/holderPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/holderPg/multiTenantPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/log2"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/sdk/ram/model/modRamAccount"
	"github.com/jinzhu/copier"
	"github.com/pangu-2/go-tools/tools/cryptPg"
	"github.com/pangu-2/go-tools/tools/strPg"
	"github.com/pangu-2/go-tools/tools/userPg"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
	"go-spring.org/log"
	"go-spring.org/spring/gs"
)

func init() {
	gs.Provide(NewAccountLoginService).Init(func(s *AccountLoginService) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

// AccountLoginService 登录
// @Description:
type AccountLoginService struct {
	dao                  *repositoryRam.RamAccountRepository                 `autowire:"?"`
	daoAuth              *repositoryRam.RamAccountAuthorizationRepository    `autowire:"?"`
	sessionAk            *repositoryRam.RamAccountSessionAccessKeyRepository `autowire:"?"`
	cacheSessionPubPrive *cacheRam.CacheSessionPubPrive                      `autowire:"?" `
	pg                   configPg.Pg                                         `value:"${pg}"`
	log                  *log2.Logger                                        `autowire:"?"`
	authLogin            pg.Auth                                             `value:"${pg.auth}"`
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
func (c *AccountLoginService) Login(ctx *gin.Context, ct modRamLogin.LoginCt, tp typeDomainPg.TypeDomain) (rt rg.Rs[modRamLogin.LoginSuccess]) {
	c.log.Infof("tp=%+v,ct=%+v", tp, ct)
	//
	client := clientPg.Browser
	//
	if !c.dao.Config().Domain.System {
		return rt.ErrorMessage("系统管理模块已禁用，不允许操作")
	}
	if strPg.IsBlank(ct.Account) {
		return rt.ErrorMessage("账号不能为空")
	}
	if strPg.IsBlank(ct.Password) {
		return rt.ErrorMessage("密码不能为空")
	}
	pwd := ct.Password
	log.Infof(context.Background(), log.TagAppDef, "authLogin=%+v", c.authLogin)
	//解密密码
	if logingEn, ok := c.authLogin.LoginEncrypt["default"]; ok && logingEn {
		login := c.cacheSessionPubPrive.DecodeByLogin(ctx, pwd)
		if login.ErrorIs() {
			return rt.ErrorMessage(login.Message)
		}
		pwd = login.Data
	}
	log.Infof(context.Background(), log.TagAppDef, "pwd=%+v", pwd)

	md5 := cryptPg.Md5(ct.Account)
	info, b, err := c.dao.FindByAccountMd5AndTypeDomain(ctx, md5, tp.Code())
	if nil != err {
		return rt.Error()
	}
	if !b {
		return rt.ErrorMessage("账号不存在")
	}
	if !enumStatePg.ENABLE.IsExistInt8(info.State) {
		return rt.ErrorMessage("账户已被禁用，不能登陆")
	}
	pwdInfo, result := c.daoAuth.FindByTypePasswordANo(ctx, info.No)
	if !result {
		return rt.ErrorMessage("用户密码未设置")
	}
	if !userPg.PasswordVerify(pwdInfo.Value, pwd, pwdInfo.ExtraData) {
		c.log.Debugf("pwd=%+v,[%+v],[value]=%+v,[extra]=%+v", pwd, userPg.PasswordSalt(pwd, pwdInfo.ExtraData), pwdInfo.Value, pwdInfo.ExtraData)
		return rt.ErrorMessage("账号密码错误")
	}
	//
	now := time.Now()
	//租户默认
	mult := multiTenantPg.MultiTenantPg{
		TenantNo: make([]string, 0),
	}
	mult.TenantNo = append(mult.TenantNo, tp.Code())
	//
	privatePubKey := authTokenPg.Result{}
	// 获取 密钥对
	key := c.cacheSessionPubPrive.PaseKey(ctx, tp, client, typeDomainPg.System.Code(), nil)
	if strPg.IsNotBlank(key.Public) {
		privatePubKey.PrivateKey = key.Private
		privatePubKey.PublicKey = key.Public
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
		TenantNo:    typeDomainPg.System.Code(),
	}
	ret := authTokenPg.MakePaseToken(param, c.pg.Jwt.System)
	if ret.ErrorIs() {
		return rt.ErrorMessage(ret.Message)
	}
	tokenResult := ret.Data
	//记录登录日志
	{
		var tmp entityRam.RamAccountEntity
		copier.Copy(&tmp, info)
		tmp.LoginTime = &now
		tmp.TenantNo = typeDomainPg.System.Code()
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
	success := modRamLogin.LoginSuccess{
		Info:        successInfo,
		Token:       tokenResult.Token,
		AccessToken: tokenResult.Token,
		AuthCode:    []string{"AC_100100", "AC_100110", "AC_100120", "AC_100010"}}
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

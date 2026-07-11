package accountSessionRamEvent

import (
	"context"
	"encoding/json"

	"github.com/hongmengzhu/xianfu-blog-go/app/core/cache/cacheRam"
	"github.com/hongmengzhu/xianfu-blog-go/infrastructure/entityRam"
	"github.com/hongmengzhu/xianfu-blog-go/infrastructure/repositoryRam"
	"github.com/hongmengzhu/xianfu-blog-go/middleware/components/authTokenPg"
	"github.com/hongmengzhu/xianfu-blog-go/middleware/components/cachePg/cacheAuthPubPrivPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/consts/constsRam/typeDomainPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/consts/constsRam/typePubPrivePg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/enum/state/clientPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/log2"
	"github.com/pangu-2/go-tools/tools/jsonPg"
	"github.com/pangu-2/go-tools/tools/strPg"
	"github.com/pangu-2/go-tools/tools/userPg"
)

// InitSessionPubPrive
// @Description: 加载密钥
type InitSessionPubPrive struct {
	log                  *log2.Logger                                        `autowire:"?"`
	sessionAk            *repositoryRam.RamAccountSessionAccessKeyRepository `autowire:"?"`
	cacheSessionPubPrive *cacheRam.CacheSessionPubPrive                      `autowire:"?"`
}

func NewInitSessionPubPrive(log *log2.Logger,
	sessionAk *repositoryRam.RamAccountSessionAccessKeyRepository,
	cacheSessionPubPrive *cacheRam.CacheSessionPubPrive,
) *InitSessionPubPrive {
	return &InitSessionPubPrive{
		log:                  log,
		sessionAk:            sessionAk,
		cacheSessionPubPrive: cacheSessionPubPrive,
	}
}

func (c *InitSessionPubPrive) Processor(ctx context.Context) error {
	typeDomain := typeDomainPg.System
	client := clientPg.Browser
	pubPriv := typePubPrivePg.Login
	mapKey := make(map[string]*entityRam.RamAccountSessionAccessKeyEntity)
	// 系统
	data, r := c.sessionAk.FindByTypeDomainInAndState(ctx, []string{typeDomain.Index(), pubPriv.Index()})
	if r {
		for _, item := range data {
			//系统 浏览器
			if clientPg.Browser.IsEqual(item.Client) {
				mapKey[item.TypeDomain+client.Index()] = item
			}
			//登录密钥
			if typePubPrivePg.Login.IsEqual(item.TypeDomain) {
				mapKey[item.TypeDomain] = item
			}
		}
	}
	//系统 浏览器
	{
		entity := mapKey[typeDomain.Index()+client.Index()]
		c.keySystem(ctx, typeDomain, client, entity)
	}
	//登录密钥
	{
		entity := mapKey[pubPriv.Index()]
		c.loginPubPriveKey(ctx, pubPriv, entity)
	}
	//
	//租户
	c.keyTenant(ctx)
	return nil
}

func (c *InitSessionPubPrive) keyTenant(ctx context.Context) {
	//是否新生成密钥
	isMakeNewKey := false
	// 获取 密钥对
	data, r := c.sessionAk.FindByTypeDomainInAndState(ctx, []string{typeDomainPg.Manage.Index()})
	if r {
		for _, item := range data {
			//跳过系统
			if typeDomainPg.System.IsEqual(item.TypeDomain) {
				continue
			}
			if typeDomainPg.System.IsEqual(item.No) {
				continue
			}
			isMakeNewKey = false
			//
			privatePubKey := authTokenPg.Result{}
			if strPg.IsNotBlank(item.Data) {
				var privatePubKeyEnt entityRam.RamAsaJsonPrivatePublicKey
				err := json.Unmarshal([]byte(item.Data), &privatePubKeyEnt)
				if err != nil {
					privatePubKey = authTokenPg.MakePublicPrivateKey()
					isMakeNewKey = true
				} else {
					privatePubKey.PrivateKey = privatePubKeyEnt.Private
					privatePubKey.PublicKey = privatePubKeyEnt.Public
				}
			}
			dataKey := entityRam.RamAsaJsonPrivatePublicKey{
				Private: privatePubKey.PrivateKey,
				Public:  privatePubKey.PublicKey,
			}
			if !isMakeNewKey {
				//缓存
				cacheAuthPubPrivPg.Set(cacheAuthPubPrivPg.KeyManage(item.TenantNo), dataKey)
			}
		}
	}
}

// 系统
func (c *InitSessionPubPrive) keySystem(ctx context.Context,
	domain typeDomainPg.TypeDomain, client clientPg.Client, entity *entityRam.RamAccountSessionAccessKeyEntity) {
	//是否新生成密钥
	isMakeNewKey := false
	privatePubKey := authTokenPg.Result{}
	// 获取 密钥对
	if nil == entity {
		privatePubKey = authTokenPg.MakePublicPrivateKey()
		isMakeNewKey = true
	} else {
		//密钥不存在，生成
		if strPg.IsBlank(entity.Data) {
			privatePubKey = authTokenPg.MakePublicPrivateKey()
			isMakeNewKey = true
		} else {
			//解析
			var privatePubKeyEnt entityRam.RamAsaJsonPrivatePublicKey
			err := json.Unmarshal([]byte(entity.Data), &privatePubKeyEnt)
			if err != nil {
				//解析失败，重新生成
				privatePubKey = authTokenPg.MakePublicPrivateKey()
				isMakeNewKey = true
			} else {
				privatePubKey.PrivateKey = privatePubKeyEnt.Private
				privatePubKey.PublicKey = privatePubKeyEnt.Public
			}
		}
	}
	dataKey := entityRam.RamAsaJsonPrivatePublicKey{
		Private: privatePubKey.PrivateKey,
		Public:  privatePubKey.PublicKey,
	}
	//判断密钥，是否需要保存
	{
		if isMakeNewKey {
			toJson, _ := jsonPg.ObjToJson(dataKey)
			save := entityRam.RamAccountSessionAccessKeyEntity{
				Ano:        "",
				Data:       toJson,
				No:         domain.Index(),
				TenantNo:   domain.Index(),
				Key:        privatePubKey.PublicKey,
				Type:       domain.Index(),
				TypeDomain: domain.Index(),
				Client:     client.Index(),
			}
			save.KindUnique = userPg.SaltMake(privatePubKey.PublicKey, toJson+save.No+save.TenantNo+save.TypeDomain)
			c.sessionAk.Create(ctx, &save)
		}
	}
	//缓存
	cacheAuthPubPrivPg.Set(cacheAuthPubPrivPg.KeySystem(), dataKey)
}

// 登录密钥
func (c *InitSessionPubPrive) loginPubPriveKey(ctx context.Context,
	pubPriv typePubPrivePg.PubPrive, entity *entityRam.RamAccountSessionAccessKeyEntity) {
	isMakeNewKey := true
	privatePubKeyEnt := entityRam.RamAsaJsonPrivatePublicKey{}
	if nil != entity && strPg.IsNotBlank(entity.Data) {
		err := json.Unmarshal([]byte(entity.Data), &privatePubKeyEnt)
		if err == nil {
			//解析成功
			isMakeNewKey = false
		}
	}
	//重新生成,判断密钥，是否需要保存
	c.cacheSessionPubPrive.LoginPubPriveKeyByNew(ctx, isMakeNewKey, privatePubKeyEnt)
}

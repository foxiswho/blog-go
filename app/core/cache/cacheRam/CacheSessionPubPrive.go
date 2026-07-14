package cacheRam

import (
	"context"
	"strings"

	"github.com/goccy/go-json"
	"github.com/hongmengzhu/xianfu-blog-go/infrastructure/entityRam"
	"github.com/hongmengzhu/xianfu-blog-go/infrastructure/repositoryRam"
	"github.com/hongmengzhu/xianfu-blog-go/middleware/components/authTokenPg"
	"github.com/hongmengzhu/xianfu-blog-go/middleware/components/cachePg/cacheAuthPubPrivPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/cachePg/rdsPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/consts/constsCachePg/constsCacheRam"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/consts/constsRam/typeDomainPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/consts/constsRam/typePubPrivePg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/enum/state/clientPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/enum/state/enumStatePg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/model"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/tools/cryptoPg"
	"github.com/pangu-2/go-tools/tools/jsonPg"
	"github.com/pangu-2/go-tools/tools/strPg"
	"github.com/pangu-2/go-tools/tools/userPg"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
	"go-spring.org/log"
	"go-spring.org/spring/gs"
)

func init() {
	gs.Provide(new(CacheSessionPubPrive))
}

type CacheSessionPubPrive struct {
	rdt         *rdsPg.BatchString                                  `autowire:"?"`
	sessionAk   *repositoryRam.RamAccountSessionAccessKeyRepository `autowire:"?"`
	pubPriveKey typePubPrivePg.PubPrive
}

func (c *CacheSessionPubPrive) LoginPubPriveKey(ctx context.Context) model.AuthPubPriveDto {
	c.pubPriveKey = typePubPrivePg.Login
	get, b := cacheAuthPubPrivPg.Get(constsCacheRam.SessionPubPrive_login())
	if b {
		return model.AuthPubPriveDto{
			PrivateKey: get.Private,
			PublicKey:  get.Public,
		}
	}
	//是否新生成密钥
	isMakeNewKey := true
	var privatePubKeyEnt entityRam.RamAsaJsonPrivatePublicKey
	info, result := c.sessionAk.FindByTypeDomainAndState(ctx, c.pubPriveKey.Index())
	if result {
		isMakeNewKey = false
		if strPg.IsNotBlank(info.Data) {
			err := json.Unmarshal([]byte(info.Data), &privatePubKeyEnt)
			if err != nil {
				//解析失败，重新生成
				isMakeNewKey = true
			}
		}
	}
	return c.LoginPubPriveKeyByNew(ctx, isMakeNewKey, privatePubKeyEnt)
}

func (c *CacheSessionPubPrive) LoginPubPriveKeyByNew(ctx context.Context, isMakeNewKey bool, json entityRam.RamAsaJsonPrivatePublicKey) model.AuthPubPriveDto {
	c.pubPriveKey = typePubPrivePg.Login
	//重新生成
	if isMakeNewKey {
		ret := cryptoPg.Sm2GenerateKey()
		if ret.ErrorIs() {
			log.Debugf(ctx, log.TagBizDef, "生成密钥失败")
			return model.AuthPubPriveDto{}
		}
		privatePubKey := ret.Data
		//先删除
		err := c.sessionAk.DeleteByNo(ctx, c.pubPriveKey.Index())
		if err != nil {
			log.Debugf(ctx, log.TagBizDef, "删除登录密钥失败")
		}
		json.Private = privatePubKey.PrivateKey
		json.Public = privatePubKey.PublicKey
		toJson, _ := jsonPg.ObjToJson(json)
		save := entityRam.RamAccountSessionAccessKeyEntity{
			Ano:        "",
			Data:       toJson,
			No:         c.pubPriveKey.Index(),
			TenantNo:   c.pubPriveKey.Index(),
			Key:        privatePubKey.PublicKey,
			Type:       c.pubPriveKey.Index(),
			TypeDomain: c.pubPriveKey.Index(),
		}
		save.KindUnique = userPg.SaltMake(privatePubKey.PublicKey, toJson+save.No+save.TenantNo+save.TypeDomain)
		err, _ = c.sessionAk.Create(ctx, &save)
		if err != nil {
			log.Debugf(ctx, log.TagBizDef, "创建新密钥失败.%+v", err)
		}
	}

	cacheAuthPubPrivPg.Set(constsCacheRam.SessionPubPrive_login(), json)
	//
	return model.AuthPubPriveDto{
		PrivateKey: json.Private,
		PublicKey:  json.Public,
	}
}

func (c *CacheSessionPubPrive) DecodeByLogin(ctx context.Context, pwd string) (rt rg.Rs[string]) {
	log.Infof(ctx, log.TagBizDef, "ct.len=%+v,=%+v", len(pwd), pwd)
	if len(pwd) < 194 {
		return rt.ErrorMessage("密文过短，可能已损坏，实际长度：" + string(rune(len(pwd))))
	}
	key := c.LoginPubPriveKey(ctx)
	if !strings.HasPrefix(pwd, "04") {
		pwd = "04" + pwd
	}
	ret := cryptoPg.NewSm2(key.PrivateKey, key.PublicKey).DecodeHex(pwd)
	if ret.ErrorIs() {
		return rt.ErrorMessage(ret.Message)
	}
	return rt.OkData(ret.Data)
}

//	PaseKeyByNew 获取密钥对
//
// @Description:
// @receiver *CacheSessionPubPrive
// @param
// @return
func (c *CacheSessionPubPrive) PaseKeyByNew(ctx context.Context, isMakeNewKey bool,
	jsonEntity *entityRam.RamAsaJsonPrivatePublicKey, typeDomain typeDomainPg.TypeDomain, client clientPg.Client, tenantNo string) entityRam.RamAsaJsonPrivatePublicKey {
	privatePubKey := authTokenPg.Result{}
	//
	// 需要重新生成
	if isMakeNewKey {
		jsonEntity = &entityRam.RamAsaJsonPrivatePublicKey{}
		privatePubKey = authTokenPg.MakePublicPrivateKey()
		jsonEntity.Private = privatePubKey.PrivateKey
		jsonEntity.Public = privatePubKey.PublicKey
		// 删除已存在的
		c.sessionAk.DeleteByTypeDomainAndClientAndState(ctx, typeDomain.Index(), []string{client.Index()})
		//
		toJson, _ := jsonPg.ObjToJson(jsonEntity)
		save := entityRam.RamAccountSessionAccessKeyEntity{
			Ano:        "",
			Data:       toJson,
			No:         tenantNo + ":" + typeDomain.Index() + ":" + client.Index(),
			TenantNo:   tenantNo,
			TypeDomain: typeDomain.Index(),
			Client:     client.Index(),
			Type:       typeDomain.Index(),
			Key:        privatePubKey.PublicKey,
			State:      enumStatePg.ENABLE.Index(),
		}
		str := toJson + save.No + save.TenantNo + save.TypeDomain + save.Client
		save.KindUnique = userPg.SaltMake(str, save.Key)
		err, _ := c.sessionAk.Create(ctx, &save)
		if err != nil {
			log.Debugf(ctx, log.TagBizDef, "创建新密钥失败.%+v", err)
		}
	} else {
		//toJson, _ := jsonPg.ObjToJson(jsonEntity)
	}
	entity := *jsonEntity
	//缓存
	cacheAuthPubPrivPg.Set(cacheAuthPubPrivPg.Key(tenantNo, typeDomain, client), entity)
	return entity
}

//	PaseKeyByNew 获取密钥对
//
// @Description:
// @receiver *CacheSessionPubPrive
// @param
// @return
func (c *CacheSessionPubPrive) PaseKey(ctx context.Context, typeDomain typeDomainPg.TypeDomain, client clientPg.Client, tenantNo string, jsonEntity *entityRam.RamAsaJsonPrivatePublicKey) entityRam.RamAsaJsonPrivatePublicKey {
	get, b := cacheAuthPubPrivPg.Get(cacheAuthPubPrivPg.Key(tenantNo, typeDomain, client))
	if b {
		return get
	}
	isMakeNewKey := true
	info, result := c.sessionAk.FindByTenantNoAndTypeDomainInAndClientAndState(ctx, tenantNo, typeDomain.Index(), client.Index())
	if result {
		if strPg.IsNotBlank(info.Data) {
			isMakeNewKey = false
			jsonEntity = &entityRam.RamAsaJsonPrivatePublicKey{}
			err := json.Unmarshal([]byte(info.Data), jsonEntity)
			if err != nil {
				//解析失败，重新生成
				isMakeNewKey = true
			}
		}
	}
	return c.PaseKeyByNew(ctx, isMakeNewKey, jsonEntity, typeDomain, client, tenantNo)
}

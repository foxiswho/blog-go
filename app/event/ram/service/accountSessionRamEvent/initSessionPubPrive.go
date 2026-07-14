package accountSessionRamEvent

import (
	"context"
	"encoding/json"

	"github.com/hongmengzhu/xianfu-blog-go/app/core/cache/cacheRam"
	"github.com/hongmengzhu/xianfu-blog-go/infrastructure/entityRam"
	"github.com/hongmengzhu/xianfu-blog-go/infrastructure/repositoryRam"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/consts/constsRam/typeDomainPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/consts/constsRam/typePubPrivePg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/enum/state/clientPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/log2"
	"github.com/pangu-2/go-tools/tools/strPg"
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
			//登录密钥
			if typePubPrivePg.Login.IsEqual(item.TypeDomain) {
				mapKey[item.TypeDomain] = item
			}
		}
	}
	//系统 浏览器
	{
		c.keySystem(ctx, typeDomainPg.System, client)
	}
	//登录密钥
	{
		entity := mapKey[pubPriv.Index()]
		c.loginPubPriveKey(ctx, pubPriv, entity)
	}
	//
	//租户
	c.keyTenant(ctx, typeDomainPg.Tenant, client)
	return nil
}

func (c *InitSessionPubPrive) keyTenant(ctx context.Context, domain typeDomainPg.TypeDomain, client clientPg.Client) {
	c.cacheSessionPubPrive.PaseKey(ctx, domain, client, "1", nil)
}

// 系统
func (c *InitSessionPubPrive) keySystem(ctx context.Context,
	domain typeDomainPg.TypeDomain, client clientPg.Client) {
	c.cacheSessionPubPrive.PaseKey(ctx, domain, client, typeDomainPg.System.Index(), nil)
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

package accountSessionRamEvent

import (
	"context"
	"encoding/json"

	"github.com/foxiswho/blog-go/infrastructure/entityRam"
	"github.com/foxiswho/blog-go/infrastructure/repositoryRam"
	"github.com/foxiswho/blog-go/middleware/components/authTokenPg"
	"github.com/foxiswho/blog-go/middleware/components/cachePg/cacheAuthPubPrivPg"
	"github.com/foxiswho/blog-go/pkg/consts/constsRam/typeDomainPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/pangu-2/go-tools/tools/jsonPg"
	"github.com/pangu-2/go-tools/tools/strPg"
	"github.com/pangu-2/go-tools/tools/userPg"
)

// InitSessionPubPrive
// @Description: 加载密钥
type InitSessionPubPrive struct {
	log       *log2.Logger                                        `autowire:"?"`
	sessionAk *repositoryRam.RamAccountSessionAccessKeyRepository `autowire:"?"`
}

func NewInitSessionPubPrive(log *log2.Logger, sessionAk *repositoryRam.RamAccountSessionAccessKeyRepository) *InitSessionPubPrive {
	return &InitSessionPubPrive{
		log:       log,
		sessionAk: sessionAk,
	}
}

func (c *InitSessionPubPrive) Processor(ctx context.Context) error {
	// 系统
	c.keySystem()
	//租户
	c.keyTenant()
	return nil
}

func (c *InitSessionPubPrive) keyTenant() {
	//是否新生成密钥
	isMakeNewKey := false
	// 获取 密钥对
	data, r := c.sessionAk.FindByTypeDomainAndState([]string{typeDomainPg.Manage.Index()})
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

func (c *InitSessionPubPrive) keySystem() {
	//是否新生成密钥
	isMakeNewKey := false
	privatePubKey := authTokenPg.Result{}
	// 获取 密钥对
	no, r := c.sessionAk.FindByNoAndState(typeDomainPg.System.Index())
	if !r {
		privatePubKey = authTokenPg.MakePublicPrivateKey()
		isMakeNewKey = true
	} else {
		//密钥不存在，生成
		if strPg.IsBlank(no.Data) {
			privatePubKey = authTokenPg.MakePublicPrivateKey()
			isMakeNewKey = true
		} else {
			//解析
			var privatePubKeyEnt entityRam.RamAsaJsonPrivatePublicKey
			err := json.Unmarshal([]byte(no.Data), &privatePubKeyEnt)
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
				No:         typeDomainPg.System.Index(),
				TenantNo:   typeDomainPg.System.Index(),
				Key:        privatePubKey.PublicKey,
				Type:       typeDomainPg.System.Index(),
				TypeDomain: typeDomainPg.System.Index(),
			}
			save.KindUnique = userPg.SaltMake(privatePubKey.PublicKey, toJson+save.No+save.TenantNo+save.TypeDomain)
			c.sessionAk.Create(&save)
		}
	}
	//缓存
	cacheAuthPubPrivPg.Set(cacheAuthPubPrivPg.KeySystem(), dataKey)
}

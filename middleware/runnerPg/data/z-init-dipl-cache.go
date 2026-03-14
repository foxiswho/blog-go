package data

import (
	"github.com/foxiswho/blog-go/infrastructure/entityApi"
	"github.com/foxiswho/blog-go/infrastructure/repositoryApi"
	"github.com/foxiswho/blog-go/middleware/components/cachePg/cacheDiplPg"
	"github.com/foxiswho/blog-go/pkg/enum/state/enumStatePg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/pangu-2/go-tools/tools/datetimePg"
	"gorm.io/gorm"
)

// ZInitDiplCache
// @Description: 初始化 dipl 缓存
type ZInitDiplCache struct {
	log *log2.Logger                              `autowire:"?"`
	sv  *repositoryApi.ApiDiplAccessKeyRepository `autowire:"?"`
}

func (b *ZInitDiplCache) Run() error {
	b.log.Infof("初始化 => 接口密钥")
	var query entityApi.ApiDiplAccessKeyEntity
	query.State = enumStatePg.ENABLE.Index()
	//过期时间 超过当前时间的数据
	infos := b.sv.FindAll(query, repositoryPg.WithCondition(func(db *gorm.DB) *gorm.DB {
		db = db.Order("create_at desc")
		db.Where("expiry_date >= ?", datetimePg.Now())
		return db
	}))
	if infos != nil && len(infos) > 0 {
		for _, info := range infos {
			// 添加缓存
			//加入缓存
			sha := cacheDiplPg.HashSha(info.Key, info.Secret)
			obj := cacheDiplPg.DiplCo{
				HashSha:  sha,
				No:       info.DiplNo,
				TenantNo: info.TenantNo,
				Key:      info.Key,
				Secret:   info.Secret,
			}
			cacheDiplPg.Set(info.Key, obj)
		}
	}
	return nil
}

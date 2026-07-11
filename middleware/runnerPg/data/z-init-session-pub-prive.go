package data

import (
	"context"

	"github.com/hongmengzhu/xianfu-blog-go/app/core/cache/cacheRam"
	"github.com/hongmengzhu/xianfu-blog-go/app/event/ram/service/accountSessionRamEvent"
	"github.com/hongmengzhu/xianfu-blog-go/infrastructure/repositoryRam"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/log2"
	"go-spring.org/log"
	_ "go-spring.org/spring/gs"
)

// 加载 密钥缓存
type InitSessionPubPrive struct {
	log                  *log2.Logger                                        `autowire:"?"`
	sessionAk            *repositoryRam.RamAccountSessionAccessKeyRepository `autowire:"?"`
	cacheSessionPubPrive *cacheRam.CacheSessionPubPrive                      `autowire:"?"`
}

func (b *InitSessionPubPrive) Run(ctx context.Context) error {
	log.Infof(context.Background(), log.TagAppDef, "[init].[密钥缓存初始化]===================")
	accountSessionRamEvent.NewInitSessionPubPrive(b.log, b.sessionAk, b.cacheSessionPubPrive).Processor(context.Background())
	return nil
}

package data

import (
	"context"

	"github.com/foxiswho/blog-go/app/event/ram/service/accountSessionRamEvent"
	"github.com/foxiswho/blog-go/infrastructure/repositoryRam"
	"github.com/foxiswho/blog-go/pkg/log2"
	syslog "github.com/go-spring/log"
)

// 加载 密钥缓存
type InitSessionPubPrive struct {
	log       *log2.Logger                                        `autowire:"?"`
	sessionAk *repositoryRam.RamAccountSessionAccessKeyRepository `autowire:"?"`
}

func (b *InitSessionPubPrive) Run() error {
	syslog.Infof(context.Background(), syslog.TagAppDef, "[init].[密钥缓存初始化]===================")
	accountSessionRamEvent.NewInitSessionPubPrive(b.log, b.sessionAk).Processor(context.Background())
	return nil
}

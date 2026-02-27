package data

import (
	"context"

	"github.com/foxiswho/blog-go/app/event/basic/service/tagsBasicEvent"
	"github.com/foxiswho/blog-go/pkg/log2"
	syslog "github.com/go-spring/log"
)

// ZInitTagsCache
// @Description: 启动后初始化一些数据
type ZInitTagsCache struct {
	log *log2.Logger       `autowire:"?"`
	sp  *tagsBasicEvent.Sp `autowire:"?"`
}

func (c *ZInitTagsCache) Run() error {
	syslog.Infof(context.Background(), syslog.TagAppDef, "[init].[标签缓存]===================")
	err := tagsBasicEvent.NewStartInit(c.sp).Processor(context.Background())
	if err != nil {
		c.log.Error("error:", err)
	}
	return nil
}

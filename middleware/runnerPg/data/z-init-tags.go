package data

import (
	"context"

	"github.com/hongmengzhu/xianfu-blog-go/app/event/basic/service/tagsBasicEvent"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/log2"
	_ "go-spring.org/spring/gs"
)

// ZInitTagsCache
// @Description: 启动后初始化一些数据
type ZInitTagsCache struct {
	log *log2.Logger       `autowire:"?"`
	sp  *tagsBasicEvent.Sp `autowire:"?"`
}

func (c *ZInitTagsCache) Run(ctx context.Context) error {
	c.log.Infof("[init].[标签缓存]===================")
	err := tagsBasicEvent.NewStartInit(c.sp).Processor(context.Background())
	if err != nil {
		c.log.Error("error:", err)
	}
	return nil
}

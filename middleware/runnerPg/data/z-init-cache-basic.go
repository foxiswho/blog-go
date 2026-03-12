package data

import (
	"context"

	"github.com/foxiswho/blog-go/app/event/basic/service/eventBasicEvent"
	"github.com/foxiswho/blog-go/app/event/basic/service/eventBasicRules"
	"github.com/foxiswho/blog-go/pkg/log2"
)

// ZInitCacheBasic
// @Description: 启动后初始化一些数据
type ZInitCacheBasic struct {
	log *log2.Logger `autowire:"?"`
}

func (c *ZInitCacheBasic) Run() error {
	c.log.Infof("[init].[启动初始化.基础.缓存]===================")
	{
		err := eventBasicEvent.NewStartInit(c.log).Processor(context.Background())
		if err != nil {
			c.log.Error("error:", err)
		}
	}
	{
		err := eventBasicRules.NewStartInit(c.log).Processor(context.Background())
		if err != nil {
			c.log.Error("error:", err)
		}
	}
	return nil
}

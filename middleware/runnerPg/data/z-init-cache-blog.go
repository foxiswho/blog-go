package data

import (
	"context"

	"github.com/foxiswho/blog-go/app/event/blog/service/articleBlogEvent"
	"github.com/foxiswho/blog-go/pkg/log2"
)

// ZInitCacheBlog
// @Description: 启动后初始化一些数据
type ZInitCacheBlog struct {
	log *log2.Logger `autowire:"?"`
}

func (c *ZInitCacheBlog) Run() error {
	c.log.Infof("[init].[博客.分类.缓存]===================")
	err := articleBlogEvent.NewStartInit(c.log).Processor(context.Background())
	if err != nil {
		c.log.Error("error:", err)
	}
	return nil
}

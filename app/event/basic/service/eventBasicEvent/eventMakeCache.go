package eventBasicEvent

import (
	"context"

	"github.com/foxiswho/blog-go/app/event/basic/model/modEventBasicEvent"
	"github.com/foxiswho/blog-go/pkg/log2"
)

// 缓存处理
type EventMakeCache struct {
	log *log2.Logger `autowire:"?"`
	Sp  Sp           `autowire:"?"`
	ct  modEventBasicEvent.EventDto
}

func (c *EventMakeCache) Processor(ctx context.Context) error {
	return nil
}

// ThisTenantAll 当前租户下全部
func (c *EventMakeCache) ThisTenantAll(ctx context.Context) error {
	return nil
}

// All 全部
func (c *EventMakeCache) All(ctx context.Context) error {
	return nil
}

// Nos 指定编号
func (c *EventMakeCache) Nos(ctx context.Context) error {
	return nil
}

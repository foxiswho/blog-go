package eventBasicEvent

import (
	"context"

	"github.com/foxiswho/blog-go/app/event/basic/model/modEventBasicEvent"
	"github.com/foxiswho/blog-go/pkg/log2"
)

type EventFieldMakeCache struct {
	Log *log2.Logger `autowire:"?"`
	Sp  *Sp          `autowire:"?"`
	ct  modEventBasicEvent.FieldDto
}

func NewEventFieldMakeCache(sp *Sp, ct modEventBasicEvent.FieldDto) *EventFieldMakeCache {
	return &EventFieldMakeCache{
		Sp:  sp,
		Log: sp.Log,
		ct:  ct,
	}
}

func (c *EventFieldMakeCache) Processor(ctx context.Context) error {
	return nil
}

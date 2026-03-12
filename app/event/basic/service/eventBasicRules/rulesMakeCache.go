package eventBasicRules

import (
	"context"

	"github.com/foxiswho/blog-go/app/event/basic/model/modEventBasicEvent"
	"github.com/foxiswho/blog-go/pkg/log2"
)

type RulesMakeCache struct {
	Log *log2.Logger `autowire:"?"`
	Sp  Sp           `autowire:"?"`
	ct  modEventBasicEvent.EventDto
}

func (c *RulesMakeCache) Processor(ctx context.Context) error {
	return nil
}

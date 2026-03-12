package listenerBasic

import (
	"context"

	"github.com/farseer-go/eventBus"
	"github.com/farseer-go/fs/core"
	"github.com/foxiswho/blog-go/app/event/basic/model/modEventBasicEvent"
	"github.com/foxiswho/blog-go/app/event/basic/service/eventBasicEvent"
	"github.com/foxiswho/blog-go/pkg/consts/constEventBusPg"
	"github.com/foxiswho/blog-go/pkg/log2"
)

type EventCacheListener struct {
	log *log2.Logger        `autowire:"?"`
	sp  *eventBasicEvent.Sp `autowire:"?"`
}

// Run 启动加载
//
//	@Description:
//	@receiver c
//	@param ctx
func (c *EventCacheListener) Run() error {
	c.log.Infof("[init].listener.[基础.模型事件.缓存]===================")
	//模型事件
	eventBus.RegisterEvent(constEventBusPg.BasicConfigEventCache).RegisterSubscribe(constEventBusPg.BasicConfigEventCache, func(message any, _ core.EventArgs) {
		c.log.Infof("listener.[基础.模型事件.缓存]22===================")
		dto := message.(modEventBasicEvent.EventDto)
		//c.log.Infof("dto=%+v", dto)
		err := eventBasicEvent.NewEventMakeCache(c.sp, dto).Processor(context.Background())
		if nil != err {
			c.sp.Log.Error("基础.模型事件.缓存:%+v", err)
		}
		message = nil
	})
	//模型事件字段
	eventBus.RegisterEvent(constEventBusPg.BasicConfigEventFieldCache).RegisterSubscribe(constEventBusPg.BasicConfigEventFieldCache, func(message any, _ core.EventArgs) {
		c.log.Infof("listener.[基础.模型事件字段.缓存]22===================")
		dto := message.(modEventBasicEvent.FieldDto)
		//c.log.Infof("dto=%+v", dto)
		err := eventBasicEvent.NewEventFieldMakeCache(c.sp, dto).Processor(context.Background())
		if nil != err {
			c.sp.Log.Error("基础.模型事件字段.缓存:%+v", err)
		}
		message = nil
	})
	return nil
}

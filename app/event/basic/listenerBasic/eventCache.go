package listenerBasic

import (
	"context"

	"github.com/farseer-go/eventBus"
	"github.com/farseer-go/fs/core"
	"github.com/hongmengzhu/xianfu-blog-go/app/event/basic/model/modEventBasicEvent"
	"github.com/hongmengzhu/xianfu-blog-go/app/event/basic/service/eventBasicEvent"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/consts/constEventBusPg"
	"go-spring.org/log"
	_ "go-spring.org/spring/gs"
)

type EventCacheListener struct {
	sp *eventBasicEvent.Sp `autowire:"?"`
}

// Run 启动加载
//
//	@Description:
//	@receiver c
//	@param ctx
func (c *EventCacheListener) Run(ctx context.Context) error {
	log.Infof(ctx, log.TagAppDef, "[init].listener.[基础.模型事件.缓存]===================")
	//模型事件
	eventBus.RegisterEvent(constEventBusPg.BasicConfigEventCache).RegisterSubscribe(constEventBusPg.BasicConfigEventCache, func(message any, _ core.EventArgs) {
		log.Infof(ctx, log.TagAppDef, "listener.[基础.模型事件.缓存]22===================")
		dto := message.(modEventBasicEvent.EventDto)
		//log.Infof(ctx, log.TagAppDef,"dto=%+v", dto)
		err := eventBasicEvent.NewEventMakeCache(c.sp, dto).Processor(context.Background())
		if nil != err {
			log.Errorf(ctx, log.TagAppDef, "基础.模型事件.缓存:%+v", err)
		}
		message = nil
	})
	//模型事件字段
	eventBus.RegisterEvent(constEventBusPg.BasicConfigEventFieldCache).RegisterSubscribe(constEventBusPg.BasicConfigEventFieldCache, func(message any, _ core.EventArgs) {
		log.Infof(ctx, log.TagAppDef, "listener.[基础.模型事件字段.缓存]22===================")
		dto := message.(modEventBasicEvent.FieldDto)
		//log.Infof(ctx, log.TagAppDef,"dto=%+v", dto)
		err := eventBasicEvent.NewEventFieldMakeCache(c.sp, dto).Processor(context.Background())
		if nil != err {
			log.Errorf(ctx, log.TagAppDef, "基础.模型事件字段.缓存:%+v", err)
		}
		message = nil
	})
	return nil
}

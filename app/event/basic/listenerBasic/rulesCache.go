package listenerBasic

import (
	"context"

	"github.com/farseer-go/eventBus"
	"github.com/farseer-go/fs/core"
	"github.com/hongmengzhu/xianfu-blog-go/app/event/basic/model/modEventBasicRules"
	"github.com/hongmengzhu/xianfu-blog-go/app/event/basic/service/eventBasicRules"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/consts/constEventBusPg"
	"go-spring.org/log"
	_ "go-spring.org/spring/gs"
)

type RulesCacheListener struct {
	sp *eventBasicRules.Sp `autowire:"?"`
}

// Run 启动加载
//
//	@Description:
//	@receiver c
//	@param ctx
func (c *RulesCacheListener) Run(ctx context.Context) error {
	log.Infof(ctx, log.TagAppDef, "[init].listener.[基础.模型字段规则.缓存]===================")
	//模型事件
	eventBus.RegisterEvent(constEventBusPg.BasicModelRulesCache).RegisterSubscribe(constEventBusPg.BasicModelRulesCache, func(message any, _ core.EventArgs) {
		log.Infof(ctx, log.TagAppDef, "listener.[基础.模型字段规则.缓存]22===================")
		dto := message.(modEventBasicRules.RulesDto)
		//log.Infof(ctx, log.TagAppDef,"dto=%+v", dto)
		err := eventBasicRules.NewRulesMakeCache(c.sp, dto).Processor(context.Background())
		if nil != err {
			log.Errorf(ctx, log.TagAppDef, "基础.模型事件.缓存:%+v", err)
		}
		message = nil
	})
	return nil
}

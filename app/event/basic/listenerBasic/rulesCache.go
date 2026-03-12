package listenerBasic

import (
	"context"

	"github.com/farseer-go/eventBus"
	"github.com/farseer-go/fs/core"
	"github.com/foxiswho/blog-go/app/event/basic/model/modEventBasicRules"
	"github.com/foxiswho/blog-go/app/event/basic/service/eventBasicRules"
	"github.com/foxiswho/blog-go/pkg/consts/constEventBusPg"
	"github.com/foxiswho/blog-go/pkg/log2"
)

type RulesCacheListener struct {
	log *log2.Logger        `autowire:"?"`
	sp  *eventBasicRules.Sp `autowire:"?"`
}

// Run 启动加载
//
//	@Description:
//	@receiver c
//	@param ctx
func (c *RulesCacheListener) Run() error {
	c.log.Infof("[init].listener.[基础.模型字段规则.缓存]===================")
	//模型事件
	eventBus.RegisterEvent(constEventBusPg.BasicModelRulesCache).RegisterSubscribe(constEventBusPg.BasicModelRulesCache, func(message any, _ core.EventArgs) {
		c.log.Infof("listener.[基础.模型字段规则.缓存]22===================")
		dto := message.(modEventBasicRules.RulesDto)
		//c.log.Infof("dto=%+v", dto)
		err := eventBasicRules.NewRulesMakeCache(c.sp, dto).Processor(context.Background())
		if nil != err {
			c.sp.Log.Error("基础.模型事件.缓存:%+v", err)
		}
		message = nil
	})
	return nil
}

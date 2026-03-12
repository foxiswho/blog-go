package eventBasicRules

import (
	"context"

	"github.com/farseer-go/eventBus"
	"github.com/foxiswho/blog-go/app/event/basic/model/modEventBasicRules"
	"github.com/foxiswho/blog-go/pkg/consts/constEventBusPg"
	"github.com/foxiswho/blog-go/pkg/log2"
)

// StartInit 启动后初始化 所有租户 分类缓存
type StartInit struct {
	log *log2.Logger `autowire:"?"`
}

func NewStartInit(log *log2.Logger) *StartInit {
	return &StartInit{
		log: log,
	}
}

func (c *StartInit) Processor(ctx context.Context) error {
	//保存到数据库
	err := eventBus.PublishEventAsync(constEventBusPg.BasicModelRulesCache, modEventBasicRules.RulesDto{
		IsAll: true,
	})
	if err != nil {
		c.log.Errorf("copier.Copy error: %+v", err)
		return nil
	}
	return nil
}

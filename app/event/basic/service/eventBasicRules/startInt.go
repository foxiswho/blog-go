package eventBasicRules

import (
	"context"

	"github.com/farseer-go/eventBus"
	"github.com/hongmengzhu/xianfu-blog-go/app/event/basic/model/modEventBasicRules"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/consts/constEventBusPg"
	"go-spring.org/log"
)

// StartInit 启动后初始化 所有租户 分类缓存
type StartInit struct {
}

func NewStartInit() *StartInit {
	return &StartInit{}
}

func (c *StartInit) Processor(ctx context.Context) error {
	//保存到数据库
	err := eventBus.PublishEventAsync(constEventBusPg.BasicModelRulesCache, modEventBasicRules.RulesDto{
		IsAll: true,
	})
	if err != nil {
		log.Errorf(ctx, log.TagAppDef, "copier.Copy error: %+v", err)
		return nil
	}
	return nil
}

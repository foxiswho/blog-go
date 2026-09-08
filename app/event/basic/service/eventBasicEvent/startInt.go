package eventBasicEvent

import (
	"context"

	"github.com/farseer-go/eventBus"
	"github.com/hongmengzhu/xianfu-blog-go/app/event/basic/model/modEventBasicEvent"
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
	{
		err := eventBus.PublishEventAsync(constEventBusPg.BasicConfigEventCache, modEventBasicEvent.EventDto{
			IsAll: true,
		})
		if err != nil {
			log.Errorf(ctx, log.TagAppDef, "copier.Copy error: %+v", err)
			return nil
		}
	}
	{
		err := eventBus.PublishEventAsync(constEventBusPg.BasicConfigEventFieldCache, modEventBasicEvent.FieldDto{
			IsAll: true,
		})
		if err != nil {
			log.Errorf(ctx, log.TagAppDef, "copier.Copy error: %+v", err)
			return nil
		}
	}
	return nil
}

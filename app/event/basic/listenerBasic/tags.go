package listenerBasic

import (
	"context"

	"github.com/farseer-go/eventBus"
	"github.com/farseer-go/fs/core"
	"github.com/hongmengzhu/xianfu-blog-go/app/event/basic/model/modEventBasicTags"
	"github.com/hongmengzhu/xianfu-blog-go/app/event/basic/service/tagsBasicEvent"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/consts/constEventBusPg"
	"github.com/pangu-2/go-tools/tools/strPg"
	"go-spring.org/log"
	_ "go-spring.org/spring/gs"
)

// TagsListener 标签处理
// @Description:
type TagsListener struct {
	sp *tagsBasicEvent.Sp `autowire:"?"`
}

// Run 启动加载
//
//	@Description:
//	@receiver c
//	@param ctx
func (c *TagsListener) Run(ctx context.Context) error {
	log.Infof(context.Background(), log.TagAppDef, "eventBus.Register=%+v", constEventBusPg.BlogArticle)
	//博客文章
	eventBus.RegisterEvent(constEventBusPg.BlogArticle).RegisterSubscribe(constEventBusPg.BlogArticle, func(message any, _ core.EventArgs) {
		log.Infof(ctx, log.TagAppDef, "SchedulerEvent.event=%+v", message)
		dto := message.(modEventBasicTags.TagsRelation)
		if strPg.IsNotBlank(dto.Category) {
			err := tagsBasicEvent.NewSaveByCategory(c.sp, dto).Processor()
			if nil != err {
				log.Errorf(ctx, log.TagAppDef, "", err)
			}
			message = nil
		}
	})
	return nil
}

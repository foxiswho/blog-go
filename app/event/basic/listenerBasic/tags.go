package listenerBasic

import (
	"context"

	"github.com/farseer-go/eventBus"
	"github.com/farseer-go/fs/core"
	"github.com/foxiswho/blog-go/app/event/basic/model/modEventBasicTags"
	"github.com/foxiswho/blog-go/app/event/basic/service/tagsBasicEvent"
	"github.com/foxiswho/blog-go/pkg/consts/constEventBusPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	syslog "github.com/go-spring/log"
	"github.com/pangu-2/go-tools/tools/strPg"
)

// TagsListener 标签处理
// @Description:
type TagsListener struct {
	log *log2.Logger       `autowire:"?"`
	sp  *tagsBasicEvent.Sp `autowire:"?"`
}

// Run 启动加载
//
//	@Description:
//	@receiver c
//	@param ctx
func (c *TagsListener) Run() error {
	syslog.Infof(context.Background(), syslog.TagAppDef, "eventBus.Register=%+v", constEventBusPg.BlogArticle)
	//博客文章
	eventBus.RegisterEvent(constEventBusPg.BlogArticle).RegisterSubscribe(constEventBusPg.BlogArticle, func(message any, _ core.EventArgs) {
		c.log.Infof("SchedulerEvent.event=%+v", message)
		dto := message.(modEventBasicTags.TagsRelation)
		if strPg.IsNotBlank(dto.Category) {
			err := tagsBasicEvent.NewSaveByCategory(c.sp, dto).Processor()
			if nil != err {
				c.log.Error("", err)
			}
			message = nil
		}
	})
	return nil
}

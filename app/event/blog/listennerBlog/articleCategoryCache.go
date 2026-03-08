package listennerBlog

import (
	"context"

	"github.com/farseer-go/eventBus"
	"github.com/farseer-go/fs/core"
	"github.com/foxiswho/blog-go/app/event/blog/model/modEventBlogArticleCategory"
	"github.com/foxiswho/blog-go/app/event/blog/service/articleBlogEvent"
	"github.com/foxiswho/blog-go/pkg/consts/constEventBusPg"
	"github.com/foxiswho/blog-go/pkg/log2"
)

// ArticleCategoryCacheListener 文章分类处理
type ArticleCategoryCacheListener struct {
	log *log2.Logger         `autowire:"?"`
	sp  *articleBlogEvent.Sp `autowire:"?"`
}

// Run 启动加载
//
//	@Description:
//	@receiver c
//	@param ctx
func (c *ArticleCategoryCacheListener) Run() error {
	c.log.Infof("[init].listener.[博客.分类.缓存]===================")
	//博客文章 分类
	eventBus.RegisterEvent(constEventBusPg.BlogArticleCategoryCache).RegisterSubscribe(constEventBusPg.BlogArticleCategoryCache, func(message any, _ core.EventArgs) {
		//c.log.Infof("[init].listener.[博客.分类.缓存]22===================")
		dto := message.(modEventBlogArticleCategory.CacheDto)
		//c.log.Infof("dto=%+v", dto)
		err := articleBlogEvent.NewCategoryCache(c.sp, dto).Processor(context.Background())
		if nil != err {
			c.sp.Log.Error("博客.分类.缓存:%+v", err)
		}
		message = nil
	})
	return nil
}

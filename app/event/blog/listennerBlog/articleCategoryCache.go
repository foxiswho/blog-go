package listennerBlog

import (
	"context"

	"github.com/farseer-go/eventBus"
	"github.com/farseer-go/fs/core"
	"github.com/foxiswho/blog-go/app/event/blog/model/modEventBlogArticleCategory"
	"github.com/foxiswho/blog-go/app/event/blog/service/articleBlogEvent"
	"github.com/foxiswho/blog-go/pkg/consts/constEventBusPg"
)

// ArticleCategoryCacheListener 文章分类处理
type ArticleCategoryCacheListener struct {
	sp *articleBlogEvent.Sp `autowire:"?"`
}

// Run 启动加载
//
//	@Description:
//	@receiver c
//	@param ctx
func (c *ArticleCategoryCacheListener) Run() error {
	//博客文章 分类
	eventBus.RegisterEvent(constEventBusPg.BlogArticleCategoryCache).RegisterSubscribe(constEventBusPg.BlogArticleCategoryCache, func(message any, _ core.EventArgs) {
		dto := message.(modEventBlogArticleCategory.CacheDto)
		err := articleBlogEvent.NewCategoryCache(c.sp, dto).Processor(context.Background())
		if nil != err {
			c.sp.Log.Error("", err)
		}
		message = nil
	})
	return nil
}

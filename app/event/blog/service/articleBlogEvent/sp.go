package articleBlogEvent

import (
	"github.com/hongmengzhu/xianfu-blog-go/infrastructure/repositoryBlog"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/cachePg/rdsPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/log2"
	"go-spring.org/spring/gs"
)

func init() {
	gs.Provide(new(Sp))
}

type Sp struct {
	Log    *log2.Logger                                  `autowire:"?"`
	rdt    *rdsPg.BatchString                            `autowire:"?"`
	catRep *repositoryBlog.BlogArticleCategoryRepository `autowire:"?"`
}

package articleBlogEvent

import (
	"github.com/hongmengzhu/xianfu-blog-go/infrastructure/repositoryBlog"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/cachePg/rdsPg"
	"go-spring.org/spring/gs"
)

func init() {
	gs.Provide(new(Sp))
}

type Sp struct {
	rdt    *rdsPg.BatchString                            `autowire:"?"`
	catRep *repositoryBlog.BlogArticleCategoryRepository `autowire:"?"`
}

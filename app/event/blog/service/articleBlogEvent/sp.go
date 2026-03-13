package articleBlogEvent

import (
	"github.com/foxiswho/blog-go/infrastructure/repositoryBlog"
	"github.com/foxiswho/blog-go/pkg/cachePg/rdsPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/go-spring/spring-core/gs"
)

func init() {
	gs.Provide(new(Sp))
}

type Sp struct {
	Log    *log2.Logger                                  `autowire:"?"`
	rdt    *rdsPg.BatchString                            `autowire:"?"`
	catRep *repositoryBlog.BlogArticleCategoryRepository `autowire:"?"`
}

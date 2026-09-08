package tagsBasicEvent

import (
	"github.com/hongmengzhu/xianfu-blog-go/infrastructure/repositoryBasic"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/cachePg/rdsPg"
	"go-spring.org/spring/gs"
)

func init() {
	gs.Provide(new(Sp))
}

type Sp struct {
	dao     *repositoryBasic.BasicAttachmentRepository   `autowire:"?"`
	TagRela *repositoryBasic.BasicTagsRelationRepository `autowire:"?"`
	TagCate *repositoryBasic.BasicTagsCategoryRepository `autowire:"?"`
	TagsDb  *repositoryBasic.BasicTagsRepository         `autowire:"?"`
	rdt     *rdsPg.BatchString                           `autowire:"?"`
}

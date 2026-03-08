package tagsBasicEvent

import (
	"github.com/foxiswho/blog-go/infrastructure/repositoryBasic"
	"github.com/foxiswho/blog-go/pkg/cachePg/rdsPg"
	"github.com/foxiswho/blog-go/pkg/log2"
)

type Sp struct {
	log     *log2.Logger                                 `autowire:"?"`
	dao     *repositoryBasic.BasicAttachmentRepository   `autowire:"?"`
	TagRela *repositoryBasic.BasicTagsRelationRepository `autowire:"?"`
	TagCate *repositoryBasic.BasicTagsCategoryRepository `autowire:"?"`
	TagsDb  *repositoryBasic.BasicTagsRepository         `autowire:"?"`
	rdt     *rdsPg.BatchString                           `autowire:"?"`
}

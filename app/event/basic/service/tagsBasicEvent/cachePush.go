package tagsBasicEvent

import (
	"context"
	"github.com/foxiswho/blog-go/app/event/basic/model/modEventBasicTags"
	"github.com/foxiswho/blog-go/infrastructure/entityBasic"
	"github.com/foxiswho/blog-go/pkg/cachePg/rdsPg"
	"github.com/foxiswho/blog-go/pkg/enum/state/enumStatePg"
	"strings"
)

// CachePush
// @Description: 缓存更新
type CachePush struct {
	rdt *rdsPg.BatchString `autowire:"?"`
	sp  *Sp                `autowire:"?"`
	dto modEventBasicTags.TagsCacheDto
}

func NewCachePush(sp *Sp, dto modEventBasicTags.TagsCacheDto) *CachePush {
	return &CachePush{
		sp:  sp,
		rdt: sp.rdt,
		dto: dto,
	}
}

func (c *CachePush) Processor(ctx context.Context) error {
	t := entityBasic.BasicTagsRelationEntity{
		State: enumStatePg.ENABLE.Index(),
	}
	if nil == c.dto.CategoryRoot {
		return nil
	}
	if len(c.dto.CategoryRoot) < 1 {
		return nil
	}
	infos, result := c.sp.TagRela.FindAllByCategoryRootIn(t, c.dto.CategoryRoot)
	if result {
		data := make(map[string]interface{})
		for _, item := range infos {
			tmp := item.CategoryRoot + ":" + strings.TrimSpace(item.Name)
			data[tmp] = strings.TrimSpace(item.TagNo)
		}
		c.rdt.SetPipeline(ctx, data)
	}
	return nil
}

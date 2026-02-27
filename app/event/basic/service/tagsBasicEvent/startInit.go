package tagsBasicEvent

import (
	"context"
	"github.com/foxiswho/blog-go/infrastructure/entityBasic"
	"github.com/foxiswho/blog-go/pkg/cachePg/rdsPg"
	"github.com/foxiswho/blog-go/pkg/enum/state/enumStatePg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/pangu-2/go-tools/tools/cryptPg"
	"strings"
)

// StartInit
// @Description: 启动后初始化一些数据
type StartInit struct {
	rdb *rdsPg.BatchString `autowire:"?"`
	log *log2.Logger       `autowire:"?"`
	sp  *Sp                `autowire:"?"`
}

func NewStartInit(sp *Sp) *StartInit {
	return &StartInit{
		sp:  sp,
		rdb: sp.rdt,
		log: sp.log,
	}
}

func (c *StartInit) Processor(ctx context.Context) error {
	t := entityBasic.BasicTagsRelationEntity{
		State: enumStatePg.ENABLE.Index(),
	}
	infos := c.sp.TagRela.FindAll(t)
	if nil != infos && len(infos) > 0 {
		data := make(map[string]interface{})
		for _, item := range infos {
			//防止过长
			md5 := cryptPg.Md5(strings.TrimSpace(item.Name))
			tmp := item.CategoryRoot + ":" + md5
			data[tmp] = strings.TrimSpace(item.TagNo)
		}
		c.rdb.SetPipeline(ctx, data)
	}
	return nil
}

package configEvent

import (
	"github.com/foxiswho/blog-go/app/manage/domainBasic/model/modBasicConfigEvent"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/pangu-2/go-tools/tools/strPg"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
)

type Detail struct {
	Sp  *Sp          `autowire:"?"`
	log *log2.Logger `autowire:"?"`
}

func NewDetail(sp *Sp) *Detail {
	return &Detail{
		Sp:  sp,
		log: sp.log,
	}
}

func (c *Detail) Process(ctx *gin.Context, id string) (rt rg.Rs[modBasicConfigEvent.CreateUpdateCt]) {
	var vo modBasicConfigEvent.CreateUpdateCt
	vo.BodyDelIds = make([]string, 0)
	if strPg.IsBlank(id) {
		return rt.ErrorMessage("模型ID不能为空")
	}
	info, result := c.Sp.repEvent.FindByIdString(ctx, id)
	if !result {
		return rt.ErrorMessage("模型不存在")
	}
	err := copier.Copy(&vo.Header, info)
	if err != nil {
		c.log.Infof("copier.Copy error: %+v", err)
	}
	//
	vo.Body = make([]modBasicConfigEvent.ItemVo, 0)
	//
	data, r := c.Sp.repEventField.FindAllByModelNo(ctx, info.No)
	if r {
		for _, item := range data {
			var obj modBasicConfigEvent.ItemVo
			copier.Copy(&obj, item)
			//
			vo.Body = append(vo.Body, obj)
		}
	}
	rt.Data = vo
	return rt.Ok()
}

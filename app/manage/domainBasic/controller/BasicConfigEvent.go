package controller

import (
	"github.com/foxiswho/blog-go/app/manage/domainBasic/model/modBasicConfigEvent"
	"github.com/foxiswho/blog-go/app/manage/domainBasic/service"
	"github.com/foxiswho/blog-go/middleware/authPg"
	"github.com/foxiswho/blog-go/middleware/validatorPg"
	"github.com/gin-gonic/gin"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
)

type BasicConfigEventController struct {
	Sp *authPg.GroupManageMiddlewareSp  `autowire:""`
	sv *service.BasicConfigEventService `autowire:"?"`
}

func (c *BasicConfigEventController) Create(ctx *gin.Context) {
	var ct modBasicConfigEvent.CreateCt
	if err := ctx.ShouldBind(&ct); err != nil {
		translate := validatorPg.Translate(err, &ct)
		if len(translate) > 0 {
			ctx.JSON(200, rg.ErrorMessageData[string](translate))
			return
		}
		ctx.JSON(200, rg.ErrorDefault[string]())
		return
	}
	ctx.JSON(200, c.sv.Create(ctx, ct))
}

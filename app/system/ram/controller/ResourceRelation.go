package controller

import (
	"github.com/foxiswho/blog-go/app/system/ram/model/modRamResourceRelation"
	"github.com/foxiswho/blog-go/app/system/ram/service"
	"github.com/foxiswho/blog-go/middleware/validatorPg"
	"github.com/foxiswho/blog-go/pkg/common/controllerPg"
	"github.com/foxiswho/blog-go/pkg/enum/state/enumStatePg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/gin-gonic/gin"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
)

func init() {

}

// ResourceRelationController 资源关系
// @Description:
type ResourceRelationController struct {
	controllerPg.SpSystemAuth
	sv  *service.RamResourceRelationService `autowire:"?"`
	log *log2.Logger                        `autowire:"?"`
}

func (c *ResourceRelationController) Query(ctx *gin.Context) {
	var ct modRamResourceRelation.QueryCt
	if err := ctx.ShouldBind(&ct); err != nil {
		//对 返回 错误进行转义 成中文
		translate := validatorPg.Translate(err, &ct)
		if len(translate) > 0 {
			ctx.JSON(200, rg.ErrorMessageData[string](translate))
			return
		}
		ctx.JSON(200, rg.ErrorDefault[string]())
		return
	}
	ctx.JSON(200, c.sv.Query(ctx, ct))
}

func (c *ResourceRelationController) SelectNodePublic(ctx *gin.Context) {
	ct := modRamResourceRelation.QueryCt{State: enumStatePg.ENABLE.IndexPg()}
	ctx.JSON(200, c.sv.SelectNodePublic(ctx, ct))
}

func (c *ResourceRelationController) SelectNodeAllPublic(ctx *gin.Context) {
	ct := modRamResourceRelation.QueryCt{State: enumStatePg.ENABLE.IndexPg()}
	ctx.JSON(200, c.sv.SelectNodeAllPublic(ctx, ct))
}

func (c *ResourceRelationController) SelectPublic(ctx *gin.Context) {
	ct := modRamResourceRelation.QueryCt{State: enumStatePg.ENABLE.IndexPg()}
	ctx.JSON(200, c.sv.SelectPublic(ctx, ct))
}

// Selected 已选中的权限
func (c *ResourceRelationController) Selected(ctx *gin.Context) {
	var ct modRamResourceRelation.QuerySelectedCt
	if err := ctx.ShouldBind(&ct); err != nil {
		//对 返回 错误进行转义 成中文
		translate := validatorPg.Translate(err, &ct)
		if len(translate) > 0 {
			ctx.JSON(200, rg.ErrorMessageData[string](translate))
			return
		}
		ctx.JSON(200, rg.ErrorDefault[string]())
		return
	}
	ctx.JSON(200, c.sv.Selected(ctx, ct.Code))
}

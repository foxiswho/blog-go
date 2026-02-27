package controller

import (
	"github.com/foxiswho/blog-go/app/system/ram/model/modRamResourceGroupRelation"
	"github.com/foxiswho/blog-go/app/system/ram/service"
	"github.com/foxiswho/blog-go/middleware/validatorPg"
	"github.com/foxiswho/blog-go/pkg/common/controllerPg"
	"github.com/foxiswho/blog-go/pkg/consts/constsRam/resourceTypeCategoryPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/gin-gonic/gin"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
)

func init() {

}

// ResourceGroupRelationController 资源关系
// @Description:
type ResourceGroupRelationController struct {
	controllerPg.SpSystemAuth
	sv  *service.RamResourceGroupRelationService `autowire:"?"`
	log *log2.Logger                             `autowire:"?"`
}

// Selected 已选中的权限
func (c *ResourceGroupRelationController) Selected(ctx *gin.Context) {
	var ct modRamResourceGroupRelation.QueryByTypeValueCt
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
	ctx.JSON(200, c.sv.Selected(ctx, ct))
}

// SelectedByRole
//
//	@Description: 角色
//	@receiver c
//	@param ctx
func (c *ResourceGroupRelationController) SelectedByRole(ctx *gin.Context) {
	var ct modRamResourceGroupRelation.QueryByTypeValueCt
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
	ct.TypeCategory = resourceTypeCategoryPg.Role.Index()
	ctx.JSON(200, c.sv.Selected(ctx, ct))
}

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hongmengzhu/xianfu-blog-go/app/system/ram/model/modRamResourceGroupRelation"
	"github.com/hongmengzhu/xianfu-blog-go/app/system/ram/service"
	"github.com/hongmengzhu/xianfu-blog-go/middleware/authPg"
	"github.com/hongmengzhu/xianfu-blog-go/middleware/serverPg/ginServer"
	"github.com/hongmengzhu/xianfu-blog-go/middleware/validatorPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/common/controllerPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/consts/constsRam/resourceTypeCategoryPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/log2"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/routerPg"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
	"go-spring.org/spring/gs"
)

func init() {
	gs.Provide(new(ResourceGroupRelationController)).Name("SystemResourceGroupRelationController").Export(gs.As[routerPg.RouteRegistrar]())
}

type ResourceGroupRelationController struct {
	routerPg.RouteRegistrar
	controllerPg.SpSystemAuth
	sv  *service.RamResourceGroupRelationService `autowire:"?"`
	log *log2.Logger                             `autowire:"?"`
}

func (c *ResourceGroupRelationController) RegisterRoutes(e *gin.Engine) {
	r := ginServer.GinServerDefault
	group := r.Group("/xianfu/sys/ram/resource-group-relation", authPg.GroupSystemMiddleware(c.Sp))
	group.POST("/selectedByRole", c.SelectedByRole)
}

func (c *ResourceGroupRelationController) Selected(ctx *gin.Context) {
	var ct modRamResourceGroupRelation.QueryByTypeValueCt
	if !routerPg.BindJson(ctx, &ct) {
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

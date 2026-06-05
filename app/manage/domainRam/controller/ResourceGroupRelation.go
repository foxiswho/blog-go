package controller

import (
	"github.com/foxiswho/blog-go/app/manage/domainRam/service"
	"github.com/foxiswho/blog-go/middleware/authPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/foxiswho/blog-go/pkg/routerPg"
	"github.com/gin-gonic/gin"
	"github.com/go-spring/spring-core/gs"
)

func init() {
	gs.Provide(new(ResourceGroupRelationController)).Name("ManageResourceGroupRelationController").Export(gs.As[routerPg.RouteRegistrar]())
}

// ResourceGroupRelationController 资源组关联
// @Description:
type ResourceGroupRelationController struct {
	routerPg.RouteRegistrar
	Sp  *authPg.GroupManageMiddlewareSp          `autowire:""`
	sv  *service.RamResourceGroupRelationService `autowire:"?"`
	log *log2.Logger                             `autowire:"?"`
}

// RegisterRoutes 注册路由
//
//	@Description:
//	@receiver c
//	@param e
func (c *ResourceGroupRelationController) RegisterRoutes(e *gin.Engine) {
	//group := e.Group("/pg2lq/manage/ram/resource-group-relation", authPg.GroupManageMiddleware(c.Sp))
	//group.POST("/selectedByRole", c.Query)
}

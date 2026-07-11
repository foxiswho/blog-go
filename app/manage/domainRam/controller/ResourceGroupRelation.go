package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hongmengzhu/xianfu-blog-go/app/manage/domainRam/service"
	"github.com/hongmengzhu/xianfu-blog-go/middleware/authPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/log2"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/routerPg"
	"go-spring.org/spring/gs"
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

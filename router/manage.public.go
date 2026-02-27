package router

import (
	"github.com/foxiswho/blog-go/app/manage/domainRam/controller"
	"github.com/foxiswho/blog-go/middleware/authPg"
	"github.com/foxiswho/blog-go/middleware/serverPg/ginServer"
	"github.com/go-spring/spring-core/gs"
)

func init() {
	// 登录后
	gs.Root(gs.Object(&controller.PublicController{}).Init(func(s *controller.PublicController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/manage/public", authPg.GroupManageMiddleware(s.Sp))
		group.GET("/info", s.Public)
		group.POST("/password", s.UpdatePassword)
	}))
}

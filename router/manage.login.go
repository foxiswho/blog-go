package router

import (
	"github.com/foxiswho/blog-go/app/manage/domainRam/controller"
	"github.com/foxiswho/blog-go/middleware/serverPg/ginServer"
	"github.com/go-spring/spring-core/gs"
)

func init() {
	//登陆
	gs.Root(gs.Object(&controller.LoginController{}).Init(func(s *controller.LoginController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/auth/manage")
		group.POST("/login", s.Login)
		group.POST("/refresh", s.RefreshToken)
	}))
	//退出
	gs.Root(gs.Object(&controller.LogoutController{}).Init(func(s *controller.LogoutController) {
		r := ginServer.GinServerDefault
		group := r.Group("/pg2lq/auth/manage")
		group.Any("/logout", s.Logout)
	}))
}

package ginServer

import (
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/routerPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/templatePg"
	"github.com/pangu-2/go-tools/tools/datetimePg"
)

// NewGinEngine 创建 gin.Engine 实例，完成静态文件、模板、路由注册后返回。
//
// 由 go-spring DI 容器装配：registrars 通过集合注入接收所有实现了
// routerPg.RouteRegistrar 接口的 Bean。返回的 *gin.Engine Bean 会被
// 官方 starter-gin 自动发现（OnBean[*gin.Engine]），进而创建
// SimpleGinServer 并导出为 gs.Server，接管 HTTP 生命周期。
func NewGinEngine(registrars []routerPg.RouteRegistrar) *gin.Engine {
	engine := gin.New()

	// 静态文件目录
	engine.Static("/assets", "./data/assets")
	engine.Static("/attachment", "./data/attachment")

	// 模板函数，需在 LoadHTMLGlob 之前设置
	engine.SetFuncMap(template.FuncMap{
		"unescaped":  templatePg.Unescaped,
		"dateformat": datetimePg.Format,
	})

	// 加载模板文件，需在路由注册之前
	engine.LoadHTMLGlob("data/templates/**/**/*")

	// 注册所有路由
	for _, r := range registrars {
		r.RegisterRoutes(engine)
	}

	return engine
}

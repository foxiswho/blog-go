package ginServer

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/foxiswho/blog-go/pkg/routerPg"
	"github.com/gin-gonic/gin"
	"github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
)

// GinServerDefault 初始化默认服务
var GinServerDefault = gin.New()

func init() {
}

// GetInstance
//
//	@Description: 获取 gin 实例
//	@return *gin.Engine
func GetInstance() *gin.Engine {
	return GinServerDefault
}

// gin 框架 整合
type GinServer struct {
	svr       *http.Server
	svrEngine *gin.Engine
	Port      string
}

// NewGinServer
//
//	@Description: 创建 GinServer 实例，并注册所有路由
//	@param port: 服务监听端口
//	@param registrars: 路由注册器集合，由 DI 容器自动注入
//	@return *GinServer
func NewGinServer(port string, registrars []routerPg.RouteRegistrar) *GinServer {
	//log.Infof(context.Background(), log.TagAppDef, "NewGinServer.port:%+v ", port)
	engine := GetInstance()
	for _, r := range registrars {
		r.RegisterRoutes(engine)
	}
	svr := &GinServer{}
	svr.Port = port
	svr.svrEngine = engine
	svr.svr = &http.Server{
		Addr:    ":" + port,
		Handler: svr.svrEngine,
	}
	return svr
}

// 启动 端口
func (s *GinServer) Run(ctx context.Context, sig gs.ReadySignal) error {
	addr := s.svr.Addr
	if addr == "" {
		addr = ":8080"
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	<-sig.TriggerAndWait() // 等待启动信号
	//
	log.Infof(context.Background(), log.TagAppDef, "starting successfully...")
	fmt.Println()
	fmt.Printf("host: %+v\n", "localhost")
	fmt.Printf("port: %+v\n", s.Port)
	fmt.Printf("url: http://localhost:%+v\n", s.Port)
	fmt.Println()
	return s.svr.Serve(ln)
}

// 关闭
func (s *GinServer) Stop() error {
	return s.svr.Shutdown(context.Background())
}

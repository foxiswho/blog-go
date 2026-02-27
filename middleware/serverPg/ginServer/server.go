package ginServer

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	syslog "github.com/go-spring/log"
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

func NewGinServer(port string) *GinServer {
	//syslog.Infof(context.Background(), syslog.TagAppDef, "NewGinServer.port:%+v ", port)
	svr := &GinServer{}
	svr.Port = port
	svr.svrEngine = GinServerDefault
	svr.svr = &http.Server{
		Addr:    ":" + port,
		Handler: svr.svrEngine,
	}
	return svr
}

// 启动 端口
func (s *GinServer) ListenAndServe(sig gs.ReadySignal) error {
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
	syslog.Infof(context.Background(), syslog.TagAppDef, "starting successfully...")
	fmt.Println()
	fmt.Printf("host: %+v\n", "localhost")
	fmt.Printf("port: %+v\n", s.Port)
	fmt.Printf("url: http://localhost:%+v\n", s.Port)
	fmt.Println()
	return s.svr.Serve(ln)
}

// 关闭
func (s *GinServer) Shutdown(ctx context.Context) error {
	return s.svr.Shutdown(ctx)
}

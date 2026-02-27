package httpServer

import (
	"context"
	"net"
	"net/http"

	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
)

func init() {
	gs.Object(NewServer()).AsServer()
}

type MyServer struct {
	svr *http.Server
}

func NewServer() *MyServer {
	return &MyServer{
		svr: &http.Server{Addr: ":18081"},
	}
}

func (s *MyServer) ListenAndServe(sig gs.ReadySignal) error {
	ln, err := net.Listen("tcp", s.svr.Addr)
	if err != nil {
		return err
	}
	<-sig.TriggerAndWait() // 等待启动信号
	syslog.Infof(context.Background(), syslog.TagBizDef, "starting successfully... Port: %+v", s.svr.Addr)
	return s.svr.Serve(ln)
}

func (s *MyServer) Shutdown(ctx context.Context) error {
	return s.svr.Shutdown(ctx)
}

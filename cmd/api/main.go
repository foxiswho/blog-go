package main

import (
	_ "github.com/foxiswho/blog-go/middleware/authPg"
	_ "github.com/foxiswho/blog-go/middleware/dbPg/postgresqlPg"
	_ "github.com/foxiswho/blog-go/middleware/runnerPg"
	_ "github.com/foxiswho/blog-go/middleware/serverPg/ginServer"
	_ "github.com/foxiswho/blog-go/router"
	"github.com/go-spring/spring-core/gs"
	"os"
)

func init() {
	//指定环境
	//gs.SetActiveProfiles("dev")
	//关闭 案例 serverPg
	gs.EnableSimplePProfServer(false)
	// 指定配置文件目录, 如果不设置，默认 conf 目录
	_ = os.Setenv("GS_SPRING_APP_CONFIG-LOCAL_DIR", "./config")
}
func main() {
	// 启动 Go-Spring 应用（自动启动 Gin 服务）
	gs.Run()
}

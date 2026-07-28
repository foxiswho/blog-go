//go:generate go run ./auto_tool/main.go -dir=./app,./infrastructure -output=auto_import.go -concurrency=8 -ignore=**/vendor/**,**/.git/**,**/*_test.go,**/testdata/**,**/node_modules/**
package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"runtime"

	"github.com/farseer-go/eventBus"
	fsE "github.com/farseer-go/fs"
	"github.com/gin-gonic/gin"
	"github.com/hongmengzhu/xianfu-blog-go/cmd"
	_ "github.com/hongmengzhu/xianfu-blog-go/middleware"
	"github.com/hongmengzhu/xianfu-blog-go/middleware/serverPg/ginServer"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/log2"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/logsPg"
	"github.com/pangu-2/go-tools/tools/ioPg"
	"go-spring.org/log"
	"go-spring.org/spring/gs"
	_ "go-spring.org/starter-gin"
)

func init() {
	//
	gin.SetMode(gin.DebugMode)
	//
	dirPath := make([]string, 0)
	// 日志目录路径
	dirPath = append(dirPath, "data/logs")
	//配置文件目录
	dirPath = append(dirPath, "data/config")
	//附件
	dirPath = append(dirPath, "data/attachment")
	if nil != dirPath {
		for _, path := range dirPath {
			// 创建目录
			err := ioPg.DirectoryCreate(path)
			if err != nil {
				fmt.Printf("创建目录失败: [%v] => %v", path, err)
			}
		}
	}
	//
	gs.Provide(log2.New(log2.LevelDebug, false))
	//指定环境
	//gs.SetActiveProfiles("dev")
}

type Controller struct{}

func (c *Controller) Echo(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello, gin!")
}

func main() {

	// 构建信息，golang版本 commit id 时间
	var version bool
	flag.BoolVar(&version, "v", false, "version")
	flag.Parse()
	if version {
		fmt.Printf("go version: %s\nBuild version: %s\nBuild commit: %s\nBuild time: %s\n",
			runtime.Version(), cmd.BuildVersion, cmd.BuildGitCommit, cmd.BuildTime)
		return
	}

	// 指定配置文件目录, 如果不设置，默认 conf 目录
	_ = os.Setenv("GS_SPRING_APP_CONFIG_DIR", "./data/config")

	log.Debugf(context.Background(), logsPg.TagAppDef, "111111111111111111111111111111111111111")
	// 提供 *gin.Engine Bean，官方 starter-gin 自动发现并创建 SimpleGinServer (gs.Server)
	gs.Provide(ginServer.NewGinEngine, gs.TagArg("?"))
	//事件监听
	fsE.Initialize[eventBus.Module]("panGu")
	gs.Configure(func(app gs.App) {
		app.Property("spring.http.server.enabled", "false")
	}).Run()
}

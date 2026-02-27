package main

import (
	"context"
	"fmt"
	"html/template"
	"os"

	"github.com/farseer-go/eventBus"
	fsE "github.com/farseer-go/fs"
	_ "github.com/foxiswho/blog-go/app"
	_ "github.com/foxiswho/blog-go/middleware"
	"github.com/foxiswho/blog-go/middleware/serverPg/ginServer"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/foxiswho/blog-go/pkg/logsPg"
	"github.com/foxiswho/blog-go/pkg/templatePg"
	"github.com/foxiswho/blog-go/pkg/tools/pathPg"
	_ "github.com/foxiswho/blog-go/router"
	"github.com/gin-gonic/gin"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
	"github.com/pangu-2/go-tools/tools/datetimePg"
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
			err := pathPg.DirectoryCreate(path)
			if err != nil {
				fmt.Printf("创建目录失败: [%v] => %v", path, err)
			}
		}
	}
	//
	gs.Object(log2.New(log2.LevelDebug, false))
	//指定环境
	//gs.SetActiveProfiles("dev")
	//关闭 案例 serverPg
	gs.EnableSimplePProfServer(false)
	// 指定配置文件目录, 如果不设置，默认 conf 目录
	_ = os.Setenv("GS_SPRING_APP_CONFIG-LOCAL_DIR", "./data/config")
	// gin 静态文件路径
	ginServer.GetInstance().Static("/assets", "./assets")
	ginServer.GetInstance().Static("/attachment", "./data/attachment")
	funcMap := template.FuncMap{
		"unescaped":  templatePg.Unescaped,
		"dateformat": datetimePg.Format,
	}
	ginServer.GetInstance().SetFuncMap(funcMap)
	//加载templates中所有模板文件, 使用不同目录下名称相同的模板,注意:一定要放在配置路由之前才得行
	ginServer.GetInstance().LoadHTMLGlob("data/templates/**/**/*")
	//html := template.Must(template.ParseFiles("file1", "file2"))
	//ginServer.GetInstance().SetHTMLTemplate(html)
}
func main() {
	syslog.Debugf(context.Background(), logsPg.TagAppDef, "111111111111111111111111111111111111111")
	//服务，传入配置 端口
	gs.Provide(ginServer.NewGinServer, gs.TagArg("${server.port}")).AsServer()
	//事件监听
	fsE.Initialize[eventBus.Module]("panGu")
	gs.Run()
}

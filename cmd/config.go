package cmd

import (
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pangu-2/go-tools/tools/wrapperPg/rg"
)

var (
	BuildVersion   string // 代码版本号 例如："v1.2.0"
	BuildGitCommit string // Git 提交哈希
	BuildTime      string // 编译时间
	BuildUser      string // 构建用户/环境标识
)

// GetVersion 获取版本信息
func GetVersion(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Writer.WriteHeader(http.StatusOK)
	mapStr := make(map[string]interface{})
	mapStr["BuildVersion"] = BuildVersion
	mapStr["BuildGitCommit"] = BuildGitCommit
	mapStr["BuildTime"] = BuildTime
	mapStr["BuildUser"] = BuildUser
	mapStr["GoVersion"] = runtime.Version()
	mapStr["OSArch"] = runtime.GOOS + "/" + runtime.GOARCH
	// 是否编译
	if !IsRelease() {
		mapStr["BuildVersion"] = "未编译"
		mapStr["BuildGitCommit"] = "未编译"
		mapStr["BuildTime"] = "未编译"
	}
	ctx.JSON(200, rg.OkData[map[string]interface{}](mapStr))
}

// IsRelease 当前执行的是否是已编译文件
func IsRelease() bool {
	arg1 := strings.ToLower(os.Args[0])
	name := filepath.Base(arg1)

	return strings.Index(name, "__") != 0 && strings.Index(arg1, "go-build") < 0
}

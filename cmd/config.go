package cmd

import (
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
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
	ctx.JSON(200, mapStr)
}

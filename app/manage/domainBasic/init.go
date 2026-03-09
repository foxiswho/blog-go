package domainBasic

import (
	"github.com/foxiswho/blog-go/app/manage/domainBasic/service/configModel"
	"github.com/go-spring/spring-core/gs"
)

// app 目录下 各个模块包，需要 统一 初始化
func init() {
	gs.Provide(new(configModel.Sp))
}

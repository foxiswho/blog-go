package domainBasic

import (
	"github.com/foxiswho/blog-go/app/manage/domainBasic/service/configEvent"
	"github.com/foxiswho/blog-go/app/manage/domainBasic/service/configModel"
	"github.com/foxiswho/blog-go/app/manage/domainBasic/service/modelRules"
	"github.com/go-spring/spring-core/gs"
)

// app 目录下 各个模块包，需要 统一 初始化
func init() {
	gs.Provide(new(configEvent.Sp))
	gs.Provide(new(configModel.Sp))
	gs.Provide(new(modelRules.Sp))
}

package basic

import (
	"github.com/foxiswho/blog-go/app/event/basic/service/eventBasicEvent"
	"github.com/foxiswho/blog-go/app/event/basic/service/eventBasicRules"
	"github.com/foxiswho/blog-go/app/event/basic/service/tagsBasicEvent"
	"github.com/go-spring/spring-core/gs"
)

// app 目录下 各个模块包，需要 统一 初始化
func init() {
	gs.Provide(new(tagsBasicEvent.Sp))
	gs.Provide(new(eventBasicEvent.Sp))
	gs.Provide(new(eventBasicRules.Sp))
}

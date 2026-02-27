package basic

import (
	"context"
	"github.com/foxiswho/blog-go/app/event/basic/service/tagsBasicEvent"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
	"reflect"
)

// app 目录下 各个模块包，需要 统一 初始化
func init() {
	gs.Provide(new(tagsBasicEvent.Sp)).Init(func(s *tagsBasicEvent.Sp) {
		syslog.Infof(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

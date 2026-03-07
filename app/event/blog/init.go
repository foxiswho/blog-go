package blog

import (
	"context"
	"reflect"

	"github.com/foxiswho/blog-go/app/event/blog/service/articleBlogEvent"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
)

// app 目录下 各个模块包，需要 统一 初始化
func init() {
	gs.Provide(new(articleBlogEvent.Sp)).Init(func(s *articleBlogEvent.Sp) {
		syslog.Infof(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

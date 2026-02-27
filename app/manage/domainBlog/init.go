package domainBlog

import (
	"context"
	"reflect"

	"github.com/foxiswho/blog-go/app/manage/domainBlog/service/blogArticle"
	"github.com/foxiswho/blog-go/app/manage/domainBlog/service/blogCollect"
	"github.com/foxiswho/blog-go/app/manage/domainBlog/service/blogTopic"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
)

// app 目录下 各个模块包，需要 统一 初始化
func init() {
	gs.Provide(new(blogArticle.Sp)).Init(func(s *blogArticle.Sp) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
	gs.Provide(new(blogCollect.Sp)).Init(func(s *blogCollect.Sp) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
	gs.Provide(new(blogTopic.Sp)).Init(func(s *blogTopic.Sp) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

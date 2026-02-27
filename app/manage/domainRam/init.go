package domainRam

import (
	"context"
	"reflect"

	"github.com/foxiswho/blog-go/app/manage/domainRam/service/ramAccount"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
)

func init() {
	gs.Provide(new(ramAccount.Sp)).Init(func(s *ramAccount.Sp) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

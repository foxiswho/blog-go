package tc

import (
	"context"
	"reflect"

	"github.com/foxiswho/blog-go/app/system/tc/service/tcAccount"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
)

func init() {
	gs.Provide(new(tcAccount.Sp)).Init(func(s *tcAccount.Sp) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

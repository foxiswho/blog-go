package repositoryApi

import (
	"context"

	"github.com/foxiswho/blog-go/infrastructure/entityApi"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/support"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"

	"reflect"
)

func init() {
	gs.Provide(new(ApiDiplRepository)).Init(func(s *ApiDiplRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[ApiDiplRepository])).Init(func(s *support.BaseService[ApiDiplRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type ApiDiplRepository struct {
	repositoryPg.BaseRepository[entityApi.ApiDiplEntity, int64]
}

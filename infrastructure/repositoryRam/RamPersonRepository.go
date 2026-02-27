package repositoryRam

import (
	"context"

	"github.com/foxiswho/blog-go/infrastructure/entityRam"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/support"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"

	"reflect"
)

func init() {
	gs.Provide(new(RamPersonRepository)).Init(func(s *RamPersonRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[RamPersonRepository])).Init(func(s *support.BaseService[RamPersonRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type RamPersonRepository struct {
	repositoryPg.BaseRepository[entityRam.RamPersonEntity, int64]
}

/*
func (c *RamPersonRepository) FindAllByTenantId2(tenantId int64, ctx ...baseContext.SpContext) (infos []entityRam.RamPersonEntity, err error) {
	var webCtx *gin.Context
	for _, option := range ctx {
		option(&webCtx)
	}

	log.Println("value.tenant=", webCtx.Get("tenant"))
	tx := c.Db().Model(&entityRam.RamPersonEntity{}).Where("tenant_id=?", tenantId).Find(&infos)
	value := tx.Statement.Context.Value("tenant")
	log.Println("value=", value)
	log.Println("value=", value)
	log.Println("value=")
	log.Println("value=", tx.Statement.Context != nil)
	v, err := knife.Load(tx.Statement.Context, "tenant")
	log.Println("err=", err)
	log.Println("value=", v)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return infos, nil
}
*/

package repositoryRam

import (
	"context"

	"github.com/hongmengzhu/xianfu-blog-go/infrastructure/entityRam"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/tools/dbHelper/support"
	"go-spring.org/log"
	"go-spring.org/spring/gs"

	"reflect"
)

func init() {
	gs.Provide(new(RamIdpCredentialRepository)).Init(func(s *RamIdpCredentialRepository) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[RamIdpCredentialRepository])).Init(func(s *support.BaseService[RamIdpCredentialRepository]) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type RamIdpCredentialRepository struct {
	repositoryPg.BaseRepository[entityRam.RamIdpCredentialEntity, int64]
	//
}

// FindBySourceNoAndCredType 按认证源编号+凭证类型查找
func (c *RamIdpCredentialRepository) FindBySourceNoAndCredType(ctx context.Context, sourceNo, credType string) (info *entityRam.RamIdpCredentialEntity, result bool) {
	tx := c.DbModel().WithContext(ctx).Where("source_no=? AND cred_type=?", sourceNo, credType).First(&info)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return nil, false
	}
	return info, true
}

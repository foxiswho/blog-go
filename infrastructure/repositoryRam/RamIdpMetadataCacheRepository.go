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
	gs.Provide(new(RamIdpMetadataCacheRepository)).Init(func(s *RamIdpMetadataCacheRepository) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[RamIdpMetadataCacheRepository])).Init(func(s *support.BaseService[RamIdpMetadataCacheRepository]) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type RamIdpMetadataCacheRepository struct {
	repositoryPg.BaseRepository[entityRam.RamIdpMetadataCacheEntity, int64]
	//
}

// FindBySourceNo 按认证源编号查找元数据缓存
func (c *RamIdpMetadataCacheRepository) FindBySourceNo(ctx context.Context, sourceNo string) (info *entityRam.RamIdpMetadataCacheEntity, result bool) {
	tx := c.DbModel().WithContext(ctx).Where("source_no=?", sourceNo).First(&info)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return nil, false
	}
	return info, true
}

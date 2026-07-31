package repositoryRam

import (
	"context"
	"reflect"

	"github.com/hongmengzhu/xianfu-blog-go/infrastructure/entityRam"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/tools/dbHelper/repositoryPg/optionsPg"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/tools/dbHelper/support"
	"go-spring.org/log"
	"go-spring.org/spring/gs"
)

func init() {
	gs.Provide(new(RamAccountAuthWebauthnRepository)).Init(func(s *RamAccountAuthWebauthnRepository) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[RamAccountAuthWebauthnRepository])).Init(func(s *support.BaseService[RamAccountAuthWebauthnRepository]) {
		log.Debugf(context.Background(), log.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type RamAccountAuthWebauthnRepository struct {
	repositoryPg.BaseRepository[entityRam.RamAccountAuthWebauthnEntity, int64]
}

// FindByAno 按账号编号查找所有 WebAuthn 凭证
func (c *RamAccountAuthWebauthnRepository) FindByAno(ctx context.Context, ano string, opts ...optionsPg.Option) (infos []*entityRam.RamAccountAuthWebauthnEntity, found bool) {
	tx := c.SetOptionScopes(c.DbModel().WithContext(ctx), opts...).Where("ano = ? AND enabled = 1", ano).Find(&infos)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if tx.RowsAffected == 0 {
		return nil, false
	}
	return infos, true
}

// FindByAnoAndCredentialID 按账号编号和凭证ID查找
func (c *RamAccountAuthWebauthnRepository) FindByAnoAndCredentialID(ctx context.Context, ano, credentialID string, opts ...optionsPg.Option) (info *entityRam.RamAccountAuthWebauthnEntity, found bool) {
	var result entityRam.RamAccountAuthWebauthnEntity
	tx := c.SetOptionScopes(c.DbModel().WithContext(ctx), opts...).Where("ano = ? AND cred_id = ?", ano, credentialID).First(&result)
	if tx.Error != nil {
		return nil, false
	}
	return &result, true
}

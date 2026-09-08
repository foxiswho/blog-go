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
	gs.Provide(new(RamAccountDenyListRepository))

	gs.Provide(new(support.BaseService[RamAccountDenyListRepository]))
}

type RamAccountDenyListRepository struct {
	repositoryPg.BaseRepository[entityRam.RamAccountDenyListEntity, int64]
}

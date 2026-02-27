package repositoryRam

import (
	"context"

	"github.com/foxiswho/blog-go/infrastructure/entityRam"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/support"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
	"github.com/pangu-2/go-tools/tools/numberPg"

	"reflect"
)

func init() {
	gs.Provide(new(RamDepartmentRepository)).Init(func(s *RamDepartmentRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[RamDepartmentRepository])).Init(func(s *support.BaseService[RamDepartmentRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type RamDepartmentRepository struct {
	repositoryPg.BaseRepository[entityRam.RamDepartmentEntity, int64]
}

func (c *RamDepartmentRepository) FindByParentId(code int64) (info *entityRam.RamDepartmentEntity, result bool, err error) {
	tx := c.Db().Where("parent_id=?", code).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false, tx.Error
	}
	if 0 == tx.RowsAffected {
		return nil, false, nil
	}
	return info, true, nil
}

func (c *RamDepartmentRepository) FindAllByIdLink(code string) (info []*entityRam.RamDepartmentEntity, result bool) {
	tx := c.Db().Where("id_link like ?", "%|"+code+"|%").Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *RamDepartmentRepository) FindAllByIdLinkArr(code []string) (info []*entityRam.RamDepartmentEntity, result bool) {
	db := c.Db()
	for index, val := range code {
		if 0 == index {
			db.Where("id_link like ?", "%|"+val+"|%")
		} else {
			db.Or("id_link like ?", "%|"+val+"|%")
		}
	}
	tx := db.Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *RamDepartmentRepository) FindAllByNoLinkArr(code []string) (info []*entityRam.RamDepartmentEntity, result bool) {
	db := c.Db()
	for index, val := range code {
		if 0 == index {
			db.Where("no_link like ?", "%|"+val+"|%")
		} else {
			db.Or("no_link like ?", "%|"+val+"|%")
		}
	}
	tx := db.Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *RamDepartmentRepository) FindAllByIdLinkReturnString(code []string) (ids []string, result bool) {
	var info []*entityRam.RamDepartmentEntity
	db := c.Db()
	for index, val := range code {
		if 0 == index {
			db.Where("id_link like ?", "%|"+val+"|%")
		} else {
			db.Or("id_link like ?", "%|"+val+"|%")
		}
	}
	tx := db.Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	ids = make([]string, 0)
	for _, entity := range info {
		ids = append(ids, numberPg.Int64ToString(entity.ID))
	}
	return ids, true
}

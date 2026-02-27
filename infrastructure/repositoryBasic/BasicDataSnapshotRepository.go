package repositoryBasic

import (
	"context"
	"reflect"

	"github.com/foxiswho/blog-go/infrastructure/entityBasic"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/support"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
	"github.com/pangu-2/go-tools/tools/jsonPg"
)

func init() {
	gs.Provide(new(BasicDataSnapshotRepository)).Init(func(s *BasicDataSnapshotRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[BasicDataSnapshotRepository])).Init(func(s *support.BaseService[BasicDataSnapshotRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type BasicDataSnapshotRepository struct {
	repositoryPg.BaseRepository[entityBasic.BasicDataSnapshotEntity, int64]
}

func (c *BasicDataSnapshotRepository) FindByNameAndIdNot(name string, id int64) (info *entityBasic.BasicDataSnapshotEntity, result bool) {
	tx := c.Db().Where("name=?", name).Where("id <> ?", id).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *BasicDataSnapshotRepository) SnapshotVersion(obj interface{}, module, tenantNo, value, version, extend string) {
	toJson, err := jsonPg.ObjToJson(obj)
	if nil != err {
		c.Log().Errorf("snapshot version error: %+v", err)
	} else {
		mark := value + "|" + value
		c.Create(&entityBasic.BasicDataSnapshotEntity{
			Module:   module,
			Data:     toJson,
			Mark:     mark,
			Name:     "",
			TenantNo: tenantNo,
			Value:    value,
			Version:  version,
			Extend:   extend,
		})
	}
}

func (c *BasicDataSnapshotRepository) SnapshotVersionAll(obj interface{}, module, tenantNo, value, version, name, extend string) {
	toJson, err := jsonPg.ObjToJson(obj)
	if nil != err {
		c.Log().Errorf("snapshot version error: %+v", err)
	} else {
		mark := value + "|" + value
		c.Create(&entityBasic.BasicDataSnapshotEntity{
			Module:   module,
			Data:     toJson,
			Mark:     mark,
			Name:     name,
			TenantNo: tenantNo,
			Value:    value,
			Version:  version,
			Extend:   extend,
		})
	}
}

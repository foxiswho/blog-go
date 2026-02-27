package repositoryBasic

import (
	"context"
	"reflect"

	"github.com/foxiswho/blog-go/infrastructure/entityBasic"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/support"
	syslog "github.com/go-spring/log"
	"github.com/go-spring/spring-core/gs"
	"github.com/pangu-2/go-tools/tools/dbPg/pagePg"
)

func init() {
	gs.Provide(new(BasicDataDictionaryRepository)).Init(func(s *BasicDataDictionaryRepository) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})

	gs.Provide(new(support.BaseService[BasicDataDictionaryRepository])).Init(func(s *support.BaseService[BasicDataDictionaryRepository]) {
		syslog.Debugf(context.Background(), syslog.TagAppDef, "%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

type BasicDataDictionaryRepository struct {
	repositoryPg.BaseRepository[entityBasic.BasicDataDictionaryEntity, int64]
}

// 分页
func (b *BasicDataDictionaryRepository) FindAllPageOwnerId0(t entityBasic.BasicDataDictionaryEntity, option pagePg.Option[entityBasic.BasicDataDictionaryEntity]) (pagePg.PaginatorPg[entityBasic.BasicDataDictionaryEntity], error) {
	var total int64
	pg := pagePg.NewPaginatorPg[entityBasic.BasicDataDictionaryEntity](option)
	countTx := b.Db().Model(b.Entity).Where(t).Where("owner_no='0'").Count(&total)
	if nil != countTx.Error {
		return pg, countTx.Error
	}
	var infos []entityBasic.BasicDataDictionaryEntity
	tx := b.Db().Model(b.Entity).Where(t).Where("owner_no='0'").Find(&infos)
	if tx.Error != nil {
		return pg, tx.Error
	}
	pg.Data = infos
	pg.Total = total
	pg.Pageable = pagePg.NewPageablePg(total, pg.PageNum, pg.PageSize)
	//if total >0 {
	//	t2 := infos[len(infos)-1]
	//	pg.OffsetId=t2
	//}
	return pg, nil
}

// 查询所有
func (b *BasicDataDictionaryRepository) FindAllOwnerId0(t entityBasic.BasicDataDictionaryEntity) (infos []*entityBasic.BasicDataDictionaryEntity, result bool) {
	tx := b.Db().Where(t).Where("owner_no='0'").Find(&infos)
	if tx.Error != nil {
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return infos, true
}

// FindAllOwners 查询所有
func (b *BasicDataDictionaryRepository) FindAllOwners(ids []string) (infos []*entityBasic.BasicDataDictionaryEntity, result bool) {
	tx := b.Db().Where("owner_no in ", ids).Find(&infos)
	if tx.Error != nil {
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return infos, true
}

// FindAllOwnersMap 查询所有
func (b *BasicDataDictionaryRepository) FindAllOwnersMap(ids []string, t entityBasic.BasicDataDictionaryEntity) (maps map[string][]*entityBasic.BasicDataDictionaryEntity, result bool) {
	if nil == maps {
		maps = make(map[string][]*entityBasic.BasicDataDictionaryEntity)
	}
	var infos []*entityBasic.BasicDataDictionaryEntity
	tx := b.Db().Where(t).Where("owner_no in ", ids).Find(&infos)
	if tx.Error != nil {
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	for _, info := range infos {
		if _, ok := maps[info.OwnerNo]; !ok {
			maps[info.OwnerNo] = make([]*entityBasic.BasicDataDictionaryEntity, 0)
		}
		maps[info.OwnerNo] = append(maps[info.OwnerNo], info)
	}
	return maps, true
}

// FindAllByCodeIn 查询所有
func (b *BasicDataDictionaryRepository) FindAllByCodeIn(ids []string) (infos []*entityBasic.BasicDataDictionaryEntity, result bool) {
	tx := b.Db().Where("code in ", ids).Find(&infos)
	if tx.Error != nil {
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return infos, true
}

func (c *BasicDataDictionaryRepository) FindByNameAndIdNot(name string, id int64) (info *entityBasic.BasicDataDictionaryEntity, result bool) {
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

func (c *BasicDataDictionaryRepository) FindByValueAndIdNotAndOwnerNo(name, id, ownerId string) (info *entityBasic.BasicDataDictionaryEntity, result bool) {
	tx := c.Db().Where("owner_no=?", ownerId).Where("value=?", name).Where("id <> ?", id).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}
func (c *BasicDataDictionaryRepository) FindByCodeAndIdNotAndOwnerNo(code, id, ownerId string) (info *entityBasic.BasicDataDictionaryEntity, result bool) {
	tx := c.Db().Where("code=?", code).Where("owner_no=?", ownerId).Where("id <> ?", id).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

// FindAllByTypeUniqueMd5AndOwnerNo 根据类型编号查询
func (c *BasicDataDictionaryRepository) FindAllByTypeUniqueMd5AndOwnerNo(code, id, ownerId string) (info []*entityBasic.BasicDataDictionaryEntity, result bool) {
	tx := c.Db().Where("owner_no=?", ownerId).Where("type_unique_md5=?", code).Where("id <> ?", id).Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

func (c *BasicDataDictionaryRepository) FindByCodeAndTypeCodeAndIdNotAndOwnerNo(code, typeCode, id, ownerId string) (info *entityBasic.BasicDataDictionaryEntity, result bool) {
	tx := c.Db().Where("code=?", code).Where("type_code=?", typeCode).Where("owner_no=?", ownerId).Where("id <> ?", id).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

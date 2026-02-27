package repositoryPg

import (
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/foxiswho/blog-go/pkg/configPg"
	"github.com/foxiswho/blog-go/pkg/consts/constContextPg"
	"github.com/foxiswho/blog-go/pkg/holderPg/multiTenantPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/gin-gonic/gin"
	"github.com/pangu-2/go-tools/tools/dbPg/genericPg"
	"github.com/pangu-2/go-tools/tools/dbPg/pagePg"
	"gorm.io/gorm"
)

type BaseCategoryRepository[T any, ID genericPg.ID] struct {
	Entity *T
	log    *log2.Logger `autowire:"?"`
	//从内部
	db *gorm.DB    `autowire:"?"`
	Pg configPg.Pg `value:"${pg}"`
	//
	ctx *gin.Context
}

func (b *BaseCategoryRepository[T, ID]) DbScopes() *gorm.DB {
	return b.db
}

func (b *BaseCategoryRepository[T, ID]) Db() *gorm.DB {
	return b.DbScopes()
}

func (b *BaseCategoryRepository[T, ID]) DbSource() *gorm.DB {
	return b.db
}
func (b *BaseCategoryRepository[T, ID]) DbModel() *gorm.DB {
	return b.DbScopes().Model(b.Entity)
}
func (b *BaseCategoryRepository[T, ID]) Log() *log2.Logger {
	return b.log
}

func (b *BaseCategoryRepository[T, ID]) SetOptionScopes(db *gorm.DB, opts ...Option) *gorm.DB {
	arg := OptionArg{}
	for _, opt := range opts {
		opt(&arg)
	}
	if nil != arg.Ctx {
		_, exists := arg.Ctx.Get(constContextPg.CTX_MULITI_TENANT)
		//b.log.Errorf("exists=xxxxxxx=%+v", exists)
		if exists {
			//解析表名称
			db.Statement.Parse(b.Entity)
			return db.Scopes(multiTenantPg.ScopeRulePgWhere(arg.Ctx, db.Statement.Schema.Table))
		}
	}
	return db
}

func (b *BaseCategoryRepository[T, ID]) Config() configPg.Pg {
	return b.Pg
}

// Create 创建
func (b *BaseCategoryRepository[T, ID]) Create(v *T) (error, int64) {
	tx := b.Db().Create(&v)
	return tx.Error, tx.RowsAffected
}

// 保存 会保存所有的字段，即使字段是零值
func (b *BaseCategoryRepository[T, ID]) Save(v *T, opts ...Option) error {
	tx := b.SetOptionScopes(b.DbModel(), opts...).Save(&v)
	return tx.Error
}

// 保存
func (b *BaseCategoryRepository[T, ID]) SaveAll(ts []*T, opts ...Option) error {
	if ts != nil && len(ts) > 0 {
		for _, info := range ts {
			if e := b.Save(info, opts...); e != nil {
				return e
			}
		}
	}
	return nil
}

// Update 更新 更新属性，只会更新非零值的字段
func (b *BaseCategoryRepository[T, ID]) Update(info T, id ID, opts ...Option) error {
	result := b.SetOptionScopes(b.DbModel(), opts...).Where("id=?", id).Updates(&info)
	if result.Error != nil {
		//log.Fatal("failed to connect database")
		return result.Error
	}
	return nil
}

// Update 更新 更新属性，只会更新非零值的字段
func (b *BaseCategoryRepository[T, ID]) UpdatePointer(info T, id ID) error {
	result := b.Db().Model(b.Entity).Where("id=?", id).Updates(info)
	if result.Error != nil {
		//log.Fatal("failed to connect database")
		return result.Error
	}
	return nil
}

// UpdateMap 更新, map里所有属性都会更新
func (b *BaseCategoryRepository[T, ID]) UpdateMap(info map[string]interface{}, id ID, opts ...Option) error {
	result := b.SetOptionScopes(b.DbModel(), opts...).Where("id=?", id).Updates(info)
	if result.Error != nil {
		//log.Fatal("failed to connect database")
		return result.Error
	}
	return nil
}

// UpdateStructMap 更新, 结构体转换为map，map里所有属性都会更新
func (b *BaseCategoryRepository[T, ID]) UpdateStructMap(info any, id ID, opts ...Option) error {
	toMap, err := convertor.StructToMap(info)
	if nil != err {
		return err
	}
	result := b.SetOptionScopes(b.DbModel(), opts...).Where("id=?", id).Updates(toMap)
	if result.Error != nil {
		//log.Fatal("failed to connect database")
		return result.Error
	}
	return nil
}

// Update 更新
func (b *BaseCategoryRepository[T, ID]) UpdateAll(ts []*T, opts ...Option) error {
	if ts != nil && len(ts) > 0 {
		for _, info := range ts {
			if e := b.Save(info); e != nil {
				return e
			}
		}
	}
	return nil
}

func (b *BaseCategoryRepository[T, ID]) DeleteById(id ID, opts ...Option) error {
	tx := b.SetOptionScopes(b.DbModel(), opts...).Delete(&b.Entity, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (b *BaseCategoryRepository[T, ID]) DeleteByIds(id []ID, opts ...Option) error {
	tx := b.SetOptionScopes(b.DbModel(), opts...).Delete(&b.Entity, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (b *BaseCategoryRepository[T, ID]) DeleteByIdsString(id []string, opts ...Option) error {
	tx := b.SetOptionScopes(b.DbModel(), opts...).Delete(&b.Entity, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (b *BaseCategoryRepository[T, ID]) DeleteAllByTenantNoAndIdsString(tenantNo string, id []string, opts ...Option) error {
	tx := b.SetOptionScopes(b.DbModel(), opts...).Where("tenant_no = ?", tenantNo).Delete(&b.Entity, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (b *BaseCategoryRepository[T, ID]) DeleteByNo(no string, opts ...Option) error {
	tx := b.SetOptionScopes(b.DbModel(), opts...).Where("no = ?", no).Delete(&b.Entity)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// FindById 根据主键查询
//
//	@Description:
//	@receiver b
//	@param id
//	@return info
//	@return result 是否查询到值
//	@return err
func (b *BaseCategoryRepository[T, ID]) FindById(id ID, opts ...Option) (info *T, result bool) {
	tx := b.SetOptionScopes(b.DbModel(), opts...).Where("id=?", id).First(&info)
	if tx.Error != nil {
		b.log.Errorf("error=%+v", tx.Error)
		return nil, false
	}
	if tx.RowsAffected == 0 {
		return nil, false
	}
	return info, true
}

// FindByIdString 根据主键查询
//
//	@Description:
//	@receiver b
//	@param id
//	@return info
//	@return result 是否查询到值
//	@return err
func (b *BaseCategoryRepository[T, ID]) FindByIdString(id string, opts ...Option) (info *T, result bool) {
	tx := b.SetOptionScopes(b.DbModel(), opts...).Where("id=?", id).First(&info)
	if tx.Error != nil {
		b.log.Errorf("error=%+v", tx.Error)
		return nil, false
	}
	if tx.RowsAffected == 0 {
		return nil, false
	}
	return info, true
}

// 查询所有
func (b *BaseCategoryRepository[T, ID]) FindAll(t T, arg ...interface{}) (infos []*T) {
	where := b.Db().Where(t)
	if nil != arg {
		for _, item := range arg {
			switch result := item.(type) {
			case Condition:
				// 条件
				where = result(where)
			case Option:
				where = b.SetOptionScopes(where, item.(Option))
			}
		}
	}
	tx := where.Find(&infos)
	if tx.Error != nil {
		return nil
	}
	return
}

// 查询所有-但是限制条数
func (b *BaseCategoryRepository[T, ID]) FindAllLimit(t T, limit int, arg ...interface{}) (infos []*T, result bool) {
	where := b.Db().Where(t).Limit(limit)
	if nil != arg {
		for _, item := range arg {
			switch result := item.(type) {
			case Condition:
				// 条件
				where = result(where)
			case Option:
				where = b.SetOptionScopes(where, item.(Option))
			}
		}
	}
	tx := where.Find(&infos)
	if tx.Error != nil {
		return nil, false
	}
	if tx.RowsAffected == 0 {
		return nil, false
	}
	return infos, true
}

// 查询所有
func (b *BaseCategoryRepository[T, ID]) FindAllData(arg ...interface{}) (infos []*T, result bool) {
	where := b.Db()
	if nil != arg {
		for _, item := range arg {
			switch result := item.(type) {
			case Condition:
				// 条件
				where = result(where)
			case Option:
				where = b.SetOptionScopes(where, item.(Option))
			}
		}
	}
	tx := where.Find(&infos)
	if tx.Error != nil {
		return nil, false
	}
	if tx.RowsAffected == 0 {
		return nil, false
	}
	return infos, true
}

// 分页
func (b *BaseCategoryRepository[T, ID]) FindAllPage(t T, option pagePg.Option[*T], opts ...Option) (pagePg.PaginatorPg[*T], error) {
	var total int64
	pg := pagePg.NewPaginatorPg[*T](option)
	countTx := b.SetOptionScopes(b.DbModel(), opts...).Where(t).Count(&total)
	if nil != countTx.Error {
		return pg, countTx.Error
	}
	var infos []*T
	tx := b.SetOptionScopes(b.DbModel(), opts...).Where(t).Find(&infos)
	//b.log.Debugf("sql=%+v", tx.Statement.SQL.String())
	if tx.Error != nil {
		return pg, tx.Error
	}
	pg.Data = infos
	pg.Total = total
	pg.Pageable = pagePg.NewPageablePg(total, pg.PageNum, pg.PageSize)
	pg.TotalPage = pg.Pageable.TotalPage
	//if total >0 {
	//	t2 := infos[len(infos)-1]
	//	pg.OffsetId=t2
	//}
	return pg, nil
}

// FindAllPageQuery 分页
func (b *BaseCategoryRepository[T, ID]) FindAllPageQuery(t T, option pagePg.OptionPageCondition[*T], opts ...Option) (pagePg.PaginatorPg[*T], error) {
	var total int64
	pg, condition := pagePg.NewOptionPageCondition[*T](option)
	if nil == condition {
		condition = b.SetOptionScopes(b.DbModel(), opts...)
	} else {
		condition = b.SetOptionScopes(condition, opts...)
	}
	countTx := condition.Where(t).Count(&total)
	if nil != countTx.Error {
		return pg, countTx.Error
	}
	var infos []*T
	tx := countTx.Scopes(pg.Scopes()).Find(&infos)
	//b.log.Infof("sql=%+v", tx.Statement.SQL.String())
	if tx.Error != nil {
		return pg, tx.Error
	}
	pg.Data = infos
	pg.Total = total
	pg.Pageable = pagePg.NewPageablePg(total, pg.PageNum, pg.PageSize)
	pg.TotalPage = pg.Pageable.TotalPage
	//if total >0 {
	//	t2 := infos[len(infos)-1]
	//	pg.OffsetId=t2
	//}
	return pg, nil
}

// FindAllByIdIn
//
//	@Description:
//	@receiver b
//	@param ids
//	@return infos
//	@return result true: 有值;    false: 错误或 没查询到
func (b *BaseCategoryRepository[T, ID]) FindAllByIdIn(ids []ID, opts ...Option) (infos []*T, result bool) {
	infos = make([]*T, 0)
	tx := b.SetOptionScopes(b.DbModel(), opts...).Where("id in (?)", ids).Find(&infos)
	if tx.Error != nil {
		return nil, false
	}
	if tx.RowsAffected == 0 {
		return nil, false
	}
	result = true
	return
}

// FindAllByIdStringIn
//
//	@Description:
//	@receiver b
//	@param ids
//	@return infos
//	@return result true: 有值;    false: 错误或 没查询到
func (b *BaseCategoryRepository[T, ID]) FindAllByIdStringIn(ids []string, opts ...Option) (infos []*T, result bool) {
	tx := b.SetOptionScopes(b.DbModel(), opts...).Where("id in (?)", ids).Find(&infos)
	if tx.Error != nil {
		return nil, false
	}
	if tx.RowsAffected == 0 {
		return nil, false
	}
	result = true
	return
}

// Count
//
//	@Description:   统计
//	@receiver b
//	@param arg
//	@return total
//	@return result
func (b *BaseCategoryRepository[T, ID]) Count(arg ...interface{}) (total int64, result bool) {
	where := b.Db()
	if nil != arg {
		for _, item := range arg {
			switch result := item.(type) {
			case Condition:
				// 条件
				where = result(where)
			case Option:
				where = b.SetOptionScopes(where, item.(Option))
			}
		}
	}
	tx := where.Count(&total)
	if tx.Error != nil {
		return 0, false
	}
	if tx.RowsAffected == 0 {
		return 0, false
	}
	return total, true
}

// FindByNo 根据主键查询
//
//	@Description:
//	@receiver b
//	@param no
//	@return info
//	@return result 是否查询到值
//	@return err
func (b *BaseCategoryRepository[T, ID]) FindByNo(no string, opts ...Option) (info *T, result bool) {
	tx := b.SetOptionScopes(b.DbModel(), opts...).Where("no=?", no).First(&info)
	if tx.Error != nil {
		b.log.Errorf("error=%+v", tx.Error)
		return nil, false
	}
	if tx.RowsAffected == 0 {
		return nil, false
	}
	return info, true
}

// FindAllByNoIn 根据主键查询
//
//	@Description:
//	@receiver b
//	@param no
//	@return info
//	@return result 是否查询到值
//	@return err
func (b *BaseCategoryRepository[T, ID]) FindAllByNoIn(no []string, opts ...Option) (info []*T, result bool) {
	tx := b.SetOptionScopes(b.DbModel(), opts...).Where("no in ?", no).Find(&info)
	if tx.Error != nil {
		b.log.Errorf("error=%+v", tx.Error)
		return nil, false
	}
	if tx.RowsAffected == 0 {
		return nil, false
	}
	return info, true
}

// FindAllByNameIn 根据主键查询
//
//	@Description:
//	@receiver b
//	@param no
//	@return info
//	@return result 是否查询到值
//	@return err
func (b *BaseCategoryRepository[T, ID]) FindAllByNameIn(no []string, opts ...Option) (info []*T, result bool) {
	tx := b.SetOptionScopes(b.DbModel(), opts...).Where("name in ?", no).Find(&info)
	if tx.Error != nil {
		b.log.Errorf("error=%+v", tx.Error)
		return nil, false
	}
	if tx.RowsAffected == 0 {
		return nil, false
	}
	return info, true
}

// FindByName 根据主键查询
//
//	@Description:
//	@receiver b
//	@param no
//	@return info
//	@return result 是否查询到值
//	@return err
func (b *BaseCategoryRepository[T, ID]) FindByName(no string, opts ...Option) (info *T, result bool) {
	tx := b.SetOptionScopes(b.DbModel(), opts...).Where("name=?", no).First(&info)
	if tx.Error != nil {
		b.log.Errorf("error=%+v", tx.Error)
		return nil, false
	}
	if tx.RowsAffected == 0 {
		return nil, false
	}
	return info, true
}

// FindByNameAndIdNot
//
//	@Description:
//	@receiver c
//	@param name
//	@param id
//	@return info
//	@return result
func (b *BaseCategoryRepository[T, ID]) FindByNameAndIdNot(name string, id string, opts ...Option) (info *T, result bool) {
	tx := b.SetOptionScopes(b.DbModel(), opts...).Where("name=?", name).Where("id <> ?", id).First(&info)
	if tx.Error != nil {
		b.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

// FindByNoAndIdNot
//
//	@Description:
//	@receiver c
//	@param name
//	@param id
//	@return info
//	@return result
func (c *BaseCategoryRepository[T, ID]) FindByNoAndIdNot(name string, id string, opts ...Option) (info *T, result bool) {
	tx := c.SetOptionScopes(c.DbModel(), opts...).Where("no=?", name).Where("id <> ?", id).First(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

// FindAllByNoLink
//
//	@Description:
//	@receiver c
//	@param name
//	@return info
//	@return result
func (c *BaseCategoryRepository[T, ID]) FindAllByNoLink(code string, opts ...Option) (info []*T, result bool) {
	tx := c.SetOptionScopes(c.DbModel(), opts...).Where("no_link like ?", "%|"+code+"|%").Find(&info)
	if tx.Error != nil {
		c.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

// FindByCode 根据主键查询
//
//	@Description:
//	@receiver b
//	@param no
//	@return info
//	@return result 是否查询到值
//	@return err
func (c *BaseCategoryRepository[T, ID]) FindByCode(no string, opts ...Option) (info *T, result bool) {
	tx := c.SetOptionScopes(c.DbModel(), opts...).Where("code=?", no).First(&info)
	if tx.Error != nil {
		c.log.Errorf("error=%+v", tx.Error)
		return nil, false
	}
	if tx.RowsAffected == 0 {
		return nil, false
	}
	return info, true
}

// FindByCodeAndIdNot
//
//	@Description:
//	@receiver c
//	@param name
//	@param id
//	@return info
//	@return result
func (b *BaseCategoryRepository[T, ID]) FindByCodeAndIdNot(name string, id string, opts ...Option) (info *T, result bool) {
	tx := b.SetOptionScopes(b.DbModel(), opts...).Where("code=?", name).Where("id <> ?", id).First(&info)
	if tx.Error != nil {
		b.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

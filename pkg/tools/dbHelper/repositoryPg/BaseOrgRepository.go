package repositoryPg

import (
	"context"

	"github.com/duke-git/lancet/v2/convertor"
	"github.com/foxiswho/blog-go/pkg/configPg"
	"github.com/foxiswho/blog-go/pkg/consts/constContextPg"
	"github.com/foxiswho/blog-go/pkg/holderPg/multiTenantPg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/pangu-2/go-tools/tools/dbPg/genericPg"
	"github.com/pangu-2/go-tools/tools/dbPg/pagePg"
	"gorm.io/gorm"
)

type BaseOrgRepository[T any, ID genericPg.ID] struct {
	Entity *T
	log    *log2.Logger `autowire:"?"`
	//从内部
	db *gorm.DB    `autowire:"?"`
	Pg configPg.Pg `value:"${pg}"`
}

func (b *BaseOrgRepository[T, ID]) DbScopes() *gorm.DB {
	return b.db
}

func (b *BaseOrgRepository[T, ID]) Db() *gorm.DB {
	return b.DbScopes()
}

func (b *BaseOrgRepository[T, ID]) DbSource() *gorm.DB {
	return b.db
}
func (b *BaseOrgRepository[T, ID]) DbModel() *gorm.DB {
	return b.DbScopes().Model(b.Entity)
}
func (b *BaseOrgRepository[T, ID]) Log() *log2.Logger {
	return b.log
}

func (b *BaseOrgRepository[T, ID]) SetOptionScopes(db *gorm.DB, opts ...OptionPg) *gorm.DB {
	if nil == opts || len(opts) == 0 {
		return db
	}
	arg := OptionParams{
		Db: db,
	}
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

func (b *BaseOrgRepository[T, ID]) SetOptionPgScopes(db *gorm.DB, opts ...OptionPg) (*gorm.DB, OptionParams) {
	arg := OptionParams{
		Db: db,
	}
	if nil == opts || len(opts) == 0 {
		return db, arg
	}
	for _, opt := range opts {
		opt(&arg)
	}
	if nil != arg.Ctx {
		_, exists := arg.Ctx.Get(constContextPg.CTX_MULITI_TENANT)
		//b.log.Errorf("exists=xxxxxxx=%+v", exists)
		if exists {
			//解析表名称
			arg.Db.Statement.Parse(b.Entity)
			return arg.Db.Scopes(multiTenantPg.ScopeRulePgWhere(arg.Ctx, db.Statement.Schema.Table)), arg
		}
	}
	return arg.Db, arg
}

func (b *BaseOrgRepository[T, ID]) Config() configPg.Pg {
	return b.Pg
}

// Create 创建
func (b *BaseOrgRepository[T, ID]) Create(ctx context.Context, v *T) (error, int64) {
	tx := b.Db().WithContext(ctx).Create(&v)
	return tx.Error, tx.RowsAffected
}

// 保存 会保存所有的字段，即使字段是零值
func (b *BaseOrgRepository[T, ID]) Save(ctx context.Context, v *T, opts ...OptionPg) error {
	tx := b.SetOptionScopes(b.DbModel().WithContext(ctx), opts...).Save(&v)
	return tx.Error
}

// 保存
func (b *BaseOrgRepository[T, ID]) SaveAll(ctx context.Context, ts []*T, opts ...OptionPg) error {
	if ts != nil && len(ts) > 0 {
		for _, info := range ts {
			if e := b.Save(ctx, info, opts...); e != nil {
				return e
			}
		}
	}
	return nil
}

// Update 更新 更新属性，只会更新非零值的字段
func (b *BaseOrgRepository[T, ID]) Update(ctx context.Context, info T, id ID, opts ...OptionPg) error {
	result := b.SetOptionScopes(b.DbModel().WithContext(ctx), opts...).Where("id=?", id).Updates(&info)
	if result.Error != nil {
		//log.Fatal("failed to connect database")
		return result.Error
	}
	return nil
}

// Update 更新 更新属性，只会更新非零值的字段
func (b *BaseOrgRepository[T, ID]) UpdatePointer(ctx context.Context, info T, id ID) error {
	result := b.Db().WithContext(ctx).Model(b.Entity).Where("id=?", id).Updates(info)
	if result.Error != nil {
		//log.Fatal("failed to connect database")
		return result.Error
	}
	return nil
}

// UpdatePointerObject 更新  更新属性，只会更新非零值的字段
func (b *BaseOrgRepository[T, ID]) UpdatePointerObject(ctx context.Context, info any, id ID) error {
	result := b.Db().WithContext(ctx).Model(b.Entity).Where("id=?", id).Updates(info)
	if result.Error != nil {
		//log.Fatal("failed to connect database")
		return result.Error
	}
	return nil
}

// UpdateMap 更新, map里所有属性都会更新
func (b *BaseOrgRepository[T, ID]) UpdateMap(ctx context.Context, info map[string]interface{}, id ID, opts ...OptionPg) error {
	result := b.SetOptionScopes(b.DbModel().WithContext(ctx), opts...).Where("id=?", id).Updates(info)
	if result.Error != nil {
		//log.Fatal("failed to connect database")
		return result.Error
	}
	return nil
}

// UpdateStructMap 更新, 结构体转换为map，map里所有属性都会更新
func (b *BaseOrgRepository[T, ID]) UpdateStructMap(ctx context.Context, info any, id ID, opts ...OptionPg) error {
	toMap, err := convertor.StructToMap(info)
	if nil != err {
		return err
	}
	result := b.SetOptionScopes(b.DbModel().WithContext(ctx), opts...).Where("id=?", id).Updates(toMap)
	if result.Error != nil {
		//log.Fatal("failed to connect database")
		return result.Error
	}
	return nil
}

// Update 更新
func (b *BaseOrgRepository[T, ID]) UpdateAll(ctx context.Context, ts []*T, opts ...OptionPg) error {
	if ts != nil && len(ts) > 0 {
		for _, info := range ts {
			if e := b.Save(ctx, info); e != nil {
				return e
			}
		}
	}
	return nil
}

func (b *BaseOrgRepository[T, ID]) DeleteById(ctx context.Context, id ID, opts ...OptionPg) error {
	tx := b.SetOptionScopes(b.DbModel().WithContext(ctx), opts...).Delete(&b.Entity, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (b *BaseOrgRepository[T, ID]) DeleteByIds(ctx context.Context, id []ID, opts ...OptionPg) error {
	if nil != opts {
		tx := b.DbModel().WithContext(ctx).Delete(&b.Entity, id)
		if tx.Error != nil {
			return tx.Error
		}
	} else {
		tx := b.SetOptionScopes(b.DbModel().WithContext(ctx), opts...).Delete(&b.Entity, id)
		if tx.Error != nil {
			return tx.Error
		}
	}
	return nil
}

func (b *BaseOrgRepository[T, ID]) DeleteByIdsString(ctx context.Context, id []string, opts ...OptionPg) error {
	tx := b.SetOptionScopes(b.DbModel().WithContext(ctx), opts...).Delete(&b.Entity, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (b *BaseOrgRepository[T, ID]) DeleteAllByTenantNoAndIdsString(ctx context.Context, tenantNo string, id []string, opts ...OptionPg) error {
	tx := b.SetOptionScopes(b.DbModel().WithContext(ctx), opts...).Where("tenant_no = ?", tenantNo).Delete(&b.Entity, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (b *BaseOrgRepository[T, ID]) DeleteByNo(ctx context.Context, no string, opts ...OptionPg) error {
	tx := b.SetOptionScopes(b.DbModel().WithContext(ctx), opts...).Where("no = ?", no).Delete(&b.Entity)
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
func (b *BaseOrgRepository[T, ID]) FindById(ctx context.Context, id ID, opts ...OptionPg) (info *T, result bool) {
	tx := b.SetOptionScopes(b.DbModel().WithContext(ctx), opts...).Where("id=?", id).First(&info)
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
func (b *BaseOrgRepository[T, ID]) FindByIdString(ctx context.Context, id string, opts ...OptionPg) (info *T, result bool) {
	tx := b.SetOptionScopes(b.DbModel().WithContext(ctx), opts...).Where("id=?", id).First(&info)
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
func (b *BaseOrgRepository[T, ID]) FindAll(ctx context.Context, t T, arg ...interface{}) (infos []*T) {
	where := b.Db().WithContext(ctx).Where(t)
	if nil != arg {
		for _, item := range arg {
			switch result := item.(type) {
			case Condition:
				where = result(where)
			case Option:
				where = b.SetOptionScopes(where, item.(OptionPg))
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
func (b *BaseOrgRepository[T, ID]) FindAllLimit(ctx context.Context, t T, limit int, arg ...interface{}) (infos []*T, result bool) {
	where := b.Db().WithContext(ctx).Where(t).Limit(limit)
	if nil != arg {
		for _, item := range arg {
			switch result := item.(type) {
			case Condition:
				where = result(where)
			case Option:
				where = b.SetOptionScopes(where, item.(OptionPg))
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
func (b *BaseOrgRepository[T, ID]) FindAllData(ctx context.Context, arg ...interface{}) (infos []*T, result bool) {
	where := b.Db().WithContext(ctx)
	if nil != arg {
		for _, item := range arg {
			switch result := item.(type) {
			case Condition:
				where = result(where)
			case Option:
				where = b.SetOptionScopes(where, item.(OptionPg))
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
func (b *BaseOrgRepository[T, ID]) FindAllPage(ctx context.Context, t T, opts ...OptionPg) (pagePg.Paginator[*T], error) {
	var total int64
	//
	condition, arg := b.SetOptionPgScopes(b.DbModel().WithContext(ctx), opts...)
	//
	pageable := arg.Pageable
	if nil == pageable {
		pageable = &pagePg.Pageable{Total: total, PageNum: 0, PageSize: 10}
	}
	pg := pagePg.NewPaginator[*T]()
	//
	countTx := condition.Where(t).Count(&total)
	if nil != countTx.Error {
		return pg, countTx.Error
	}
	var infos []*T
	tx := countTx.Scopes(Scopes(pageable)).Find(&infos)
	//b.log.Infof("sql=%+v", tx.Statement.SQL.String())
	if tx.Error != nil {
		return pg, tx.Error
	}
	pg.Data = infos
	pg.Total = total
	pg.Pageable = pagePg.NewPageable(total, pageable.PageNum, pageable.PageSize)
	pg.TotalPage = pg.Pageable.TotalPage
	//if total >0 {
	//	t2 := infos[len(infos)-1]
	//	pg.OffsetId=t2
	//}
	return pg, nil
}

// FindAllPageQuery 分页
func (b *BaseOrgRepository[T, ID]) FindAllPageQuery(ctx context.Context, t T, opts ...OptionPg) (pagePg.Paginator[*T], error) {
	var total int64
	//
	condition, arg := b.SetOptionPgScopes(b.DbModel().WithContext(ctx), opts...)
	//
	pageable := arg.Pageable
	if nil == pageable {
		pageable = &pagePg.Pageable{Total: total, PageNum: 0, PageSize: 10}
	}
	pg := pagePg.NewPaginator[*T]()
	//
	countTx := condition.Where(t).Count(&total)
	if nil != countTx.Error {
		return pg, countTx.Error
	}
	var infos []*T
	tx := countTx.Scopes(Scopes(pageable)).Find(&infos)
	//b.log.Infof("sql=%+v", tx.Statement.SQL.String())
	if tx.Error != nil {
		return pg, tx.Error
	}
	pg.Data = infos
	pg.Total = total
	pg.Pageable = pagePg.NewPageable(total, pageable.PageNum, pageable.PageSize)
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
func (b *BaseOrgRepository[T, ID]) FindAllByIdIn(ctx context.Context, ids []ID, opts ...OptionPg) (infos []*T, result bool) {
	infos = make([]*T, 0)
	tx := b.SetOptionScopes(b.DbModel().WithContext(ctx), opts...).Where("id in (?)", ids).Find(&infos)
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
func (b *BaseOrgRepository[T, ID]) FindAllByIdStringIn(ctx context.Context, ids []string, opts ...OptionPg) (infos []*T, result bool) {
	tx := b.SetOptionScopes(b.DbModel().WithContext(ctx), opts...).Where("id in (?)", ids).Find(&infos)
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
func (b *BaseOrgRepository[T, ID]) Count(ctx context.Context, arg ...interface{}) (total int64, result bool) {
	where := b.Db().WithContext(ctx)
	if nil != arg {
		for _, item := range arg {
			switch result := item.(type) {
			case Condition:
				// 条件
				where = result(where)
			case Option:
				where = b.SetOptionScopes(where, item.(OptionPg))
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
func (b *BaseOrgRepository[T, ID]) FindByNo(ctx context.Context, no string, opts ...OptionPg) (info *T, result bool) {
	tx := b.SetOptionScopes(b.DbModel().WithContext(ctx), opts...).Where("no=?", no).First(&info)
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
func (b *BaseOrgRepository[T, ID]) FindAllByNoIn(ctx context.Context, no []string, opts ...OptionPg) (info []*T, result bool) {
	tx := b.SetOptionScopes(b.DbModel().WithContext(ctx), opts...).Where("no in ?", no).Find(&info)
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
func (b *BaseOrgRepository[T, ID]) FindAllByNameIn(ctx context.Context, no []string, opts ...OptionPg) (info []*T, result bool) {
	tx := b.SetOptionScopes(b.DbModel().WithContext(ctx), opts...).Where("name in ?", no).Find(&info)
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
func (b *BaseOrgRepository[T, ID]) FindByName(ctx context.Context, no string, opts ...OptionPg) (info *T, result bool) {
	tx := b.SetOptionScopes(b.DbModel().WithContext(ctx), opts...).Where("name=?", no).First(&info)
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
func (b *BaseOrgRepository[T, ID]) FindByNameAndIdNot(ctx context.Context, name string, id string, opts ...OptionPg) (info *T, result bool) {
	tx := b.SetOptionScopes(b.DbModel().WithContext(ctx), opts...).Where("name=?", name).Where("id <> ?", id).First(&info)
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
func (c *BaseOrgRepository[T, ID]) FindByNoAndIdNot(ctx context.Context, name string, id string, opts ...OptionPg) (info *T, result bool) {
	tx := c.SetOptionScopes(c.DbModel().WithContext(ctx), opts...).Where("no=?", name).Where("id <> ?", id).First(&info)
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
func (c *BaseOrgRepository[T, ID]) FindAllByNoLink(ctx context.Context, code string, opts ...OptionPg) (info []*T, result bool) {
	tx := c.SetOptionScopes(c.DbModel().WithContext(ctx), opts...).Where("no_link like ?", "%|"+code+"|%").Find(&info)
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
func (c *BaseOrgRepository[T, ID]) FindByCode(ctx context.Context, no string, opts ...OptionPg) (info *T, result bool) {
	tx := c.SetOptionScopes(c.DbModel().WithContext(ctx), opts...).Where("code=?", no).First(&info)
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
func (b *BaseOrgRepository[T, ID]) FindByCodeAndIdNot(ctx context.Context, name string, id string, opts ...OptionPg) (info *T, result bool) {
	tx := b.SetOptionScopes(b.DbModel().WithContext(ctx), opts...).Where("code=?", name).Where("id <> ?", id).First(&info)
	if tx.Error != nil {
		b.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

// FindByCodeAndNoNot
//
//	@Description:
//	@receiver c
//	@param name
//	@param no
//	@return info
//	@return result
func (b *BaseOrgRepository[T, ID]) FindByCodeAndNoNot(ctx context.Context, name string, no string, opts ...OptionPg) (info *T, result bool) {
	tx := b.SetOptionScopes(b.DbModel().WithContext(ctx), opts...).Where("code=?", name).Where("no != ?", no).First(&info)
	if tx.Error != nil {
		b.Log().Error("", tx.Error)
		return nil, false
	}
	if 0 == tx.RowsAffected {
		return nil, false
	}
	return info, true
}

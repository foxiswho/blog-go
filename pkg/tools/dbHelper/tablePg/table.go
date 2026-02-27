package tablePg

import (
	"gorm.io/gorm/schema"
	"reflect"
)

type TableCommenter interface {
	TableComment() string
}

// 获取表名
func GetTableName(model interface{}) string {
	if tabler, ok := model.(schema.Tabler); ok {
		return tabler.TableName()
	}
	return schema.NamingStrategy{}.TableName(reflect.TypeOf(model).Elem().Name())
}

// 获取表注释
func GetTableComment(model interface{}) string {
	if commenter, ok := model.(TableCommenter); ok {
		return commenter.TableComment()
	}
	return ""
}

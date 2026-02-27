package noPg

import (
	"github.com/foxiswho/blog-go/pkg/consts/constNoPg"
	"github.com/pangu-2/go-tools/tools/strPg"
)

// No
//
//	@Description: 编号
//	@return string
func No() string {
	return strPg.GenerateNumberId22ByPrefix("")
}

// MakeNo
//
//	@Description: 生成指定前缀的编号
//	@param prefix
//	@return string
func MakeNo(prefix string) string {
	return strPg.GenerateNumberId22ByPrefix(prefix)
}

// TenantNo
//
//	@Description:  租户编号
//	@return string
func TenantNo() string {
	return MakeNo(constNoPg.Tenant)
}

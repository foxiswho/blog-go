package noPg2

import (
	"github.com/foxiswho/blog-go/pkg/consts/constNoPg"
	"github.com/pangu-2/go-tools/tools/noPg"
)

// TenantNo
//
//	@Description:  租户编号
//	@return string
func TenantNo() string {
	return noPg.MakeNo(constNoPg.Tenant)
}

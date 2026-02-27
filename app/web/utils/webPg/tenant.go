package webPg

import (
	"github.com/foxiswho/blog-go/pkg/consts/constHeaderPg"
	"github.com/gin-gonic/gin"
)

// GetTenantNo
//
//	@Description:
//	@param ctx
//	@return string
func GetTenantNo(ctx *gin.Context) string {
	value, exists := ctx.Get(constHeaderPg.WebTenantNo)
	if exists {
		return value.(string)
	}
	return "-1"
}

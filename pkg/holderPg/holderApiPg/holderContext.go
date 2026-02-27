package holderApiPg

import (
	"context"
	"github.com/foxiswho/blog-go/pkg/consts/constContextPg"
	"github.com/gin-gonic/gin"
)

func GetContextAccount(ctx *gin.Context) (holder HolderPg) {
	get, is := ctx.Get(constContextPg.AUTH_LOGIN_API)
	if !is || nil == get {
		return
	}
	holder = get.(HolderPg)
	return
}

func SetContextValue(ctx *gin.Context) context.Context {
	return context.WithValue(ctx.Request.Context(), constContextPg.CTX, ctx)
}

func SetContextValueGs(ctx *gin.Context) context.Context {
	return context.WithValue(ctx.Request.Context(), constContextPg.CTX, ctx)
}

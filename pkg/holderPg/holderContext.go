package holderPg

import (
	"context"

	"github.com/foxiswho/blog-go/pkg/consts/constContextPg"
	"github.com/gin-gonic/gin"
	"github.com/go-spring/log"
)

func GetContextAccount(ctx *gin.Context) (holder HolderPg) {
	get, is := ctx.Get(constContextPg.AUTH_LOGIN)
	log.Infof(context.Background(), log.TagAppDef, "holder-get:%+v", get)
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

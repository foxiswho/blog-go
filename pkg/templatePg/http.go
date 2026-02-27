package templatePg

import (
	"github.com/gin-gonic/gin"
	"net/url"
)

type HttpPg struct {
	Path string   `json:"path"`
	URL  *url.URL `json:"URL"`
}

func NewHttpPg(ctx *gin.Context) *HttpPg {
	return &HttpPg{Path: ctx.Request.URL.Path, URL: ctx.Request.URL}
}

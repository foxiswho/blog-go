package modBlogArticleCategory

import (
	"github.com/foxiswho/blog-go/pkg/model"
)

type QueryPublicCt struct {
	model.BaseQueryNodeCt
	ParentNo string `json:"parentNo" label:"上级" `
}

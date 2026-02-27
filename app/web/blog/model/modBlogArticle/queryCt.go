package modBlogArticle

import (
	"github.com/foxiswho/blog-go/pkg/model"
	"github.com/foxiswho/blog-go/pkg/tools/typePg"
)

type QueryCt struct {
	model.BaseQueryCt
	ID        typePg.Uint64String `json:"id" label:"" `
	Name      string              `json:"name" label:"名称" `
	No        string              `json:"no" label:"编号代号" `
	Q         string              `json:"q" label:"关键词"  form:"q"`
	TagsQuery []string            `json:"tagsQuery" label:"标签" form:"tags"`
}

package modBlogTopicRelation

import (
	"github.com/foxiswho/blog-go/pkg/model"
	"github.com/foxiswho/blog-go/pkg/tools/typePg"
)

type QueryCt struct {
	model.BaseQueryCt
	ID           typePg.Uint64String `json:"id" label:"" `
	Name         string              `json:"name" label:"名称" `
	Version      string              `json:"version" label:"版本" `
	TenantNo     string              `json:"tenantNo" label:"租户编号" `
	TenantNoName string              `json:"tenantNoName" label:"租户" `
	Description  string              `json:"description" label:"描述" `
	TopicNo      string              `json:"topicNo" label:"话题编号" `
	ArticleNo    string              `json:"articleNo" label:"文章编号" `
}

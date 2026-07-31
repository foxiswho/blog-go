package modRamIdpMetadataCache

import (
	"github.com/hongmengzhu/xianfu-blog-go/pkg/model"
	"github.com/hongmengzhu/xianfu-blog-go/pkg/tools/typePg"
	"time"
)

type QueryCt struct {
	model.BaseQueryCt
	ID          typePg.Uint64String `json:"id" label:"" `
	IdpNo       string              `json:"idpNo" label:"身份提供商编号" ` // 身份提供商编号
	SourceNo    string              `json:"sourceNo" label:"认证源编号" `  // 认证源编号
	Protocol    string              `json:"protocol" label:"协议" `        // 协议
	State       typePg.Int8         `json:"state" label:"状态:1启用;2失效" ` // 状态:1启用;2失效
	CreateAt    *time.Time          `json:"createAt" label:"创建时间" `     // 创建时间
}

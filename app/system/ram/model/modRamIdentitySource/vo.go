package modRamIdentitySource

import (
	"github.com/hongmengzhu/xianfu-blog-go/pkg/tools/typePg"
	"time"
)

type Vo struct {
	ID                 typePg.Uint64String `json:"id" label:"id" `
	Name               string              `json:"name" label:"名称" `     // 名称
	NameFl             string              `json:"nameFl" label:"名称外文" ` // 名称外文
	Code               string              `json:"code" label:"标志" `
	NameFull           string              `json:"nameFull" label:"全称" `      // 全称
	State              typePg.Int8         `json:"state" label:"状态:1启用;2禁用" ` // 状态:1启用;2禁用
	Description        string              `json:"description" label:"描述" `   // 描述
	Icon               string              `json:"icon" label:"图标" `
	Protocol           string              `json:"protocol" label:"协议" `
	Idp                string              `json:"idp" label:"身份提供商" `
	Configured         typePg.Int8         `json:"configured" label:"是否配置" `
	EnableSlo          typePg.Int8         `json:"enableSlo" label:"单点登出" `
	AutoCreateUser     typePg.Int8         `json:"autoCreateUser" label:"自动创建用户" `
	CreateAt           *time.Time          `json:"createAt" label:"创建时间" ` // 创建时间
	UpdateAt           *time.Time          `json:"updateAt" label:"更新时间" ` // 更新时间
	CreateBy           string              `json:"createBy" label:"创建人" `  // 创建人
	UpdateBy           string              `json:"updateBy" label:"更新人" `  // 更新人
}

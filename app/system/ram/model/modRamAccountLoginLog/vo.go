package modRamAccountLoginLog

import (
	"github.com/foxiswho/blog-go/pkg/tools/typePg"
)

type Vo struct {
	ID          typePg.Uint64String `json:"id" label:"id" `
	CreateAt    *typePg.Time        `json:"createAt" label:"创建时间" `
	Ano         string              `json:"ano" label:"编号" `
	LoginSource string              `json:"loginSource" label:"登录来源" `
	AppNo       string              `json:"appNo" label:"应用编号" `
	Client      string              `json:"client" label:"客户端" `
	TenantNo    string              `json:"tenantNo" label:"租户编号" `
	System      string              `json:"system" label:"系统" `
	Ip          string              `json:"ip" label:"ip" `
	Device      string              `json:"device" label:"设备" `
	DeviceNo    string              `json:"deviceNo" label:"设备编号" `
	Version     string              `json:"version" label:"版本" `
	UserAgent   string              `json:"userAgent" label:"用户代理" `
	ExtraData   map[string]any      `json:"extraData"`
}

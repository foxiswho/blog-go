package modSysPub

type RamMenuRouterVo struct {
	OtherAuth string                `json:"otherAuth" label:"其他权限"`
	DataCodes []string              `json:"dataCodes" label:"按钮权限"`
	Data      []RamMenuRouterMetaVo `json:"data" label:"路由"`
}

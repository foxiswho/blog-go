package settingPg

// ManageConfig
// @Description: 管理设置
type ManageConfig struct {
	SiteOpen bool `json:"siteOpen"` //打开还是关闭
}

func NewManageConfig(open bool) *ManageConfig {
	return &ManageConfig{SiteOpen: open}
}

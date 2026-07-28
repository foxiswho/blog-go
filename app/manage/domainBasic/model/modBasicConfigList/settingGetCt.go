package modBasicConfigList

type SettingGet struct {
	EventNo string `json:"eventNo" validate:"required" label:"事件编号"`
}

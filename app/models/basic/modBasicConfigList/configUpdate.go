package modBasicConfigList

type ConfigUpdateCt struct {
	Data    map[string]interface{} `json:"data"`
	EventNo string                 `json:"eventNo" label:"事件编号"`
}

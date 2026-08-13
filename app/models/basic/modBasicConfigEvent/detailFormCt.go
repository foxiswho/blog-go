package modBasicConfigEvent

type DetailFormCt struct {
	EventNo string `json:"eventNo" validate:"required" label:"事件编号"`
}

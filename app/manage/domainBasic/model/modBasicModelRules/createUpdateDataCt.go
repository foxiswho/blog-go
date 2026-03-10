package modBasicModelRules

type CreateUpdateDataCt struct {
	ValueNo    string   `json:"valueNo" label:"值编号/模块编号"`
	Body       []ItemVo `json:"body" label:"表体"`
	BodyDelIds []string `json:"bodyDelIds" label:"表体删除的id"`
}

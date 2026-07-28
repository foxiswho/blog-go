package modBasicModelRules

type CreateUpdateDataCt struct {
	FieldNo    string   `json:"fieldNo" label:"值编号/模块编号"`
	Body       []ItemVo `json:"body" label:"表体"`
	BodyDelIds []string `json:"bodyDelIds" label:"表体删除的id"`
}

package modBasicConfigModel

type CreateUpdateCt struct {
	Header     HeaderVo `json:"header" label:"表头"`
	Body       []ItemVo `json:"body" label:"表体"`
	BodyDelIds []string `json:"bodyDelIds" label:"表体删除的id"`
}

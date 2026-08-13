package modRamResourceRelation

type QuerySelectedCt struct {
	Code string `json:"code" label:"码值" validate:"required,min=1,max=255"`
}

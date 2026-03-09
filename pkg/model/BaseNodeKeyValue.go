package model

type BaseNodeKeyValue struct {
	Key      string      `json:"key" label:"键"`
	Label    string      `json:"label" label:"值"`
	ParentId string      `json:"parentId" label:"上级"`
	Extend   interface{} `json:"extend"`
}

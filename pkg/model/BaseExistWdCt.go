package model

// BaseExistWdCt 基础 详情
type BaseExistWdCt[ID any] struct {
	Id ID     `json:"id"`
	Wd string `json:"wd"`
}

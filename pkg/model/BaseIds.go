package model

import (
	"github.com/foxiswho/blog-go/pkg/interfaces"
)

// BaseIds 基础 详情
type BaseIds[ID any] struct {
	Ids    []ID                 `json:"ids"`
	Holder interfaces.IHolderPg `json:"holder"` // 会话信息
}

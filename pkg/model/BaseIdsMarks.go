package model

import (
	"github.com/foxiswho/blog-go/pkg/interfaces"
)

// BaseIdMarks 基础 详情
type BaseIdsMarks[ID, MARK any] struct {
	Ids    []ID                 `json:"ids"`
	Marks  []MARK               `json:"marks"`
	Holder interfaces.IHolderPg `json:"holder"` // 会话信息
}

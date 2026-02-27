package model

import (
	"github.com/foxiswho/blog-go/pkg/interfaces"
)

// BaseWdArr 基础 详情
type BaseWdArr struct {
	Wd     []string             `json:"wd"`
	Holder interfaces.IHolderPg `json:"holder"` // 会话信息
}

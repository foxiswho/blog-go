package model

import (
	"github.com/foxiswho/blog-go/pkg/interfaces"
)

// BaseWd 基础 详情
type BaseWd struct {
	Wd     string               `json:"wd"`
	Holder interfaces.IHolderPg `json:"holder"` // 会话信息
}

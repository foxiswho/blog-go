package model

import (
	"encoding/json"

	"github.com/foxiswho/blog-go/pkg/interfaces"
)

type BaseQueryByIdNotIdVdo[ID any] struct {
	Id     ID                   `json:"id"`
	IdNot  ID                   `json:"idNot"`
	Holder interfaces.IHolderPg `json:"holder"` // 会话信息
}

func (c BaseQueryByIdNotIdVdo[ID]) ToJsonString() (string, error) {
	marshal, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(marshal), nil
}

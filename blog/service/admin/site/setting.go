package site

import "github.com/foxiswho/blog-go/blog/model"

type Setting struct {
	*model.Type
	SettingsRadio []SettingRadio
	TypeForm      string
}

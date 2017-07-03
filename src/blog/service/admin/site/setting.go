package site

import "blog/model"

type Setting struct {
	*model.Type
	SettingsRadio []SettingRadio
	TypeForm      string
}

package modPublic

import (
	"github.com/foxiswho/blog-go/app/system/pub/model/modSysPub"
	"github.com/foxiswho/blog-go/app/system/ram/model/modRamAccount"
)

type InfoPublicVo struct {
	Info modRamAccount.AccountPub  `json:"info"`
	Menu modSysPub.RamMenuRouterVo `json:"menuRouter" label:"菜单路由"`
}

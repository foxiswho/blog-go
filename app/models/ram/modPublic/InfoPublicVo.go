package modPublic

import (
	"github.com/hongmengzhu/xianfu-blog-go/app/models/pub/model/modSysPub"
	"github.com/hongmengzhu/xianfu-blog-go/app/models/ram/modRamAccount"
)

type InfoPublicVo struct {
	Info modRamAccount.AccountPub  `json:"info"`
	Menu modSysPub.RamMenuRouterVo `json:"menuRouter" label:"菜单路由"`
}

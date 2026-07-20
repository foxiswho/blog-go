package modPublic

import (
	"github.com/hongmengzhu/xianfu-blog-go/app/manage/domainRam/model/modRamAccount"
	"github.com/hongmengzhu/xianfu-blog-go/app/system/pub/model/modSysPub"
)

type InfoPublicVo struct {
	Info modRamAccount.AccountPub  `json:"info"`
	Menu modSysPub.RamMenuRouterVo `json:"menuRouter" label:"菜单路由"`
}

package app

import (
	_ "github.com/foxiswho/blog-go/app/event/basic"
	_ "github.com/foxiswho/blog-go/app/event/ram"
	_ "github.com/foxiswho/blog-go/app/manage/domainBasic"
	_ "github.com/foxiswho/blog-go/app/manage/domainBlog"
	_ "github.com/foxiswho/blog-go/app/manage/domainRam"
	_ "github.com/foxiswho/blog-go/app/manage/domainTc"
	_ "github.com/foxiswho/blog-go/app/system/ram"
	_ "github.com/foxiswho/blog-go/app/system/tc"
)

// app 目录下 各个模块包，需要 统一 初始化

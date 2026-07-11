package middleware

import (
	_ "github.com/hongmengzhu/xianfu-blog-go/middleware/authPg"
	_ "github.com/hongmengzhu/xianfu-blog-go/middleware/cachePg/redisPg"
	_ "github.com/hongmengzhu/xianfu-blog-go/middleware/components/attachmentPg"
	_ "github.com/hongmengzhu/xianfu-blog-go/middleware/dbPg/postgresqlPg"
	_ "github.com/hongmengzhu/xianfu-blog-go/middleware/runnerPg"
)

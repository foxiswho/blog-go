package middleware

import (
	_ "github.com/foxiswho/blog-go/middleware/authPg"
	_ "github.com/foxiswho/blog-go/middleware/cachePg/redisPg"
	_ "github.com/foxiswho/blog-go/middleware/components/attachmentPg"
	_ "github.com/foxiswho/blog-go/middleware/dbPg/postgresqlPg"
	_ "github.com/foxiswho/blog-go/middleware/runnerPg"
)

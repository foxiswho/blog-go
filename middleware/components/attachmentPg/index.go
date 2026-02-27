package attachmentPg

import (
	"github.com/foxiswho/blog-go/middleware/components/attachmentPg/drive"
	"github.com/foxiswho/blog-go/middleware/components/attachmentPg/types"
	"github.com/go-spring/spring-core/gs"
)

func init() {
	gs.Object(new(drive.Local)).
		Export(gs.As[types.FileProvider]())
	// 本地存储，当 minio 不存在时才注册
	// 可以添加其它判断条件，例如 aliyun 等
	//On(cond.Group(cond.And, cond.OnMissingBean((*minio.Client)(nil))))

	//gs.Object(new(Minio)).
	//	Export((*FileProvider)(nil)).
	//	On(cond.OnBean((*minio.Client)(nil)))
}

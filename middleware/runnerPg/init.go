package runnerPg

import (
	"github.com/foxiswho/blog-go/app/event/basic/listenerBasic"
	"github.com/foxiswho/blog-go/app/event/ram/listenerRam"
	"github.com/foxiswho/blog-go/middleware/runnerPg/data"
	"github.com/foxiswho/blog-go/middleware/runnerPg/table"
	"github.com/go-spring/spring-core/gs"
)

func init() {
	//初始化 表
	gs.Object(new(table.AInitTable)).AsRunner()
	//初始化租户域名
	gs.Object(new(data.InitTenantDomain)).AsRunner()
	//初始化 基础数据
	gs.Object(new(data.IBasicData)).AsRunner()
	//

	//
	//附件 上传
	gs.Object(new(listenerBasic.AttachmentListener)).AsRunner()
	gs.Object(new(listenerBasic.TagsListener)).AsRunner()
	// ram相关
	gs.Object(new(listenerRam.RamListener)).AsRunner()
	//
	// 初始化会话密钥
	gs.Object(new(data.InitSessionPubPrive)).AsRunner()
	// 初始化标签缓存
	gs.Object(new(data.ZInitTagsCache)).AsRunner()
	//超管账号初始化
	gs.Object(new(data.ZInitAccountAdmin)).AsRunner()
	//初始化Dipl缓存
	gs.Object(new(data.ZInitDiplCache)).AsRunner()
	//显示服务启动信息
	gs.Object(new(data.ZzBootstrap)).AsRunner()
}

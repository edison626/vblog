package apps

// 业务控制器(负责导入所有的业务实现: 注册到 Ioc里面的Controller区域)
import (
	// 到包的先后顺序 就对象注册的先后顺序
	_ "gitee.com/go-course/go12/vblog/apps/token/impl"
	_ "gitee.com/go-course/go12/vblog/apps/user/impl"

	// Api Handler注册
	_ "gitee.com/go-course/go12/vblog/apps/token/api"
)

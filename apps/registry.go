package apps

// 业务控制器 （负责倒入说有的业务实现： 注册到Ioc里面的controller区域）

import (
	//倒入包就是先后顺序 就对注册的先后顺序
	_ "github.com/edison626/vblog/apps/token/impl"
	_ "github.com/edison626/vblog/apps/user/impl"

	// Api Handler注册
	_ "github.com/edison626/vblog/apps/token/api"
)

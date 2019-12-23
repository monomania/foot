package main

import (
	_ "tesou.io/platform/foot-parent/foot-core/common/fliters"
	_ "tesou.io/platform/foot-parent/foot-core/common/routers"
	"tesou.io/platform/foot-parent/foot-core/module/leisu/service"
)

func main() {
	//开启SQL输出

	pubService := new(service.PubService)
	pubService.PubBJDC()

}

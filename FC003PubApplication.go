package main

import (
	"tesou.io/platform/foot-parent/foot-core/module/leisu/service"
	_ "tesou.io/platform/foot-parent/foot-core/common/fliters"
	_ "tesou.io/platform/foot-parent/foot-core/common/routers"
)

func main() {

	pubService := new(service.PubService)
	pubService.PubBJDC()

}

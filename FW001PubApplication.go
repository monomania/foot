package main

import (
	"tesou.io/platform/foot-parent/foot-core/module/leisu/service"
	_ "tesou.io/platform/foot-parent/foot-web/common/fliters"
	_ "tesou.io/platform/foot-parent/foot-web/common/routers"
)

func main() {

	pubService := new(service.PubService)
	pubService.PubBJDC()

}

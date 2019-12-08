package main

import (
	_ "tesou.io/platform/foot-parent/foot-web/common/fliters"
	_ "tesou.io/platform/foot-parent/foot-web/common/routers"
	"tesou.io/platform/foot-parent/foot-web/module/leisu/service"
)

func main() {

	pubService := new(service.PubService)
	pubService.PubBJDC()

}

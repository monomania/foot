package main

import (
	"github.com/astaxie/beego"
	_ "tesou.io/platform/foot-parent/foot-web/common/fliters"
	_ "tesou.io/platform/foot-parent/foot-web/common/routers"
)

func main() {

	beeRun()

}

func beeRun() {
	beego.LoadAppConfig("ini", "conf/app.conf")
	beego.SetLevel(beego.LevelDebug)
	beego.SetLogger("file", `{"filename":"/home/logs/foot-web/foot-web.log"}`)
	beego.Run()
}

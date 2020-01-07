package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"os"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-core/common/routers"
)

func init() {
	router := &routers.FootRouter{}
	router.Hello()
}

func main() {
	base.Log.Info(1)
	base.Log.Info(2)
	base.Log.Info(3)
	base.Log.Info(4)
	base.Log.Info(5)
	_, err := os.Stat("abss")
	if err != nil {
		base.Log.Error(err)
	}

	beeRun()

}

func beeRun() {
	beego.LoadAppConfig("ini", "conf/app.conf")
	logs.SetLogger(logs.AdapterConsole, `{"level":1,"color":true}`)
	//logs.SetLogger(logs.AdapterFile,`{"filename":"foot-web.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
	//输出文件名和行号
	logs.EnableFuncCallDepth(true)
	//异步输出日志
	//logs.Async(1e3)
	//启动
	beego.Run()
}
package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strconv"
	"time"
)

func main() {

	beeRun()

}

func beeRun() {
	hours, _ := strconv.Atoi(time.Now().Format("15"))
	fmt.Println(time.Duration(int64(24-hours)))
	beego.LoadAppConfig("ini", "conf/app.conf")
	logs.SetLogger(logs.AdapterConsole, `{"level":1,"color":true}`)
	//logs.SetLogger(logs.AdapterFile,`{"filename":"foot-web.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
	//输出文件名和行号
	logs.EnableFuncCallDepth(true)
	//异步输出日志
	logs.Async(1e3)
	//启动
	beego.Run()
}
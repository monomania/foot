package main

import (
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	"tesou.io/platform/foot-parent/foot-core/module/analy/service"
)

func main() {
	//关闭SQL输出
	mysql.ShowSQL(false)
	base.Log.Info("---------------------------------------------------------------")
	base.Log.Info("---------------C1模型--------------")
	base.Log.Info("---------------------------------------------------------------")
	c1 := new(service.C1Service)
	c1.MaxLetBall = 1
	c1.Analy(false)
	//关闭SQL输出
	mysql.ShowSQL(true)
}


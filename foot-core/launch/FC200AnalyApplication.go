package launch

import (
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	"tesou.io/platform/foot-parent/foot-core/module/analy/service"
)

func Analy() {
	//关闭SQL输出
	maxLetBall := 1.0
	printOddData := false
	mysql.ShowSQL(false)
	base.Log.Info("---------------------------------------------------------------")
	base.Log.Info("----------------E2模型-------------------")
	base.Log.Info("---------------------------------------------------------------")
	e2 := new(service.E2Service)
	e2.MaxLetBall = maxLetBall
	e2.PrintOddData = printOddData
	e2.Analy()
	base.Log.Info("---------------------------------------------------------------")
	base.Log.Info("----------------Q1模型-------------------")
	base.Log.Info("---------------------------------------------------------------")
	q1 := new(service.Q1Service)
	q1.MaxLetBall = maxLetBall
	q1.PrintOddData = printOddData
	q1.Analy()
	base.Log.Info("---------------------------------------------------------------")
	base.Log.Info("----------------E1模型-------------------")
	base.Log.Info("---------------------------------------------------------------")
	e1 := new(service.E1Service)
	e1.MaxLetBall = maxLetBall
	e1.PrintOddData = printOddData
	e1.Analy()
	base.Log.Info("---------------------------------------------------------------")
	base.Log.Info("---------------A1模型--------------")
	base.Log.Info("---------------------------------------------------------------")
	a1 := new(service.A1Service)
	a1.MaxLetBall = maxLetBall
	a1.Analy()
	base.Log.Info("---------------------------------------------------------------")
	base.Log.Info("---------------处理结果--------------")
	base.Log.Info("---------------------------------------------------------------")
	analyService := new(service.AnalyService)
	analyService.ModifyResult()
	mysql.ShowSQL(true)
}

func Analy_Near() {
	//关闭SQL输出
	mysql.ShowSQL(true)
	base.Log.Info("---------------------------------------------------------------")
	base.Log.Info("----------------E2模型-------------------")
	base.Log.Info("---------------------------------------------------------------")
	e2 := new(service.E2Service)
	e2.MaxLetBall = 1
	e2.PrintOddData = false
	e2.Analy_Near()
	base.Log.Info("---------------------------------------------------------------")
	base.Log.Info("----------------Q1模型-------------------")
	base.Log.Info("---------------------------------------------------------------")
	q1 := new(service.Q1Service)
	q1.MaxLetBall = 1
	q1.PrintOddData = false
	q1.Analy_Near()
	base.Log.Info("---------------------------------------------------------------")
	base.Log.Info("----------------E1模型-------------------")
	base.Log.Info("---------------------------------------------------------------")
	e1 := new(service.E1Service)
	e1.MaxLetBall = 1
	e1.PrintOddData = false
	e1.Analy_Near()
	base.Log.Info("---------------------------------------------------------------")
	base.Log.Info("---------------A1模型--------------")
	base.Log.Info("---------------------------------------------------------------")
	a1 := new(service.A1Service)
	a1.MaxLetBall = 1
	a1.Analy_Near()
	mysql.ShowSQL(true)
}

package launch

import (
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	"tesou.io/platform/foot-parent/foot-core/module/analy/service"
)


func Analy() {
	//关闭SQL输出
	mysql.ShowSQL(false)
	base.Log.Info("---------------------------------------------------------------")
	base.Log.Info("----------------Euro20191212Service-------------------")
	base.Log.Info("---------------------------------------------------------------")
	euro20191212Service := new(service.Euro20191212Service)
	euro20191212Service.MaxLetBall = 1
	euro20191212Service.PrintOddData = false
	euro20191212Service.Analy()
	base.Log.Info("---------------------------------------------------------------")
	base.Log.Info("----------------Euro20191226Service-------------------")
	base.Log.Info("---------------------------------------------------------------")
	euro20191226Service := new(service.Euro20191226Service)
	euro20191226Service.MaxLetBall = 1
	euro20191226Service.PrintOddData = false
	euro20191226Service.Analy()
	base.Log.Info("---------------------------------------------------------------")
	base.Log.Info("----------------Euro20191206Service-------------------")
	base.Log.Info("---------------------------------------------------------------")
	euro20191206Service := new(service.Euro20191206Service)
	euro20191206Service.MaxLetBall = 1
	euro20191206Service.PrintOddData = false
	euro20191206Service.Analy()
	base.Log.Info("---------------------------------------------------------------")
	base.Log.Info("---------------Asia20191206Service--------------")
	base.Log.Info("---------------------------------------------------------------")
	asia20191206Service := new(service.Asia20191206Service)
	asia20191206Service.MaxLetBall = 1
	asia20191206Service.Analy()
	base.Log.Info("---------------------------------------------------------------")
	base.Log.Info("---------------处理结果--------------")
	base.Log.Info("---------------------------------------------------------------")
	analyService := new(service.AnalyService)
	analyService.ModifyResult()
	mysql.ShowSQL(true)
}

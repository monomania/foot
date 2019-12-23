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
	base.Log.Info("----------------Euro20191206Service-------------------")
	base.Log.Info("---------------------------------------------------------------")
	euroService := new(service.Euro20191206Service)
	euroService.MaxLetBall = 1
	euroService.PrintOddData = false
	euroService.Analy()

	base.Log.Info("---------------------------------------------------------------")
	base.Log.Info("---------------处理结果--------------")
	base.Log.Info("---------------------------------------------------------------")
	analyService := new(service.AnalyService)
	analyService.ModifyResult()
	mysql.ShowSQL(true)
}

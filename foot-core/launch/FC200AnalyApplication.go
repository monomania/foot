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
	euro616_104Service := new(service.Euro20191212Service)
	euro616_104Service.MaxLetBall = 0.5
	euro616_104Service.PrintOddData = false
	euro616_104Service.Analy()
	base.Log.Info("---------------------------------------------------------------")
	base.Log.Info("----------------Euro20191206Service-------------------")
	base.Log.Info("---------------------------------------------------------------")
	euro_81_616Service := new(service.Euro20191206Service)
	euro_81_616Service.MaxLetBall = 0.75
	euro_81_616Service.PrintOddData = false
	euro_81_616Service.Analy()
	base.Log.Info("---------------------------------------------------------------")
	base.Log.Info("---------------Asia20191206Service--------------")
	base.Log.Info("---------------------------------------------------------------")
	euro81616Service := new(service.Asia20191206Service)
	euro81616Service.Analy()
	base.Log.Info("---------------------------------------------------------------")
	base.Log.Info("---------------处理结果--------------")
	base.Log.Info("---------------------------------------------------------------")
	analyService := new(service.AnalyService)
	analyService.ModifyResult()
	mysql.ShowSQL(true)
}

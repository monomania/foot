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
	base.Log.Info("----------------计算欧AnalyEuro_81_616Service-------------------")
	base.Log.Info("---------------------------------------------------------------")
	euro_81_616Service := new(service.Euro81_616Service)
	euro_81_616Service.MaxLetBall = 0.75
	euro_81_616Service.PrintOddData = false
	euro_81_616Service.Analy()
	base.Log.Info("---------------------------------------------------------------")
	base.Log.Info("---------------计算AnalyAsia_18_Euro_81_616Service--------------")
	base.Log.Info("---------------------------------------------------------------")
	euro81616Service := new(service.Asia18EuroUDReverseService)
	euro81616Service.Analy()
}

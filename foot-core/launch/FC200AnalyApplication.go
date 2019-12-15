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
	base.Log.Info("----------------Euro616_104Service-------------------")
	base.Log.Info("---------------------------------------------------------------")
	euro616_104Service := new(service.Euro616_104Service)
	euro616_104Service.MaxLetBall = 0.5
	euro616_104Service.PrintOddData = false
	euro616_104Service.Analy()
	base.Log.Info("---------------------------------------------------------------")
	base.Log.Info("----------------Euro81_616_104Service-------------------")
	base.Log.Info("---------------------------------------------------------------")
	euro_81_616Service := new(service.Euro81_616_104Service)
	euro_81_616Service.MaxLetBall = 0.75
	euro_81_616Service.PrintOddData = false
	euro_81_616Service.Analy()
	base.Log.Info("---------------------------------------------------------------")
	base.Log.Info("---------------Asia18EuroUDReverseService--------------")
	base.Log.Info("---------------------------------------------------------------")
	euro81616Service := new(service.Asia18EuroUDReverseService)
	euro81616Service.Analy()
}

package launch

import (
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	"tesou.io/platform/foot-parent/foot-core/module/analy/service"
)

func init(){
	//关闭SQL输出
	mysql.ShowSQL(false)
}


func Analy() []interface{} {
	analysisService := new(service.AnalyService)
	analysisService.MaxLetBall = 0.75
	analysisService.PrintOddData = false
	base.Log.Info("-----------------------------------------------")
	base.Log.Info("----------------计算欧86之差-------------------")
	base.Log.Info("-----------------------------------------------")
	r1 := analysisService.Euro_Calc()
	base.Log.Info("-----------------------------------------------")
	base.Log.Info("---------------计算亚欧之差--------------------")
	base.Log.Info("-----------------------------------------------")
	r2 := analysisService.Euro_Asia_Diff()

	i := append(r1, r2)
	return i

}

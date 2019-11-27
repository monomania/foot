package main

import (
	"tesou.io/platform/foot-parent/foot-core/module/core/service"
)

func main() {
	//生成数据库表
	GenTable()
	//gui()
	//清空数据表
	TruncateTable()
	//运行分析
	//analysisService := new(service2.AnalyService)
	//analysisService.AnalyToday()

}

func TruncateTable() {
	opsService := new(service.DBOpsService)
	//指定需要清空的数据表
	opsService.TruncateTable([]string{"t_match_last", "t_match_last_config", "t_euro_last", "t_euro_his", "t_asia_last"})
}

func GenTable() {
	generateService := new(service.DBOpsService)
	generateService.SyncTableStruct()
}
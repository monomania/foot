package launch

import (
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)


func init(){
	//开启SQL输出
	mysql.ShowSQL(true)
}

func TruncateTable() {
	opsService := new(mysql.DBOpsService)
	//指定需要清空的数据表
	opsService.TruncateTable([]string{"t_match_last", "t_match_last_config", "t_euro_last", "t_euro_his", "t_asia_last"})
}

func GenTable() {
	generateService := new(mysql.DBOpsService)
	generateService.SyncTableStruct()
}
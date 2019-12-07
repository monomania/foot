package main

import (
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	"tesou.io/platform/foot-parent/foot-spider/launch"
)

func init() {
	//开启SQL输出
	mysql.ShowSQL(true)
}

//抓取欧赔数据少于两条的不完整的比赛
func main() {
	launch.Spider_euroHis_Incomplete(2)
}


package main

import (
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	"tesou.io/platform/foot-parent/foot-spider/launch"
)


func main() {
	//开启SQL输出
	mysql.ShowSQL(true)
	launch.Spider_euroLast()
	launch.Spider_euroHis()
}

package main

import (
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	"tesou.io/platform/foot-parent/foot-spider/launch"
)

func init(){
	//开启SQL输出
	mysql.ShowSQL(true)
}


func main() {
	launch.Spider_euroLast()
	launch.Spider_euroHis()
}

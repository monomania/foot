package main

import (
	"tesou.io/platform/foot-parent/foot-spider/launch"
)


//抓取欧赔数据少于两条的不完整的比赛
func main() {
	//开启SQL输出

	launch.Spider_euroHis_Incomplete()
}


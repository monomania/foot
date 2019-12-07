package main

import (
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	"tesou.io/platform/foot-parent/foot-spider/launch"
)

func init() {
	//开启SQL输出
	mysql.ShowSQL(true)
}

func main() {
	Spider(4)
}

func Spider(matchLevel int) {
	launch.Before_spider_match()
	launch.Before_spider_asiaLast()
	launch.Before_spider_euroLast()
	launch.Before_spider_euroHis()
	//执行抓取比赛数据
	//执行抓取比赛欧赔数据
	//执行抓取亚赔数据
	//执行抓取欧赔历史
	launch.Spider_match(matchLevel)
	//launch.Spider_asiaLast()
	launch.Spider_asiaLastNew()
	launch.Spider_euroLast()
	launch.Spider_euroHis()
	//再对欧赔数据不完整的比赛进行两次抓取
	launch.Spider_euroHis_Incomplete(2)
	launch.Spider_euroHis_Incomplete(2)
}

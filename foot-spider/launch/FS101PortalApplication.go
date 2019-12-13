package launch

import (
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	"tesou.io/platform/foot-parent/foot-core/launch"
)

func Spider(matchLevel int) {
	//开启SQL输出
	mysql.ShowSQL(true)
	//清空数据表
	launch.TruncateTable()
	Before_spider_match()
	Before_spider_asiaLast()
	Before_spider_euroLast()
	Before_spider_euroHis()
	//执行抓取比赛数据
	//执行抓取比赛欧赔数据
	//执行抓取亚赔数据
	//执行抓取欧赔历史
	Spider_match(matchLevel)
	//Spider_asiaLast()
	Spider_asiaLastNew()
	Spider_euroLast()
	Spider_euroHis()
	//再对欧赔数据不完整的比赛进行两次抓取
	Spider_euroHis_Incomplete(3)
	Spider_euroHis_Incomplete(3)
}
package launch

import (
	"strconv"
	"tesou.io/platform/foot-parent/foot-core/common/utils"
	"tesou.io/platform/foot-parent/foot-core/launch"
	"tesou.io/platform/foot-parent/foot-core/module/spider/constants"
	"time"
)

func Clean(){
	//清空数据表
	launch.TruncateTable()
	Before_spider_match()
	Before_spider_asiaLast()
	Before_spider_euroLast()
}


func Spider() {
	//执行抓取比赛数据
	//执行抓取比赛欧赔数据
	//执行抓取亚赔数据
	//执行抓取欧赔历史
	matchLevelStr := utils.GetVal("spider", "match_level")
	if len(matchLevelStr) <= 0 {
		matchLevelStr = "4"
	}
	matchLevel, _ := strconv.Atoi(matchLevelStr)
	Spider_match(matchLevel)
	Spider_baseFace()
	//Spider_asiaLast()
	Spider_asiaLastNew()
	Spider_euroLast()
	Spider_euroHis()
	//再对欧赔数据不完整的比赛进行两次抓取
	Spider_euroHis_Incomplete()
	//记录数据爬取时间
	constants.SpiderDateStr = time.Now().Format("2006-01-02 15:04:05")
}


func Spider_Near() {
	Spider_baseFace_near()
	Spider_asiaLastNew_near()
	Spider_euroLast_near()
	Spider_euroHis_near()
	//记录数据爬取时间
	constants.SpiderDateStr = time.Now().Format("2006-01-02 15:04:05")
}
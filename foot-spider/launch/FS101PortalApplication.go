package launch

import (
	"strconv"
	"tesou.io/platform/foot-parent/foot-core/common/utils"
	"tesou.io/platform/foot-parent/foot-core/module/spider/constants"
	"time"
)

func Clean(){
	//清空数据表
	//Before_spider_match()
	//Before_spider_baseFace()
	Before_spider_asiaLast()
	Before_spider_euroLast()
}


func Spider() {
	//记录数据爬取时间
	constants.SpiderDateStr = time.Now().Format("2006-01-02 15:04:05")
	constants.FullSpiderDateStr = constants.SpiderDateStr
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
	Spider_asiaLastNew(false)
	//Spider_euroLast()
	//Spider_euroHis()
	//再对欧赔数据不完整的比赛进行两次抓取
	//Spider_euroHis_Incomplete()
}


func Spider_Near() {
	//记录数据爬取时间
	constants.SpiderDateStr = time.Now().Format("2006-01-02 15:04:05")

	matchLevelStr := utils.GetVal("spider", "match_level")
	if len(matchLevelStr) <= 0 {
		matchLevelStr = "4"
	}
	matchLevel, _ := strconv.Atoi(matchLevelStr)
	Spider_match(matchLevel)
	//基本面不会改变
	//Spider_baseFace_near()
	Spider_asiaLastNew_near()
	Spider_euroLast_near()
	Spider_euroHis_near()
}
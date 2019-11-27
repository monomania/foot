package launch

import (
	"tesou.io/platform/foot-parent/foot-core/module/match/entity"
	"tesou.io/platform/foot-parent/foot-spider/module/win007"
	"tesou.io/platform/foot-parent/foot-spider/module/win007/proc"
)

/*func main() {
	//执行抓取比赛数据
	spider_match_analy()
}*/


//抓取比赛分析数据
func spider_match_analy() {
	matchLastConfig := new(entity.MatchLastConfig)
	matchLastConfig.S = win007.MODULE_FLAG
	matchLastConfig.EuroSpided = false
	matchLastConfigs := matchLastConfig.Query()

	processer := proc.GetMatchAnalyProcesser()
	processer.MatchLastConfig_list = matchLastConfigs
	processer.Startup()
}
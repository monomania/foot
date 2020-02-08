package launch

import (
	"tesou.io/platform/foot-parent/foot-core/module/match/service"
	"tesou.io/platform/foot-parent/foot-spider/module/win007/proc"
)



//抓取比赛分析数据
func spider_match_analy() {
	matchLastService := new(service.MatchLastService)
	matchLasts := matchLastService.FindNotFinished()

	processer := proc.GetMatchAnalyProcesser()
	processer.MatchLastList = matchLasts
	processer.Startup()
}
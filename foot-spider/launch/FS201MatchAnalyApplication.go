package launch

import (
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	"tesou.io/platform/foot-parent/foot-core/module/match/service"
	"tesou.io/platform/foot-parent/foot-spider/module/win007"
	"tesou.io/platform/foot-parent/foot-spider/module/win007/proc"
)

/*func main() {
	//执行抓取比赛数据
	spider_match_analy()
}*/


//抓取比赛分析数据
func spider_match_analy() {
	matchLastConfigService := new(service.MatchLastConfigService)
	config := &pojo.MatchLastConfig{}
	config.S = win007.MODULE_FLAG
	config.EOSpider = false
	matchLastConfigs := matchLastConfigService.Query(config)

	processer := proc.GetMatchAnalyProcesser()
	processer.MatchLastConfig_list = matchLastConfigs
	processer.Startup()
}
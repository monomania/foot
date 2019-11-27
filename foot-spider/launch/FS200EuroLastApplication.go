package launch

import (
	"tesou.io/platform/foot-parent/foot-core/module/core/service"
	"tesou.io/platform/foot-parent/foot-core/module/match/entity"
	"tesou.io/platform/foot-parent/foot-spider/module/win007"
	"tesou.io/platform/foot-parent/foot-spider/module/win007/proc"
)

/*func main() {
	//执行抓取比赛欧赔数据
	Before_spider_euroLast()
	Spider_euroLast()
}*/

func Before_spider_euroLast(){
	//抓取前清空当前比较表
	opsService := new(service.DBOpsService)
	//指定需要清空的数据表
	opsService.TruncateTable([]string{"t_euro_last", "t_comp", "t_comp_config"})
}

//查询标识为win007,且欧赔未抓取的配置数据,指定菠菜公司
func Spider_euroLast() {
	matchLastConfig := new(entity.MatchLastConfig)
	matchLastConfig.S = win007.MODULE_FLAG
	matchLastConfig.EuroSpided = false
	matchLastConfigs := matchLastConfig.Query()
	//281 -- bet 365  18 -- 12BET 976 -- 18Bet 81 -- 伟德 616 -- 888Sport
	betCompWin007Ids := []string{"81","616"}
	//为空会抓取所有,这里没有必要配置所有的波菜公司ID
	//betCompWin007Ids := new(entity2.Comp).FindAllIds()

	processer := proc.GetEuroLastProcesser()
	processer.MatchLastConfig_list = matchLastConfigs
	processer.BetCompWin007Ids = betCompWin007Ids
	processer.Startup()
}

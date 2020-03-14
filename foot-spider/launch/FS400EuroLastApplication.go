package launch

import (
	"strings"
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	"tesou.io/platform/foot-parent/foot-core/common/utils"
	"tesou.io/platform/foot-parent/foot-core/module/elem/service"
	service2 "tesou.io/platform/foot-parent/foot-core/module/match/service"
	"tesou.io/platform/foot-parent/foot-spider/module/win007/proc"
)

/*func main() {
	//执行抓取比赛欧赔数据
	Before_spider_euroLast()
	Spider_euroLast()
}*/

func Before_spider_euroLast() {
	//抓取前清空当前比较表
	opsService := new(mysql.DBOpsService)
	//指定需要清空的数据表
	opsService.TruncateTable([]string{"t_euro_last"})
}

//查询标识为win007,且欧赔未抓取的配置数据,指定菠菜公司
func Spider_euroLast() {
	matchLastService := new(service2.MatchLastService)
	matchLasts := matchLastService.FindNotFinished()

	var compIds []string
	val := utils.GetVal("spider", "euro_comp_ids")
	if len(val) < 0 {
		//为空会抓取所有,这里没有必要配置所有的波菜公司ID
		compService := new(service.CompService)
		compIds = compService.FindEuroIds()
	} else {
		compIds = strings.Split(val, ",")
	}

	processer := proc.GetEuroLastProcesser()
	processer.MatchLastList = matchLasts
	processer.CompWin007Ids = compIds
	processer.SingleThread = true
	processer.Startup()
}

func Spider_euroLast_his(season string) {
	matchLastService := new(service2.MatchHisService)
	var matchLasts []*pojo.MatchLast
	matchLasts = matchLastService.FindBySeason(season)

	var compIds []string
	//为空会抓取所有,这里没有必要配置所有的波菜公司ID
	compService := new(service.CompService)
	compIds = compService.FindEuroIds()

	processer := proc.GetEuroLastProcesser()
	processer.MatchLastList = matchLasts
	processer.CompWin007Ids = compIds
	processer.Startup()
}

//查询标识为win007,且欧赔未抓取的配置数据,指定菠菜公司
func Spider_euroLast_near() {
	matchLastService := new(service2.MatchLastService)
	matchLasts := matchLastService.FindNear()
	if len(matchLasts) <= 0 {
		return
	}

	var compIds []string
	val := utils.GetVal("spider", "euro_comp_ids")
	if len(val) < 0 {
		//为空会抓取所有,这里没有必要配置所有的波菜公司ID
		compService := new(service.CompService)
		compIds = compService.FindEuroIds()
	} else {
		compIds = strings.Split(val, ",")
	}

	processer := proc.GetEuroLastProcesser()
	processer.MatchLastList = matchLasts
	processer.CompWin007Ids = compIds
	processer.Startup()
}

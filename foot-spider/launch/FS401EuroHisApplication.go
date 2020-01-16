package launch

import (
	"strings"
	"tesou.io/platform/foot-parent/foot-core/common/utils"
	"tesou.io/platform/foot-parent/foot-core/module/elem/service"
	service2 "tesou.io/platform/foot-parent/foot-core/module/match/service"
	"tesou.io/platform/foot-parent/foot-spider/module/win007/proc"
)

/*func main() {
	//遍历十次
	count := 10
	for i := 0; i < count; i++ {
		spider_euroHis()
	}
}
*/
//func Before_spider_euroHis() {
//	//抓取前清空当前比较表
//	opsService := new(mysql.DBOpsService)
//	//指定需要清空的数据表
//	opsService.TruncateTable([]string{"t_euro_his"})
//}

//查询标识为win007,且欧赔未抓取的配置数据,指定菠菜公司
func Spider_euroHis() {
	matchLastService := new(service2.MatchLastService)
	matchLasts := matchLastService.FindNotFinished()

	var compIds []string
	val := utils.GetVal("spider", "euro_comp_ids")
	if len(val) < 0 {
		//为空会抓取所有,这里没有必要配置所有的波菜公司ID
		compService := new(service.CompService)
		compIds = compService.FindAllIds()
	}else{
		compIds = strings.Split(val, ",")
	}

	processer := proc.GetEuroTrackProcesser()
	processer.CompWin007Ids = compIds
	processer.MatchLastList = matchLasts
	processer.Startup()

}


//查询标识为win007,且欧赔未抓取的配置数据,指定菠菜公司
func Spider_euroHis_near() {
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
		compIds = compService.FindAllIds()
	}else{
		compIds = strings.Split(val, ",")
	}

	processer := proc.GetEuroTrackProcesser()
	processer.CompWin007Ids = compIds
	processer.MatchLastList = matchLasts
	processer.Startup()

}

func Spider_euroHis_Incomplete() {
	var compIds []string
	val := utils.GetVal("spider", "euro_comp_ids")
	if len(val) < 0 {
		//为空会抓取所有,这里没有必要配置所有的波菜公司ID
		compService := new(service.CompService)
		compIds = compService.FindAllIds()
	}else{
		compIds = strings.Split(val, ",")
	}

	matchLastService := new(service2.MatchLastService)
	matchLasts := matchLastService.FindEuroIncomplete(len(compIds))

	processer := proc.GetEuroTrackProcesser()
	processer.CompWin007Ids = compIds
	processer.MatchLastList = matchLasts
	processer.Startup()

}

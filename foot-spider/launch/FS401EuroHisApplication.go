package launch

import (
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
	matchLasts := matchLastService.FindAll()

	//设置要抓取的波菜公司id
	betCompWin007Ids := []string{"151","1129"}
	//betCompWin007Ids := new(entity2.Comp).FindAllIds()

	processer := proc.GetEuroHisProcesser()
	processer.CompWin007Ids = betCompWin007Ids
	processer.MatchLastList = matchLasts
	processer.Startup()

}

func Spider_euroHis_Incomplete(count int) {
	matchLastService := new(service2.MatchLastService)
	matchLasts := matchLastService.FindEuroIncomplete(count)

	//设置要抓取的波菜公司id
	compWin007Ids := []string{"115", "1129", "432"}

	processer := proc.GetEuroHisProcesser()
	processer.CompWin007Ids = compWin007Ids
	processer.MatchLastList = matchLasts
	processer.Startup()

}

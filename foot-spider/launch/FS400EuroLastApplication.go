package launch

import (
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	service2 "tesou.io/platform/foot-parent/foot-core/module/match/service"
	"tesou.io/platform/foot-parent/foot-spider/module/win007/proc"
)

func Before_spider_euroLast() {
	//抓取前清空当前比较表
	opsService := new(mysql.DBOpsService)
	//指定需要清空的数据表
	opsService.TruncateTable([]string{"t_euro_last"})
}

//查询标识为win007,且欧赔未抓取的配置数据,指定菠菜公司
func Spider_euroLast() {
	matchLastService := new(service2.MatchLastService)
	matchLasts := matchLastService.FindAll()
	betCompWin007Ids := []string{"115", "1129", "432","281"}
	processer := proc.GetEuroLastProcesser()
	processer.MatchLastList = matchLasts
	processer.CompWin007Ids = betCompWin007Ids
	processer.Startup()
}

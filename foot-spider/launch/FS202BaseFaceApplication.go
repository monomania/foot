package launch

import (
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	service2 "tesou.io/platform/foot-parent/foot-core/module/match/service"
	"tesou.io/platform/foot-parent/foot-spider/module/win007/proc"
)


func Before_spider_baseFace(){
	//抓取前清空当前比较表
	opsService := new(mysql.DBOpsService)
	//指定需要清空的数据表
	opsService.TruncateTable([]string{"t_b_f_battle","t_b_f_future_event","t_b_f_score"})
}

//查询标识为win007,且欧赔未抓取的配置数据,指定菠菜公司
func Spider_baseFace(spiderAll bool) {
	matchLastService := new(service2.MatchLastService)
	var matchLasts []*pojo.MatchLast
	if spiderAll{
		matchLasts = matchLastService.FindAll()
	}else{
		matchLasts = matchLastService.FindNotFinished()
	}

	processer := proc.GetBaseFaceProcesser()
	processer.MatchLastList = matchLasts
	processer.Startup()
}



func Spider_baseFace_near() {
	matchLastService := new(service2.MatchLastService)
	matchLasts := matchLastService.FindNear()

	processer := proc.GetBaseFaceProcesser()
	processer.MatchLastList = matchLasts
	processer.Startup()
}


package launch

import (
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	"tesou.io/platform/foot-parent/foot-spider/module/win007/proc"
)

func Before_spider_match() {
	//抓取前清空当前比较表
	opsService := new(mysql.DBOpsService)
	//指定需要清空的数据表
	opsService.TruncateTable([]string{"t_match_last"})
}

//抓取比赛数据
func Spider_match(flag int) {
	//开始抓取比赛数据
	strings := make([]string, 0)

	strings = append(strings, "http://m.win007.com/phone/Schedule_0_4.txt")
	for _, v := range strings {
		processer := proc.GetMatchPageProcesser()
		processer.MatchlastUrl = v
		processer.Startup()
	}
}

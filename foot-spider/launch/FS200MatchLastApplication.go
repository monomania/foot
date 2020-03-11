package launch

import (
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	"tesou.io/platform/foot-parent/foot-spider/module/win007/proc"
)

/*func main() {
	//执行抓取比赛数据
	Before_spider_match()
	Spider_match()
}*/

func Before_spider_match(){
	//抓取前清空当前比较表
	opsService := new(mysql.DBOpsService)
	//指定需要清空的数据表
	opsService.TruncateTable([]string{"t_match_last"})
}

//抓取比赛数据
func Spider_match(matchLevel int) {
	//开始抓取比赛数据
	strings := make([]string, 0)
	//0全部，1一级，2足彩，3竞猜，4单场
	if matchLevel == 0{
	 strings = append(strings, "http://m.win007.com/phone/Schedule_0_0.txt")
	}
	//1一级
	//strings = append(strings, "http://m.win007.com/phone/Schedule_0_1.txt")
	//2足彩
	//strings = append(strings, "http://m.win007.com/phone/Schedule_0_2.txt")
	//3竞猜
	//strings = append(strings, "http://m.win007.com/phone/Schedule_0_3.txt")
	//4单场
	if matchLevel == 4 {
		strings = append(strings, "http://m.win007.com/phone/Schedule_0_4.txt")
	}
	for _, v := range strings {
		processer := proc.GetMatchLastProcesser()
		processer.MatchlastUrl = v
		processer.MatchLevel = matchLevel
		processer.Startup()
	}
}


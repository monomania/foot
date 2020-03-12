package launch

import (
	"tesou.io/platform/foot-parent/foot-spider/module/win007/proc"
)

//抓取比赛数据
func Spider_match_his(matchLevel int) {

	processer := proc.GetMatchHisProcesser()
	//                                  /联赛时间/联赛id_联赛子id_第几轮.htm
	//http://m.win007.com/info/fixture/2019-2020/36_0_1.htm
	processer.MatchlastUrl = "http://m.win007.com/info/fixture/2019-2020/36_0_1.htm"
	processer.MatchLevel = matchLevel
	processer.Startup()
}


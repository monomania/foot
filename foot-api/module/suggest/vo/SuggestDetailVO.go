package vo

import (
	pojo3 "tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	pojo2 "tesou.io/platform/foot-parent/foot-api/module/odds/pojo"
	"tesou.io/platform/foot-parent/foot-api/module/suggest/pojo"
)

type SuggStubDetailVO struct {
	MatchDateStr string
	//联赛
	LeagueName string
	//主队
	MainTeam string
	//客队
	GuestTeam string
	//主队进球
	MainTeamGoal string
	//客队进球
	GuestTeamGoal string
	//指数
	HitCount int
	//开始时间
	BeginDateStr string
	//结束时间
	EndDateStr string
	//是否倒序
	IsDesc bool
	//算法标识
	AlFlags []string

	pojo.SuggStub `xorm:"extends"`

	//欧赔
	EuroOdd *pojo2.EuroLast
	//亚赔
	AsiaOdd *pojo2.AsiaLast

	//主队积分排名
	//客队积分排名
	BFSMainZong  *pojo3.BFScore
	BFSMainZhu   *pojo3.BFScore
	BFSMainJin   *pojo3.BFScore
	BFSGuestZong *pojo3.BFScore
	BFSGuestKe   *pojo3.BFScore
	BFSGuestJin  *pojo3.BFScore

	//过往战绩对阵
	BattleCount         int
	BattleMainWinCount  int
	BattleDrawCount  int
	BattleGuestWinCount int

	//下一场比赛
	MainNextMainTeam  string
	GuestNextMainTeam string
}

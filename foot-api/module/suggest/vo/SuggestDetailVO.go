package vo

import (
	pojo3 "tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	pojo2 "tesou.io/platform/foot-parent/foot-api/module/odds/pojo"
	"tesou.io/platform/foot-parent/foot-api/module/suggest/pojo"
)

type SuggestDetailVO struct {
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

	pojo.Suggest `xorm:"extends"`

	//欧赔
	EuroOdd *pojo2.EuroLast
	//亚赔
	AsiaOdd *pojo2.AsiaLast

	//主队积分排名
	BFSMainZong *pojo3.BFScore
	BFSMainZhu  *pojo3.BFScore
	BFSMainJin  *pojo3.BFScore

	//客队积分排名
	BFSGuestZong *pojo3.BFScore
	BFSGuestKe   *pojo3.BFScore
	BFSGuestJin  *pojo3.BFScore

	//基本面
	BFBList  []*pojo3.BFBattle

	BFFEList []*pojo3.BFFutureEvent
}

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
	EOddList []*pojo2.EuroLast
	//亚赔
	AOddList []*pojo2.AsiaLast

	//基本面
	BFSList []*pojo3.BFScore
	BFBList []*pojo3.BFBattle
	BFFEList []*pojo3.BFFutureEvent



}

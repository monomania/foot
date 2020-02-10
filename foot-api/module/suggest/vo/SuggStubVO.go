package vo

import (
	"tesou.io/platform/foot-parent/foot-api/module/suggest/pojo"
)

type SuggStubVO struct {
	//比赛时间
	MatchDateStr string
	//联赛
	LeagueId string
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
}

package vo

import (
	"tesou.io/platform/foot-parent/foot-api/module/suggest/pojo"
)

type SuggestVO struct {
	MatchDateStr string
	//联赛
	LeagueName string
	//主队
	MainTeam string
	//主球
	LetBall float64
	//客队
	GuestTeam string
	//主队进球
	MainTeamGoal string
	//客队进球
	GuestTeamGoal string

	//开始时间
	BeginDateStr string
	//结束时间
	EndDateStr string

	pojo.Suggest `xorm:"extends"`
}

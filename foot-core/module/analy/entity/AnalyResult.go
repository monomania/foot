package entity

import "tesou.io/platform/foot-parent/foot-core/common/base/entity"

type AnalyResult struct {
	//比赛id
	MatchId   string
	MatchDate string
	//结果标识
	PreResult string
	Result    string

	/**
	 * 联赛Id
	 */
	LeagueId string
	/**
	 * 主队id
	 */
	MainTeamId string
	/**
	 * 主队进球数
	 */
	MainTeamGoals int
	/**
	 * 客队id
	 */
	GuestTeamId string
	/**
	 * 客队进球数
	 */
	GuestTeamGoals int

	//算法标识
	Al_Flag string
	//结果
	Context string

	entity.Base `xorm:"extends"`
}

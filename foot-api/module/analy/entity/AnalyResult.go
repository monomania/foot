package entity

import "tesou.io/platform/foot-parent/foot-api/common/base/entity"

type AnalyResult struct {
	//比赛id
	MatchId   string `xorm:" comment('比赛id') index"`
	MatchDate string `xorm:" comment('比赛时间') index"`
	//结果标识
	PreResult string `xorm:" comment('预测结果') index"`
	Result    string `xorm:" comment('实际结果') index"`

	/**
	 * 联赛Id
	 */
	LeagueId string `xorm:" comment('联赛Id') index"`
	/**
	 * 主队id
	 */
	MainTeamId string `xorm:" comment('主队id') index"`
	/**
	 * 主队进球数
	 */
	MainTeamGoals int `xorm:" comment('主队进球数') index"`
	/**
	 * 客队id
	 */
	GuestTeamId string `xorm:" comment('客队id') index"`
	/**
	 * 客队进球数
	 */
	GuestTeamGoals int `xorm:" comment('客队进球数') index"`

	//算法标识
	AlFlag string `xorm:" comment('算法标识') index"`
	//结果
	Context string

	entity.Base `xorm:"extends"`
}

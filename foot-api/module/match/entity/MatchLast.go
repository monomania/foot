package entity

import (
	"tesou.io/platform/foot-parent/foot-api/common/base/entity"
)

//足球比赛信息
type MatchLast struct {
	/**
	 * 比赛时间
	 */
	MatchDate string `xorm:"unique(MatchDate_MainTeamId_GuestTeamId)"`

	/**
	数据时间
	*/
	DataDate string	`xorm:" comment('数据时间') index"`
	/**
	 * 联赛Id
	 */
	LeagueId string	`xorm:" comment('联赛Id') index"`
	/**
	 * 主队id
	 */
	MainTeamId string `xorm:"unique(MatchDate_MainTeamId_GuestTeamId)"`
	/**
	 * 主队进球数
	 */
	MainTeamGoals int	`xorm:" comment('主队进球数') index"`
	/**
	 * 客队id
	 */
	GuestTeamId string `xorm:"unique(MatchDate_MainTeamId_GuestTeamId)"`
	/**
	 * 客队进球数
	 */
	GuestTeamGoals int	`xorm:" comment('客队进球数') index"`

	entity.Base `xorm:"extends"`
}


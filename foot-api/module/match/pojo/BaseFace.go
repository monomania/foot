package pojo

import (
	"tesou.io/platform/foot-parent/foot-api/common/base/pojo"
)

/**
基本面
 */
type BaseFace struct {
	//比赛id
	MatchId string `xorm:"comment('比赛ID')  unique(MatchId)"`

	/**
	数据时间
	*/
	DataDate string	`xorm:" comment('数据时间') index"`
	/**
	 * 联赛Id
	 */
	LeagueId string	`xorm:" comment('联赛Id') index"`
	/**
	 * 主队id,目前为主队名称
	 */
	MainTeamId string `xorm:"unique(MatchDate_MainTeamId_GuestTeamId)"`
	/**
	 * 主队进球数
	 */
	MainTeamGoals int	`xorm:" comment('主队进球数') index"`
	/**
	 * 客队id,目前为客队名称
	 */
	GuestTeamId string `xorm:"unique(MatchDate_MainTeamId_GuestTeamId)"`
	/**
	 * 客队进球数
	 */
	GuestTeamGoals int	`xorm:" comment('客队进球数') index"`

	pojo.BasePojo `xorm:"extends"`
}


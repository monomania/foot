package pojo

import (
	"tesou.io/platform/foot-parent/foot-api/common/base/pojo"
	"time"
)

/**
比赛当前数据表,仅存放当前的比赛
 */
//足球比赛信息
type MatchLast struct {
	/**
	 * 比赛时间
	 */
	MatchDate time.Time `xorm:"unique(MatchDate_MainTeamId_GuestTeamId) index"`

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
	MainTeamId string `xorm:"unique(MatchDate_MainTeamId_GuestTeamId) index"`
	/**
	 * 主队进球数
	 */
	MainTeamGoals int	`xorm:" comment('主队进球数') index"`
	/**
	 * 客队id,目前为客队名称
	 */
	GuestTeamId string `xorm:"unique(MatchDate_MainTeamId_GuestTeamId) index"`
	/**
	 * 客队进球数
	 */
	GuestTeamGoals int	`xorm:" comment('客队进球数') index"`

	pojo.BasePojo `xorm:"extends"`
}


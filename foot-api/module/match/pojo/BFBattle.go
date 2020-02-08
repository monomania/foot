package pojo

import (
	"tesou.io/platform/foot-parent/foot-api/common/base/pojo"
	"time"
)

/**
历史对战
*/
type BFBattle struct {
	//比赛id
	MatchId string `xorm:"comment('比赛ID') index"`

	/**
	 * 比赛时间
	 */
	BattleMatchDate time.Time `xorm:"unique(BattleMatchDate_MainTeamId_GuestTeamId) index"`
	/**
	 * 联赛Id
	 */
	BattleLeagueId string	`xorm:" comment('联赛Id') index"`
	/**
	 * 主队id,目前为主队名称
	 */
	BattleMainTeamId string `xorm:"unique(BattleMatchDate_MainTeamId_GuestTeamId) index"`
	/**
	 * 主队进球数
	 */
	BattleMainTeamHalfGoals int	`xorm:" comment('主队半场进球数') index"`
	BattleMainTeamGoals int	`xorm:" comment('主队进球数') index"`
	/**
	 * 客队id,目前为客队名称
	 */
	BattleGuestTeamId string `xorm:"unique(BattleMatchDate_MainTeamId_GuestTeamId) index"`
	/**
	 * 客队进球数
	 */
	BattleGuestTeamHalfGoals int	`xorm:" comment('客队半场进球数') index"`
	BattleGuestTeamGoals int	`xorm:" comment('客队进球数') index"`



	pojo.BasePojo `xorm:"extends"`
}

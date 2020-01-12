package pojo

import (
	"tesou.io/platform/foot-parent/foot-api/common/base/pojo"
	"time"
)

/**
未来赛事
*/
type BFFutureEvent struct {
	//比赛id
	MatchId string `xorm:"comment('比赛ID') index"`

	/**
	 * 比赛时间
	 */
	EventMatchDate time.Time `xorm:"unique(EventMatchDate_MainTeamId_GuestTeamId)"`
	/**
	 * 联赛Id
	 */
	EventLeagueId string `xorm:" comment('联赛Id') index"`
	/**
	 * 主队id,目前为主队名称
	 */
	EventMainTeamId string `xorm:"unique(EventMatchDate_MainTeamId_GuestTeamId)"`
	/**
	 * 客队id,目前为客队名称
	 */
	EventGuestTeamId string `xorm:"unique(EventMatchDate_MainTeamId_GuestTeamId)"`
	/**
	间隔天数
	*/
	IntervalDay int `xorm:" comment('间隔天数') index"`


	pojo.BasePojo `xorm:"extends"`
}

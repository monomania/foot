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
	MatchId string `xorm:"comment('比赛ID')  unique(MatchId_TeamId_EventMatchDate) index"`
	/**
	 * 队伍名称
	 */
	TeamId string `xorm:" comment('队伍名称') unique(MatchId_TeamId_EventMatchDate) index"`
	/**
	 * 比赛时间
	 */
	EventMatchDate time.Time `xorm:" comment('比赛时间') unique(MatchId_TeamId_EventMatchDate) index"`
	/**
	 * 联赛Id
	 */
	EventLeagueId string `xorm:" comment('联赛Id') index"`
	/**
	 * 主队id,目前为主队名称
	 */
	EventMainTeamId string `xorm:"  comment('主队id,目前为主队名称') index "`
	/**
	 * 客队id,目前为客队名称
	 */
	EventGuestTeamId string `xorm:" comment('客队id,目前为客队名称') index "`
	/**
	间隔天数
	*/
	IntervalDay int `xorm:" comment('间隔天数') index"`

	pojo.BasePojo `xorm:"extends"`
}

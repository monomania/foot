package vo

import "tesou.io/platform/foot-parent/foot-api/module/analy/pojo"

type AnalyResultVO struct {
	pojo.AnalyResult `xorm:"extends"`
	/**
	 * 主队id
	 */
	MainTeamId string
	/**
	 * 客队id
	 */
	GuestTeamId string

}

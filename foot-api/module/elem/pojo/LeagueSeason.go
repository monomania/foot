package pojo

import (
	"tesou.io/platform/foot-parent/foot-api/common/base/pojo"
)

/**
联赛赛季表,,
 */
type LeagueSeason struct {
	//联赛id
	LeagueId string  `xorm:" comment('LeagueId') index"`
	//赛季
	Season string  `xorm:" comment('Season') index"`
	//赛季开始的月份
	BeginMonth int `xorm:" comment('BeginMonth') index"`

	//最大的回合数
	Round int `xorm:" comment('最大的回合数')"`

	pojo.BasePojo `xorm:"extends"`
}



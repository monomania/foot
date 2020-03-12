package pojo

import (
	"tesou.io/platform/foot-parent/foot-api/common/base/pojo"
)

/**
联赛赛季次级表,,  如联赛,升级附加赛,降级附加赛
 */
type LeagueSub struct {
	//联赛id
	LeagueId string  `xorm:" comment('LeagueId') index"`
	//赛季
	Season string `xorm:" comment('Season') index"`
	//赛季开始的月份
	BeginMonth int `xorm:" comment('BeginMonth') index"`

	SubId string  `xorm:" comment('SubId') index"`
	SubName string  `xorm:" comment('SubName') index"`
	//最大的回合数
	Round int `xorm:" comment('最大的回合数')"`

	pojo.BasePojo `xorm:"extends"`
}



package pojo

import (
	"tesou.io/platform/foot-parent/foot-api/common/base/pojo"
)

/**
联赛赛季次级表,,  如联赛,升级附加赛,降级附加赛
 */
type LeagueSub struct {
	//联赛id
	LeagueId string  `xorm:" comment('LeagueId') index varchar(20)"`

	LeagueName string  `xorm:" comment('LeagueName') index varchar(50)"`
	//赛季
	Season string `xorm:" comment('Season') index varchar(50)"`
	//赛季开始的月份
	BeginMonth int `xorm:" comment('BeginMonth') index"`

	SubId string  `xorm:" comment('SubId') index varchar(20)"`
	SubName string  `xorm:" comment('SubName') index varchar(50)"`
	//最大的回合数
	Round int `xorm:" comment('最大的回合数')"`

	pojo.BasePojo `xorm:"extends"`
}



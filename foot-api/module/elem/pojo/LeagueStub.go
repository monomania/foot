package pojo

import (
	"tesou.io/platform/foot-parent/foot-api/common/base/pojo"
)

/**
联赛子表表
 */
//不管是从哪个平台抓取的数据，都使用win007的联赛的ID数据
type LeagueStub struct {
	//联赛id
	LeagueId string  `xorm:" comment('LeagueId') index"`
	//年份,使用,间隔
	Season string `xorm:" comment('Season')"`
	//subs
	SubIds []string `xorm:" comment('Subs')"`
	//--
	Round int `xorm:" comment('最大的回合数')"`

	pojo.BasePojo `xorm:"extends"`
}



package pojo

import (
	"tesou.io/platform/foot-parent/foot-api/common/base/pojo"
)

/**
联赛子表表
 */
//不管是从哪个平台抓取的数据，都使用win007的联赛的ID数据
type LeagueSub struct {
	//联赛id
	LeagueId string  `xorm:" comment('LeagueId') index"`
	//次级名称
	Name string `xorm:" comment('次级名称') index"`


	pojo.BasePojo `xorm:"extends"`
}


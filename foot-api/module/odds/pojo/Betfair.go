package pojo

import (
	"tesou.io/platform/foot-parent/foot-api/common/base/pojo"
)

/**
必发数据表
*/
type Betfair struct {
	//比赛id
	MatchId string `xorm:"unique(MatchId_Kind) index varchar(20)"`
	//主 ,和 ,客
	Kind string `xorm:"unique(MatchId_Kind) comment('主 ,和 ,客') varchar(20)" `
	//价位
	Price float64 `xorm:" comment('价位') "`
	//成交量
	Volume int `xorm:" comment('成交量')" `
	//比例
	Scale float64 `xorm:" comment('比例')" `
	//盈亏
	Profit int `xorm:" comment('盈亏')" `
	//冷热
	Hot int `xorm:" comment('冷热') "`

	pojo.BasePojo `xorm:"extends"`
}

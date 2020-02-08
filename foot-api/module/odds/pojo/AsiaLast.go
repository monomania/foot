package pojo

import (
	"tesou.io/platform/foot-parent/foot-api/common/base/pojo"
)
/**
亚赔当前表,仅初盘，即时盘
*/
type AsiaLast struct {
	//博彩公司id
	CompId string `xorm:"unique(CompId_MatchId) index"`
	//比赛id
	MatchId string `xorm:"unique(CompId_MatchId) index"`

	/**
	初主队盘口赔率
	*/
	Sp3 float64	`xorm:" comment('Sp3') index"`
	Sp0 float64	`xorm:" comment('Sp0') index"`
	//让球
	SLetBall float64 `xorm:" comment('s让球') index"`

	/**
	即时客队盘口赔率
	*/
	Ep3 float64	`xorm:" comment('Ep3') index"`
	Ep0 float64	`xorm:" comment('Ep0') index"`
	//让球
	ELetBall float64 `xorm:" comment('e让球') index"`

	//数据时间
	OddDate string	`xorm:" comment('数据时间') index"`

	pojo.BasePojo `xorm:"extends"`
}

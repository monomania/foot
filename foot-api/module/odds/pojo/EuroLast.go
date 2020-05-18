package pojo

import (
	"tesou.io/platform/foot-parent/foot-api/common/base/pojo"
)
/**
欧赔当前表,仅初盘，即时盘
*/
type EuroLast struct {
	//博彩公司id
	CompId int `xorm:"unique(CompId_MatchId)"`
	CompName string `xorm:"varchar(50)"`
	//比赛id
	MatchId string `xorm:"unique(CompId_MatchId) varchar(20)"`

	/**
	初盘胜平负赔率
	*/
	Sp3 float64	`xorm:" comment('Sp3') index"`
	Sp1 float64	`xorm:" comment('Sp1') index"`
	Sp0 float64	`xorm:" comment('Sp0') index"`

	/**
	即时盘胜平负赔率
	*/
	Ep3 float64	`xorm:" comment('Ep3') index"`
	Ep1 float64	`xorm:" comment('Ep1') index"`
	Ep0 float64	`xorm:" comment('Ep0') index"`

	//赔付率
	Payout float64	`xorm:" comment('赔付率') index"`

	//数据时间
	OddDate string	`xorm:" comment('数据时间') index varchar(20)"`

	/**
	胜平负凯利指数
	*/
	Kelly3 float64
	Kelly1 float64
	Kelly0 float64

	pojo.BasePojo `xorm:"extends"`
}

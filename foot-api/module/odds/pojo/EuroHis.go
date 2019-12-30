package pojo

import "tesou.io/platform/foot-parent/foot-api/common/base/pojo"

/**
欧赔历史表
 */
type EuroHis struct {
	/**
	初盘胜平负赔率
	*/
	Sp3 float64
	Sp1 float64
	Sp0 float64

	/**
	即时盘胜平负赔率
	*/
	Ep3 float64
	Ep1 float64
	Ep0 float64

	//博彩公司id
	CompId string `xorm:"unique(CompId_MatchId_OddDate)"`
	//比赛id
	MatchId string `xorm:"unique(CompId_MatchId_OddDate)"`
	//数据时间
	OddDate string	`xorm:"unique(CompId_MatchId_OddDate)"`

	//赔付率
	Payout float64

	/**
	胜平负凯利指数
	*/
	Kelly3 float64
	Kelly1 float64
	Kelly0 float64

	pojo.BasePojo `xorm:"extends"`
}

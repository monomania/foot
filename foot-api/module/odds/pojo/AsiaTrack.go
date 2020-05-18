package pojo

import "tesou.io/platform/foot-parent/foot-api/common/base/pojo"

/**
亚赔历史,变化过程表
*/
type AsiaTrack struct {
	/**
	初主队盘口赔率
	*/
	Sp3 float64
	Sp0 float64
	//让球
	SPanKou float64 `xorm:" comment('s让球') index"`

	//博彩公司id
	CompId   int `xorm:"unique(CompId_MatchId_OddDate_Num)"`
	//比赛id
	MatchId string `xorm:"unique(CompId_MatchId_OddDate_Num)  varchar(20)"`
	//数据时间
	OddDate string `xorm:"unique(CompId_MatchId_OddDate_Num)  varchar(20)"`

	Num int `xorm:"unique(CompId_MatchId_OddDate_Num)"`

	pojo.BasePojo `xorm:"extends"`
}

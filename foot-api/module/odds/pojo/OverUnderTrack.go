package pojo

import "tesou.io/platform/foot-parent/foot-api/common/base/pojo"

type OverUnderTrack struct {
	//博彩公司id
	CompId int `xorm:"unique(CompId_MatchId_OddDate_Num) "`
	//比赛id
	MatchId string `xorm:"unique(CompId_MatchId_OddDate_Num)  varchar(20)"`
	//数据时间
	OddDate string `xorm:"unique(CompId_MatchId_OddDate_Num) varchar(20)"`

	Num int `xorm:"unique(CompId_MatchId_OddDate_Num)"`

	/**
	初主队盘口赔率
	*/
	Sp3 float64
	Sp0 float64
	//大小球
	PanKou float64 `xorm:" comment('大小球') index"`



	pojo.BasePojo `xorm:"extends"`
}

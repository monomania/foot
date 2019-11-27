package entity

import (
	"log"
	"tesou.io/platform/foot-parent/foot-core/common/base/entity"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

type EuroHis struct {
	//博彩公司id
	CompId string
	//比赛id
	MatchId string

	/**
	初盘胜平负赔率
	*/
	Sp3 float64
	Sp1 float64
	Sp0 float64

	//赔付率
	Payout float64

	//数据时间
	OddDate string

	/**
	胜平负凯利指数
	*/
	Kelly3 float64
	Kelly1 float64
	Kelly0 float64

	entity.Base `xorm:"extends"`
}

func (this *EuroHis) FindExists() bool {
	exist, err := mysql.GetEngine().Exist(&EuroHis{MatchId: this.MatchId, CompId: this.CompId, OddDate: this.OddDate})
	if err != nil {
		log.Println("错误:", err)
	}
	return exist
}

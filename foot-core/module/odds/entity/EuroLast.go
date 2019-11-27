package entity

import (
	"log"
	"strings"
	"tesou.io/platform/foot-parent/foot-core/common/base/entity"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

type EuroLast struct {
	//博彩公司id
	CompId string `xorm:"unique(CompId,MatchId,OddDate)"`
	//比赛id
	MatchId string `xorm:"unique(CompId,MatchId,OddDate)"`

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

	//赔付率
	Payout float64

	//数据时间
	OddDate string `xorm:"unique(CompId,MatchId,OddDate)"`

	entity.Base `xorm:"extends"`
}

func (this *EuroLast) FindExists() bool {
	exist, err := mysql.GetEngine().Get(this)
	if err != nil {
		log.Println("错误:", err)
	}
	return exist
}

//根据比赛ID查找欧赔
func (this *EuroLast) FindByMatchId() []*EuroLast {
	dataList := make([]*EuroLast, 0)
	err := mysql.GetEngine().Where(" MatchId = ? ", this.MatchId).Find(dataList)
	if err != nil {
		log.Println("错误:", err)
	}
	return dataList
}

//根据比赛ID和波菜公司ID查找欧赔
func (this *EuroLast) FindByMatchIdCompId(matchId string, compIds ...string) []*EuroLast {
	dataList := make([]*EuroLast, 0)
	sql_build := strings.Builder{}
	sql_build.WriteString(" MatchId = '"+matchId+"' AND CompId in ( 0 ")
	for _, v := range compIds {
		sql_build.WriteString(" , ")
		sql_build.WriteString(v)
	}
	sql_build.WriteString(")")
	err := mysql.GetEngine().Where(sql_build.String()).Find(&dataList)
	if err != nil {
		log.Println("错误:", err)
	}
	return dataList
}

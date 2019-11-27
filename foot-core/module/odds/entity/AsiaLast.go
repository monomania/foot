package entity

import (
	"log"
	"strings"
	"tesou.io/platform/foot-parent/foot-core/common/base/entity"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

type AsiaLast struct {
	//博彩公司id
	CompId string `xorm:"unique(CompId_MatchId)"`
	//比赛id
	MatchId string `xorm:"unique(CompId_MatchId)"`

	/**
	初上下盘赔率
	*/
	Sp3 float64
	Sp0 float64
	//让球
	SLetBall string

	/**
	即时上下盘赔率
	*/
	Ep3 float64
	Ep0 float64
	//让球
	ELetBall string

	//数据时间
	OddDate string

	entity.Base `xorm:"extends"`
}

//查看数据是否已经存在
func (this *AsiaLast) FindExists() bool {
	exist, err := mysql.GetEngine().Get(this)
	if err != nil {
		log.Println("错误:", err)
	}
	return exist
}

//根据比赛ID查找亚赔
func (this *AsiaLast) FindByMatchId() []*AsiaLast {
	dataList := make([]*AsiaLast, 0)
	err := mysql.GetEngine().Where(" MatchId = ? ", this.MatchId).Find(dataList)
	if err != nil {
		log.Println("错误:", err)
	}
	return dataList
}

//根据比赛ID和波菜公司ID查找欧赔
func (this *AsiaLast) FindByMatchIdCompId(matchId string, compIds ...string) []*AsiaLast {
	dataList := make([]*AsiaLast, 0)
	sql_build := strings.Builder{}
	sql_build.WriteString(" MatchId = '"+matchId+"' AND CompId in ( '0' ")
	for _, v := range compIds {
		sql_build.WriteString(" ,'")
		sql_build.WriteString(v)
		sql_build.WriteString("'")
	}
	sql_build.WriteString(")")
	err := mysql.GetEngine().Where(sql_build.String()).Find(&dataList)
	if err != nil {
		log.Println("错误:", err)
	}
	return dataList
}


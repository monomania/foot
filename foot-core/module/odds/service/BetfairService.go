package service

import (
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/odds/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

type BetfairService struct {
	mysql.BaseService
}

//查看数据是否已经存在
func (this *BetfairService) Exist(v *pojo.Betfair) (string, bool) {
	temp := &pojo.Betfair{MatchId: v.MatchId,Kind:v.Kind}
	var id string
	exist, err := mysql.GetEngine().Get(temp)
	if err != nil {
		base.Log.Error("Exist:", err)
	}
	if exist {
		id = temp.Id
	}
	return id, exist
}

//根据比赛ID查找亚赔
func (this *BetfairService) FindByMatchId(matchId string) []*pojo.Betfair {
	dataList := make([]*pojo.Betfair, 0)
	err := mysql.GetEngine().Where(" MatchId = ? ", matchId).Find(dataList)
	if err != nil {
		base.Log.Error("FindByMatchId:", err)
	}
	return dataList
}

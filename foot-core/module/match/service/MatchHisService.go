package service

import (
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

type MatchHisService struct {
	mysql.BaseService
}

func (this *MatchHisService) FindExists(v *pojo.MatchHis) bool {
	has, err := mysql.GetEngine().Table("`t_match_his`").Where(" `Id` = ?  ", v.Id).Exist()
	if err != nil {
		base.Log.Error("FindExists", err)
	}
	return has
}

func (this *MatchHisService) FindById(matchId string) *pojo.MatchHis {
	data := new(pojo.MatchHis)
	data.Id = matchId
	_, err := mysql.GetEngine().Get(data)
	if err != nil {
		base.Log.Error("FindById:", err)
	}
	return data
}

package service

import (
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

type BFScoreService struct {
	mysql.BaseService
}

func (this *BFScoreService) Exist(e *pojo.BFScore) (string, bool)  {
	sql_build := strings.Builder{}
	sql_build.WriteString(" MatchId = '" + e.MatchId + "' AND TeamId = '" + e.TeamId + "' AND Type = '" + e.Type + "'")
	temp := &pojo.BFScore{}
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

func (this *BFScoreService) FindByMatchId(matchId string) []*pojo.BFScore {
	dataList := make([]*pojo.BFScore, 0)
	sql_build := strings.Builder{}
	sql_build.WriteString(" MatchId = '" + matchId + "'")
	err := mysql.GetEngine().Where(sql_build.String()).Find(&dataList)
	if err != nil {
		base.Log.Error("FindByMatchId:", err)
	}
	return dataList
}



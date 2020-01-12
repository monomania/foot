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

func (this *BFScoreService) FindByMatchId(matchId string) []*pojo.BFScore {
	dataList := make([]*pojo.BFScore, 0)
	sql_build := strings.Builder{}
	sql_build.WriteString(" MatchId = '" + matchId+"'")
	err := mysql.GetEngine().Where(sql_build.String()).Find(&dataList)
	if err != nil {
		base.Log.Error("FindByMatchId:", err)
	}
	return dataList
}

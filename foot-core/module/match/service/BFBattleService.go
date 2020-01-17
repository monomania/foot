package service

import (
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

type BFBattleService struct {
	mysql.BaseService
}

func (this *BFBattleService) Exist(e *pojo.BFBattle) (string, bool) {
	sql_build := strings.Builder{}
	battleMatchDateStr := e.BattleMatchDate.Format("2006-01-02 15:04:05")
	sql_build.WriteString(" BattleMatchDate = '" + battleMatchDateStr + "' AND BattleMainTeamId = '" + e.BattleMainTeamId + "' AND BattleGuestTeamId = '" + e.BattleGuestTeamId + "'")
	temp := &pojo.BFBattle{}
	var id string
	exist, err := mysql.GetEngine().Where(sql_build.String()).Get(temp)
	if err != nil {
		base.Log.Error("Exist:", err)
	}
	if exist {
		id = temp.Id
	}
	return id, exist
}

func (this *BFBattleService) FindByMatchId(matchId string) []*pojo.BFBattle {
	dataList := make([]*pojo.BFBattle, 0)
	sql_build := strings.Builder{}
	sql_build.WriteString(" MatchId = '" + matchId + "'")
	err := mysql.GetEngine().Where(sql_build.String()).Find(&dataList)
	if err != nil {
		base.Log.Error("FindByMatchId:", err)
	}
	return dataList
}

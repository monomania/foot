package service

import (
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	"time"
)

type BFBattleService struct {
	mysql.BaseService
}

func (this *BFBattleService) FindByMatchId(matchId string) []*pojo.BFBattle {
	dataList := make([]*pojo.BFBattle, 0)
	sql_build := strings.Builder{}
	sql_build.WriteString(" MatchId = '" + matchId+"'")
	err := mysql.GetEngine().Where(sql_build.String()).Find(&dataList)
	if err != nil {
		base.Log.Error("FindByMatchId:", err)
	}
	return dataList
}

func (this *BFBattleService) Exist(battleMatchDate time.Time, battleMainTeamId string, battleGuestTeamId string) bool {
	sql_build := strings.Builder{}
	battleMatchDateStr := battleMatchDate.Format("2006-01-02 15:04:05")
	sql_build.WriteString(" BattleMatchDate = '" + battleMatchDateStr + "' AND BattleMainTeamId = '" + battleMainTeamId + "' AND BattleGuestTeamId = '" + battleGuestTeamId + "'")
	result, err := mysql.GetEngine().Where(sql_build.String()).Exist(new(pojo.BFBattle))
	if err != nil {
		base.Log.Error("Exist:", err)
	}
	return result
}
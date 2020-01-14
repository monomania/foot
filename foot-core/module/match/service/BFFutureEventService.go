package service

import (
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	"time"
)

type BFFutureEventService struct {
	mysql.BaseService
}

func (this *BFFutureEventService) FindByMatchId(matchId string) []*pojo.BFFutureEvent {
	dataList := make([]*pojo.BFFutureEvent, 0)
	sql_build := strings.Builder{}
	sql_build.WriteString(" MatchId = '" + matchId + "'")
	err := mysql.GetEngine().Where(sql_build.String()).Find(&dataList)
	if err != nil {
		base.Log.Error("FindByMatchId:", err)
	}
	return dataList
}

func (this *BFFutureEventService) FindNextBattle(matchId string, mainId string) *pojo.BFFutureEvent {
	data := new(pojo.BFFutureEvent)
	sql_build := strings.Builder{}
	sql_build.WriteString(" MatchId = '" + matchId + "' and TeamId = '" + mainId + "'")
	_, err := mysql.GetEngine().Where(sql_build.String()).OrderBy(" EventMatchDate ASC").Get(data)
	if err != nil {
		base.Log.Error("FindNextBattle:", err)
	}
	return data
}

func (this *BFFutureEventService) Exist(matchId string, teamId string, eventMatchDate time.Time) bool {
	sql_build := strings.Builder{}
	eventMatchDateStr := eventMatchDate.Format("2006-01-02 15:04:05")
	sql_build.WriteString(" MatchId = '" + matchId + "' AND TeamId = '" + teamId + "' AND EventMatchDate = '" + eventMatchDateStr + "'")
	result, err := mysql.GetEngine().Where(sql_build.String()).Exist(new(pojo.BFFutureEvent))
	if err != nil {
		base.Log.Error("Exist:", err)
	}
	return result
}

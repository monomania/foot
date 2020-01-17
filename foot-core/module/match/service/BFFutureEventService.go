package service

import (
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

type BFFutureEventService struct {
	mysql.BaseService
}

func (this *BFFutureEventService) Exist(e *pojo.BFFutureEvent) (string, bool) {
	sql_build := strings.Builder{}
	eventMatchDateStr := e.EventMatchDate.Format("2006-01-02 15:04:05")
	sql_build.WriteString(" MatchId = '" + e.MatchId + "' AND TeamId = '" + e.TeamId + "' AND EventMatchDate = '" + eventMatchDateStr + "'")
	temp := &pojo.BFFutureEvent{}
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



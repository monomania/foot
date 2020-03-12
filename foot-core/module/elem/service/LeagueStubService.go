package service

import (
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/elem/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

//联赛表
type LeagueStubService struct {
	mysql.BaseService
}

func (this *LeagueStubService) Exist(e *pojo.LeagueStub) (string, bool) {
	//sql_build := strings.Builder{}
	//eventMatchDateStr := e.EventMatchDate.Format("2006-01-02 15:04:05")
	//sql_build.WriteString(" MatchId = '" + e.MatchId + "' AND TeamId = '" + e.TeamId + "' AND EventMatchDate = '" + eventMatchDateStr + "'")
	//temp := &pojo.BFFutureEvent{}
	//var id string
	//exist, err := mysql.GetEngine().Where(sql_build.String()).Get(temp)
	//if err != nil {
	//	base.Log.Error("Exist:", err)
	//}
	//if exist {
	//	id = temp.Id
	//}
	//return id, exist
	return "",false
}

func (this *LeagueStubService) FindById(id string) *pojo.League {
	league := new(pojo.League)
	league.Id = id
	_, err := mysql.GetEngine().Get(league)
	if err != nil {
		base.Log.Info("FindById:", err)
	}
	return league
}

func (this *LeagueStubService) FindByName(name string) *pojo.League {
	league := new(pojo.League)
	league.Name = name
	_, err := mysql.GetEngine().Get(league)
	if err != nil {
		base.Log.Info("FindByName:", err)
	}
	return league
}

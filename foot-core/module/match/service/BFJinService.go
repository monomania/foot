package service

import (
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

type BFJinService struct {
	mysql.BaseService
}

func (this *BFJinService) Exist(e *pojo.BFJin) (string, bool) {
	sql_build := strings.Builder{}
	sql_build.WriteString(" MatchTimeStr = '" + e.MatchTimeStr + "' AND HomeTeam = '" + e.HomeTeam + "' AND GuestTeam = '" + e.GuestTeam + "'")
	temp := &pojo.BFJin{}
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

func (this *BFJinService) FindByMatchId(matchId string) []*pojo.BFJin {
	dataList := make([]*pojo.BFJin, 0)
	sql_build := strings.Builder{}
	sql_build.WriteString(" ScheduleID = '" + matchId + "'")
	err := mysql.GetEngine().Where(sql_build.String()).Find(&dataList)
	if err != nil {
		base.Log.Error("FindByMatchId:", err)
	}
	return dataList
}
func (this *BFJinService) FindNearByMatchId(matchId string,count int) []*pojo.BFJin {
	dataList := make([]*pojo.BFJin, 0)
	sql_build := strings.Builder{}
	sql_build.WriteString(" ScheduleID = '" + matchId + "'")
	err := mysql.GetEngine().Where(sql_build.String()).OrderBy("MatchTimeStr DESC").Limit(count,0).Find(&dataList)
	if err != nil {
		base.Log.Error("FindByMatchId:", err)
	}
	return dataList
}

func (this *BFJinService) FindNearByTeamName(teamName string,count int) []*pojo.BFJin {
	dataList := make([]*pojo.BFJin, 0)
	sql_build := strings.Builder{}
	sql_build.WriteString(" HomeTeam = '" + teamName + "' OR GuestTeam = '" + teamName + "'")
	err := mysql.GetEngine().Where(sql_build.String()).OrderBy("MatchTimeStr DESC").Limit(count,0).Find(&dataList)
	if err != nil {
		base.Log.Error("FindNearByTeamName:", err)
	}
	return dataList
}

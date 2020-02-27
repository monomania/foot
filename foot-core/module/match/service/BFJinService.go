package service

import (
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	"time"
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

func (this *BFJinService) FindNearByTeamName(matchDate time.Time,teamName string,count int) []*pojo.BFJin {
	matchDateStr := matchDate.Format("2006-01-02 15:04:05")
	dataList := make([]*pojo.BFJin, 0)
	sql_build := strings.Builder{}
	sql_build.WriteString("  STR_TO_DATE(MatchTimeStr, '%Y%m%d%H%i%s') < '"+ matchDateStr +"' AND ( HomeTeam = '" + teamName + "' OR  GuestTeam = '" + teamName + "') ")
	err := mysql.GetEngine().Where(sql_build.String()).OrderBy("MatchTimeStr DESC").Limit(count,0).Find(&dataList)
	if err != nil {
		base.Log.Error("FindNearByTeamName:", err)
	}
	return dataList
}

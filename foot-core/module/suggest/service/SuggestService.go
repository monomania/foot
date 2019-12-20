package service

import (
	vo2 "tesou.io/platform/foot-parent/foot-api/module/suggest/vo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

/**
发布推荐
*/
type SuggestService struct {
	mysql.BaseService
}

func (this *SuggestService) Query(param *vo2.SuggestVO) []*vo2.SuggestVO {
	sql := `
SELECT 
  l.Name AS LeagueName,
  ml.MainTeamId AS MainTeam,
  ml.GuestTeamId AS GuestTeam,
  ml.MainTeamGoals AS MainTeamGoal,
  al.ELetBall AS LetBall,
  ml.GuestTeamGoals AS GuestTeamGoal,
  ar.* 
FROM
  foot.t_league l,
  foot.t_match_last ml,
  foot.t_asia_last al,
  foot.t_analy_result ar 
WHERE ml.LeagueId = l.Id 
  AND ml.Id = al.MatchId 
  AND al.CompId = '韦德' 
  AND ml.Id = ar.MatchId 
	`
	if len(param.AlFlag) > 0 {
		sql += " AND ar.AlFlag = '" + param.AlFlag + "' "
	}
	if len(param.BeginDateStr) > 0 {
		sql += " AND ml.`MatchDate` >= '" + param.BeginDateStr + "' "
	}

	if len(param.EndDateStr) > 0 {
		sql += " AND ml.`MatchDate` <= '" + param.EndDateStr + "' "
	}

	sql += " ORDER BY ar.MatchDate ASC,  ar.PreResult DESC "

	//结果值
	entitys := make([]*vo2.SuggestVO, 0)
	//执行查询
	this.FindBySQL(sql, &entitys)

	return entitys

}

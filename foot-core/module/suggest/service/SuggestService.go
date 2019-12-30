package service

import (
	"strconv"
	vo2 "tesou.io/platform/foot-parent/foot-api/module/suggest/vo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

/**
发布推荐
*/
type SuggestService struct {
	mysql.BaseService
}

/**
查询待选池中的比赛
*/
func (this *SuggestService) QueryTbs(param *vo2.SuggestVO) []*vo2.SuggestVO {
	sql := `
SELECT 
  l.Name AS LeagueName,
  mh.MainTeamId AS MainTeam,
  mh.GuestTeamId AS GuestTeam,
  mh.MainTeamGoals AS MainTeamGoal,
  mh.GuestTeamGoals AS GuestTeamGoal,
  ar.* 
FROM
  foot.t_league l,
  foot.t_match_his mh,
  foot.t_analy_result ar 
WHERE mh.LeagueId = l.Id 
  AND mh.Id = ar.MatchId
  AND (ar.HitCount < THitCount or ar.HitCount = 0) 
	`
	if len(param.AlFlag) > 0 {
		sql += " AND ar.AlFlag = '" + param.AlFlag + "' "
	}
	if len(param.BeginDateStr) > 0 {
		sql += " AND mh.`MatchDate` >= '" + param.BeginDateStr + "' "
	}

	if len(param.EndDateStr) > 0 {
		sql += " AND mh.`MatchDate` <= '" + param.EndDateStr + "' "
	}
	if param.IsDesc {
		sql += " ORDER BY ar.MatchDate DESC, l.id ASC,mh.MainTeamId asc, ar.PreResult DESC "
	} else {
		sql += " ORDER BY ar.MatchDate ASC,  l.id ASC,mh.MainTeamId asc,ar.PreResult DESC "
	}
	//结果值
	entitys := make([]*vo2.SuggestVO, 0)
	//执行查询
	this.FindBySQL(sql, &entitys)

	return entitys

}


func (this *SuggestService) Query(param *vo2.SuggestVO) []*vo2.SuggestVO {
	sql := `
SELECT 
  l.Name AS LeagueName,
  mh.MainTeamId AS MainTeam,
  mh.GuestTeamId AS GuestTeam,
  mh.MainTeamGoals AS MainTeamGoal,
  mh.GuestTeamGoals AS GuestTeamGoal,
  ar.* 
FROM
  foot.t_league l,
  foot.t_match_his mh,
  foot.t_analy_result ar 
WHERE mh.LeagueId = l.Id 
  AND mh.Id = ar.MatchId
  AND ar.HitCount > 0 
  AND ar.HitCount >= ar.THitCount
	`
	if param.HitCount > 0 {
		sql += " AND ar.HitCount >= '" + strconv.Itoa(param.HitCount) + "' "
	}

	if len(param.AlFlag) > 0 {
		sql += " AND ar.AlFlag in (" + param.AlFlag + ") "
	}
	if len(param.BeginDateStr) > 0 {
		sql += " AND mh.`MatchDate` >= '" + param.BeginDateStr + "' "
	}

	if len(param.EndDateStr) > 0 {
		sql += " AND mh.`MatchDate` <= '" + param.EndDateStr + "' "
	}
	if param.IsDesc {
		sql += " ORDER BY ar.MatchDate DESC, l.id ASC,mh.MainTeamId asc, ar.PreResult DESC "
	} else {
		sql += " ORDER BY ar.MatchDate ASC,  l.id ASC,mh.MainTeamId asc,ar.PreResult DESC "
	}
	//结果值
	entitys := make([]*vo2.SuggestVO, 0)
	//执行查询
	this.FindBySQL(sql, &entitys)

	return entitys

}

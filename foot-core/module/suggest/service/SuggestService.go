package service

import (
	"strconv"
	vo2 "tesou.io/platform/foot-parent/foot-api/module/suggest/vo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	"tesou.io/platform/foot-parent/foot-core/common/utils"
	"tesou.io/platform/foot-parent/foot-core/module/analy/constants"
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
  AND ar.HitCount < THitCount 
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
		sql += " ORDER BY ar.MatchDate DESC, l.id ASC, ar.PreResult DESC "
	} else {
		sql += " ORDER BY ar.MatchDate ASC,  l.id ASC,ar.PreResult DESC "
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
	`
	if param.HitCount > 0 {
		sql += "  AND ar.HitCount >= THitCount AND ar.HitCount >= '" + strconv.Itoa(param.HitCount) + "' "
	} else {
		hit_count_str := utils.GetVal(constants.SECTION_NAME, "hit_count")
		sql += " AND ar.HitCount >= THitCount AND ar.HitCount >= '" + hit_count_str + "' "
	}

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
		sql += " ORDER BY ar.MatchDate DESC, l.id ASC, ar.PreResult DESC "
	} else {
		sql += " ORDER BY ar.MatchDate ASC,  l.id ASC,ar.PreResult DESC "
	}
	//结果值
	entitys := make([]*vo2.SuggestVO, 0)
	//执行查询
	this.FindBySQL(sql, &entitys)

	return entitys

}

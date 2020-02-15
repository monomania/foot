package service

import (
	"strconv"
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	vo2 "tesou.io/platform/foot-parent/foot-api/module/suggest/vo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	service2 "tesou.io/platform/foot-parent/foot-core/module/match/service"
	"tesou.io/platform/foot-parent/foot-core/module/odds/service"
)

/**
发布推荐
*/
type SuggestService struct {
	mysql.BaseService
	service.EuroLastService
	service.AsiaLastService
	service2.BFScoreService
	service2.BFBattleService
	service2.BFFutureEventService
}


/**
查询待选池中的比赛
 */
func (this *SuggestService) QueryGuts(param *vo2.SuggStubVO) []*vo2.SuggStubVO {
	sql := `
SELECT 
  l.Id AS LeagueId,
  l.Name AS LeagueName,
  mh.MainTeamId AS MainTeam,
  mh.GuestTeamId AS GuestTeam,
  mh.MainTeamGoals AS MainTeamGoal,
  mh.GuestTeamGoals AS GuestTeamGoal,
  ar.* 
FROM
  foot.t_league l,
  foot.t_match_his mh,
  foot.t_analy_result ar,
  (SELECT 
    ar1.MatchId 
  FROM
    foot.t_analy_result ar1,
    foot.t_analy_result ar2 
  WHERE ar1.MatchId = ar2.MatchId 
    AND ar1.AlFlag = 'E2' 
    AND ar2.AlFlag = 'C1' 
    AND ar1.PreResult = ar2.PreResult
    ) temp 
WHERE mh.LeagueId = l.Id 
  AND mh.Id = ar.MatchId 
  AND ar.MatchId = temp.MatchId 
  AND ar.AlFlag IN ('E2' ,'C1')
	`
	if len(param.BeginDateStr) > 0 {
		sql += " AND mh.`MatchDate` >= '" + param.BeginDateStr + "' "
	}

	if len(param.EndDateStr) > 0 {
		sql += " AND mh.`MatchDate` <= '" + param.EndDateStr + "' "
	}
	if param.IsDesc {
		sql += " ORDER BY ar.MatchDate DESC, l.id ASC,mh.MainTeamId asc, ar.`AlFlag` DESC,ar.PreResult DESC "
	} else {
		sql += " ORDER BY ar.MatchDate ASC,  l.id ASC,mh.MainTeamId asc,ar.`AlFlag` DESC,ar.PreResult DESC "
	}
	//结果值
	dataList := make([]*vo2.SuggStubVO, 0)
	//执行查询
	this.FindBySQL(sql, &dataList)

	return dataList
}

/**
查询待选池中的比赛
 */
func (this *SuggestService) QueryTbs(param *vo2.SuggStubVO) []*vo2.SuggStubVO {
	sql := `
SELECT 
  l.Id AS LeagueId,
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
	if len(param.AlFlags) > 0 {
		sql += " AND ar.AlFlag in (''"
		for _, v := range param.AlFlags {
			sql += ",'" + v + "'"
		}
		sql += " ) "
	}
	if len(param.BeginDateStr) > 0 {
		sql += " AND mh.`MatchDate` >= '" + param.BeginDateStr + "' "
	}

	if len(param.EndDateStr) > 0 {
		sql += " AND mh.`MatchDate` <= '" + param.EndDateStr + "' "
	}
	if param.IsDesc {
		sql += " ORDER BY ar.`AlFlag` DESC,ar.MatchDate DESC, l.id ASC,mh.MainTeamId asc, ar.PreResult DESC "
	} else {
		sql += " ORDER BY ar.`AlFlag` DESC,ar.MatchDate ASC,  l.id ASC,mh.MainTeamId asc,ar.PreResult DESC "
	}
	//结果值
	entitys := make([]*vo2.SuggStubVO, 0)
	//执行查询
	this.FindBySQL(sql, &entitys)

	return entitys
}

func (this *SuggestService) Query(param *vo2.SuggStubVO) []*vo2.SuggStubVO {
	sql := `
SELECT 
  l.Id AS LeagueId,
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
  AND ar.TOVoid IS FALSE
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
		sql += " ORDER BY ar.`AlFlag` DESC,ar.MatchDate DESC, l.id ASC,mh.MainTeamId asc, ar.PreResult DESC "
	} else {
		sql += " ORDER BY ar.`AlFlag` ASC,ar.MatchDate ASC,  l.id ASC,mh.MainTeamId asc,ar.PreResult DESC "
	}
	//结果值
	entitys := make([]*vo2.SuggStubVO, 0)
	//执行查询
	this.FindBySQL(sql, &entitys)

	return entitys
}

func (this *SuggestService) QueryDetail(param *vo2.SuggStubDetailVO) []*vo2.SuggStubDetailVO {
	sql := `
SELECT 
  l.Id AS LeagueId,
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
		sql += " ORDER BY ar.`AlFlag` DESC,ar.MatchDate DESC, l.id ASC,mh.MainTeamId asc, ar.PreResult DESC "
	} else {
		sql += " ORDER BY ar.`AlFlag` DESC,ar.MatchDate ASC,  l.id ASC,mh.MainTeamId asc,ar.PreResult DESC "
	}

	//结果值
	dataList := make([]*vo2.SuggStubDetailVO, 0)
	//执行查询
	this.FindBySQL(sql, &dataList)
	if len(dataList) <= 0 {
		return dataList
	}

	//for _, e := range dataList {
	//	matchId := e.MatchId
	//	eOddList := this.EuroLastService.FindByMatchId(matchId)
	//	aOddList := this.AsiaLastService.FindByMatchId(matchId)
	//	bfsList := this.BFScoreService.FindByMatchId(matchId)
	//	bfbList := this.BFBattleService.FindByMatchId(matchId)
	//	bffeList := this.BFFutureEventService.FindByMatchId(matchId)
	//
	//}
	for _, v := range dataList {
		matchId := v.MatchId
		//积分,排名
		bfsList := this.BFScoreService.FindByMatchId(matchId)
		if len(bfsList) > 0 {
			for _, e := range bfsList { //bfs_arr有多语言版本,条数很多
				if e.TeamId == v.MainTeam {
					if e.Type == "总" {
						v.BFSMainZong = e
					}
					if e.Type == "主" {
						v.BFSMainZhu = e
					}
					if e.Type == "近" {
						v.BFSMainJin = e
					}
				} else if e.TeamId == v.GuestTeam {
					if e.Type == "总" {
						v.BFSGuestZong = e
					}
					if e.Type == "客" {
						v.BFSGuestKe = e
					}
					if e.Type == "近" {
						v.BFSGuestJin = e
					}
				}
			}
		}
		if nil == v.BFSMainZong {
			v.BFSMainZong = new(pojo.BFScore)
		}
		if nil == v.BFSMainZhu {
			v.BFSMainZhu = new(pojo.BFScore)
		}
		if nil == v.BFSMainJin {
			v.BFSMainJin = new(pojo.BFScore)
		}
		if nil == v.BFSGuestZong {
			v.BFSGuestZong = new(pojo.BFScore)
		}
		if nil == v.BFSGuestKe {
			v.BFSGuestKe = new(pojo.BFScore)
		}
		if nil == v.BFSGuestJin {
			v.BFSGuestJin = new(pojo.BFScore)
		}

		//过往战绩
		bfbList := this.BFBattleService.FindByMatchId(matchId)
		battleCount := 0
		mainWin := 0
		drawCount := 0
		guestWin := 0
		for _, e := range bfbList {
			if e.BattleMainTeamId == v.MainTeam || e.BattleMainTeamId == v.GuestTeam {
				battleCount++
			}
			if e.BattleMainTeamId == v.MainTeam && e.BattleMainTeamGoals > e.BattleGuestTeamGoals {
				mainWin++
			}
			if e.BattleGuestTeamId == v.MainTeam && e.BattleGuestTeamGoals > e.BattleMainTeamGoals {
				mainWin++
			}
			if e.BattleMainTeamId == v.GuestTeam && e.BattleMainTeamGoals > e.BattleGuestTeamGoals {
				guestWin++
			}
			if e.BattleGuestTeamId == v.GuestTeam && e.BattleGuestTeamGoals > e.BattleMainTeamGoals {
				guestWin++
			}
			if (e.BattleMainTeamId == v.MainTeam || e.BattleMainTeamId == v.GuestTeam) && e.BattleMainTeamGoals == e.BattleGuestTeamGoals {
				drawCount++
			}
		}
		v.BattleCount = battleCount
		v.BattleMainWinCount = mainWin
		v.BattleDrawCount = drawCount
		v.BattleGuestWinCount = guestWin
		//未来赛事
		bffe_main := this.BFFutureEventService.FindNextBattle(matchId, v.MainTeam)
		if nil != bffe_main {
			v.MainNextMainTeam = bffe_main.EventMainTeamId
		}
		bffe_guest := this.BFFutureEventService.FindNextBattle(matchId, v.GuestTeam)
		if nil != bffe_main {
			v.GuestNextMainTeam = bffe_guest.EventMainTeamId
		}
	}
	return dataList
}


func (this *SuggestService) QueryLeisu(param *vo2.SuggStubDetailVO) []*vo2.SuggStubDetailVO {
	sql := `
SELECT 
  l.Id AS LeagueId,
  l.Name AS LeagueName,
  mh.MainTeamId AS MainTeam,
  mh.GuestTeamId AS GuestTeam,
  mh.MainTeamGoals AS MainTeamGoal,
  mh.GuestTeamGoals AS GuestTeamGoal,
  ar.* 
FROM
  foot.t_league l,
  foot.t_match_his mh,
  foot.t_analy_result ar,
  (SELECT 
    ar1.MatchId 
  FROM
    foot.t_analy_result ar1,
    foot.t_analy_result ar2 
  WHERE ar1.MatchId = ar2.MatchId 
    AND ar1.AlFlag = 'E2' 
    AND ar2.AlFlag = 'C1' 
    AND ar2.TOVoid IS FALSE 
    AND ar1.PreResult = ar2.PreResult) temp 
WHERE mh.LeagueId = l.Id 
  AND mh.Id = ar.MatchId 
  AND ar.MatchId = temp.MatchId 
  AND ar.AlFlag IN ('E2', 'C1') 
  AND mh.Id NOT IN 
  (SELECT 
    matchId 
  FROM
    foot.t_pub p 
  WHERE p.CreateTime > DATE_SUB(NOW(), INTERVAL 2 DAY))
	`
	if len(param.BeginDateStr) > 0 {
		sql += " AND mh.`MatchDate` >= '" + param.BeginDateStr + "' "
	}

	if len(param.EndDateStr) > 0 {
		sql += " AND mh.`MatchDate` <= '" + param.EndDateStr + "' "
	}
	if param.IsDesc {
		sql += " ORDER BY ar.MatchDate DESC, l.id ASC,mh.MainTeamId asc, ar.`AlFlag` DESC,ar.PreResult DESC "
	} else {
		sql += " ORDER BY ar.MatchDate ASC,  l.id ASC,mh.MainTeamId asc,ar.`AlFlag` DESC,ar.PreResult DESC "
	}

	//结果值
	dataList := make([]*vo2.SuggStubDetailVO, 0)
	//执行查询
	this.FindBySQL(sql, &dataList)
	if len(dataList) <= 0 {
		return dataList
	}

	for _, v := range dataList {
		matchId := v.MatchId
		//积分,排名
		bfsList := this.BFScoreService.FindByMatchId(matchId)
		if len(bfsList) > 0 {
			for _, e := range bfsList { //bfs_arr有多语言版本,条数很多
				if e.TeamId == v.MainTeam {
					if e.Type == "总" {
						v.BFSMainZong = e
					}
					if e.Type == "主" {
						v.BFSMainZhu = e
					}
					if e.Type == "近" {
						v.BFSMainJin = e
					}
				} else if e.TeamId == v.GuestTeam {
					if e.Type == "总" {
						v.BFSGuestZong = e
					}
					if e.Type == "客" {
						v.BFSGuestKe = e
					}
					if e.Type == "近" {
						v.BFSGuestJin = e
					}
				}
			}
		}
		if nil == v.BFSMainZong {
			v.BFSMainZong = new(pojo.BFScore)
		}
		if nil == v.BFSMainZhu {
			v.BFSMainZhu = new(pojo.BFScore)
		}
		if nil == v.BFSMainJin {
			v.BFSMainJin = new(pojo.BFScore)
		}
		if nil == v.BFSGuestZong {
			v.BFSGuestZong = new(pojo.BFScore)
		}
		if nil == v.BFSGuestKe {
			v.BFSGuestKe = new(pojo.BFScore)
		}
		if nil == v.BFSGuestJin {
			v.BFSGuestJin = new(pojo.BFScore)
		}

		//过往战绩
		bfbList := this.BFBattleService.FindByMatchId(matchId)
		battleCount := 0
		mainWin := 0
		drawCount := 0
		guestWin := 0
		for _, e := range bfbList {
			if e.BattleMainTeamId == v.MainTeam || e.BattleMainTeamId == v.GuestTeam {
				battleCount++
			}
			if e.BattleMainTeamId == v.MainTeam && e.BattleMainTeamGoals > e.BattleGuestTeamGoals {
				mainWin++
			}
			if e.BattleGuestTeamId == v.MainTeam && e.BattleGuestTeamGoals > e.BattleMainTeamGoals {
				mainWin++
			}
			if e.BattleMainTeamId == v.GuestTeam && e.BattleMainTeamGoals > e.BattleGuestTeamGoals {
				guestWin++
			}
			if e.BattleGuestTeamId == v.GuestTeam && e.BattleGuestTeamGoals > e.BattleMainTeamGoals {
				guestWin++
			}
			if (e.BattleMainTeamId == v.MainTeam || e.BattleMainTeamId == v.GuestTeam) && e.BattleMainTeamGoals == e.BattleGuestTeamGoals {
				drawCount++
			}
		}
		v.BattleCount = battleCount
		v.BattleMainWinCount = mainWin
		v.BattleDrawCount = drawCount
		v.BattleGuestWinCount = guestWin
		//未来赛事
		bffe_main := this.BFFutureEventService.FindNextBattle(matchId, v.MainTeam)
		if nil != bffe_main {
			v.MainNextMainTeam = bffe_main.EventMainTeamId
		}
		bffe_guest := this.BFFutureEventService.FindNextBattle(matchId, v.GuestTeam)
		if nil != bffe_main {
			v.GuestNextMainTeam = bffe_guest.EventMainTeamId
		}
	}
	return dataList
}
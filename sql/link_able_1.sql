SELECT 
  ar.MatchDate,
  ar.AlFlag,
  ar.PreResult,
  ar.Result,
  ar.HitCount,
  l.Name AS LeagueName,
  mh.MainTeamId AS MainTeam,
  mh.GuestTeamId AS GuestTeam,
  mh.MainTeamGoal AS MainTeamGoal,
  mh.GuestTeamGoal AS GuestTeamGoal
FROM
  t_league l,
  t_match_his mh,
  t_analy_result ar,
  (SELECT 
    ar1.MatchId
  FROM
    t_analy_result ar1,
    t_analy_result ar2
  WHERE ar1.MatchId = ar2.MatchId
    AND ar1.AlFlag = 'Q1'
    AND ar2.AlFlag = 'A1'
    ) temp 
WHERE mh.LeagueId = l.Id 
  AND mh.Id = ar.MatchId 
  AND ar.MatchId = temp.MatchId
ORDER BY ar.MatchDate DESC,
  l.id ASC,
  mh.MainTeamId ASC,
  ar.AlFlag DESC ,
  ar.PreResult DESC
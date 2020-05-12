SELECT 
  ar.`MatchDate`,
  ar.`AlFlag`,
  ar.`PreResult`,
  ar.`Result`,
  ar.`HitCount`,
  l.Name AS LeagueName,
  mh.MainTeamId AS MainTeam,
  mh.GuestTeamId AS GuestTeam,
  mh.MainTeamGoals AS MainTeamGoal,
  mh.GuestTeamGoals AS GuestTeamGoal 
FROM
  foot.t_league l,
  foot.t_match_his mh,
  foot.t_analy_result ar,
  (SELECT 
    ar1.`MatchId` 
  FROM
    foot.`t_analy_result` ar1,
    foot.`t_analy_result` ar2 
  WHERE ar1.`MatchId` = ar2.`MatchId` 
    AND ar1.`AlFlag` = 'Q1' 
    AND ar2.`AlFlag` = 'A1' 
    ) temp 
WHERE mh.LeagueId = l.Id 
  AND mh.Id = ar.MatchId 
  AND ar.`MatchId` = temp.MatchId 
ORDER BY ar.MatchDate DESC,
  l.id ASC,
  mh.MainTeamId ASC,
  ar.`AlFlag` DESC ,
  ar.`PreResult` DESC
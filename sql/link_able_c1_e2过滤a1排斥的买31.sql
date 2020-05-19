SELECT 
  ar.MatchDate,
  ar.TOVoid,
  ar.TOVoidDesc,
  ar.AlFlag,
  ar.PreResult,
  ar.Result,
  mh.MainTeamGoal AS MainTeamGoal,
  mh.GuestTeamGoal AS GuestTeamGoal,
  ar.HitCount,
  ar.LetBall,
  ar.MyLetBall,
  l.Name AS LeagueName,
  mh.MainTeamId AS MainTeam,
  mh.GuestTeamId AS GuestTeam
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
    AND ar1.AlFlag = 'E2'
    AND ar2.AlFlag = 'C1'
    AND ar1.PreResult = ar2.PreResult
    AND ar2.TOVoidDesc = '') temp
WHERE mh.LeagueId = l.Id
  AND mh.Id = ar.MatchId
  #AND mh.MainTeamGoal != mh.GuestTeamGoal
  AND ar.MatchId = temp.MatchId
  AND ar.AlFlag IN ('E2', 'C1')
ORDER BY ar.MatchDate DESC,
  l.id ASC,
  mh.MainTeamId ASC,
  ar.AlFlag DESC,
  ar.PreResult DESC
#-------------------------------------
#-------------------------------------
#-------------------------------------
SELECT
  COUNT(1) / 2,
  ar.Result
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
    AND ar1.AlFlag = 'E2'
    AND ar2.AlFlag = 'C1'
    AND ar1.PreResult = ar2.PreResult
    AND ar2.TOVoidDesc = '') temp
WHERE mh.LeagueId = l.Id
  AND mh.Id = ar.MatchId
  #AND mh.MainTeamGoal != mh.GuestTeamGoal
  AND ar.MatchId = temp.MatchId
  AND ar.AlFlag IN ('E2', 'C1')
GROUP BY ar.Result
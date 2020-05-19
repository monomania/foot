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
    temp.MatchId
  FROM
    (SELECT
      ar1.MatchId,
      ar1.PreResult
    FROM
      t_analy_result ar1,
      t_analy_result ar2
    WHERE ar1.MatchId = ar2.MatchId
      AND ar1.AlFlag = 'E2'
      AND ar2.AlFlag = 'C1'
      AND ar1.PreResult = ar2.PreResult
      AND ar2.TOVoidDesc = '') temp,
    t_analy_result a1
  WHERE a1.MatchId = temp.MatchId
    AND a1.AlFlag = 'A1'
    AND a1.PreResult = temp.PreResult) temp
WHERE mh.LeagueId = l.Id
  AND mh.Id = ar.MatchId
  AND ar.MatchId = temp.MatchId
  AND ar.AlFlag IN ('E2', 'C1','A1')
ORDER BY ar.MatchDate DESC,
  l.id ASC,
  mh.MainTeamId ASC,
  ar.AlFlag DESC,
  ar.PreResult DESC
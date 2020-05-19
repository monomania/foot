
#用于查询稳定的场次
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
    ar1.Id
  FROM
    t_analy_result ar1
  WHERE ar1.MatchId IN
    (SELECT
      temp.matchId
    FROM
      t_analy_result temp
    GROUP BY temp.MatchId
    HAVING COUNT(1) >= 3)) temp
WHERE mh.LeagueId = l.Id
  AND mh.Id = ar.MatchId
  AND ar.Id = temp.id
ORDER BY ar.MatchDate DESC,
  l.id ASC,
  mh.MainTeamId ASC,
  ar.AlFlag DESC
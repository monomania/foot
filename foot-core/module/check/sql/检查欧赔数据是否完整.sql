SELECT
  l.`CompId`,
  le.`Name`,
  la.`MainTeamId`,
  la.`GuestTeamId`
FROM
  foot.`t_euro_last` l,
  foot.`t_match_last` la,
  foot.`t_league` le
WHERE la.`LeagueId` = le.`Id`
  AND l.`MatchId` = la.`Id`
GROUP BY l.`MatchId`
HAVING COUNT(1) < 2
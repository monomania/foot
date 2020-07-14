SELECT
  l.`Name`,
  h.`Id` AS "比赛Id",
  h.`MatchDate` AS "比赛时间",
  h.`MainTeam` AS "主队",
  h.`GuestTeam` AS "客队",
  h.`MainTeamGoal` AS "主队全场得分",
  h.`GuestTeamGoal` AS "客队全场得分",
  h.`MainTeamHalfGoal` AS "主队半场得分",
  h.`GuestTeamHalfGoal` AS "客队半场得分",
  ca.`Name` AS "菠菜公司",
  track.`Sp3`,
  track.`SPanKou`,
  track.`Sp0`,
  track.`Num` AS "变赔顺序",
  track.`OddDate` AS "赔率时间"
FROM
  t_league l,
  t_match_his h,
  (SELECT
    t1.*
  FROM
    t_asia_track_202005 t1
  UNION
  ALL
  SELECT
    t2.*
  FROM
    t_asia_track_202006 t2
  UNION
  ALL
  SELECT
    t3.*
  FROM
    t_asia_track_202007 t3) track,
  t_comp_asia ca
WHERE 1 = 1
  AND h.`LeagueId` = l.`Id`
  AND h.`Id` = track.`MatchId`
  AND track.`CompId` = ca.`Id`
ORDER BY h.`Id`,
  track.compId,
  track.num
CREATE OR REPLACE VIEW v_c4_0_novoid AS
SELECT
  l.`Name` AS "联赛",
  h.`Id`,
  h.`MatchDate`,
  CONCAT(
    h.`MainTeam`,
    "(",
    r.`SLetBall`,
    ")",
    "(",
    r.`LetBall`,
    ")",
    h.`GuestTeam`
  ) AS "主(初)(即时)客",
  CONCAT(
    h.`MainTeamGoal`,
    ":",
    h.`GuestTeamGoal`,
    "(",
    h.`MainTeamHalfGoal`,
    ":",
    h.`GuestTeamHalfGoal`,
    ")"
  ) AS "比分(半场)",
  r.`AlFlag`,
  r.`Result`,
  r.`PreResult`,
  r.`TOVoid`
FROM
  t_league l,
  t_match_his h,
  t_analy_result r
WHERE 1 = 1
  AND h.`LeagueId` = l.`Id`
  AND h.`Id` = r.`MatchId`
  AND r.`AlFlag` = "C4"
  AND r.`TOVoid` IS FALSE
  and r.preResult = 0
ORDER BY r.`MatchDate` DESC,
  r.Result,
  r.preresult DESC
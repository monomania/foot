SELECT 
  *
FROM
  (SELECT
    ar.`AlFlag` AS 'flag',
    DATE_FORMAT(ar.`MatchDate`, '%w') AS 'week',
    ar.`LetBall` AS 'letball',
    COUNT(1) AS 'total',
    SUM(IF (ar.`Result` = "命中", 1, 0)) AS 'right',
    SUM(IF (ar.`Result` = "错误", 1, 0)) AS 'error',
    SUM(IF (ar.`Result` = "走盘", 1, 0)) AS 'none',
    SUM(
      IF (
        mh.`MainTeamGoals` > mh.`GuestTeamGoals`,
        1,
        0
      )
    ) AS '3',
    SUM(
      IF (
        mh.`MainTeamGoals` = mh.`GuestTeamGoals`,
        1,
        0
      )
    ) AS '1',
    SUM(
      IF (
        mh.`MainTeamGoals` < mh.`GuestTeamGoals`,
        1,
        0
      )
    ) AS '0'
  FROM
    foot.`t_analy_result` ar,
    foot.`t_match_his` mh,
    foot.`t_league` l
  WHERE ar.`MatchId` = mh.`Id`
    AND mh.`LeagueId` = l.`Id` #and ar.`TOVoidDesc` = ''
    AND ar.`Result` != "待定"
    AND ar.`AlFlag` IN ('A1', 'A3')
   #AND DATE_FORMAT(ar.`MatchDate`, '%w') = 0
  GROUP BY ar.`AlFlag`,
    DATE_FORMAT(ar.`MatchDate`, '%w'),
    ar.`LetBall`
  ORDER BY ar.`AlFlag`,
    DATE_FORMAT(ar.`MatchDate`, '%w'),
    ar.`LetBall`) t
WHERE t.total > 10
  AND (
    t.error / t.total >= 0.8
    OR t.right / t.total >= 0.8
    OR t.none / t.total >= 0.8
    OR t.1/t.total > 0.5
  ) ORDER BY t.week ASC
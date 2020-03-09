SELECT 
  ar.`AlFlag` AS '模型',
  DATE_FORMAT(ar.`MatchDate`, '%w') AS '星期',
  ar.`LetBall` AS '让球',
  COUNT(1) AS '总场次',
  SUM(IF (ar.`Result` = "命中", 1, 0)) AS '命中场次',
  SUM(IF (ar.`Result` = "走盘", 1, 0)) AS '走盘场次',
  SUM(IF (ar.`Result` = "错误", 1, 0)) AS '错误场次',
  SUM(
    IF (
      mh.`MainTeamGoals` > mh.`GuestTeamGoals`,
      1,
      0
    )
  ) AS '主胜场次',
  SUM(
    IF (
      mh.`MainTeamGoals` = mh.`GuestTeamGoals`,
      1,
      0
    )
  ) AS '平局场次',
  SUM(
    IF (
      mh.`MainTeamGoals` < mh.`GuestTeamGoals`,
      1,
      0
    )
  ) AS '客胜场次'
FROM
  foot.`t_analy_result` ar,
  foot.`t_match_his` mh,
  foot.`t_league` l
WHERE ar.`MatchId` = mh.`Id`
  AND mh.`LeagueId` = l.`Id`
  AND ar.`TOVoidDesc` = ''
  AND ar.`AlFlag` IN ('C1', 'C2')
GROUP BY ar.`AlFlag`,
  DATE_FORMAT(ar.`MatchDate`, '%w'),
  ar.`LetBall`
ORDER BY ar.`AlFlag`,
  DATE_FORMAT(ar.`MatchDate`, '%w'),
  ar.`LetBall`
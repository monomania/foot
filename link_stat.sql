SELECT 
  ar.`AlFlag` AS '模型',
  DATE_FORMAT(ar.`MatchDate`, '%w') AS '星期',
  ar.`LetBall` AS '让球',
  COUNT(1) AS '总场次',
  SUM(IF (ar.`Result` = "命中", 1, 0)) AS '命中',
  SUM(IF (ar.`Result` = "走盘", 1, 0)) AS '走盘',
  SUM(IF (ar.`Result` = "错误", 1, 0)) AS '错误',
  SUM(
    IF (
      mh.`MainTeamGoals` > mh.`GuestTeamGoals`,
      1,
      0
    )
  ) AS '主胜',
  SUM(
    IF (
      mh.`MainTeamGoals` = mh.`GuestTeamGoals`,
      1,
      0
    )
  ) AS '平局',
  SUM(
    IF (
      mh.`MainTeamGoals` < mh.`GuestTeamGoals`,
      1,
      0
    )
  ) AS '客胜'
FROM
  foot.`t_analy_result` ar,
  foot.`t_match_his` mh,
  foot.`t_league` l
WHERE ar.`MatchId` = mh.`Id`
  AND mh.`LeagueId` = l.`Id`
  #and ar.`TOVoidDesc` = ''
  AND ar.`Result` != "待定"
  AND ar.`AlFlag` IN ('A1', 'A3')
GROUP BY ar.`AlFlag`,
  DATE_FORMAT(ar.`MatchDate`, '%w'),
  ar.`LetBall`
ORDER BY ar.`AlFlag`,
  DATE_FORMAT(ar.`MatchDate`, '%w'),
  ar.`LetBall`
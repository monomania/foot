SELECT 
  ar.`AlFlag` AS '模型',
  DATE_FORMAT(ar.`MatchDate`, '%w') AS '星期',
  ar.`LetBall` AS '盘口',
  ar.`Result` AS '结果',
  COUNT(1) AS '场次' 
FROM
  foot.`t_analy_result` ar,
  foot.`t_match_his` mh 
WHERE ar.`MatchId` = mh.`Id` #and mh.`MainTeamGoals` = mh.`GuestTeamGoals`
  AND ar.`TOVoidDesc` = '' 
  AND ar.`AlFlag` IN ('C1','C2')
GROUP BY ar.`AlFlag`,
  ar.`LetBall`,
  ar.`Result`,
  DATE_FORMAT(ar.`MatchDate`, '%w') 
ORDER BY ar.`AlFlag`,
  DATE_FORMAT(ar.`MatchDate`, '%w'),
  ar.`LetBall`,
  ar.`Result` 
SELECT 
  ar.`AlFlag` AS '模型',
  DATE_FORMAT(ar.`MatchDate`, '%w') AS '星期',
  ar.`LetBall` AS '盘口',
  COUNT(1) AS '场次',
  ar.`Result` AS '结果' 
FROM
  foot.`t_analy_result` ar,
  foot.`t_match_his` mh 
WHERE ar.`MatchId` = mh.`Id` 
  #and mh.`MainTeamGoals` = mh.`GuestTeamGoals` 
  AND ar.`TOVoidDesc` = '' 
GROUP BY ar.`AlFlag`,
  ar.`Result`,
  DATE_FORMAT(ar.`MatchDate`, '%w') 
ORDER BY ar.`AlFlag`,
  DATE_FORMAT(ar.`MatchDate`, '%w'),
  ar.`LetBall`,
  ar.`Result` 
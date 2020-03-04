SELECT 

  CONCAT(
    mh.`MainTeamGoals`,
    ":",
    mh.`GuestTeamGoals`
  ) AS '比分',
  COUNT(1) AS '次数' 
FROM
  foot.`t_analy_result` ar,
  foot.`t_match_his` mh 
WHERE ar.`MatchId` = mh.`Id` 
  AND ar.`Result` != "待定" 
  AND ar.`TOVoidDesc` = '' 
  AND ar.`AlFlag` IN ('C1', 'C2') #and DATE_FORMAT(ar.`MatchDate`, '%w') = 4 
  AND mh.`MainTeamGoals` = mh.`GuestTeamGoals` 
  AND ar.`LetBall` IN (0, 0.25, 0.75) 
GROUP BY 
  mh.`MainTeamGoals`,
  mh.`GuestTeamGoals` 
  
  ORDER BY mh.`MainTeamGoals`
  
  

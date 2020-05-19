SELECT 
  COUNT(IF(r.Result = "命中", 1, 0)) AS hitCount,
  COUNT(IF(r.Result = "错误", 1, 0)) AS failCount,
  DATE_FORMAT(r.MatchDate, '%Y-%m') AS md 
FROM
  t_analy_result r 
WHERE r.AlFlag = 'C1' 
GROUP BY md ORDER BY md ASC 
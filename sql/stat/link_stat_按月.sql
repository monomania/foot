SELECT 
  t.right / (t.total-t.none) AS '命中率',
    t.1 / (t.total) AS '平局率',
  t.*
FROM
  (SELECT
    DATE_FORMAT(ar.MatchDate, '%m') AS 'month',
    ar.AlFlag AS 'flag',
        IF(ar.preresult = 3 , '主','客') AS '推荐' ,
    COUNT(1) AS 'total',
    SUM(IF (ar.Result = "命中", 1, 0)) AS 'right',
    SUM(IF (ar.Result = "错误", 1, 0)) AS 'error',
    SUM(
      IF (
        ar.Result = "走盘"
        OR ar.Result = "未知"
        OR ar.Result = "",
        1,
        0
      )
    ) AS 'none',
    SUM(
      IF (
        mh.MainTeamGoal > mh.GuestTeamGoal,
        1,
        0
      )
    ) AS '3',
    SUM(
      IF (
        mh.MainTeamGoal = mh.GuestTeamGoal,
        1,
        0
      )
    ) AS '1',
    SUM(
      IF (
        mh.MainTeamGoal < mh.GuestTeamGoal,
        1,
        0
      )
    ) AS '0'
  FROM
    t_analy_result ar,
    t_match_his mh,
    t_league l
  WHERE ar.MatchId = mh.Id
    AND mh.LeagueId = l.Id
    AND ar.Result != "待定"
  GROUP BY ar.AlFlag,
    DATE_FORMAT(ar.MatchDate, '%m'),ar.preresult
  ORDER BY ar.AlFlag,
    DATE_FORMAT(ar.MatchDate, '%m'),ar.preresult) t
WHERE 1 = 1
ORDER BY t.flag,t.month ASC
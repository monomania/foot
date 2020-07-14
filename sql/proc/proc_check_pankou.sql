DELIMITER $$

USE `foot_001`$$

DROP PROCEDURE  IF EXISTS `proc_check_pankou`$$

CREATE DEFINER=`root`@`%` PROCEDURE  `proc_check_pankou`(
  IN matchId VARCHAR(20),
  IN CompId INT,
  OUT pick INT
)
BEGIN
    DECLARE upCount  INT DEFAULT 0 ;
    DECLARE downCount  INT DEFAULT 0 ;
    DECLARE done INT DEFAULT 0;
    DECLARE l_matchid VARCHAR(20);
    DECLARE l_spankou,temp_spankou DOUBLE;
    DECLARE l_num INT;
    DECLARE matchDate DATETIME;

    -- 声明游标
    DECLARE mc CURSOR FOR (SELECT t.matchId,t.SPanKou,t.Num FROM v_asia_track_temp t  WHERE 1=1 );
    -- 指定游标循环结束时的返回值
    DECLARE CONTINUE HANDLER FOR NOT FOUND SET done=10000;


    SELECT h.MatchDate INTO matchDate FROM t_match_his h WHERE h.id = matchid LIMIT 0,1;

    SET @sqlstr = "CREATE OR REPLACE  VIEW v_asia_track_temp as select t.* from t_asia_track_";
    SET @sqlstr = CONCAT(@sqlstr , DATE_FORMAT(matchDate,'%Y%m')," t where 1=1 ");
    SET @sqlstr = CONCAT(@sqlstr , " AND t.matchid = '",matchId,"'");
    SET @sqlstr = CONCAT(@sqlstr , " AND t.OddDate < '",matchDate,"'");
    SET @sqlstr = CONCAT(@sqlstr , " AND t.compId = ",compid);
    SET @sqlstr = CONCAT(@sqlstr , " ORDER BY t.num DESC ");



    PREPARE stmt FROM @sqlstr;
    EXECUTE stmt;
    DEALLOCATE PREPARE stmt;


    OPEN mc;

    SET temp_spankou = -10000;
    xxx:LOOP
        -- 根据游标当前指向的一条数据
        FETCH mc INTO l_matchid,l_spankou,l_num;

        -- 当 游标的返回值为 1 时 退出 loop循环
        IF done = 10000 THEN
            LEAVE xxx;
        END IF;

        IF temp_spankou = -10000 THEN
	   SET temp_spankou = l_spankou;
        END IF;

	IF l_spankou < temp_spankou THEN
	   SET temp_spankou = l_spankou;
	   SET downCount = downCount+1;
        END IF;

        IF l_spankou > temp_spankou THEN
	   SET temp_spankou = l_spankou;
	   SET upCount = upCount+1;
        END IF;
    END LOOP;
    CLOSE mc;

   SET pick = upCount * 10000 + downCount ;


   SELECT
  h.`Id`,
  l.`Name`,
  h.`MatchDate` AS "比赛时间",
  h.`MainTeam` AS "主队",
  h.`GuestTeam` AS "客队",
  h.`MainTeamGoal` AS "主得分",
  h.`GuestTeamGoal` AS "客得分",
  h.`MainTeamHalfGoal` AS "主半得分",
  h.`GuestTeamHalfGoal` AS "客半得分",
  upCount AS "升盘",
  downCount AS "降盘",
  ah.sp3,
  ah.spankou AS "初盘",
  ah.sp0,
  ah.ep3,
  ah.epankou AS "终盘",
  ah.ep0
FROM
  t_league l,
  t_match_his h,
  t_asia_his_2020 ah
WHERE 1 = 1
  AND h.`LeagueId` = l.`Id`
  AND h.id = matchId
  AND ah.matchId = matchId
  AND ah.compid = compid
  ;


END$$

DELIMITER ;
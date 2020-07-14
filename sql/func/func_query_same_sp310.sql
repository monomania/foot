DELIMITER $$

USE `foot_001`$$

DROP FUNCTION IF EXISTS `func_query_same_sp310`$$

CREATE DEFINER=`root`@`%` FUNCTION `func_query_same_sp310`(
  matchDate DATETIME,
  compId INT,
  sp3 DOUBLE,
  sp1 DOUBLE,
  sp0 DOUBLE
)  RETURNS INT
BEGIN

  DECLARE pick  INT ;
  DECLARE done INT DEFAULT 0;
  DECLARE mw,mwc INT DEFAULT 0;
  DECLARE ml,mlc INT DEFAULT 0;

  -- 声明游标
DECLARE mc CURSOR FOR SELECT
  IF(
    mh.MainTeamGoal > mh.GuestTeamGoal,
    1,
    0
  ) AS mainWin,
  IF(
    mh.MainTeamGoal < mh.GuestTeamGoal,
    1,
    0
  ) AS mainLoss
FROM
  t_match_his mh,
  t_euro_his_2020 eh
WHERE 1 = 1
  AND mh.Id = eh.MatchId
  AND eh.CompId = compId
  AND eh.Sp3 = sp3
  AND eh.Sp1 = sp1
  AND eh.Sp0 = sp0
  AND eh.MatchDate < matchDate ;

  -- 指定游标循环结束时的返回值
  DECLARE CONTINUE HANDLER FOR NOT FOUND SET done=10000;

 OPEN mc;


    xxx:LOOP
        -- 根据游标当前指向的一条数据
        FETCH mc INTO mw,ml;
        -- 当 游标的返回值为 1 时 退出 loop循环
        IF done = 10000 THEN
            LEAVE xxx;
        END IF;

        IF mw > 0 THEN
	   SET mwc = mwc+1;
        END IF;


        IF ml > 0 THEN
	   SET mlc = mlc+1;
        END IF;

    END LOOP;
    CLOSE mc;



  SET pick = mwc * 10000 + mlc;
  RETURN pick ;
END$$

DELIMITER ;
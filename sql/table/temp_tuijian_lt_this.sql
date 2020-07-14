
DROP table IF EXISTS `temp_tuijian_lt_this`;

CREATE TABLE temp_tuijian_lt_this AS
SELECT
  t.*
FROM
  (SELECT
    l.Id AS LeagueId,
    l.Name AS LeagueName,
    mh.MainTeam AS MainTeam,
    mh.GuestTeam AS GuestTeam,
    mh.MainTeamGoal AS MainTeamGoal,
    mh.GuestTeamGoal AS GuestTeamGoal,
    mh.MainTeamHalfGoal,
    mh.GuestTeamHalfGoal,
    func_istuijian_lt_this (
      mh.`LeagueId`,
      ar.`AlFlag`,
      ar.`PreResult`,
      mh.`MatchDate`
    ) AS istuijian,
    ar.*
  FROM
    t_league l,
    t_match_his mh,
    t_analy_result ar
  WHERE 1 = 1
    AND mh.LeagueId = l.Id
    AND mh.Id = ar.MatchId) t
WHERE t.istuijian != 0 ;


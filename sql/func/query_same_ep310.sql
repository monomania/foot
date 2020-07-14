SELECT
 *
FROM
 (SELECT
  (t.s1idx + t.s2idx + t.s3idx) sidx,
  FLOOR((t.s1idx + t.s2idx + t.s3idx) / 10000) sidx_3,
  FLOOR((t.s1idx + t.s2idx + t.s3idx) % 10000) sidx_0,
  IF(
   FLOOR((t.s1idx + t.s2idx + t.s3idx) / 10000) > FLOOR((t.s1idx + t.s2idx + t.s3idx) % 10000),
   "6666666666",
   "败败败败败"
  ) sidx_result,
  (t.e1idx + t.e2idx + t.e3idx) eidx,
  FLOOR((t.e1idx + t.e2idx + t.e3idx) / 10000) eidx_3,
  FLOOR((t.e1idx + t.e2idx + t.e3idx) % 10000) eidx_0,
  IF(
   FLOOR((t.e1idx + t.e2idx + t.e3idx) / 10000) > FLOOR((t.e1idx + t.e2idx + t.e3idx) % 10000),
   "6666666666",
   "败败败败败"
  ) eidx_result,
  t.*
 FROM
  (SELECT
   IF(
    mh.MainTeamGoal > mh.GuestTeamGoal,
    "6666666666",
    "败败败败败"
   ) AS mainResult,
   mh.id,
   mh.`MatchDate`,
   mh.`MainTeam`,
   CONCAT(
    "(",
    ah.spankou,
    ")",
    "(",
    ah.epankou,
    ")"
   ) AS "(初)(即时)",
   CONCAT(
    mh.`MainTeamGoal`,
    ":",
    mh.`GuestTeamGoal`,
    "(",
    mh.`MainTeamHalfGoal`,
    ":",
    mh.`GuestTeamHalfGoal`,
    ")"
   ) AS score,
   mh.`GuestTeam`,
   #bc1.compId,
   bc1.compName AS compName1,
   func_query_same_sp310 (
    mh.`MatchDate`,
    "281",
    bc1.sp3,
    bc1.sp1,
    bc1.sp0
   ) AS s1idx,
   #bc1.spayout,
   func_query_same_ep310 (
    mh.`MatchDate`,
    "281",
    bc1.ep3,
    bc1.ep1,
    bc1.ep0
   ) AS e1idx,
   #bc1.epayout,
   #bc2.compId,
   bc2.compName AS compName2,
   func_query_same_ep310 (
    mh.`MatchDate`,
    "115",
    bc2.sp3,
    bc2.sp1,
    bc2.sp0
   ) AS s2idx,
   #bc2.spayout,
   func_query_same_ep310 (
    mh.`MatchDate`,
    "115",
    bc2.ep3,
    bc2.ep1,
    bc2.ep0
   ) AS e2idx,
   #bc2.epayout,
   #bc3.compId,
   bc3.compName AS compName3,
   func_query_same_ep310 (
    mh.`MatchDate`,
    "16",
    bc3.sp3,
    bc3.sp1,
    bc3.sp0
   ) AS s3idx,
   #bc3.spayout,
   func_query_same_ep310 (
    mh.`MatchDate`,
    "16",
    bc3.ep3,
    bc3.ep1,
    bc3.ep0
   ) AS e3idx #bc3.epayout
  FROM
   t_match_his mh,
   (SELECT
    ah.*
   FROM
    t_asia_his_2020 ah
   WHERE ah.compid = "8") ah,
   (SELECT
    eh.*
   FROM
    t_euro_his_2020 eh
   WHERE eh.compid = "281") bc1,
   (SELECT
    eh.*
   FROM
    t_euro_his_2020 eh
   WHERE eh.compid = "115") bc2,
   (SELECT
    eh.*
   FROM
    t_euro_his_2020 eh
   WHERE eh.compid = "16") bc3
  WHERE 1 = 1
   AND mh.id = ah.matchId

   AND mh.id = bc1.matchId
   AND mh.id = bc2.matchId
   AND mh.id = bc3.matchId) t) t
WHERE 1 = 1
 AND ABS(t.sidx_3 - t.sidx_0) >= 3
 AND ABS(t.eidx_3 - t.eidx_0) >= 3
LIMIT 0, 50
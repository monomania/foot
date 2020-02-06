package service

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	entity5 "tesou.io/platform/foot-parent/foot-api/module/analy/pojo"
	"tesou.io/platform/foot-parent/foot-api/module/analy/vo"
	entity2 "tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	entity3 "tesou.io/platform/foot-parent/foot-api/module/odds/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	"tesou.io/platform/foot-parent/foot-core/module/analy/constants"
	service3 "tesou.io/platform/foot-parent/foot-core/module/elem/service"
	service2 "tesou.io/platform/foot-parent/foot-core/module/match/service"
	"tesou.io/platform/foot-parent/foot-core/module/odds/service"
	"time"
)

type AnalyService struct {
	mysql.BaseService
	service.EuroHisService
	service.AsiaHisService
	service2.MatchLastService
	service2.MatchHisService
	service3.LeagueService
	//是否打印赔率数据
	PrintOddData bool
}

func (this *AnalyService) Find(matchId string, alFlag string) *entity5.AnalyResult {
	data := entity5.AnalyResult{MatchId: matchId, AlFlag: alFlag}
	mysql.GetEngine().Get(&data)
	return &data
}

func (this *AnalyService) FindAll() []*entity5.AnalyResult {
	dataList := make([]*entity5.AnalyResult, 0)
	mysql.GetEngine().OrderBy("CreateTime Desc").Find(&dataList)
	return dataList
}

/**
C1使用的检查是否存在其他模型存在互斥选项
 */
func (this *AnalyService) FindOtherAlFlag(matchId string, alFlag string, preResult int) bool {
	sql_build := `
SELECT 
  ar.* 
FROM
  foot.t_analy_result ar 
WHERE ar.MatchId = ? 
  AND ar.AlFlag != ? 
  AND ar.PreResult != ?
     `
	//结果值
	entitys := make([]*vo.AnalyResultVO, 0)
	//执行查询
	mysql.GetEngine().SQL(sql_build, matchId, alFlag, preResult).Find(&entitys)
	if len(entitys) > 0 {
		return true
	}
	return false
}

/**
更新结果
 */
func (this *AnalyService) ModifyAllResult() {
	sql_build := `
SELECT 
  ar.* 
FROM
  foot.t_analy_result ar 
     `
	//结果值
	entitys := make([]*entity5.AnalyResult, 0)
	//执行查询
	this.FindBySQL(sql_build, &entitys)

	if len(entitys) <= 0 {
		return
	}
	for _, e := range entitys {
		aList := this.AsiaHisService.FindByMatchIdCompId(e.MatchId, constants.DEFAULT_REFER_ASIA)
		if nil == aList || len(aList) < 1 {
			aList = make([]*entity3.AsiaHis, 1)
			aList[0] = new(entity3.AsiaHis)
		}
		his := this.MatchHisService.FindById(e.MatchId)
		if nil == his {
			continue
		}
		last := new(entity2.MatchLast)
		last.MatchDate = his.MatchDate
		last.DataDate = his.DataDate
		last.LeagueId = his.LeagueId
		last.MainTeamId = his.MainTeamId
		last.MainTeamGoals = his.MainTeamGoals
		last.GuestTeamId = his.GuestTeamId
		last.GuestTeamGoals = his.GuestTeamGoals
		if strings.EqualFold(e.AlFlag, "E2") || strings.EqualFold(e.AlFlag, "C1") {
			//E2使用特别自身的验证结果方法
			e.Result = this.IsRight2Option(last, e)
		} else {
			e.Result = this.IsRight(last, e)
		}
		this.Modify(e)
	}
}

/**
更新结果
 */
func (this *AnalyService) ModifyResult() {
	sql_build := `
SELECT 
  ar.* 
FROM
  foot.t_analy_result ar 
WHERE DATE_ADD(ar.MatchDate, INTERVAL 3 HOUR) > NOW()
     `
	//结果值
	entitys := make([]*entity5.AnalyResult, 0)
	//执行查询
	this.FindBySQL(sql_build, &entitys)

	if len(entitys) <= 0 {
		return
	}
	for _, e := range entitys {
		aList := this.AsiaHisService.FindByMatchIdCompId(e.MatchId, constants.DEFAULT_REFER_ASIA)
		if nil == aList || len(aList) < 1 {
			aList = make([]*entity3.AsiaHis, 1)
			aList[0] = new(entity3.AsiaHis)
		}
		his := this.MatchHisService.FindById(e.MatchId)
		if nil == his {
			continue
		}
		last := new(entity2.MatchLast)
		last.MatchDate = his.MatchDate
		last.DataDate = his.DataDate
		last.LeagueId = his.LeagueId
		last.MainTeamId = his.MainTeamId
		last.MainTeamGoals = his.MainTeamGoals
		last.GuestTeamId = his.GuestTeamId
		last.GuestTeamGoals = his.GuestTeamGoals
		if strings.EqualFold(e.AlFlag, "E2") || strings.EqualFold(e.AlFlag, "C1") {
			//E2使用特别自身的验证结果方法
			e.Result = this.IsRight2Option(last, e)
		} else {
			e.Result = this.IsRight(last, e)
		}
		this.Modify(e)
	}
}

/**
获取可发布的数据项
1.预算结果是主队
2.比赛未开始
3.比赛未结束
4.alName 算法名称，默认为Euro81_616Service ;
5.option 3(只筛选主队),1(只筛选平局),0(只筛选客队)选项
*/
func (this *AnalyService) List(alName string, hitCount int, option int) []*vo.AnalyResultVO {
	sql_build := `
SELECT 
  l.Name as LeagueName,
  ml.MainTeamId,
  ml.GuestTeamId,
  ar.* 
FROM
  foot.t_match_last ml,
  foot.t_league l,
  foot.t_analy_result ar 
WHERE ml.LeagueId = l.Id 
  AND ml.Id = ar.MatchId 
  AND ar.HitCount >= THitCount
  AND ar.LeisuPubd IS FALSE 
  AND ar.MatchDate > NOW()
     `

	if len(alName) > 0 {
		sql_build += " AND ar.AlFlag = '" + alName + "' "
	}
	if hitCount > 0 {
		sql_build += " AND ar.HitCount >= " + strconv.Itoa(hitCount)
	} else {
		sql_build += " AND ar.HitCount > 0 "
	}
	if option >= 0 {
		sql_build += " AND ar.PreResult = " + strconv.Itoa(option) + " "
	}
	sql_build += " ORDER BY ar.MatchDate ASC ,ar.PreResult DESC  "
	//结果值
	entitys := make([]*vo.AnalyResultVO, 0)
	//执行查询
	this.FindBySQL(sql_build, &entitys)
	return entitys
}

//测试加载数据
func (this *AnalyService) LoadByMatchId(matchId string) []*entity5.AnalyResult {
	sql_build := `
SELECT 
  ml.*,
  bc.id,
  bc.name AS compName,
  el.* 
FROM
  t_match_last ml,
  t_euro_last el,
  t_comp bc 
WHERE ml.id = el.matchid 
  AND el.compid = bc.id 
	`
	sql_build += "  AND ml.id = '" + matchId + "' "
	//结果值
	entitys := make([]*entity5.AnalyResult, 0)
	//执行查询
	this.FindBySQL(sql_build, &entitys)
	return entitys
}

func (this *AnalyService) IsRight2Option(v *entity2.MatchLast, analy *entity5.AnalyResult) string {
	//比赛结果
	var globalResult int
	h2, _ := time.ParseDuration("129m")
	//比赛结束的时间点
	matchEndDate := v.MatchDate.Add(h2)
	if matchEndDate.Before(time.Now()) {
		//比赛未结束
		globalResult = -1
	} else {
		if v.MainTeamGoals > v.GuestTeamGoals {
			globalResult = 3
		} else if v.MainTeamGoals < v.GuestTeamGoals {
			globalResult = 0
		} else {
			globalResult = 1
		}
	}
	var resultFlag string
	if globalResult == -1 {
		resultFlag = constants.UNCERTAIN
	} else if globalResult == analy.PreResult || globalResult == 1 {
		resultFlag = constants.HIT
	} else {
		resultFlag = constants.UNHIT
	}

	//打印数据
	league := this.LeagueService.FindById(v.LeagueId)
	matchDateStr := v.MatchDate.Format("2006-01-02 15:04:05")
	base.Log.Info("比赛Id:" + v.Id + ",比赛时间:" + matchDateStr + ",联赛:" + league.Name + ",对阵:" + v.MainTeamId + "(" + strconv.FormatFloat(analy.LetBall, 'f', -1, 64) + ")" + v.GuestTeamId + ",预算结果:" + strconv.Itoa(analy.PreResult) + ",已得结果:" + strconv.Itoa(v.MainTeamGoals) + "-" + strconv.Itoa(v.GuestTeamGoals) + " (" + resultFlag + ")")
	return resultFlag
}

func (this *AnalyService) IsRight(last *entity2.MatchLast, analy *entity5.AnalyResult) string {
	//比赛结果
	globalResult := this.ActualResult(last, analy)
	var resultFlag string
	if globalResult == -1 {
		resultFlag = constants.UNCERTAIN
	} else if globalResult == analy.PreResult {
		resultFlag = constants.HIT
	} else if globalResult == 1 {
		resultFlag = constants.WALKING_PLATE
	} else {
		resultFlag = constants.UNHIT
	}

	//打印数据
	league := this.LeagueService.FindById(last.LeagueId)
	matchDate := last.MatchDate.Format("2006-01-02 15:04:05")
	base.Log.Info("比赛Id:" + last.Id + ",比赛时间:" + matchDate + ",联赛:" + league.Name + ",对阵:" + last.MainTeamId + "(" + strconv.FormatFloat(analy.LetBall, 'f', -1, 64) + ")" + last.GuestTeamId + ",预算结果:" + strconv.Itoa(analy.PreResult) + ",已得结果:" + strconv.Itoa(last.MainTeamGoals) + "-" + strconv.Itoa(last.GuestTeamGoals) + " (" + resultFlag + ")")
	return resultFlag
}

/**
比赛的实际结果计算
*/
func (this *AnalyService) ActualResult(last *entity2.MatchLast, analy *entity5.AnalyResult) int {
	var result int
	h2, _ := time.ParseDuration("129m")
	matchEndDate := last.MatchDate.Add(h2)
	if matchEndDate.Before(time.Now()) {
		//比赛未结束
		return -1
	}

	var mainTeamGoals float64
	elb_sum := analy.LetBall
	if elb_sum > 0 {
		mainTeamGoals = float64(last.MainTeamGoals) - math.Abs(elb_sum)
	} else {
		mainTeamGoals = float64(last.MainTeamGoals) + math.Abs(elb_sum)
	}
	//diff_goals := float64(last.MainTeamGoals-last.GuestTeamGoals) - elb_sum
	//if diff_goals <= 0.25 && diff_goals >= -0.25 {
	//	result = 1
	//}
	if mainTeamGoals > float64(last.GuestTeamGoals) {
		result = constants.WIN
	} else if mainTeamGoals < float64(last.GuestTeamGoals) {
		result = constants.LOST
	} else {
		result = constants.DRAW
	}
	return result
}

/**
1.欧赔是主降还是主升 主降为true
*/
func EuroMainDown(e81data *entity3.EuroHis, e616data *entity3.EuroHis) int {
	if e81data.Ep3 <= e81data.Sp3 && e616data.Ep3 <= e616data.Sp3 {
		return 3
	} else if e81data.Ep0 <= e81data.Sp0 && e616data.Ep0 <= e616data.Sp0 {
		return 0
	}
	return 1
}

/**
2.亚赔是主降还是主升 主降为true
*/
func AsiaMainDown(his *entity3.AsiaHis) bool {
	slb_sum := his.SLetBall
	elb_sum := his.ELetBall

	if elb_sum > slb_sum {
		return true
	} else if elb_sum < slb_sum {
		return false
	} else {
		//初始让球和即时让球一致
		if his.Ep3 < his.Sp3 {
			return true
		}
	}
	return false
}

/**
float64保留两位小数
 */
func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

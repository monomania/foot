package service

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/common/base/pojo"
	entity5 "tesou.io/platform/foot-parent/foot-api/module/analy/pojo"
	"tesou.io/platform/foot-parent/foot-api/module/analy/vo"
	entity2 "tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	entity3 "tesou.io/platform/foot-parent/foot-api/module/odds/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	"tesou.io/platform/foot-parent/foot-core/common/utils"
	"tesou.io/platform/foot-parent/foot-core/module/analy/constants"
	service3 "tesou.io/platform/foot-parent/foot-core/module/elem/service"
	service2 "tesou.io/platform/foot-parent/foot-core/module/match/service"
	"tesou.io/platform/foot-parent/foot-core/module/odds/service"
	"time"
)

type AnalyInterface interface {
	ModelName() string

	Analy_Near()

	Analy(analyAll bool)

	analyStub(v *entity2.MatchLast) (int, *entity5.AnalyResult)
}

type AnalyService struct {
	mysql.BaseService
	service.EuroHisService
	service.EuroTrackService
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

func (this *AnalyService) Analy(analyAll bool, thiz AnalyInterface) {
	var matchList []*entity2.MatchLast
	if analyAll {
		var currentPage,pageSize int64 = 1,1000
		var page *pojo.Page
		page = new(pojo.Page)
		page.PageSize = pageSize
		page.CurPage = currentPage
		matchList = make([]*entity2.MatchLast, 0)
		err := this.MatchHisService.PageSql("select mh.* from foot.t_match_his mh ", page, &matchList)
		if nil != err {
			base.Log.Error(err)
			return
		}
		go this.Analy_Process(matchList, thiz)
		for currentPage <= page.TotalPage {
			currentPage++
			page = new(pojo.Page)
			page.PageSize = pageSize
			page.CurPage = currentPage
			matchList = make([]*entity2.MatchLast, 0)
			err := this.MatchHisService.PageSql("select mh.* from foot.t_match_his mh", page, &matchList)
			if nil != err {
				base.Log.Error(err)
				continue
			}
			go this.Analy_Process(matchList, thiz)
		}
	} else {
		matchList = this.MatchLastService.FindNotFinished()
		this.Analy_Process(matchList, thiz)
	}

}

func (this *AnalyService) Analy_Near(thiz AnalyInterface) {
	matchList := this.MatchLastService.FindNear()
	this.Analy_Process(matchList, thiz)
}

/**
汇总结果并输出，且持久化
 */
func (this *AnalyService) Analy_Process(matchList []*entity2.MatchLast, thiz AnalyInterface) {
	hit_count_str := utils.GetVal(constants.SECTION_NAME, "hit_count")
	hit_count, _ := strconv.Atoi(hit_count_str)
	data_list_slice := make([]interface{}, 0)
	data_modify_list_slice := make([]interface{}, 0)
	var rightCount = 0
	var errorCount = 0

	for _, v := range matchList {
		stub, temp_data := thiz.analyStub(v)

		if nil != temp_data {
			if strings.EqualFold(temp_data.Result, "命中") {
				rightCount++
			}
			if strings.EqualFold(temp_data.Result, "错误") {
				errorCount++
			}
		}

		if stub == 0 || stub == 1 {
			temp_data.TOVoid = false
			hours := v.MatchDate.Sub(time.Now()).Hours()
			if hours > 0 {
				temp_data.THitCount = hit_count
			} else {
				temp_data.THitCount = 1
			}
			if stub == 0 {
				data_list_slice = append(data_list_slice, temp_data)
			} else if stub == 1 {
				data_modify_list_slice = append(data_modify_list_slice, temp_data)
			}
		} else if nil != temp_data {
			temp_data.TOVoid = true
			if len(temp_data.Id) > 0 {
				if temp_data.HitCount >= hit_count {
					temp_data.HitCount = (hit_count / 2) - 1
				} else {
					temp_data.HitCount = 0
				}
				data_modify_list_slice = append(data_modify_list_slice, temp_data)
			}
		}
	}

	base.Log.Info("------------------")
	base.Log.Info("------------------")
	base.Log.Info("------------------")
	base.Log.Info("GOOOO场次:", rightCount)
	base.Log.Info("X0000场次:", errorCount)
	base.Log.Info("------------------")

	this.SaveList(data_list_slice)
	this.ModifyList(data_modify_list_slice)
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
		last.Id = his.Id
		last.MatchDate = his.MatchDate
		last.DataDate = his.DataDate
		last.LeagueId = his.LeagueId
		last.MainTeamId = his.MainTeamId
		last.MainTeamGoals = his.MainTeamGoals
		last.GuestTeamId = his.GuestTeamId
		last.GuestTeamGoals = his.GuestTeamGoals
		if strings.EqualFold(e.AlFlag, "E2") || strings.EqualFold(e.AlFlag, "C1") || strings.EqualFold(e.AlFlag, "C2") {
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
WHERE DATE_ADD(ar.MatchDate, INTERVAL 6 HOUR) > NOW()
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
		last.Id = his.Id
		last.MatchDate = his.MatchDate
		last.DataDate = his.DataDate
		last.LeagueId = his.LeagueId
		last.MainTeamId = his.MainTeamId
		last.MainTeamGoals = his.MainTeamGoals
		last.GuestTeamId = his.GuestTeamId
		last.GuestTeamGoals = his.GuestTeamGoals
		if strings.EqualFold(e.AlFlag, "E2") || strings.EqualFold(e.AlFlag, "C1") || strings.EqualFold(e.AlFlag, "C2") {
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
  AND ar.MatchDate > NOW()4e43
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

//测试加载数据
func (this *AnalyService) DelTovoidData() {
	//E2 C1 不可删除
	sql_build := `
DELETE FROM foot.t_analy_result  WHERE AlFlag IN ("E1","Q1") AND TOVoid IS TRUE
	`
	_, err := mysql.GetEngine().Exec(sql_build)
	if nil != err {
		base.Log.Error("DelTovoidData" + err.Error())
	}
}

func (this *AnalyService) IsRight2Option(last *entity2.MatchLast, analy *entity5.AnalyResult) string {
	if strings.EqualFold(analy.MatchId, "1826976") {
		fmt.Println("--")
	}
	//比赛结果
	var globalResult int
	if utils.GetHourDiffer(time.Now(), last.MatchDate) < 2 {
		//比赛未结束
		globalResult = -1
	} else {
		if last.MainTeamGoals > last.GuestTeamGoals {
			globalResult = 3
		} else if last.MainTeamGoals < last.GuestTeamGoals {
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
	analy.Result = resultFlag
	league := this.LeagueService.FindById(last.LeagueId)
	if this.IsCupMatch(league.Name) {
		analy.TOVoid = true
		analy.TOVoidDesc = "杯赛"
	} else {
		analy.TOVoidDesc = ""
	}

	//打印数据
	matchDateStr := last.MatchDate.Format("2006-01-02 15:04:05")
	base.Log.Info("比赛Id:" + last.Id + ",比赛时间:" + matchDateStr + ",联赛:" + league.Name + ",对阵:" + last.MainTeamId + "(" + strconv.FormatFloat(analy.LetBall, 'f', -1, 64) + ")" + last.GuestTeamId + ",预算结果:" + strconv.Itoa(analy.PreResult) + ",已得结果:" + strconv.Itoa(last.MainTeamGoals) + "-" + strconv.Itoa(last.GuestTeamGoals) + " (" + resultFlag + ")")
	return resultFlag
}

func (this *AnalyService) IsCupMatch(leagueName string) bool {
	if strings.Contains(leagueName, "杯") || strings.Contains(leagueName, "锦") {
		return true;
	}
	return false;
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
	analy.Result = resultFlag

	league := this.LeagueService.FindById(last.LeagueId)
	if this.IsCupMatch(league.Name) {
		analy.TOVoid = true
		analy.TOVoidDesc = "杯赛"
	} else {
		analy.TOVoidDesc = ""
	}

	//打印数据
	matchDate := last.MatchDate.Format("2006-01-02 15:04:05")
	base.Log.Info("比赛Id:" + last.Id + ",比赛时间:" + matchDate + ",联赛:" + league.Name + ",对阵:" + last.MainTeamId + "(" + strconv.FormatFloat(analy.LetBall, 'f', -1, 64) + ")" + last.GuestTeamId + ",预算结果:" + strconv.Itoa(analy.PreResult) + ",已得结果:" + strconv.Itoa(last.MainTeamGoals) + "-" + strconv.Itoa(last.GuestTeamGoals) + " (" + resultFlag + ")")
	return resultFlag
}

/**
比赛的实际结果计算
*/
func (this *AnalyService) ActualResult(last *entity2.MatchLast, analy *entity5.AnalyResult) int {
	var result int
	if utils.GetHourDiffer(time.Now(), last.MatchDate) < 2 {
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
func (this *AnalyService) EuroDirection(e81 *entity3.EuroHis, e616 *entity3.EuroHis) int {
	//val_diff := 0.3
	//e81_3_diff := math.Abs(e81.Ep3 - e81.Sp3)
	//e81_0_diff := math.Abs(e81.Ep0 - e81.Sp0)
	e81_ep3_small := e81.Ep3 <= e81.Sp3
	e81_ep0_small := e81.Ep0 <= e81.Sp0
	//e616_3_diff := math.Abs(e616.Ep3 - e616.Sp3)
	//e616_0_diff := math.Abs(e616.Ep0 - e616.Sp0)
	e616_ep3_small := e616.Ep3 <= e616.Sp3
	e616_ep0_small := e616.Ep0 <= e616.Sp0

	if e616_ep3_small && e81_ep3_small && e616.Ep3 <= e81.Ep3 {
		return 3
	}
	if e616_ep0_small && e81_ep0_small && e616.Ep0 <= e81.Ep0 {
		return 0
	}
	return -1
}

/**
2.亚赔是主降还是主升 主降为true
*/
func (this *AnalyService) AsiaDirectionMulti(matchId string) int {
	aList := this.AsiaHisService.FindByMatchIdCompId(matchId, "Crown", "明陞", "金宝博", "12bet", "盈禾", "18Bet")
	if len(aList) < 3 {
		return -1
	}

	var mainCount, guestCount int
	for _, e := range aList {
		direction := this.AsiaDirection(e)
		if direction == 3 {
			mainCount++
		} else if direction == 0 {
			guestCount++
		}
	}

	if mainCount-guestCount > 1 {
		return 3
	}
	if mainCount-guestCount < -1 {
		return 0
	}
	return -1
}

/**
2.亚赔是主降还是主升 主降为true
*/
func (this *AnalyService) AsiaDirection(ahis *entity3.AsiaHis) int {
	mark := -1
	slb := ahis.SLetBall
	elb := ahis.ELetBall
	ep3_small := ahis.Ep3 < ahis.Sp3
	ep0_small := ahis.Ep0 < ahis.Sp0
	if elb > 0 {
		if elb > slb {
			mark = 3
		} else if elb < slb {
			mark = 0
		} else {
			//初始让球和即时让球一致
			if ep3_small && !ep0_small {
				mark = 3
			} else if !ep3_small && ep0_small {
				mark = 0
			}
		}
	} else {
		if elb < slb {
			mark = 0
		} else if elb > slb {
			mark = 3
		} else {
			//初始让球和即时让球一致
			if ep3_small && !ep0_small {
				mark = 3
			} else if !ep3_small && ep0_small {
				mark = 0
			}
		}
	}
	return mark
}

/**
float64保留两位小数
 */
func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

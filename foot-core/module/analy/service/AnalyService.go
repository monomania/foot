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
	"tesou.io/platform/foot-parent/foot-core/common/utils"
	service3 "tesou.io/platform/foot-parent/foot-core/module/elem/service"
	service2 "tesou.io/platform/foot-parent/foot-core/module/match/service"
	"tesou.io/platform/foot-parent/foot-core/module/odds/service"
	"time"
)

type AnalyService struct {
	mysql.BaseService
	service.EuroLastService
	service.EuroHisService
	service.AsiaLastService
	service2.MatchLastService
	service3.LeagueService
	MaxLetBall   float64
	PrintOddData bool
}

//测试加载数据
func (this *AnalyService) LoadData(matchId string) []*vo.AnalyResult {
	sql_build := strings.Builder{}
	sql_build.WriteString("SELECT ml.*,  bc.id,  bc.name AS compName, el.*  FROM t_match_last ml, t_euro_last el,  t_comp bc  WHERE ml.id = el.matchid  AND el.compid = bc.id AND ml.id =  '" + matchId + "' ")
	//结果值
	entitys := make([]*vo.AnalyResult, 0)
	//执行查询
	this.FindBySQL(sql_build.String(), &entitys)
	return entitys
}

/**
计算欧赔81 616的即时盘,和初盘的差异
*/
func (this *AnalyService) Euro_Calc() []interface{} {

	matchList := this.MatchLastService.FindAll()
	data_list_slice := make([]interface{}, 0)
	for _, v := range matchList {
		matchId := v.Id
		//声明使用变量
		var e81data *entity3.EuroLast
		var e616data *entity3.EuroLast
		var a18betData *entity3.AsiaLast
		//81 -- 伟德
		eList := this.EuroLastService.FindByMatchIdCompId(matchId, "81", "616")
		if len(eList) < 2 {
			continue
		}
		for _, ev := range eList {
			if strings.EqualFold(ev.CompId, "81") {
				e81data = ev
				continue
			}
			if strings.EqualFold(ev.CompId, "616") {
				e616data = ev
				continue
			}
		}
		//0.没有变化则跳过
		if e81data.Ep3 == e81data.Sp3 || e81data.Ep0 == e81data.Sp0 {
			continue
		}
		if e616data.Ep3 == e616data.Sp3 || e616data.Ep0 == e616data.Sp0 {
			continue
		}
		//1.有变化,进行以下逻辑
		//亚赔
		aList := this.AsiaLastService.FindByMatchIdCompId(matchId, "18Bet")
		if len(aList) < 1 {
			continue
		}
		a18betData = aList[0]
		if math.Abs(a18betData.ELetBall) > this.MaxLetBall {
			continue
		}
		//2.亚赔是主降还是主升 主降为true
		//得出结果
		var result string
		asiaMainDown := AsiaMainDown(a18betData)
		if asiaMainDown { //主降
			if (e616data.Sp3-e616data.Ep3 > e81data.Sp3-e81data.Ep3) && (e616data.Ep0 > e616data.Sp0) && (e616data.Ep0-e616data.Sp0 > e81data.Ep0-e81data.Sp0) { //主队有希望
				result = "主队"
			} else { //主队希望不大
				continue
			}
		} else { //主升
			if (e616data.Sp0-e616data.Ep0 > e81data.Sp0-e81data.Ep0) && (e616data.Ep3 > e616data.Sp3) && (e616data.Ep3-e616data.Sp3 > e81data.Ep3-e81data.Sp3) { //客队有希望
				result = "客队"
			} else { //客队希望不大
				continue
			}
		}

		//1.1准备数据
		//联赛数据
		league := this.LeagueService.FindById(v.LeagueId)
		//比赛结果
		globalResult := MatchResult(a18betData, v)
		if this.PrintOddData {
			base.Log.Info("比赛Id:" + e81data.MatchId + " e81data\tSp3:" + strconv.FormatFloat(e81data.Sp3, 'f', -1, 64) + "\t\tSp0:" + strconv.FormatFloat(e81data.Sp0, 'f', -1, 64))
			base.Log.Info("比赛Id:" + e81data.MatchId + " e81data\tEp3:" + strconv.FormatFloat(e81data.Ep3, 'f', -1, 64) + "\t\tEp0:" + strconv.FormatFloat(e81data.Ep0, 'f', -1, 64))
			base.Log.Info("比赛Id:" + e616data.MatchId + " e616data\tSp3:" + strconv.FormatFloat(e616data.Sp3, 'f', -1, 64) + "\t\tEp0:" + strconv.FormatFloat(e616data.Sp0, 'f', -1, 64))
			base.Log.Info("比赛Id:" + e616data.MatchId + " e616data\tEp3:" + strconv.FormatFloat(e616data.Ep3, 'f', -1, 64) + "\t\tEp0:" + strconv.FormatFloat(e616data.Ep0, 'f', -1, 64))
		}
		logStr := "比赛Id:" + v.Id + ",比赛时间:" + v.MatchDate + ",联赛:" + league.Name + ",对阵:" + v.MainTeamId + "(" + strconv.FormatFloat(a18betData.ELetBall,'f', -1, 64) + ")" + v.GuestTeamId + ",预算结果:" + result + ",已得结果:" + globalResult
		var resultFlag string
		if strings.Contains(globalResult, result) {
			resultFlag = "正确"
		} else if strings.Contains(globalResult, "走盘") {
			resultFlag = "走盘"
		} else if strings.Contains(globalResult, "未得") {
			resultFlag = "未得"
		} else {
			resultFlag = "错误"
		}
		logStr += "," + resultFlag
		base.Log.Info(logStr)
		analyResult := new(entity5.AnalyResult)
		analyResult.MatchId = v.Id
		analyResult.MatchDate = v.MatchDate
		analyResult.LeagueId = v.LeagueId
		analyResult.MainTeamId = v.MainTeamId
		analyResult.MainTeamGoals = v.MainTeamGoals
		analyResult.GuestTeamId = v.GuestTeamId
		analyResult.GuestTeamGoals = v.GuestTeamGoals
		format := time.Now().Format("1504")
		analyResult.AlFlag = utils.RunFuncName() + "-" + format
		analyResult.Context = logStr
		analyResult.PreResult = result
		analyResult.Result = resultFlag
		data_list_slice = append(data_list_slice, analyResult)
	}
	this.SaveList(data_list_slice)
	return data_list_slice
}

/**
分析比赛数据,, 结合亚赔 赔赔差异
( 1.欧赔降水,亚赔反之,以亚赔为准)
( 2.欧赔升水,亚赔反之,以亚赔为准)
*/
func (this *AnalyService) Euro_Asia_Diff() []interface{} {
	matchList := this.MatchLastService.FindAll()
	data_list_slice := make([]interface{}, 0)
	for _, v := range matchList {
		matchId := v.Id
		//声明使用变量
		var e81data *entity3.EuroLast
		var e616data *entity3.EuroLast
		var a18betData *entity3.AsiaLast
		//81 -- 伟德
		eList := this.EuroLastService.FindByMatchIdCompId(matchId, "81", "616")
		if len(eList) < 2 {
			continue
		}
		for _, ev := range eList {
			if strings.EqualFold(ev.CompId, "81") {
				e81data = ev
				continue
			}
			if strings.EqualFold(ev.CompId, "616") {
				e616data = ev
				continue
			}
		}

		//亚赔
		aList := this.AsiaLastService.FindByMatchIdCompId(matchId, "18Bet")
		if len(aList) < 1 {
			continue
		}
		a18betData = aList[0]
		if a18betData.ELetBall > this.MaxLetBall {
			continue
		}

		//判断分析logic
		//1.欧赔是主降还是主升 主降为true
		euroMainDown := EuroMainDown(e81data, e616data)
		//2.亚赔是主降还是主升 主降为true
		asiaMainDown := AsiaMainDown(a18betData)
		//得出结果
		var result string
		if euroMainDown == 3 && !asiaMainDown {
			result = "客队"
		} else if euroMainDown == 0 && asiaMainDown {
			result = "主队"
		} else {
			continue
		}
		//联赛数据
		league := this.LeagueService.FindById(v.LeagueId)
		//比赛结果
		globalResult := MatchResult(a18betData, v)

		if this.PrintOddData {
			base.Log.Info("比赛Id:" + e81data.MatchId + " e81data\tSp3:" + strconv.FormatFloat(e81data.Sp3, 'f', -1, 64) + "\t\tSp0:" + strconv.FormatFloat(e81data.Sp0, 'f', -1, 64))
			base.Log.Info("比赛Id:" + e81data.MatchId + " e81data\tEp3:" + strconv.FormatFloat(e81data.Ep3, 'f', -1, 64) + "\t\tEp0:" + strconv.FormatFloat(e81data.Ep0, 'f', -1, 64))
			base.Log.Info("比赛Id:" + e616data.MatchId + " e616data\tSp3:" + strconv.FormatFloat(e616data.Sp3, 'f', -1, 64) + "\t\tEp0:" + strconv.FormatFloat(e616data.Sp0, 'f', -1, 64))
			base.Log.Info("比赛Id:" + e616data.MatchId + " e616data\tEp3:" + strconv.FormatFloat(e616data.Ep3, 'f', -1, 64) + "\t\tEp0:" + strconv.FormatFloat(e616data.Ep0, 'f', -1, 64))
		}
		logStr := "比赛Id:" + v.Id + ",比赛时间:" + v.MatchDate + ",联赛:" + league.Name + ",对阵:" + v.MainTeamId + "(" + strconv.FormatFloat(a18betData.ELetBall,'f', -1, 64) + ")" + v.GuestTeamId + ",预算结果:" + result + ",已得结果:" + globalResult
		var resultFlag string
		if strings.Contains(globalResult, result) {
			resultFlag = "正确"
		} else if strings.Contains(globalResult, "走盘") {
			resultFlag = "走盘"
		} else if strings.Contains(globalResult, "未得") {
			resultFlag = "未得"
		} else {
			resultFlag = "错误"
		}
		logStr += "," + resultFlag
		base.Log.Info(logStr)
		analyResult := new(entity5.AnalyResult)
		analyResult.MatchId = v.Id
		analyResult.MatchDate = v.MatchDate
		analyResult.LeagueId = v.LeagueId
		analyResult.MainTeamId = v.MainTeamId
		analyResult.MainTeamGoals = v.MainTeamGoals
		analyResult.GuestTeamId = v.GuestTeamId
		analyResult.GuestTeamGoals = v.GuestTeamGoals
		format := time.Now().Format("1504")
		analyResult.AlFlag = utils.RunFuncName() + "-" + format
		analyResult.Context = logStr
		analyResult.PreResult = result
		analyResult.Result = resultFlag
		data_list_slice = append(data_list_slice, analyResult)
	}
	this.SaveList(data_list_slice)
	return data_list_slice
}

func MatchResult(last *entity3.AsiaLast, v *entity2.MatchLast) string {
	var result string
	h2, _ := time.ParseDuration("2h")
	local, _ := time.LoadLocation("Local")
	matchDate, _ := time.ParseInLocation("2006-01-02 15:04:05", v.MatchDate, local)
	matchDate = matchDate.Add(h2)
	nowDate := time.Now()
	if matchDate.After(nowDate) { //比赛是否已经结束
		result = "未得(" + strconv.Itoa(v.MainTeamGoals) + "-" + strconv.Itoa(v.GuestTeamGoals) + ")"
		return result
	}

	elb_sum := last.ELetBall
	var mainTeamGoals float64
	if elb_sum > 0 {
		mainTeamGoals = float64(v.MainTeamGoals) - elb_sum
	} else {
		mainTeamGoals = float64(v.MainTeamGoals) + math.Abs(elb_sum)
	}
	diff_goals := float64((v.MainTeamGoals - v.GuestTeamGoals)) - elb_sum
	if diff_goals <= 0.25 && diff_goals >= -0.25 {
		result = "走盘(" + strconv.Itoa(v.MainTeamGoals) + "-" + strconv.Itoa(v.GuestTeamGoals) + ")"
	} else if mainTeamGoals > float64(v.GuestTeamGoals) {
		result = "主队(" + strconv.Itoa(v.MainTeamGoals) + "-" + strconv.Itoa(v.GuestTeamGoals) + ")"
	} else if mainTeamGoals < float64(v.GuestTeamGoals) {
		result = "客队(" + strconv.Itoa(v.MainTeamGoals) + "-" + strconv.Itoa(v.GuestTeamGoals) + ")"
	} else {
		result = "走盘(" + strconv.Itoa(v.MainTeamGoals) + "-" + strconv.Itoa(v.GuestTeamGoals) + ")"
	}
	return result
}

/**
1.欧赔是主降还是主升 主降为true
*/
func EuroMainDown(e81data *entity3.EuroLast, e616data *entity3.EuroLast) int {
	if e81data.Ep3 < e81data.Sp3 && e616data.Ep3 < e616data.Sp3 {
		return 3
	} else if e81data.Ep0 < e81data.Sp0 && e616data.Ep0 < e616data.Sp0 {
		return 0
	}
	return 1
}


/**
2.亚赔是主降还是主升 主降为true
*/
func AsiaMainDown(a18betData *entity3.AsiaLast) bool {
	slb_sum := a18betData.SLetBall
	elb_sum := a18betData.ELetBall

	if elb_sum > slb_sum {
		return true
	} else if elb_sum < slb_sum {
		return false
	} else { //初始让球和即时让球一致
		if a18betData.Ep3 < a18betData.Sp3 {
			return true
		}
	}
	return false
}

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

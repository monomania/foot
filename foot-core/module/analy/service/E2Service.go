package service

import (
	"math"
	"strconv"
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	entity5 "tesou.io/platform/foot-parent/foot-api/module/analy/pojo"
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	entity3 "tesou.io/platform/foot-parent/foot-api/module/odds/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/utils"
	"tesou.io/platform/foot-parent/foot-core/module/analy/constants"
	"time"
)

type E2Service struct {
	AnalyService
	//最大让球数据
	MaxLetBall float64
}

func (this *E2Service) ModelName() string{
	return "E2"
}

/**
计算欧赔81 616的即时盘,和初盘的差异
*/
func (this *E2Service) Analy() {
	matchList := this.MatchLastService.FindNotFinished()
	data_list_slice := make([]interface{}, 0)
	data_modify_list_slice := make([]interface{}, 0)
	for _, v := range matchList {
		stub, data := this.analyStub(v)

		if stub == 0 || stub == 1 {
			hours := v.MatchDate.Sub(time.Now()).Hours()
			if hours > 0 {
				hours = math.Abs(hours * 0.5)
				data.THitCount = int(hours)
			} else {
				data.THitCount = 1
			}
			if stub == 0 {
				data_list_slice = append(data_list_slice, data)
			} else if stub == 1 {
				data_modify_list_slice = append(data_modify_list_slice, data)
			}
		} else {
			temp_data := this.Find(v.Id, this.ModelName())
			if len(temp_data.Id) > 0 {
				hit_count_str := utils.GetVal(constants.SECTION_NAME, "hit_count")
				hit_count, _ := strconv.Atoi(hit_count_str)
				if temp_data.HitCount >= hit_count {
					temp_data.HitCount = (hit_count / 2) - 1
				} else {
					temp_data.HitCount = 0
				}
				this.AnalyService.Modify(temp_data)
			}
		}
	}
	this.AnalyService.SaveList(data_list_slice)
	this.AnalyService.ModifyList(data_modify_list_slice)

}

func (this *E2Service) analyStub(v *pojo.MatchLast) (int, *entity5.AnalyResult) {
	matchId := v.Id
	//声明使用变量
	var e616data *entity3.EuroLast
	var e104data *entity3.EuroLast
	var a18betData *entity3.AsiaLast
	//81 -- 伟德
	eList := this.EuroLastService.FindByMatchIdCompId(matchId, "616", "104")
	if len(eList) < 2 {
		return -1, nil
	}
	for _, ev := range eList {
		if strings.EqualFold(ev.CompId, "616") {
			e616data = ev
			continue
		}
		if strings.EqualFold(ev.CompId, "104") {
			e104data = ev
			continue
		}
	}
	//0.没有变化则跳过
	if e104data.Ep3 == e104data.Sp3 || e104data.Ep0 == e104data.Sp0 {
		return -3, nil
	}
	if e616data.Ep3 == e616data.Sp3 || e616data.Ep0 == e616data.Sp0 {
		return -3, nil
	}

	//1.有变化,进行以下逻辑
	//亚赔
	aList := this.AsiaLastService.FindByMatchIdCompId(matchId, "18Bet")
	if len(aList) < 1 {
		return -1, nil
	}
	a18betData = aList[0]
	if math.Abs(a18betData.ELetBall) > this.MaxLetBall {
		return -2, nil
	}

	//得出结果
	if e616data.Ep0 < e616data.Sp0 && e616data.Ep0 < e104data.Sp0 {
		return -3, nil
	}
	var preResult int
	if e616data.Ep3 > (e616data.Sp3+0.01) && e104data.Ep3 < e104data.Sp3 {
		preResult = 3
	} else if e616data.Ep3 < e616data.Sp3 && e104data.Ep3 < e104data.Sp3 && e616data.Ep3 < e104data.Ep3 {
		preResult = 3
	} else {
		return -3, nil
	}

	var data *entity5.AnalyResult
	temp_data := this.Find(matchId, this.ModelName())
	if len(temp_data.Id) > 0 {
		temp_data.PreResult = preResult
		temp_data.HitCount = temp_data.HitCount + 1
		temp_data.LetBall = a18betData.ELetBall
		data = temp_data
		//比赛结果
		data.Result = this.IsRight(a18betData, v, data)
		return 1, data
	} else {
		data = new(entity5.AnalyResult)
		data.MatchId = v.Id
		data.MatchDate = v.MatchDate
		data.LetBall = a18betData.ELetBall
		data.AlFlag = this.ModelName()
		format := time.Now().Format("0102150405")
		data.AlSeq = format
		data.PreResult = preResult
		data.HitCount = 1
		data.LetBall = a18betData.ELetBall
		//比赛结果
		data.Result = this.IsRight(a18betData, v, data)
		return 0, data
	}
}

func (this *E2Service) IsRight(last *entity3.AsiaLast, v *pojo.MatchLast, analy *entity5.AnalyResult) string {
	//比赛结果
	var globalResult int
	h2, _ := time.ParseDuration("148m")
	matchDate := v.MatchDate.Add(h2)
	if matchDate.After(time.Now()) {
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
		resultFlag = "待定"
	} else if globalResult == analy.PreResult || globalResult == 1{
		resultFlag = "正确"
	} else {
		resultFlag = "错误"
	}

	//打印数据
	league := this.LeagueService.FindById(v.LeagueId)
	matchDateStr := v.MatchDate.Format("2006-01-02 15:04:05")
	base.Log.Info("比赛Id:" + v.Id + ",比赛时间:" + matchDateStr + ",联赛:" + league.Name + ",对阵:" + v.MainTeamId + "(" + strconv.FormatFloat(last.ELetBall, 'f', -1, 64) + ")" + v.GuestTeamId + ",预算结果:" + strconv.Itoa(analy.PreResult) + ",已得结果:" + strconv.Itoa(v.MainTeamGoals) + "-" + strconv.Itoa(v.GuestTeamGoals) + " (" + resultFlag + ")")
	return resultFlag
}

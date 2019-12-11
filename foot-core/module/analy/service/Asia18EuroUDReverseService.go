package service

import (
	"reflect"
	"strconv"
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	entity5 "tesou.io/platform/foot-parent/foot-api/module/analy/pojo"
	entity2 "tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	entity3 "tesou.io/platform/foot-parent/foot-api/module/odds/pojo"
	"time"
)

/**
亚赔与欧赔up down 颠倒
 */
type Asia18EuroUDReverseService struct {
	AnalyService
	//最大让球数据
	MaxLetBall float64
}

/**
分析比赛数据,, 结合亚赔 赔赔差异
( 1.欧赔降水,亚赔反之,以亚赔为准)
( 2.欧赔升水,亚赔反之,以亚赔为准)
*/
func (this *Asia18EuroUDReverseService) Analy() []interface{} {
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
		var result int
		if euroMainDown == 3 && !asiaMainDown {
			result = 0
		} else if euroMainDown == 0 && asiaMainDown {
			result = 3
		} else {
			continue
		}
		data := this.buildData(v, a18betData, e81data, e616data, result)
		temp_data := this.Find(data.MatchId)
		if len(temp_data.Id) > 0 {
			data.LeisuPubd = temp_data.LeisuPubd
		}
		data_list_slice = append(data_list_slice, data)
	}
	this.AnalyService.SaveList(data_list_slice)
	return data_list_slice
}

func (this *Asia18EuroUDReverseService) buildData(v *entity2.MatchLast, a18betData *entity3.AsiaLast, e81data *entity3.EuroLast, e616data *entity3.EuroLast, result int) *entity5.AnalyResult {
	//比赛结果
	globalResult := this.ActualResult(a18betData, v)
	if this.PrintOddData {
		base.Log.Info("比赛Id:" + e81data.MatchId + " e81data\tSp3:" + strconv.FormatFloat(e81data.Sp3, 'f', -1, 64) + "\t\tSp0:" + strconv.FormatFloat(e81data.Sp0, 'f', -1, 64))
		base.Log.Info("比赛Id:" + e81data.MatchId + " e81data\tEp3:" + strconv.FormatFloat(e81data.Ep3, 'f', -1, 64) + "\t\tEp0:" + strconv.FormatFloat(e81data.Ep0, 'f', -1, 64))
		base.Log.Info("比赛Id:" + e616data.MatchId + " e616data\tSp3:" + strconv.FormatFloat(e616data.Sp3, 'f', -1, 64) + "\t\tEp0:" + strconv.FormatFloat(e616data.Sp0, 'f', -1, 64))
		base.Log.Info("比赛Id:" + e616data.MatchId + " e616data\tEp3:" + strconv.FormatFloat(e616data.Ep3, 'f', -1, 64) + "\t\tEp0:" + strconv.FormatFloat(e616data.Ep0, 'f', -1, 64))
	}
	var resultFlag string
	if globalResult == -1 {
		resultFlag = "待定"
	} else if globalResult == result {
		resultFlag = "正确"
	} else if globalResult == 1 {
		resultFlag = "走盘"
	} else {
		resultFlag = "错误"
	}
	analyResult := new(entity5.AnalyResult)
	analyResult.MatchId = v.Id
	analyResult.MatchDate = v.MatchDate
	format := time.Now().Format("1504")
	analyResult.AlFlag = reflect.TypeOf(*this).Name() + "-" + format
	analyResult.PreResult = result
	analyResult.Result = resultFlag

	//打印数据
	league := this.LeagueService.FindById(v.LeagueId)
	matchDate := v.MatchDate.Format("2006-01-02 15:04:05")
	base.Log.Info("比赛Id:" + v.Id + ",比赛时间:" + matchDate + ",联赛:" + league.Name + ",对阵:" + v.MainTeamId + "(" + strconv.FormatFloat(a18betData.ELetBall, 'f', -1, 64) + ")" + v.GuestTeamId + ",预算结果:" + strconv.Itoa(result) + ",已得结果:" + strconv.Itoa(v.MainTeamGoals) + "-" + strconv.Itoa(v.GuestTeamGoals) + " (" + resultFlag + ")")
	return analyResult
}

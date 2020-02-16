package service

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	entity5 "tesou.io/platform/foot-parent/foot-api/module/analy/pojo"
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	entity3 "tesou.io/platform/foot-parent/foot-api/module/odds/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/utils"
	"tesou.io/platform/foot-parent/foot-core/module/analy/constants"
	"tesou.io/platform/foot-parent/foot-core/module/match/service"
	"time"
)

type C2Service struct {
	AnalyService
	service.BFBattleService
	service.BFJinService

	//最大让球数据
	MaxLetBall float64
}

func (this *C2Service) ModelName() string {
	//alFlag := reflect.TypeOf(*this).Name()
	return "C2"
}

func (this *C2Service) Analy(analyAll bool) {
	var matchLasts []*pojo.MatchLast
	if analyAll {
		matchHis := this.MatchHisService.FindAll()
		for _, e := range matchHis {
			matchLasts = append(matchLasts, &e.MatchLast)
		}
		//matchLasts = this.MatchLastService.FindAll()
	} else {
		matchLasts = this.MatchLastService.FindNotFinished()
	}
	this.Analy_Process(matchLasts)

}

func (this *C2Service) Analy_Near() {
	matchList := this.MatchLastService.FindNear()
	this.Analy_Process(matchList)
}

func (this *C2Service) Analy_Process(matchList []*pojo.MatchLast) {
	hit_count_str := utils.GetVal(constants.SECTION_NAME, "hit_count")
	hit_count, _ := strconv.Atoi(hit_count_str)
	data_list_slice := make([]interface{}, 0)
	data_modify_list_slice := make([]interface{}, 0)
	var rightCount = 0
	var errorCount = 0

	for _, v := range matchList {
		stub, data := this.analyStub(v)
		if nil != data {
			if strings.EqualFold(data.Result, "命中") {
				rightCount++
			}
			if strings.EqualFold(data.Result, "错误") {
				errorCount++
			}
		}

		if stub == 0 || stub == 1 {
			data.TOVoid = false
			hours := v.MatchDate.Sub(time.Now()).Hours()
			if hours > 0 {
				data.THitCount = hit_count
			} else {
				data.THitCount = 1
			}
			//如其他模型存在互斥选项，设置为作废
			diff_preResult := this.FindOtherAlFlag(data.MatchId, data.AlFlag, data.PreResult)
			if diff_preResult {
				data.TOVoid = true
				data.TOVoidDesc = "与其他模型互斥"
			}

			if stub == 0 {
				data_list_slice = append(data_list_slice, data)
			} else if stub == 1 {
				data_modify_list_slice = append(data_modify_list_slice, data)
			}
		} else {
			if stub != -2 {
				data = this.Find(v.Id, this.ModelName())
			}
			data.TOVoid = true
			if len(data.Id) > 0 {
				if data.HitCount >= hit_count {
					data.HitCount = (hit_count / 2) - 1
				} else {
					data.HitCount = 0
				}
				this.AnalyService.Modify(data)
			}
		}
	}
	base.Log.Info("------------------")
	base.Log.Info("------------------")
	base.Log.Info("------------------")
	base.Log.Info("GOOOO场次:", rightCount)
	base.Log.Info("X0000场次:", errorCount)
	base.Log.Info("------------------")

	this.AnalyService.SaveList(data_list_slice)
	this.AnalyService.ModifyList(data_modify_list_slice)
}

/**
  -1 参数错误
  -2 不符合让球数
  -3 计算分析错误
  0  新增的分析结果
  1  需要更新结果
 */
func (this *C2Service) analyStub(v *pojo.MatchLast) (int, *entity5.AnalyResult) {
	matchId := v.Id
	if matchId == "1770574" {
		base.Log.Info("-")
	}
	//声明使用变量
	var a18betData *entity3.AsiaHis
	//亚赔
	aList := this.AsiaHisService.FindByMatchIdCompId(matchId, constants.C1_REFER_ASIA)
	if len(aList) < 1 {
		return -1, nil
	}
	a18betData = aList[0]
	if math.Abs(a18betData.ELetBall) > this.MaxLetBall {
		temp_data := this.Find(v.Id, this.ModelName())
		temp_data.LetBall = a18betData.ELetBall
		return -2, temp_data
	}
	if matchId == "1770574"{
		fmt.Println("--------------")
	}
	preResult := -1
	bfb_arr := this.BFBattleService.FindNearByMatchId(matchId, 1)
	mainWin := 0
	guestWin := 0
	for _, e := range bfb_arr {
		if e.BattleMainTeamId == v.MainTeamId && e.BattleMainTeamGoals > e.BattleGuestTeamGoals {
			mainWin++
		}
		if e.BattleGuestTeamId == v.MainTeamId && e.BattleGuestTeamGoals > e.BattleMainTeamGoals {
			mainWin++
		}
		if e.BattleMainTeamId == v.GuestTeamId && e.BattleMainTeamGoals > e.BattleGuestTeamGoals {
			guestWin++
		}
		if e.BattleGuestTeamId == v.GuestTeamId && e.BattleGuestTeamGoals > e.BattleMainTeamGoals {
			guestWin++
		}
	}
	if mainWin > guestWin {
		preResult = 3
	} else if guestWin > mainWin {
		preResult = 0
	} else {
		return -3, nil
	}

	preResult2 := -1
	if preResult == 3 {
		bfj_main := this.BFJinService.FindNearByTeamName(v.MainTeamId, 1)
		for _, e := range bfj_main {
			if e.HomeTeam == v.MainTeamId && e.HomeScore > e.GuestScore {
				preResult2 = 3
			}
			if e.GuestTeam == v.MainTeamId && e.GuestScore > e.HomeScore {
				preResult2 = 3
			}
		}
	} else {
		bfj_guest := this.BFJinService.FindNearByTeamName(v.GuestTeamId, 1)
		for _, e := range bfj_guest {
			if e.HomeTeam == v.GuestTeamId && e.HomeScore > e.GuestScore {
				preResult2 = 0
			}
			if e.GuestTeam == v.GuestTeamId && e.GuestScore > e.HomeScore {
				preResult2 = 0
			}
		}
	}

	if preResult != preResult2 {
		return -3, nil
	}

	var data *entity5.AnalyResult
	temp_data := this.Find(v.Id, this.ModelName())
	if len(temp_data.Id) > 0 {
		temp_data.PreResult = preResult
		temp_data.HitCount = temp_data.HitCount + 1
		temp_data.LetBall = a18betData.ELetBall
		data = temp_data
		//比赛结果
		data.Result = this.IsRight(v, data)
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
		data.HitCount = 10
		data.LetBall = a18betData.ELetBall
		//比赛结果
		data.Result = this.IsRight(v, data)
		return 0, data
	}
}

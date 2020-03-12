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
	"tesou.io/platform/foot-parent/foot-core/module/match/service"
	"time"
)

type E3Service struct {
	AnalyService
	service.BFScoreService
	service.BFBattleService
	service.BFJinService
	service.BFFutureEventService
	//最大让球数据
	MaxLetBall float64
}

func (this *E3Service) ModelName() string {
	return "E3"
}

/**
计算欧赔81 616的即时盘,和初盘的差异
*/
func (this *E3Service) Analy(analyAll bool) {
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

/**
计算欧赔81 616的即时盘,和初盘的差异
*/
func (this *E3Service) Analy_Near() {
	matchList := this.MatchLastService.FindNear()
	this.Analy_Process(matchList)
}

func (this *E3Service) Analy_Process(matchList []*pojo.MatchLast) {
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
func (this *E3Service) analyStub(v *pojo.MatchLast) (int, *entity5.AnalyResult) {
	matchId := v.Id

	//未来赛事
	bffe_main := this.BFFutureEventService.FindNextBattle(matchId, v.MainTeamId)
	bffe_guest := this.BFFutureEventService.FindNextBattle(matchId, v.GuestTeamId)

	if this.IsCupMatch(bffe_main.EventLeagueId) || this.IsCupMatch(bffe_guest.EventLeagueId) {
		//下一场打杯赛
		return -3, nil
	}

	//声明使用变量
	var e281_1 *entity3.EuroTrack
	var e281_2 *entity3.EuroTrack
	var aBet365 *entity3.AsiaHis
	eList := this.EuroTrackService.FindByMatchIdCompId(matchId, "281")
	if len(eList) < 2 {
		return -1, nil
	}
	e281_1 = eList[0]
	e281_2 = eList[1]

	//1.有变化,进行以下逻辑
	//亚赔
	aList := this.AsiaHisService.FindByMatchIdCompId(matchId, "Bet365")
	if len(aList) < 1 {
		return -1, nil
	}
	aBet365 = aList[0]
	if math.Abs(aBet365.ELetBall) > this.MaxLetBall {
		temp_data := this.Find(v.Id, this.ModelName())
		temp_data.LetBall = aBet365.ELetBall
		return -2, temp_data
	}


	//得出结果
	//1.0判断主队是否是让球方
	mainLetball := true
	if aBet365.ELetBall > 0 {
		mainLetball = true
	} else if aBet365.ELetBall < 0 {
		mainLetball = false
	} else {
		//EletBall == 0
		//通过赔率确立
		if aBet365.Ep3 > aBet365.Ep0 {
			mainLetball = false
		} else {
			mainLetball = true
		}
	}

	preResult := -1
	if mainLetball{
		if aBet365.Ep3 > aBet365.Sp3 && e281_1.Ep3 < e281_1.Sp3 && e281_1.Kelly3 > e281_2.Kelly3{
			preResult = 0
		}
		if aBet365.Ep3 > aBet365.Sp3 && e281_1.Ep3 < e281_1.Sp3 && e281_1.Kelly3 < e281_2.Kelly3{
			preResult = 3
		}
		if aBet365.Ep3 < aBet365.Sp3 && e281_1.Ep3 < e281_1.Sp3 && e281_1.Kelly3 > e281_2.Kelly3{
			preResult = 0
		}
	}else{
		if aBet365.Ep0 > aBet365.Sp0 && e281_1.Ep0 < e281_1.Sp0 && e281_1.Kelly0 > e281_2.Kelly0{
			preResult = 3
		}
		if aBet365.Ep0 > aBet365.Sp0 && e281_1.Ep0 < e281_1.Sp0 && e281_1.Kelly0 < e281_2.Kelly0{
			preResult = 0
		}
		if aBet365.Ep0 < aBet365.Sp0 && e281_1.Ep0 < e281_1.Sp0 && e281_1.Kelly0 > e281_2.Kelly0{
			preResult = 3
		}
	}

	var data *entity5.AnalyResult
	temp_data := this.Find(matchId, this.ModelName())
	if len(temp_data.Id) > 0 {
		temp_data.PreResult = preResult
		temp_data.HitCount = temp_data.HitCount + 1
		temp_data.LetBall = aBet365.ELetBall
		data = temp_data
		//比赛结果
		data.Result = this.IsRight(v, data)
		return 1, data
	} else {
		data = new(entity5.AnalyResult)
		data.MatchId = v.Id
		data.MatchDate = v.MatchDate
		data.LetBall = aBet365.ELetBall
		data.AlFlag = this.ModelName()
		format := time.Now().Format("0102150405")
		data.AlSeq = format
		data.PreResult = preResult
		data.HitCount = 3
		data.LetBall = aBet365.ELetBall
		//比赛结果
		data.Result = this.IsRight(v, data)
		return 0, data
	}
}


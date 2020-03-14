package service

import (
	"math"
	entity5 "tesou.io/platform/foot-parent/foot-api/module/analy/pojo"
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	entity3 "tesou.io/platform/foot-parent/foot-api/module/odds/pojo"
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
	this.AnalyService.Analy(analyAll,this)
}

/**
计算欧赔81 616的即时盘,和初盘的差异
*/
func (this *E3Service) Analy_Near() {
	this.AnalyService.Analy_Near(this)
}

/**
  -1 参数错误
  -2 不符合让球数
  -3 计算分析错误
  0  新增的分析结果
  1  需要更新结果
 */
func (this *E3Service) analyStub(v *pojo.MatchLast) (int, *entity5.AnalyResult) {
	temp_data := this.Find(v.Id, this.ModelName())
	matchId := v.Id

	//未来赛事
	bffe_main := this.BFFutureEventService.FindNextBattle(matchId, v.MainTeamId)
	bffe_guest := this.BFFutureEventService.FindNextBattle(matchId, v.GuestTeamId)

	if this.IsCupMatch(bffe_main.EventLeagueId) || this.IsCupMatch(bffe_guest.EventLeagueId) {
		//下一场打杯赛
		return -3, temp_data
	}

	//声明使用变量
	var e281_1 *entity3.EuroTrack
	var e281_2 *entity3.EuroTrack
	var aBet365 *entity3.AsiaHis
	eList := this.EuroTrackService.FindByMatchIdCompId(matchId, "281")
	if len(eList) < 2 {
		return -1, temp_data
	}
	e281_1 = eList[0]
	e281_2 = eList[1]

	//1.有变化,进行以下逻辑
	//亚赔
	aList := this.AsiaHisService.FindByMatchIdCompId(matchId, "Bet365")
	if len(aList) < 1 {
		return -1, temp_data
	}
	aBet365 = aList[0]
	temp_data.LetBall = aBet365.ELetBall
	if math.Abs(aBet365.ELetBall) > this.MaxLetBall {
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
		data.SLetBall = a18Bet.SLetBall
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


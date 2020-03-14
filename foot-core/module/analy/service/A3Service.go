package service

import (
	"math"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	entity5 "tesou.io/platform/foot-parent/foot-api/module/analy/pojo"
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	entity3 "tesou.io/platform/foot-parent/foot-api/module/odds/pojo"
	"tesou.io/platform/foot-parent/foot-core/module/analy/constants"
	"tesou.io/platform/foot-parent/foot-core/module/match/service"
	"time"
)

type A3Service struct {
	AnalyService
	service.BFScoreService
	service.BFBattleService
	service.BFJinService
	service.BFFutureEventService

	//最大让球数据
	MaxLetBall float64
}

func (this *A3Service) ModelName() string {
	//alFlag := reflect.TypeOf(*this).Name()
	return "A3"
}

func (this *A3Service) Analy(analyAll bool) {
	this.AnalyService.Analy(analyAll,this)
}

func (this *A3Service) Analy_Near() {
	this.AnalyService.Analy_Near(this)
}

/**
  -1 参数错误
  -2 不符合让球数
  -3 计算分析错误
  0  新增的分析结果
  1  需要更新结果
 */
func (this *A3Service) analyStub(v *pojo.MatchLast) (int, *entity5.AnalyResult) {
	temp_data := this.Find(v.Id, this.ModelName())
	matchId := v.Id
	if matchId == "1836932" {
		base.Log.Info("-")
	}
	//声明使用变量
	var a18bet *entity3.AsiaHis
	//亚赔
	aList := this.AsiaHisService.FindByMatchIdCompId(matchId, constants.C1_REFER_ASIA)
	if len(aList) < 1 {
		return -1, temp_data
	}
	a18bet = aList[0]
	temp_data.LetBall = a18bet.ELetBall
	if math.Abs(a18bet.ELetBall) > this.MaxLetBall {
		//temp_data.Result =""
		//return -2, temp_data
		//return -2, nil
	}

	//限制初盘,即时盘让球在0.25以内
	sLetBall := math.Abs(a18bet.SLetBall)
	eLetBall := math.Abs(a18bet.ELetBall)
	if math.Abs(sLetBall-eLetBall) > 0.25 {
		//temp_data.Result =""
		//return -2, temp_data
		//return -2, nil
	}

	//得出结果
	preResult := -1
	//------
	bfs_arr := this.BFScoreService.FindByMatchId(matchId)
	if len(bfs_arr) < 1 {
		return -1, temp_data
	}

	var mainZongBfs *pojo.BFScore
	var mainZhuBfs *pojo.BFScore
	var guestZongBfs *pojo.BFScore
	var guestKeBfs *pojo.BFScore
	for _, e := range bfs_arr { //bfs_arr有多语言版本,条数很多
		if e.TeamId == v.MainTeamId {
			if e.Type == "总" {
				mainZongBfs = e
			}
			if e.Type == "主" {
				mainZhuBfs = e
			}
		} else if e.TeamId == v.GuestTeamId {
			if e.Type == "总" {
				guestZongBfs = e
			}
			if e.Type == "客" {
				guestKeBfs = e
			}
		}
	}
	if mainZongBfs == nil || guestZongBfs == nil || mainZhuBfs == nil || guestKeBfs == nil {
		return -1, temp_data
	}

	//3-主队强 0-客队强
	mainStrong := -1
	rankDiff := 5.0
	if mainZongBfs.MatchCount >= 8 && guestZongBfs.MatchCount >= 8 {
		var temp_val_1 float64
		var temp_val_2 float64
		//排名越小越强
		temp_val_1 = float64(mainZongBfs.Ranking - guestZongBfs.Ranking)
		temp_val_2 = float64(mainZhuBfs.Ranking - guestKeBfs.Ranking)
		if temp_val_1 >= rankDiff && temp_val_2 >= rankDiff{
			mainStrong = 0
		}
		temp_val_1 = float64(guestZongBfs.Ranking - mainZongBfs.Ranking)
		temp_val_2 = float64(guestKeBfs.Ranking - mainZhuBfs.Ranking)
		if temp_val_1 >= rankDiff && temp_val_2 >= rankDiff{
			mainStrong = 3
		}
	} else {
		//return -1, temp_data
	}

	//------
	//未来赛事
	bffe_main := this.BFFutureEventService.FindNextBattle(matchId, v.MainTeamId)
	bffe_guest := this.BFFutureEventService.FindNextBattle(matchId, v.GuestTeamId)

	if this.IsCupMatch(bffe_main.EventLeagueId) || this.IsCupMatch(bffe_guest.EventLeagueId) {
		//下一场打杯赛
		return -3, temp_data
	}

	//1.0判断主队是否是让球方
	mainLetball := true
	if a18bet.ELetBall > 0 {
		mainLetball = true
	} else if a18bet.ELetBall < 0 {
		mainLetball = false
	} else {
		//EletBall == 0
		//通过赔率确立
		if a18bet.Ep3 > a18bet.Ep0 {
			mainLetball = false
		} else {
			mainLetball = true
		}
	}

	odd_flag_1 := a18bet.Ep3 >= 1.00 || (a18bet.Ep3 >= 0.94 && a18bet.Ep3 < 0.96)
	//odd_flag_1 := a18bet.Ep3 < 0.94 && a18bet.Sp3 < 0.94
	odd_flag_2 := a18bet.Ep0 >= 1.00 || (a18bet.Ep0 >= 0.94 && a18bet.Ep0 < 0.96)
	//odd_flag_2 := a18bet.Ep0 < 0.94 && a18bet.Sp0 < 0.94
	if mainStrong == 3 && mainLetball && odd_flag_1 && a18bet.ELetBall == a18bet.SLetBall{
		preResult = 0
	}
	if mainStrong == 0 && !mainLetball &&  odd_flag_2 && a18bet.ELetBall == a18bet.SLetBall{
		preResult = 3
	}

	if preResult < 0 {
		return -3, temp_data
	}
	base.Log.Info(a18bet.Sp3," ",a18bet.SLetBall," ",a18bet.Sp0,"   ",a18bet.Ep3," ",a18bet.ELetBall," ",a18bet.Ep0)

	var data *entity5.AnalyResult
	if len(temp_data.Id) > 0 {
		temp_data.PreResult = preResult
		temp_data.HitCount = temp_data.HitCount + 1
		temp_data.LetBall = a18bet.ELetBall
		data = temp_data
		//比赛结果
		data.Result = this.IsRight(v, data)
		return 1, data
	} else {
		data = new(entity5.AnalyResult)
		data.MatchId = v.Id
		data.MatchDate = v.MatchDate
		data.LetBall = a18bet.ELetBall
		data.AlFlag = this.ModelName()
		format := time.Now().Format("0102150405")
		data.AlSeq = format
		data.PreResult = preResult
		data.HitCount = 3
		data.LetBall = a18bet.ELetBall
		//比赛结果
		data.Result = this.IsRight(v, data)
		return 0, data
	}
}

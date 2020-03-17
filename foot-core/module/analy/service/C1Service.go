package service

import (
	"fmt"
	"math"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	entity5 "tesou.io/platform/foot-parent/foot-api/module/analy/pojo"
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	entity3 "tesou.io/platform/foot-parent/foot-api/module/odds/pojo"
	"tesou.io/platform/foot-parent/foot-core/module/analy/constants"
	"tesou.io/platform/foot-parent/foot-core/module/match/service"
	"time"
)

type C1Service struct {
	AnalyService
	service.BFScoreService
	service.BFBattleService
	service.BFFutureEventService
	service.BFJinService

	//最大让球数据
	MaxLetBall float64
}

func (this *C1Service) ModelName() string {
	//alFlag := reflect.TypeOf(*this).Name()
	return "C1"
}

func (this *C1Service) AnalyTest() {
	this.AnalyService.AnalyTest(this)
}

func (this *C1Service) Analy(analyAll bool) {
	this.AnalyService.Analy(analyAll,this)

}

func (this *C1Service) Analy_Near() {
	this.AnalyService.Analy_Near(this)
}

/**
  -1 参数错误
  -2 不符合让球数
  -3 计算分析错误
  0  新增的分析结果
  1  需要更新结果
 */
func (this *C1Service) analyStub(v *pojo.MatchLast) (int, *entity5.AnalyResult) {
	temp_data := this.Find(v.Id, this.ModelName())
	matchId := v.Id
	if matchId == "1779520" {
		fmt.Println("-------------------")
	}
	//声明使用变量
	var a18Bet *entity3.AsiaHis
	//亚赔
	aList := this.AsiaHisService.FindByMatchIdCompId(matchId, constants.C1_REFER_ASIA)
	if len(aList) < 1 {
		return -1, temp_data
	}
	a18Bet = aList[0]
	temp_data.LetBall = a18Bet.ELetBall
	if math.Abs(a18Bet.ELetBall) > this.MaxLetBall {
		//temp_data.Result = ""
		return -2, temp_data
	}

	//限制初盘,即时盘让球在0.25以内
	if math.Abs(a18Bet.SLetBall-a18Bet.ELetBall) > 0.25 {
		//temp_data.Result = ""
		//return -2, temp_data
	}

	//得出结果
	preResult := -1
	letBall := 0.00
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
	baseVal := 0.25
	xishu := 5.0
	//积分排名----排名越小越好
	if mainZongBfs.MatchCount >= 8 && guestZongBfs.MatchCount >= 8 {
		temp_val := math.Abs(float64(mainZongBfs.Ranking - guestZongBfs.Ranking))
		if temp_val > xishu {
			if mainZongBfs.Ranking < guestZongBfs.Ranking {
				letBall += baseVal * (temp_val / xishu)
			} else {
				letBall -= baseVal * (temp_val / xishu)
			}
		}
		temp_val = math.Abs(float64(guestZongBfs.Ranking - mainZongBfs.Ranking))
		if temp_val > xishu {
			if mainZongBfs.Ranking < guestZongBfs.Ranking {
				letBall += baseVal * (temp_val / xishu)
			} else {
				letBall -= baseVal * (temp_val / xishu)
			}
		}
	} else {
		//return -1, temp_data
	}

	//两队对战
	bfb_arr := this.BFBattleService.FindNearByMatchId(matchId, 4)
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
		letBall += baseVal
	}
	if guestWin > mainWin {
		letBall -= baseVal
	}
	//近期战绩
	bfj_main := this.BFJinService.FindNearByTeamName(v.MatchDate,v.MainTeamId, 4)
	bfj_guest := this.BFJinService.FindNearByTeamName(v.MatchDate,v.GuestTeamId, 4)
	bfj_mainWin := 0
	bfj_guestWin := 0
	for _, e := range bfj_main {
		if e.HomeTeam == v.MainTeamId && e.HomeScore > e.GuestScore {
			bfj_mainWin++
		}
		if e.GuestTeam == v.MainTeamId && e.GuestScore > e.HomeScore {
			bfj_mainWin++
		}
	}
	for _, e := range bfj_guest {
		if e.HomeTeam == v.GuestTeamId && e.HomeScore > e.GuestScore {
			bfj_guestWin++
		}
		if e.GuestTeam == v.GuestTeamId && e.GuestScore > e.HomeScore {
			bfj_guestWin++
		}
	}
	if bfj_mainWin > (bfj_guestWin + 1) {
		letBall += baseVal
	}
	if bfj_guestWin > (bfj_mainWin + 1) {
		letBall -= baseVal
	}

	//未来赛事
	bffe_main := this.BFFutureEventService.FindNextBattle(matchId, v.MainTeamId)
	bffe_guest := this.BFFutureEventService.FindNextBattle(matchId, v.GuestTeamId)
	if this.IsCupMatch(bffe_main.EventLeagueId){
		//下一场打杯赛
		return -3, temp_data
	} else {
		//如果主队下一场打客场,战意充足
		if v.MainTeamId == bffe_main.EventGuestTeamId {
			letBall += 0.125
		}
	}
	if this.IsCupMatch(bffe_guest.EventLeagueId){
		//下一场打杯赛
		return -3, temp_data
	} else {
		//如果客队下一场打主场，战意懈怠
		if v.GuestTeamId == bffe_guest.EventMainTeamId {
			letBall += 0.125
		}
	}

	//1.0判断主队是否是让球方
	mainLetball := this.AnalyService.mainLetball(a18Bet)

	val_range := 0.375
	if mainLetball {
		if letBall > 0 {
			if a18Bet.ELetBall > 0 {
				tempLetball1 := math.Abs(a18Bet.ELetBall - letBall)
				if tempLetball1 < val_range {
					preResult = 3
				} else {
					//preResult = 0
				}
			} else if a18Bet.ELetBall < 0 {
				//preResult = 0
			}
		}
		if letBall < 0 {
			if a18Bet.ELetBall < 0 {
				tempLetball1 := math.Abs(a18Bet.ELetBall - letBall)
				if tempLetball1 < val_range {
					//preResult = 0
				} else {
					//preResult = 3
				}
			} else if a18Bet.ELetBall > 0 {
				//preResult = 3
			}
		}
	}
	if !mainLetball {
		if letBall < 0 {
			if a18Bet.ELetBall < 0 {
				tempLetball1 := math.Abs(a18Bet.ELetBall - letBall)
				if tempLetball1 < val_range {
					//preResult = 0
				} else {
					//preResult = 3
				}
			} else if a18Bet.ELetBall > 0 {
				//preResult = 3
			}
		}
		if letBall > 0 {
			if a18Bet.ELetBall > 0 {
				tempLetball1 := math.Abs(a18Bet.ELetBall - letBall)
				if tempLetball1 < val_range {
					preResult = 3
				} else {
					//preResult = 0
				}
			} else if a18Bet.ELetBall < 0 {
				//preResult = 0
			}
		}
	}

	if preResult < 0 {
		return -3, temp_data
	}
	temp_bfb_arr := this.BFBattleService.FindNearByMatchId(matchId, 1)
	for _, e := range temp_bfb_arr {
		if e.BattleMainTeamId == v.MainTeamId && e.BattleMainTeamGoals > e.BattleGuestTeamGoals {
			continue
		}
		if e.BattleGuestTeamId == v.MainTeamId && e.BattleGuestTeamGoals > e.BattleMainTeamGoals {
			continue
		}
		if e.BattleMainTeamId == v.GuestTeamId && e.BattleMainTeamGoals > e.BattleGuestTeamGoals {
			return -3, temp_data
		}
		if e.BattleGuestTeamId == v.GuestTeamId && e.BattleGuestTeamGoals > e.BattleMainTeamGoals {
			return -3, temp_data
		}
	}
	temp_bfj_main := this.BFJinService.FindNearByTeamName(v.MatchDate,v.MainTeamId, 1)
	temp_bfj_guest := this.BFJinService.FindNearByTeamName(v.MatchDate,v.GuestTeamId, 1)
	for _, e := range temp_bfj_main {
		if e.HomeTeam == v.MainTeamId && e.HomeScore > e.GuestScore {
			continue
		}else if e.GuestTeam == v.MainTeamId && e.GuestScore > e.HomeScore {
			continue
		}else{
			return -3, temp_data
		}
	}
	for _, e := range temp_bfj_guest {
		if e.HomeTeam == v.GuestTeamId && e.HomeScore > e.GuestScore {
			return -3, temp_data
		}else if e.GuestTeam == v.GuestTeamId && e.GuestScore > e.HomeScore {
			return -3, temp_data
		}else{
			continue
		}
	}

	base.Log.Info("所属于区间:", ",对阵", v.MainTeamId+":"+v.GuestTeamId, ",计算得出让球为:", letBall, ",初盘让球:", a18Bet.SLetBall, ",即时盘让球:", a18Bet.ELetBall)
	var data *entity5.AnalyResult
	if len(temp_data.Id) > 0 {
		temp_data.PreResult = preResult
		temp_data.HitCount = temp_data.HitCount + 1
		temp_data.LetBall = a18Bet.ELetBall
		temp_data.MyLetBall = Decimal(letBall)
		data = temp_data
		//比赛结果
		data.Result = this.IsRight2Option(v, data)
		return 1, data
	} else {
		data = new(entity5.AnalyResult)
		data.MatchId = v.Id
		data.MatchDate = v.MatchDate
		data.SLetBall = a18Bet.SLetBall
		data.LetBall = a18Bet.ELetBall
		data.AlFlag = this.ModelName()
		format := time.Now().Format("0102150405")
		data.AlSeq = format
		data.PreResult = preResult
		data.HitCount = 3
		data.MyLetBall = Decimal(letBall)
		//比赛结果
		data.Result = this.IsRight2Option(v, data)
		return 0, data
	}
}

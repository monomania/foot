package service

import (
	"fmt"
	"math"
	"strconv"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	entity5 "tesou.io/platform/foot-parent/foot-api/module/analy/pojo"
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	entity3 "tesou.io/platform/foot-parent/foot-api/module/odds/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/utils"
	"tesou.io/platform/foot-parent/foot-core/module/analy/constants"
	"tesou.io/platform/foot-parent/foot-core/module/match/service"
	"time"
)

type C1Service struct {
	AnalyService
	service.BFScoreService
	service.BFBattleService
	service.BFFutureEventService

	//最大让球数据
	MaxLetBall float64
}

func (this *C1Service) ModelName() string {
	//alFlag := reflect.TypeOf(*this).Name()
	return "C1"
}

func (this *C1Service) Analy(analyAll bool) {
	var matchLasts []*pojo.MatchLast
	if analyAll {
		//matchHis := this.MatchHisService.FindAll()
		//for _, e := range matchHis {
		//	matchLasts = append(matchLasts, &e.MatchLast)
		//}
		matchLasts = this.MatchLastService.FindAll()
	} else {
		matchLasts = this.MatchLastService.FindNotFinished()
	}
	this.Analy_Process(matchLasts)

}

func (this *C1Service) Analy_Near() {
	matchList := this.MatchLastService.FindNear()
	this.Analy_Process(matchList)
}

func (this *C1Service) Analy_Process(matchList []*pojo.MatchLast) {
	data_list_slice := make([]interface{}, 0)
	data_modify_list_slice := make([]interface{}, 0)
	for _, v := range matchList {
		stub, data := this.analyStub(v)

		if stub == 0 || stub == 1 {
			hours := v.MatchDate.Sub(time.Now()).Hours()
			if hours > 0 {
				hit_count_str := utils.GetVal(constants.SECTION_NAME, "hit_count")
				hit_count, _ := strconv.Atoi(hit_count_str)
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

func (this *C1Service) analyStub(v *pojo.MatchLast) (int, *entity5.AnalyResult) {
	matchId := v.Id
	//声明使用变量
	var a18betData *entity3.AsiaHis
	//亚赔
	//aList := this.AsiaHisService.FindByMatchIdCompId(matchId, "澳门")
	aList := this.AsiaHisService.FindByMatchIdCompId(matchId, "Bet365")
	//aList := this.AsiaHisService.FindByMatchIdCompId(matchId, "18Bet")
	if len(aList) < 1 {
		return -1, nil
	}
	a18betData = aList[0]
	if math.Abs(a18betData.ELetBall) > this.MaxLetBall {
		return -2, nil
	}

	//得出结果
	var preResult int
	letBall := 0.00
	//------
	bfs_arr := this.BFScoreService.FindByMatchId(matchId)
	if len(bfs_arr) < 1 {
		return -1, nil
	}
	if matchId == "1723728" {
		fmt.Println("-")
	}
	var temp_val float64
	var mainZongBfs *pojo.BFScore
	var mainZhuBfs *pojo.BFScore
	var guestZongBfs *pojo.BFScore
	var guestKeBfs *pojo.BFScore
	for _, e := range bfs_arr {//bfs_arr有多语言版本,条数很多
		if e.TeamId == v.MainTeamId {
			if e.Type == "总" {
				mainZongBfs = e
			}
			if e.Type == "主" {
				mainZhuBfs = e
			}
		} else if e.TeamId == v.GuestTeamId{
			if e.Type == "总" {
				guestZongBfs = e
			}
			if e.Type == "客" {
				guestKeBfs = e
			}
		}
	}
	if mainZongBfs == nil || guestZongBfs == nil || mainZhuBfs == nil || guestKeBfs == nil {
		return -1, nil
	}

	//排名越小越好
	rankDiff := 5.0
	temp_val = float64(mainZongBfs.Ranking - guestZongBfs.Ranking)
	if temp_val >= rankDiff {
		letBall += -0.125 - (temp_val/rankDiff)*0.125
	}
	temp_val = float64(guestZongBfs.Ranking - mainZongBfs.Ranking)
	if temp_val >= rankDiff {
		letBall += 0.125 + (temp_val/rankDiff)*0.125
	}
	temp_val = float64(mainZhuBfs.Ranking - guestKeBfs.Ranking)
	if temp_val >= rankDiff {
		letBall += -0.125 - (temp_val/rankDiff)*0.125
	}
	temp_val = float64(guestKeBfs.Ranking - mainZhuBfs.Ranking)
	if temp_val >= rankDiff {
		letBall += 0.125 + (temp_val/rankDiff)*0.125
	}

	//------
	bfb_arr := this.BFBattleService.FindByMatchId(matchId)
	winCountDiff := 2
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
	if mainWin > (guestWin + 1) {
		letBall += 0.125 + float64(mainWin/winCountDiff)*0.125
	}
	if guestWin > (mainWin + 1) {
		letBall += -0.125 - float64(guestWin/winCountDiff)*0.125
	}
	//------
	bffe_main := this.BFFutureEventService.FindNextBattle(matchId, v.MainTeamId)
	bffe_guest := this.BFFutureEventService.FindNextBattle(matchId, v.GuestTeamId)
	//如果主队下一场打客场,战意充足
	if v.MainTeamId == bffe_main.EventGuestTeamId {
		letBall += 0.075
	}
	//如果客队下一场打主场，战意懈怠
	if v.GuestTeamId == bffe_guest.EventMainTeamId {
		letBall += 0.075
	}

	//判断主队是否是让球方
	mainLetball := true
	if a18betData.ELetBall > 0 {
		mainLetball = true
	} else if a18betData.ELetBall < 0 {
		mainLetball = false
	} else {
		if a18betData.SLetBall > 0 && a18betData.ELetBall <= 0 {
			mainLetball = false
		} else {
			if letBall >= 0 {
				mainLetball = true
			} else {
				mainLetball = false
			}
		}
	}

	if a18betData.ELetBall > 0 && letBall < 0 {
		preResult = 3
	} else if a18betData.ELetBall < 0 && letBall > 0 {
		preResult = 0
	} else {
		sLetBall := math.Abs(a18betData.SLetBall)
		eLetBall := math.Abs(a18betData.ELetBall)
		tLetBall := math.Abs(letBall)
		ableUpDown := false
		ableOdd := 0.075
		if math.Abs(sLetBall-eLetBall) <= 0.25 && sLetBall != 0 && eLetBall != 0 {
			ableUpDown = true
			ableOdd = 0.25
		}

		if (sLetBall >= tLetBall && eLetBall >= tLetBall) || (sLetBall < tLetBall && eLetBall >= tLetBall) {
			var seLetBall float64
			if ableUpDown {
				if sLetBall > eLetBall {
					seLetBall = eLetBall
				} else {
					seLetBall = sLetBall
				}
			} else {
				seLetBall = eLetBall
			}
			if math.Abs(seLetBall-tLetBall) <= ableOdd {
				if mainLetball {
					preResult = 3
				} else {
					preResult = 0
				}
			} else {
				if mainLetball {
					preResult = 0
				} else {
					preResult = 3
				}
			}
		} else if (sLetBall < tLetBall && eLetBall < tLetBall) || (sLetBall >= tLetBall && eLetBall < tLetBall) {
			var seLetBall float64
			if ableUpDown {
				if sLetBall > eLetBall {
					seLetBall = sLetBall
				} else {
					seLetBall = eLetBall
				}
			} else {
				seLetBall = eLetBall
			}
			if math.Abs(tLetBall-seLetBall) <= ableOdd {
				if mainLetball {
					preResult = 3
				} else {
					preResult = 0
				}
			} else {
				if mainLetball {
					preResult = 0
				} else {
					preResult = 3
				}
			}
		} else {
			base.Log.Error("C1分析模型出现错误!!!")
		}
	}
	base.Log.Info("计算得出让球为:", letBall, ",初盘让球:", a18betData.SLetBall, ",即时盘让球:", a18betData.ELetBall)
	var data *entity5.AnalyResult
	temp_data := this.Find(v.Id, this.ModelName())
	if len(temp_data.Id) > 0 {
		temp_data.PreResult = preResult
		temp_data.HitCount = temp_data.HitCount + 1
		temp_data.LetBall = a18betData.ELetBall
		temp_data.MyLetBall = Decimal(letBall)
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
		data.HitCount = 1
		data.LetBall = a18betData.ELetBall
		data.MyLetBall = Decimal(letBall)
		//比赛结果
		data.Result = this.IsRight(v, data)
		return 0, data
	}

}

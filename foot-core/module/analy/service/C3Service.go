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

type C3Service struct {
	AnalyService
	service.BFScoreService
	service.BFBattleService
	service.BFFutureEventService
	service.BFJinService

	//最大让球数据
	MaxLetBall float64
}

func (this *C3Service) ModelName() string {
	return "C3"
}

func (this *C3Service) AnalyTest() {
	this.AnalyService.AnalyTest(this)
}

func (this *C3Service) Analy(analyAll bool) {
	this.AnalyService.Analy(analyAll, this)
}

func (this *C3Service) Analy_Near() {
	this.AnalyService.Analy_Near(this)
}

/**
  -1 参数错误
  -2 不符合让球数
  -3 计算分析错误
  0  新增的分析结果
  1  需要更新结果
 */
func (this *C3Service) analyStub(v *pojo.MatchLast) (int, *entity5.AnalyResult) {
	temp_data := this.Find(v.Id, this.ModelName())
	matchId := v.Id
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
		//temp_data.Result =""
		return -2, temp_data
	}

	//限制初盘,即时盘让球在0.25以内
	if math.Abs(a18Bet.SLetBall-a18Bet.ELetBall) > 0.25 {
		//temp_data.Result =""
		return -2, temp_data
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
	var mainZhuJinBfs *pojo.BFScore
	var guestZongBfs *pojo.BFScore
	var guestKeBfs *pojo.BFScore
	var guestKeJinBfs *pojo.BFScore
	for _, e := range bfs_arr { //bfs_arr有多语言版本,条数很多
		if e.TeamId == v.MainTeamId {
			if e.Type == "总" {
				mainZongBfs = e
			}
			if e.Type == "主" {
				mainZhuBfs = e
			}
			if e.Type == "近" {
				mainZhuJinBfs = e
			}
		} else if e.TeamId == v.GuestTeamId {
			if e.Type == "总" {
				guestZongBfs = e
			}
			if e.Type == "客" {
				guestKeBfs = e
			}
			if e.Type == "近" {
				guestKeJinBfs = e
			}
		}
	}
	if mainZongBfs == nil || mainZhuBfs == nil || mainZhuJinBfs == nil || guestZongBfs == nil || guestKeBfs == nil || guestKeJinBfs == nil {
		return -1, temp_data
	}

	var temp_zong_val_1, temp_zong_val_2 float64 = 0, 0
	var temp_jin_val_1, temp_jin_val_2 float64 = 0, 0
	if mainZongBfs.MatchCount >= 8 && guestZongBfs.MatchCount >= 8 {
		//排名越小越好
		temp_zong_val_1 = float64(mainZongBfs.Ranking - guestZongBfs.Ranking)
		temp_zong_val_2 = float64(mainZongBfs.GetGoal - guestZongBfs.GetGoal)

		temp_jin_val_1 = float64(mainZhuJinBfs.Score - guestKeJinBfs.Score)
		temp_jin_val_2 = float64(mainZhuJinBfs.GetGoal - guestKeJinBfs.GetGoal)
	} else {
		return -1, temp_data
	}

	//保证积分榜上,进球能力有支持
	if (temp_zong_val_1 > 0 && temp_zong_val_2 < 0) || (temp_zong_val_1 < 0 && temp_zong_val_2 > 0) {
		return -2, temp_data
	}

	//2.0区间段
	var sectionBlock1, sectionBlock2 int = -1, -1

	var sectionValArr = []float64{0.0, 0.25, 0.50, 0.75, 1, 1.25, 1.5, 1.75, 2}
	//总积分的让球值
	tempLetball1 := math.Abs(temp_zong_val_1) / 10
	for i, e := range sectionValArr {
		if tempLetball1 <= e {
			sectionBlock1 = i
			break;
		}
	}

	tempLetball2 := math.Abs(temp_zong_val_2) / 10
	for i, e := range sectionValArr {
		if tempLetball2 <= e {
			sectionBlock2 = i
			break;
		}
	}

	if sectionBlock1 == -1 || sectionBlock2 == -1 {
		return -3, temp_data
	}

	//3.0即时盘赔率大于等于初盘赔率
	sLetBall := math.Abs(a18Bet.SLetBall)
	eLetBall := math.Abs(a18Bet.ELetBall)
	endUp := eLetBall >= sLetBall
	if !endUp {
		return -3, temp_data
	}


	if matchId == "1721267" {
		fmt.Println("")
	}

	//
	var bf_begin, bf_end float64
	if sectionBlock1 == sectionBlock2 {
		if sectionBlock1 == 0 {
			return -3, temp_data
		}
		bf_begin = sectionValArr[sectionBlock1] - 0.25
		bf_end = sectionValArr[sectionBlock1] + 0.25
	} else if sectionBlock1-sectionBlock2 == 1 || sectionBlock1-sectionBlock2 == -1 {
		if sectionBlock1 < sectionBlock2 {
			bf_begin = sectionValArr[sectionBlock1] - 0.25
			bf_end = sectionValArr[sectionBlock2] + 0.25
		} else {
			bf_begin = sectionValArr[sectionBlock2] - 0.25
			bf_end = sectionValArr[sectionBlock1] + 0.25
		}
	}

	//存在合理区间里
	var smallLetBall float64
	if eLetBall-sLetBall >= 0 {
		smallLetBall = sLetBall
	} else {
		smallLetBall = eLetBall
	}
	if bf_end-smallLetBall >= 0.5 {
		return -3, temp_data
	}
	var reasonable = false
	if bf_begin <= eLetBall && eLetBall <= bf_end {
		reasonable = true
	}

	//1.0判断主队是否是让球方
	mainLetball := this.AnalyService.mainLetball(a18Bet)
	if reasonable {
		if mainLetball {
			preResult = 3
		} else {
			preResult = 0
		}
	} else {

	}

	if preResult == 3 && (temp_jin_val_1 < 0 || temp_jin_val_2 < 0) {
		preResult = -1
	}
	if preResult == 3 && (temp_jin_val_1 > 0 || temp_jin_val_2 > 0) {
		preResult = -1
	}

	if preResult < 0 {
		return -3, temp_data
	}

	base.Log.Info("所属于区间: ", bf_begin, " - ", bf_end, " ,对阵", v.MainTeamId+":"+v.GuestTeamId, ",初盘让球:", a18Bet.SLetBall, ",即时盘让球:", a18Bet.ELetBall)
	var data *entity5.AnalyResult
	if len(temp_data.Id) > 0 {
		temp_data.PreResult = preResult
		temp_data.HitCount = temp_data.HitCount + 1
		temp_data.LetBall = a18Bet.ELetBall
		data = temp_data
		//比赛结果
		data.Result = this.IsRight(v, data)
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
		//比赛结果
		data.Result = this.IsRight(v, data)
		return 0, data
	}
}

package service

import (
	"math"
	"strings"
	entity5 "tesou.io/platform/foot-parent/foot-api/module/analy/pojo"
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	entity3 "tesou.io/platform/foot-parent/foot-api/module/odds/pojo"
	"time"
)

type E2Service struct {
	AnalyService
	//最大让球数据
	MaxLetBall float64
}

func (this *E2Service) ModelName() string {
	return "E2"
}

/**
计算欧赔81 616的即时盘,和初盘的差异
*/
func (this *E2Service) Analy(analyAll bool) {
	this.AnalyService.Analy(analyAll,this)
}

/**
计算欧赔81 616的即时盘,和初盘的差异
*/
func (this *E2Service) Analy_Near() {
	this.AnalyService.Analy_Near(this)
}

/**
  -1 参数错误
  -2 不符合让球数
  -3 计算分析错误
  0  新增的分析结果
  1  需要更新结果
 */
func (this *E2Service) analyStub(v *pojo.MatchLast) (int, *entity5.AnalyResult) {
	temp_data := this.Find(v.Id, this.ModelName())
	matchId := v.Id
	//声明使用变量
	var e616data *entity3.EuroHis
	var e104data *entity3.EuroHis
	var a18Bet *entity3.AsiaHis
	//81 -- 伟德
	eList := this.EuroHisService.FindByMatchIdCompId(matchId, "616", "104")
	if len(eList) < 2 {
		return -1, temp_data
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

	if e616data == nil || e104data == nil  {
		return -1, temp_data
	}

	//0.没有变化则跳过
	if e104data.Ep3 == e104data.Sp3 || e104data.Ep0 == e104data.Sp0 {
		return -3, temp_data
	}
	if e616data.Ep3 == e616data.Sp3 || e616data.Ep0 == e616data.Sp0 {
		return -3, temp_data
	}

	//1.有变化,进行以下逻辑
	//亚赔
	aList := this.AsiaHisService.FindByMatchIdCompId(matchId, "18Bet")
	if len(aList) < 1 {
		return -1, temp_data
	}
	a18Bet = aList[0]
	temp_data.LetBall = a18Bet.ELetBall
	if math.Abs(a18Bet.ELetBall) > this.MaxLetBall {
		return -2, temp_data
	}

	//得出结果
	if e616data.Ep0 < e616data.Sp0 && e616data.Ep0 < e104data.Sp0 {
		return -3, temp_data
	}
	var preResult int
	if e616data.Ep3 > (e616data.Sp3+0.01) && e104data.Ep3 < e104data.Sp3 {
		preResult = 3
	} else if e616data.Ep3 < e616data.Sp3 && e104data.Ep3 < e104data.Sp3 && e616data.Ep3 < e104data.Ep3 {
		preResult = 3
	} else {
		return -3, temp_data
	}

	var data *entity5.AnalyResult
	if len(temp_data.Id) > 0 {
		temp_data.PreResult = preResult
		temp_data.HitCount = temp_data.HitCount + 1
		temp_data.LetBall = a18Bet.ELetBall
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
		//比赛结果
		data.Result = this.IsRight2Option(v, data)
		return 0, data
	}
}


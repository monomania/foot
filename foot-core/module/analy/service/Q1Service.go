package service

import (
	"math"
	"strings"
	entity5 "tesou.io/platform/foot-parent/foot-api/module/analy/pojo"
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	entity3 "tesou.io/platform/foot-parent/foot-api/module/odds/pojo"
	"time"
)

type Q1Service struct {
	AnalyService
	//最大让球数据
	MaxLetBall float64
}

func (this *Q1Service) ModelName() string {
	return "Q1"
}

func (this *Q1Service) Analy(analyAll bool) {
	this.AnalyService.Analy(analyAll,this)
}

func (this *Q1Service) Analy_Near() {
	this.AnalyService.Analy_Near(this)
}

/**
  -1 参数错误
  -2 不符合让球数
  -3 计算分析错误
  0  新增的分析结果
  1  需要更新结果
 */
func (this *Q1Service) analyStub(v *pojo.MatchLast) (int, *entity5.AnalyResult) {
	temp_data := this.Find(v.Id, this.ModelName())
	matchId := v.Id
	//声明使用变量
	var e281data *entity3.EuroHis
	var e1129data *entity3.EuroHis
	var a18bet *entity3.AsiaHis
	eList := this.EuroHisService.FindByMatchIdCompId(matchId, "281", "1129")
	if len(eList) < 2 {
		return -1, temp_data
	}
	for _, ev := range eList {
		if strings.EqualFold(ev.CompId, "281") {
			e281data = ev
			continue
		}
		if strings.EqualFold(ev.CompId, "1129") {
			e1129data = ev
			continue
		}
	}
	if e281data == nil || e1129data == nil  {
		return -1, temp_data
	}

	//0.没有变化则跳过
	if e281data.Ep3 == e281data.Sp3 || e281data.Ep0 == e281data.Sp0 {
		return -3, temp_data
	}
	//if e1129data.Ep3 == e1129data.Sp3 || e1129data.Ep0 == e1129data.Sp0 {
	//	return -3, nil
	//}

	//1.有变化,进行以下逻辑
	//亚赔
	aList := this.AsiaHisService.FindByMatchIdCompId(matchId, "18Bet")
	if len(aList) < 1 {
		return -1, temp_data
	}
	a18bet = aList[0]
	temp_data.LetBall = a18bet.ELetBall
	if math.Abs(a18bet.ELetBall) > this.MaxLetBall {
		return -2, temp_data
	}

	//得出结果
	var preResult int
	if e1129data.Ep3 > e281data.Ep3 && e1129data.Ep0 < e281data.Ep0 {
		preResult = 0
	} else if e1129data.Ep0 > e281data.Ep0 && e1129data.Ep3 < e281data.Ep3 {
		preResult = 3
	} else {
		return -3, temp_data
	}

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
		data.SLetBall = a18Bet.SLetBall
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

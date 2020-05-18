package service

import (
	"math"
	entity5 "tesou.io/platform/foot-parent/foot-api/module/analy/pojo"
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	entity3 "tesou.io/platform/foot-parent/foot-api/module/odds/pojo"
	"tesou.io/platform/foot-parent/foot-core/module/analy/constants"
	"time"
)

type E1Service struct {
	AnalyService
	//最大让球数据
	MaxLetBall float64
}

func (this *E1Service) ModelName() string {
	return "E1"
}

func (this *E1Service) AnalyTest() {
	this.AnalyService.AnalyTest(this)
}

/**
计算欧赔81 616的即时盘,和初盘的差异
*/
func (this *E1Service) Analy(analyAll bool) {
	this.AnalyService.Analy(analyAll,this)
}

func (this *E1Service) Analy_Near() {
	this.AnalyService.Analy_Near(this)
}

/**
  -1 参数错误
  -2 不符合让球数
  -3 计算分析错误
  0  新增的分析结果
  1  需要更新结果
 */
func (this *E1Service) analyStub(v *pojo.MatchLast) (int, *entity5.AnalyResult) {
	temp_data := this.Find(v.Id, this.ModelName())
	matchId := v.Id
	//声明使用变量
	var e81 *entity3.EuroHis
	var e616 *entity3.EuroHis
	var e104 *entity3.EuroHis
	var a18Bet *entity3.AsiaHis
	//81 -- 伟德
	eList := this.EuroHisService.FindByMatchIdCompId(matchId, "81", "616", "104")
	if len(eList) < 3 {
		return -1, temp_data
	}
	for _, ev := range eList {
		if ev.CompId == 81 {
			e81 = ev
			continue
		}
		if ev.CompId == 616{
			e616 = ev
			continue
		}
		if ev.CompId == 104{
			e104 = ev
			continue
		}
	}

	if e81 == nil || e616 == nil || e104 == nil {
		return -1, temp_data
	}
	//0.没有变化则跳过
	if e81.Ep3 == e81.Sp3 || e81.Ep0 == e81.Sp0 {
		return -3, temp_data
	}
	if e616.Ep3 == e616.Sp3 || e616.Ep0 == e616.Sp0 {
		return -3, temp_data
	}

	//1.有变化,进行以下逻辑
	//亚赔
	aList := this.AsiaHisService.FindByMatchIdCompId(matchId,  constants.DEFAULT_REFER_ASIA)
	if len(aList) < 1 {
		return -1, temp_data
	}
	a18Bet = aList[0]
	temp_data.LetBall = a18Bet.EPanKou
	if math.Abs(a18Bet.EPanKou) > this.MaxLetBall {
		//temp_data.Result = ""
		return -2, temp_data
	}
	//2.亚赔是主降还是主升 主降为true
	//得出结果
	var preResult int
	asiaMainDown := this.AsiaDirection(a18Bet)
	e616_s3 := e616.Sp3 / (e616.Sp3 + e616.Sp1 + e616.Sp0)
	e616_e3 := e616.Ep3 / (e616.Ep3 + e616.Ep1 + e616.Ep0)
	e616_s0 := e616.Sp0 / (e616.Sp3 + e616.Sp1 + e616.Sp0)
	e616_e0 := e616.Ep0 / (e616.Ep3 + e616.Ep1 + e616.Ep0)

	e104_s3 := e104.Sp3 / (e104.Sp3 + e104.Sp1 + e104.Sp0)
	e104_e3 := e104.Ep3 / (e104.Ep3 + e104.Ep1 + e104.Ep0)
	e104_s0 := e104.Sp3 / (e104.Sp3 + e104.Sp1 + e104.Sp0)
	e104_e0 := e104.Ep3 / (e104.Ep3 + e104.Ep1 + e104.Ep0)

	if asiaMainDown == 3 && e616_e3 < e616_s3 &&  e104_e3 < e104_s3 && e616_e3 < e104_e3 {
		preResult = 3
	} else if asiaMainDown == 0 && e616_e0 < e616_s0 &&  e104_e0 < e104_s0 && e616_e0 < e104_e0{
		preResult = 0
	} else {
		return -3, temp_data
	}

	var data *entity5.AnalyResult
	if len(temp_data.Id) > 0 {
		temp_data.MatchDate = v.MatchDate
		temp_data.PreResult = preResult
		temp_data.HitCount = temp_data.HitCount + 1
		temp_data.LetBall = a18Bet.EPanKou
		data = temp_data
		//比赛结果
		data.Result = this.IsRight(v, data)
		return 1, data
	} else {
		data = new(entity5.AnalyResult)
		data.MatchId = v.Id
		data.MatchDate = v.MatchDate
		data.SLetBall = a18Bet.SPanKou
		data.LetBall = a18Bet.EPanKou
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

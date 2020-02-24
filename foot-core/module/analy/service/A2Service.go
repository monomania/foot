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
	"time"
)

/**
亚赔与欧赔up down 颠倒
 */
type A2Service struct {
	AnalyService
	//最大让球数据
	MaxLetBall float64
}

func (this *A2Service) ModelName() string {
	//alFlag := reflect.TypeOf(*this).Name()
	return "A2"
}

/**
分析比赛数据,, 结合亚赔 赔赔差异
( 1.欧赔降水,亚赔反之,以亚赔为准)
( 2.欧赔升水,亚赔反之,以亚赔为准)
*/
func (this *A2Service) Analy(analyAll bool) {
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

func (this *A2Service) Analy_Near() {
	matchList := this.MatchLastService.FindNear()
	this.Analy_Process(matchList)
}

func (this *A2Service) Analy_Process(matchList []*pojo.MatchLast) {
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
func (this *A2Service) analyStub(v *pojo.MatchLast) (int, *entity5.AnalyResult) {
	matchId := v.Id
	//声明使用变量
	var e81 *entity3.EuroHis
	var e616 *entity3.EuroHis
	//var e104data *entity3.EuroHis
	var a18Bet *entity3.AsiaHis
	//81 -- 伟德
	eList := this.EuroHisService.FindByMatchIdCompId(matchId, "81", "616", "281")
	if len(eList) < 3 {
		return -1, nil
	}
	for _, ev := range eList {
		if strings.EqualFold(ev.CompId, "81") {
			e81 = ev
			continue
		}
		if strings.EqualFold(ev.CompId, "616") {
			e616 = ev
			continue
		}
		//if strings.EqualFold(ev.CompId, "104") {
		//	e104data = ev
		//	continue
		//}
	}

	if e81 == nil || e616 == nil {
		return -1, nil
	}
	if matchId == "1835087" {
		base.Log.Info("-")
	}

	//亚赔
	aList := this.AsiaHisService.FindByMatchIdCompId(matchId, "18Bet")
	if len(aList) < 1 {
		return -1, nil
	}
	a18Bet = aList[0]
	if math.Abs(a18Bet.ELetBall) > this.MaxLetBall {
		temp_data := this.Find(v.Id, this.ModelName())
		temp_data.LetBall = a18Bet.ELetBall
		//temp_data.Result = ""
		return -2, temp_data
	}

	//判断分析logic
	//主降为3 客降为0
	euroDirection := this.EuroDirection(e81, e616)
	//2.亚赔是主降还是主升 主降为true
	asiaDirection := this.AsiaDirectionMulti(matchId)
	//asiaDirection := this.AsiaDirection(a18Bet)
	//得出结果
	var preResult int
	if euroDirection == 3 && asiaDirection == 0 {
		preResult = 0
	} else if euroDirection == 0 && asiaDirection == 3 {
		preResult = 3
	} else {
		return -3, nil
	}

	////增加104 --Interwetten过滤
	//if preResult == 3 && (e616.Ep3 > e104data.Ep3 || e104data.Ep0 < e104data.Sp0) {
	//	return -3, nil
	//}
	//if preResult == 0 && (e616.Ep0 > e104data.Ep0 || e104data.Ep3 < e104data.Sp3) {
	//	return -3, nil
	//}

	var data *entity5.AnalyResult
	temp_data := this.Find(v.Id, this.ModelName())
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
		data.LetBall = a18Bet.ELetBall
		data.AlFlag = this.ModelName()
		format := time.Now().Format("0102150405")
		data.AlSeq = format
		data.PreResult = preResult
		data.HitCount = 3
		data.LetBall = a18Bet.ELetBall
		//比赛结果
		data.Result = this.IsRight(v, data)
		return 0, data
	}

}

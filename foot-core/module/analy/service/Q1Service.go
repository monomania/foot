package service

import (
	"math"
	"strconv"
	"strings"
	entity5 "tesou.io/platform/foot-parent/foot-api/module/analy/pojo"
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	entity3 "tesou.io/platform/foot-parent/foot-api/module/odds/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/utils"
	"tesou.io/platform/foot-parent/foot-core/module/analy/constants"
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

func (this *Q1Service) Analy() {
	matchList := this.MatchLastService.FindNotFinished()
	this.Analy_Process(matchList)

}

func (this *Q1Service) Analy_Near() {
	matchList := this.MatchLastService.FindNear()
	this.Analy_Process(matchList)

}

func (this *Q1Service) Analy_Process(matchList []*pojo.MatchLast) {
	hit_count_str := utils.GetVal(constants.SECTION_NAME, "hit_count")
	hit_count, _ := strconv.Atoi(hit_count_str)
	data_list_slice := make([]interface{}, 0)
	data_modify_list_slice := make([]interface{}, 0)
	for _, v := range matchList {
		stub, data := this.analyStub(v)

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
func (this *Q1Service) analyStub(v *pojo.MatchLast) (int, *entity5.AnalyResult) {
	matchId := v.Id
	//声明使用变量
	var e281data *entity3.EuroHis
	var e1129data *entity3.EuroHis
	var a18betData *entity3.AsiaHis
	eList := this.EuroHisService.FindByMatchIdCompId(matchId, "281", "1129")
	if len(eList) < 2 {
		return -1, nil
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
		return -1, nil
	}

	//0.没有变化则跳过
	if e281data.Ep3 == e281data.Sp3 || e281data.Ep0 == e281data.Sp0 {
		return -3, nil
	}
	//if e1129data.Ep3 == e1129data.Sp3 || e1129data.Ep0 == e1129data.Sp0 {
	//	return -3, nil
	//}

	//1.有变化,进行以下逻辑
	//亚赔
	aList := this.AsiaHisService.FindByMatchIdCompId(matchId, "18Bet")
	if len(aList) < 1 {
		return -1, nil
	}
	a18betData = aList[0]
	if math.Abs(a18betData.ELetBall) > this.MaxLetBall {
		temp_data := this.Find(v.Id, this.ModelName())
		temp_data.LetBall = a18betData.ELetBall
		return -2, temp_data
	}

	//得出结果
	var preResult int
	if e1129data.Ep3 > e281data.Ep3 && e1129data.Ep0 < e281data.Ep0 {
		preResult = 0
	} else if e1129data.Ep0 > e281data.Ep0 && e1129data.Ep3 < e281data.Ep3 {
		preResult = 3
	} else {
		return -3, nil
	}

	var data *entity5.AnalyResult
	temp_data := this.Find(matchId, this.ModelName())
	if len(temp_data.Id) > 0 {
		temp_data.PreResult = preResult
		temp_data.HitCount = temp_data.HitCount + 1
		temp_data.LetBall = a18betData.ELetBall
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
		data.HitCount = 3
		data.LetBall = a18betData.ELetBall
		//比赛结果
		data.Result = this.IsRight(v, data)
		return 0, data
	}
}

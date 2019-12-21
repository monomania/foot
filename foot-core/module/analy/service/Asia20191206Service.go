package service

import (
	"reflect"
	"strconv"
	"strings"
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
type Asia20191206Service struct {
	AnalyService
	//最大让球数据
	MaxLetBall float64
}

/**
分析比赛数据,, 结合亚赔 赔赔差异
( 1.欧赔降水,亚赔反之,以亚赔为准)
( 2.欧赔升水,亚赔反之,以亚赔为准)
*/
func (this *Asia20191206Service) Analy() {
	matchList := this.MatchLastService.FindAll()
	data_list_slice := make([]interface{}, 0)
	data_modify_list_slice := make([]interface{}, 0)
	for _, v := range matchList {
		stub, result := this.analyStub(v)
		if nil == result {
			continue
		}
		if stub == 0 {
			data_list_slice = append(data_list_slice, result)
		} else if stub == 1 {
			data_modify_list_slice = append(data_modify_list_slice, result)
		} else {
			temp_data := this.Find(v.Id, result.AlFlag)
			if len(temp_data.Id) > 0 {
				hit_count_str := utils.GetVal(constants.SECTION_NAME, "hit_count")
				hit_count, _ := strconv.Atoi(hit_count_str)
				if temp_data.HitCount >= hit_count{
					continue
				}
				this.AnalyService.Del(temp_data)
			}
		}
	}
	this.AnalyService.SaveList(data_list_slice)
	this.AnalyService.ModifyList(data_modify_list_slice)
}

func (this *Asia20191206Service) analyStub(v *pojo.MatchLast) (int, *entity5.AnalyResult) {
	matchId := v.Id
	//声明使用变量
	var e81data *entity3.EuroLast
	var e616data *entity3.EuroLast
	var e104data *entity3.EuroLast
	var a18betData *entity3.AsiaLast
	//81 -- 伟德
	eList := this.EuroLastService.FindByMatchIdCompId(matchId, "81", "616", "104")
	if len(eList) < 3 {
		return -1, nil
	}
	for _, ev := range eList {
		if strings.EqualFold(ev.CompId, "81") {
			e81data = ev
			continue
		}
		if strings.EqualFold(ev.CompId, "616") {
			e616data = ev
			continue
		}
		if strings.EqualFold(ev.CompId, "104") {
			e104data = ev
			continue
		}
	}

	//亚赔
	aList := this.AsiaLastService.FindByMatchIdCompId(matchId, "18Bet")
	if len(aList) < 1 {
		return -1, nil
	}
	a18betData = aList[0]
	if a18betData.ELetBall > this.MaxLetBall {
		return -2, nil
	}

	//判断分析logic
	//1.欧赔是主降还是主升 主降为true
	euroMainDown := EuroMainDown(e81data, e616data)
	//2.亚赔是主降还是主升 主降为true
	asiaMainDown := AsiaMainDown(a18betData)
	//得出结果
	var preResult int
	if euroMainDown == 3 && !asiaMainDown {
		preResult = 0
	} else if euroMainDown == 0 && asiaMainDown {
		preResult = 3
	} else {
		return -3, nil
	}

	//增加104 --Interwetten过滤
	if preResult == 3 && (e616data.Ep3 > e104data.Ep3 || e104data.Ep0 < e104data.Sp0) {
		return -3, nil
	}
	if preResult == 0 && (e616data.Ep0 > e104data.Ep0 || e104data.Ep3 < e104data.Sp3) {
		return -3, nil
	}

	alFlag := reflect.TypeOf(*this).Name()
	var data *entity5.AnalyResult
	temp_data := this.Find(v.Id, alFlag)
	if len(temp_data.Id) > 0 {
		temp_data.PreResult = preResult
		temp_data.HitCount = temp_data.HitCount + 1
		data = temp_data
		data.LetBall = a18betData.ELetBall
		//比赛结果
		data.Result = this.IsRight(a18betData, v, data)
		return 1, data
	} else {
		data = new(entity5.AnalyResult)
		data.MatchId = v.Id
		data.MatchDate = v.MatchDate
		data.LetBall = a18betData.ELetBall
		data.AlFlag = alFlag
		format := time.Now().Format("0102150405")
		data.AlSeq = format
		data.PreResult = preResult
		data.HitCount = 1
		//比赛结果
		data.Result = this.IsRight(a18betData, v, data)
		return 0, data
	}

}

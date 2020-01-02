package service

import (
	"fmt"
	"math"
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

type Euro20191206Service struct {
	AnalyService
	//最大让球数据
	MaxLetBall float64
}

/**
计算欧赔81 的即时盘,和初盘的差异
*/
func (this *Euro20191206Service) Analy() {
	matchList := this.MatchLastService.FindNotFinished()
	data_list_slice := make([]interface{}, 0)
	data_modify_list_slice := make([]interface{}, 0)
	for _, v := range matchList {
		stub, data := this.analyStub(v)

		if stub == 0 || stub == 1 {
			hours := v.MatchDate.Sub(time.Now()).Hours()
			if hours > 0 {
				hours = math.Abs(hours * 0.7)
				data.THitCount = int(hours)
			}
			if stub == 0 {
				data_list_slice = append(data_list_slice, data)
			} else if stub == 1 {
				data_modify_list_slice = append(data_modify_list_slice, data)
			}
		} else {
			alFlag := reflect.TypeOf(*this).Name()
			temp_data := this.Find(v.Id, alFlag)
			if len(temp_data.Id) > 0 {
				hit_count_str := utils.GetVal(constants.SECTION_NAME, "hit_count")
				hit_count, _ := strconv.Atoi(hit_count_str)
				if temp_data.HitCount >= hit_count {
					temp_data.HitCount = (hit_count / 2) - 1
					this.AnalyService.Modify(temp_data)
					continue
				}
				this.AnalyService.Del(temp_data)
			}
		}
	}
	this.AnalyService.SaveList(data_list_slice)
	this.AnalyService.ModifyList(data_modify_list_slice)

}

func (this *Euro20191206Service) analyStub(v *pojo.MatchLast) (int, *entity5.AnalyResult) {
	matchId := v.Id
	//声明使用变量
	var e1data *entity3.EuroLast
	var e2data *entity3.EuroLast
	var e3data *entity3.EuroLast
	eList := this.EuroLastService.FindByMatchIdCompId(matchId, "115", "1129", "432")
	if len(eList) < 3 {
		return -1, nil
	}
	for _, ev := range eList {
		if strings.EqualFold(ev.CompId, "115") {
			e1data = ev
			continue
		}
		if strings.EqualFold(ev.CompId, "1129") {
			e2data = ev
			continue
		}
		if strings.EqualFold(ev.CompId, "432") {
			e3data = ev
			continue
		}
	}

	//2.亚赔是主降还是主升 主降为true
	//得出结果
	var preResult int
	//主降
	if (e2data.Sp3-e2data.Ep3 > e1data.Sp3-e1data.Ep3) && (e2data.Ep0 > e2data.Sp0) && (e2data.Ep0-e2data.Sp0 > e1data.Ep0-e1data.Sp0) {
		//主队有希望
		preResult = 3
	} else {
		//主队希望不大
		return -3, nil
	}
	if preResult < 0 {
		//主升
		if (e2data.Sp0-e2data.Ep0 > e1data.Sp0-e1data.Ep0) && (e2data.Ep3 > e2data.Sp3) && (e2data.Ep3-e2data.Sp3 > e1data.Ep3-e1data.Sp3) {
			//客队有希望
			preResult = 0
		} else {
			//客队希望不大
			return -3, nil
		}
	}

	fmt.Println(e3data.MatchId)

	alFlag := reflect.TypeOf(*this).Name()
	var data *entity5.AnalyResult
	temp_data := this.Find(v.Id, alFlag)
	if len(temp_data.Id) > 0 {
		temp_data.PreResult = preResult
		temp_data.HitCount = temp_data.HitCount + 1
		data = temp_data
		//比赛结果
		data.Result = this.IsRight(v, data)
		return 1, data
	} else {
		data = new(entity5.AnalyResult)
		data.MatchId = v.Id
		data.MatchDate = v.MatchDate
		data.AlFlag = alFlag
		format := time.Now().Format("0102150405")
		data.AlSeq = format
		data.PreResult = preResult
		data.HitCount = 1
		//比赛结果
		data.Result = this.IsRight(v, data)
		return 0, data
	}
}

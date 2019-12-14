package service

import (
	"math"
	"reflect"
	"strings"
	entity5 "tesou.io/platform/foot-parent/foot-api/module/analy/pojo"
	entity3 "tesou.io/platform/foot-parent/foot-api/module/odds/pojo"
	"time"
)

type Euro81_616_104Service struct {
	AnalyService
	//最大让球数据
	MaxLetBall float64
}

/**
计算欧赔81 616的即时盘,和初盘的差异
*/
func (this *Euro81_616_104Service) Analy() {
	matchList := this.MatchLastService.FindAll()
	data_list_slice := make([]interface{}, 0)
	data_modify_list_slice := make([]interface{}, 0)
	for _, v := range matchList {
		matchId := v.Id
		//声明使用变量
		var e81data *entity3.EuroLast
		var e616data *entity3.EuroLast
		var e104data *entity3.EuroLast
		var a18betData *entity3.AsiaLast
		//81 -- 伟德
		eList := this.EuroLastService.FindByMatchIdCompId(matchId, "81", "616", "104")
		if len(eList) < 3 {
			continue
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
		//0.没有变化则跳过
		if e81data.Ep3 == e81data.Sp3 || e81data.Ep0 == e81data.Sp0 {
			continue
		}
		if e616data.Ep3 == e616data.Sp3 || e616data.Ep0 == e616data.Sp0 {
			continue
		}

		//1.有变化,进行以下逻辑
		//亚赔
		aList := this.AsiaLastService.FindByMatchIdCompId(matchId, "18Bet")
		if len(aList) < 1 {
			continue
		}
		a18betData = aList[0]
		if math.Abs(a18betData.ELetBall) > this.MaxLetBall {
			continue
		}
		//2.亚赔是主降还是主升 主降为true
		//得出结果
		var preResult int
		asiaMainDown := AsiaMainDown(a18betData)
		if asiaMainDown {
			//主降
			if (e616data.Sp3-e616data.Ep3 > e81data.Sp3-e81data.Ep3) && (e616data.Ep0 > e616data.Sp0) && (e616data.Ep0-e616data.Sp0 > e81data.Ep0-e81data.Sp0) {
				//主队有希望
				preResult = 3
			} else {
				//主队希望不大
				continue
			}
		} else {
			//主升
			if (e616data.Sp0-e616data.Ep0 > e81data.Sp0-e81data.Ep0) && (e616data.Ep3 > e616data.Sp3) && (e616data.Ep3-e616data.Sp3 > e81data.Ep3-e81data.Sp3) {
				//客队有希望
				preResult = 0
			} else {
				//客队希望不大
				continue
			}
		}

		//增加104 --Interwetten过滤
		if preResult == 3 && (e616data.Ep3 > e104data.Ep3 || e104data.Ep0 < e104data.Sp0){
			continue
		}
		if preResult == 0 && (e616data.Ep0 > e104data.Ep0 || e104data.Ep3 < e104data.Sp3){
			continue
		}

		var data *entity5.AnalyResult
		temp_data := this.Find(v.Id)
		if len(temp_data.Id) > 0 {
			temp_data.PreResult = preResult
			temp_data.HitCount = temp_data.HitCount + 1
			data = temp_data
			data_modify_list_slice = append(data_modify_list_slice,data)
		} else {
			data = new(entity5.AnalyResult)
			data.MatchId = v.Id
			data.MatchDate = v.MatchDate
			data.AlFlag = reflect.TypeOf(*this).Name()
			format := time.Now().Format("0102150405")
			data.AlSeq = format
			data.PreResult = preResult
			data.HitCount =  1
			data_list_slice = append(data_list_slice, data)
		}
		//比赛结果
		data.Result = this.IsRight(a18betData, v, preResult)

	}
	this.AnalyService.SaveList(data_list_slice)
	this.AnalyService.ModifyList(data_modify_list_slice)
}

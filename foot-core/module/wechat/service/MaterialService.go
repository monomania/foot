package service

import (
	"github.com/chanxuehong/wechat/mp/core"
)

type MaterialService struct {
	SuggestTodayService
	SuggestWeekService
	SuggestMonthService
}

func (this *MaterialService) ModifyNews(wcClient *core.Client) {
	this.SuggestTodayService.ModifyToday(wcClient)
	this.SuggestTodayService.ModifyTodayDetail(wcClient)
	this.SuggestTodayService.ModifyTodayTbs(wcClient)
	this.SuggestTodayService.ModifyTodayA1(wcClient)
	this.SuggestTodayService.ModifyTodayC1(wcClient)

	this.SuggestWeekService.ModifyWeek(wcClient)
	this.SuggestMonthService.ModifyMonth(wcClient)
}

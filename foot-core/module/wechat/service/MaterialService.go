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
	this.SuggestTodayService.ModifyTodayTbs(wcClient)
	this.SuggestTodayService.ModifyTodayA1(wcClient)
	this.SuggestTodayService.ModifyTodayC2(wcClient)
	this.SuggestTodayService.ModifyToday(wcClient)
	this.SuggestTodayService.ModifyTodayGuts(wcClient)

	this.SuggestWeekService.ModifyWeek(wcClient)
	this.SuggestMonthService.ModifyMonth(wcClient)
	this.SuggestMonthService.ModifyGutsWeek(wcClient)
	this.SuggestMonthService.ModifyGutsMonth(wcClient)
}

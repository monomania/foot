package service

import (
	"github.com/chanxuehong/wechat/mp/core"
)

type MaterialService struct {
	MatchService
}


func (this *MaterialService) ModifyNews(wcClient *core.Client){
	this.MatchService.ModifyToday(wcClient)
	this.MatchService.ModifyTodayDetail(wcClient)
	this.MatchService.ModifyTodayTbs(wcClient)
	this.MatchService.ModifyWeek(wcClient)
	this.MatchService.ModifyMonth(wcClient)
}
package controller

import (
	"tesou.io/platform/foot-parent/foot-api/module/match/entity"
	"log"
	"tesou.io/platform/foot-parent/foot-core/module/match/service"
	"tesou.io/platform/foot-parent/foot-web/common/base/controller"
)

type MatchLastConfigController struct {
	controller.BaseController
	service.MatchLastConfigService
}

func (this *MatchLastConfigController) Query() {
	matchLastConfig := new(entity.MatchLastConfig)

	//处理参数
	this.ParseForm(matchLastConfig)
	values := this.Ctx.Request.Form
	log.Println(values.Encode())
	flag := values.Get("S")
	log.Println(flag)
	//执行查询
	result := this.MatchLastConfigService.Query(&entity.MatchLastConfig{})

	//封装结果
	this.Data["json"] = result
	this.ServeJSON()
}

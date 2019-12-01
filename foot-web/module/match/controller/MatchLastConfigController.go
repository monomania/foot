package controller

import (
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"

	"tesou.io/platform/foot-parent/foot-core/module/match/service"
	"tesou.io/platform/foot-parent/foot-web/common/base/controller"
)

type MatchLastConfigController struct {
	controller.BaseController
	service.MatchLastConfigService
}

func (this *MatchLastConfigController) Query() {
	matchLastConfig := new(pojo.MatchLastConfig)

	//处理参数
	this.ParseForm(matchLastConfig)
	values := this.Ctx.Request.Form
	base.Log.Info(values.Encode())
	flag := values.Get("S")
	base.Log.Info(flag)
	//执行查询
	result := this.MatchLastConfigService.Query(&pojo.MatchLastConfig{})

	//封装结果
	this.Data["json"] = result
	this.ServeJSON()
}

package controller

import (
	"tesou.io/platform/foot-parent/foot-core/common/base/controller"
	"tesou.io/platform/foot-parent/foot-core/module/match/entity"
	"log"
)

type MatchLastConfigController struct {
	controller.BaseController
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
	result := matchLastConfig.Query()

	//封装结果
	this.Data["json"] = result
	this.ServeJSON()
}

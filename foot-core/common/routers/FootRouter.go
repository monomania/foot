package routers

import (
	"github.com/astaxie/beego"
	"tesou.io/platform/foot-parent/foot-core/module/index/controller"
	controller2 "tesou.io/platform/foot-parent/foot-core/module/match/controller"
)

type FootRouter struct {

}

func init() {
	beego.Router("/", &controller.IndexController{})

	//match
	beego.AutoRouter(&controller2.MatchController{})
	beego.AutoRouter(&controller2.MatchLastController{})

}

func (this *FootRouter) Hello(){

}

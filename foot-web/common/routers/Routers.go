package routers

import (
	"github.com/astaxie/beego"
	"tesou.io/platform/foot-parent/foot-web/module/index/controller"
	controller2 "tesou.io/platform/foot-parent/foot-web/module/match/controller"
)

type Routers struct {
}

func init() {
	beego.Router("/", &controller.IndexController{})

	//match
	beego.AutoRouter(&controller2.MatchController{})
	beego.AutoRouter(&controller2.MatchLastConfigController{})

}

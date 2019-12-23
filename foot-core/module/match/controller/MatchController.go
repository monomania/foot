package controller

import "tesou.io/platform/foot-parent/foot-core/common/base/controller"

type MatchController struct {
	controller.BaseController
}

func (th *MatchController) Get() {
	th.Data["json"] = "match"
	th.ServeJSON()
}

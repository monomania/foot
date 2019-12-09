package controller

import "tesou.io/platform/foot-parent/foot-core/common/base/controller"

type MatchLastController struct {
	controller.BaseController
}

func (th *MatchLastController) Get() {
	th.Data["json"] = "match"
	th.ServeJSON()
}

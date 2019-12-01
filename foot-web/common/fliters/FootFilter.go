package fliters

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"tesou.io/platform/foot-parent/foot-api/common/base"
)

func init() {
	beego.InsertFilter("*", beego.BeforeStatic, BeforeStatic, false)
	beego.InsertFilter("*", beego.BeforeRouter, BeforeRouter)
	beego.InsertFilter("*", beego.BeforeExec, BeforeExec)
	beego.InsertFilter("*", beego.AfterExec, AfterExec, false)
	beego.InsertFilter("*", beego.FinishRouter, FinishRouter, false)
}

//BeforeStatic 静态地址之前
func BeforeStatic(ctx *context.Context) {
	base.Log.Info("BeforeStatic 静态地址之前")
}

//BeforeRouter 寻找路由之前
func BeforeRouter(ctx *context.Context) {
	base.Log.Info("BeforeRouter 寻找路由之前")
}

//BeforeExec 找到路由之后，开始执行相应的 Controller 之前
func BeforeExec(ctx *context.Context) {
	base.Log.Info("BeforeExec 找到路由之后，开始执行相应的 Controller 之前")
}

//AfterExec 执行完 Controller 逻辑之后执行的过滤器
func AfterExec(ctx *context.Context) {
	base.Log.Info("AfterExec 执行完 Controller 逻辑之后执行的过滤器")
}

//FinishRouter 执行完逻辑之后执行的过滤器
func FinishRouter(ctx *context.Context) {
	base.Log.Info("FinishRouter 执行完逻辑之后执行的过滤器")
}

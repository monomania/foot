package fliters

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"log"
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
	log.Println("BeforeStatic 静态地址之前")
}

//BeforeRouter 寻找路由之前
func BeforeRouter(ctx *context.Context) {
	log.Println("BeforeRouter 寻找路由之前")
}

//BeforeExec 找到路由之后，开始执行相应的 Controller 之前
func BeforeExec(ctx *context.Context) {
	log.Println("BeforeExec 找到路由之后，开始执行相应的 Controller 之前")
}

//AfterExec 执行完 Controller 逻辑之后执行的过滤器
func AfterExec(ctx *context.Context) {
	log.Println("AfterExec 执行完 Controller 逻辑之后执行的过滤器")
}

//FinishRouter 执行完逻辑之后执行的过滤器
func FinishRouter(ctx *context.Context) {
	log.Println("FinishRouter 执行完逻辑之后执行的过滤器")
}

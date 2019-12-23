package controller

import (
	_ "github.com/astaxie/beego"
	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/mp/menu"
	"github.com/chanxuehong/wechat/mp/message/callback/request"
	"github.com/chanxuehong/wechat/mp/message/callback/response"
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-core/common/base/controller"
	"tesou.io/platform/foot-parent/foot-core/common/utils"
	"tesou.io/platform/foot-parent/foot-core/module/wechat/service"
)

type WechatController struct {
	controller.BaseController
	service.MessageService
}

var (
	wcServer *core.Server
	wcClient *core.Client
)

func init() {
	msgHandler := core.NewServeMux()
	msgHandler.DefaultMsgHandleFunc(defaultMsgHandler)
	msgHandler.DefaultEventHandleFunc(defaultEventHandler)
	msgHandler.MsgHandleFunc(request.MsgTypeText, textMsgHandler)
	msgHandler.EventHandleFunc(menu.EventTypeClick, menuClickEventHandler)

	section_map := utils.GetSectionMap("wechat")
	wcServer = core.NewServer(section_map["OriId"], section_map["AppID"], section_map["Token"], section_map["EncodingAESKey"], msgHandler, nil)

	accessTokenServer := core.NewDefaultAccessTokenServer(section_map["AppID"], section_map["AppSecret"], nil)
	wcClient = core.NewClient(accessTokenServer, nil)
}

func textMsgHandler(ctx *core.Context) {
	base.Log.Info("收到文本消息:\n%s\n", ctx.MsgPlaintext)
	msg := request.GetText(ctx.MixedMsg)
	var resp interface{}
	if strings.EqualFold("今日推荐", msg.Content) || strings.EqualFold("推荐", msg.Content) {
		messageService := new(service.MessageService)
		today := messageService.Today()
		resp = response.NewNews(msg.FromUserName, msg.ToUserName, msg.CreateTime, today)
	} else {
		resp = response.NewText(msg.FromUserName, msg.ToUserName, msg.CreateTime, msg.Content)
	}
	ctx.RawResponse(resp) // 明文回复
	//ctx.AESResponse(resp, 0, "", nil) // aes密文回复
}

func defaultMsgHandler(ctx *core.Context) {
	base.Log.Info("收到消息:\n%s\n", ctx.MsgPlaintext)
	ctx.NoneResponse()
}

func menuClickEventHandler(ctx *core.Context) {
	base.Log.Info("收到菜单 click 事件:\n%s\n", ctx.MsgPlaintext)

	event := menu.GetClickEvent(ctx.MixedMsg)
	resp := response.NewText(event.FromUserName, event.ToUserName, event.CreateTime, "收到 click 类型的事件")
	//ctx.RawResponse(resp) // 明文回复
	ctx.AESResponse(resp, 0, "", nil) // aes密文回复
}

func defaultEventHandler(ctx *core.Context) {
	base.Log.Info("收到事件:\n%s\n", ctx.MsgPlaintext)
	ctx.NoneResponse()
}

/**
消息接收处理
*/
func (this *WechatController) Portable() {
	wcServer.ServeHTTP(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
}

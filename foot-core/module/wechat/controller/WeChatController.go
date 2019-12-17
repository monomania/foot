package controller

import (
	"fmt"
	_ "github.com/astaxie/beego"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/cache"
	"github.com/silenceper/wechat/message"
	"tesou.io/platform/foot-parent/foot-core/common/base/controller"
	"tesou.io/platform/foot-parent/foot-core/common/utils"
	"tesou.io/platform/foot-parent/foot-core/module/wechat/service"
)

type WeChatController struct {
	controller.BaseController
	service.MessageService
}

var (
	section_map map[string]string
)

func init() {
	section_map = utils.GetSectionMap("wechat")
}

func (this *WeChatController) Portal() {
	memory := cache.NewMemory()
	//配置微信参数
	config := &wechat.Config{
		AppID:          section_map["AppID"],
		AppSecret:      section_map["AppSecret"],
		Token:          section_map["Token"],
		EncodingAESKey: section_map["EncodingAESKey"],
		Cache:          memory,
	}
	wc := wechat.NewWechat(config)
	// 传入request和responseWriter
	server := wc.GetServer(this.Ctx.Request, this.Ctx.ResponseWriter)
	//设置接收消息的处理方法
	//server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
	//	//回复消息：演示回复用户发送的消息
	//	text := message.NewText(msg.Content)
	//	return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	//})
	server.SetMessageHandler(func(v message.MixMessage) *message.Reply {
		reply := this.MessageService.Handle(v)
		return reply
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
	//发送回复的消息
	server.Send()
}

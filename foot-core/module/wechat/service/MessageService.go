package service

import (
	"github.com/silenceper/wechat/message"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

type MessageService struct {
	mysql.BaseService
}

/**
消息管理
 */
func (this *MessageService) Handle(v message.MixMessage) *message.Reply {
	base.Log.Info("请求内容:",v)
	switch v.MsgType {
	//文本消息
	case message.MsgTypeText:
		//do something
		text := message.NewText(v.Content)
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
		//图片消息
	case message.MsgTypeImage:
		//do something
		return nil
		//语音消息
	case message.MsgTypeVoice:
		//do something
		return nil
		//视频消息
	case message.MsgTypeVideo:
		//do something
		return nil
		//小视频消息
	case message.MsgTypeShortVideo:
		//do something
		return nil
		//地理位置消息
	case message.MsgTypeLocation:
		//do something
		return nil
		//链接消息
	case message.MsgTypeLink:
		//do something
		return nil
		//事件推送消息
	case message.MsgTypeEvent:
		return this.handleMsgTypeEvent(v)
	}
	text := message.NewText(v.Content)
	return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
}

/**
事件推送消息
 */
func (this *MessageService) handleMsgTypeEvent(v message.MixMessage) *message.Reply {
	switch v.Event {
	//EventSubscribe 订阅
	case message.EventSubscribe:
		//do something
		return nil
		//取消订阅
	case message.EventUnsubscribe:
		//do something
		return nil
		//用户已经关注公众号，则微信会将带场景值扫描事件推送给开发者
	case message.EventScan:
		//do something
		return nil
		// 上报地理位置事件
	case message.EventLocation:
		//do something
		return nil
		// 点击菜单拉取消息时的事件推送
	case message.EventClick:
		//do something
		return nil
		// 点击菜单跳转链接时的事件推送
	case message.EventView:
		//do something
		return nil
		// 扫码推事件的事件推送
	case message.EventScancodePush:
		//do something
		return nil
		// 扫码推事件且弹出“消息接收中”提示框的事件推送
	case message.EventScancodeWaitmsg:
		//do something
		return nil
		// 弹出系统拍照发图的事件推送
	case message.EventPicSysphoto:
		//do something
		return nil
		// 弹出拍照或者相册发图的事件推送
	case message.EventPicPhotoOrAlbum:
		//do something
		return nil
		// 弹出微信相册发图器的事件推送
	case message.EventPicWeixin:
		//do something
		return nil
		// 弹出地理位置选择器的事件推送
	case message.EventLocationSelect:
		//do something
		return nil
	}
	return nil
}

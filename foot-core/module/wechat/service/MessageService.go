package service

import (
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/message"
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	"tesou.io/platform/foot-parent/foot-core/module/analy/service"
)

type MessageService struct {
	mysql.BaseService
	service.RecommendService
}

/**
消息管理
 */
func (this *MessageService) Handle(v message.MixMessage) *message.Reply {
	base.Log.Info("请求内容:", v)
	switch v.MsgType {
	//文本消息
	case message.MsgTypeText:
		//do something
		text := message.NewText(v.Content)
		textStr := string(text.Content)
		if strings.EqualFold("今日推荐", textStr) || strings.EqualFold("推荐", textStr) {
			return this.Today()
		} else {
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
		}
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
		return this.Today()
	}
	text := message.NewText(v.Content)
	return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
}

func (this *MessageService) Today() *message.Reply {
	listData := this.RecommendService.ListData()
	articles := make([]*message.Article, len(listData))
	for i, e := range listData {
		bytes, _ := json.Marshal(e)
		base.Log.Warn("比赛信息:" + string(bytes))
		matchDateStr := e.MatchDate.Format("01月02日15点04分")
		article := new(message.Article)
		article.Title = fmt.Sprintf("%v", matchDateStr)
		article.Description = fmt.Sprintf("%v %v vs %v", e.LeagueName, e.MainTeamId, e.GuestTeamId)
		article.PicURL = "http://mmbiz.qpic.cn/sz_mmbiz_jpg/BePaFicK2B5QZuw0bf1HsiarnqQXzuWxE9XYC25oe2mVLeguvo6Rd1j1D2ibRibfmpu8eDqs0lfXaEfXR2bhslrPKQ/0?wx_fmt=jpeg"
		article.URL = "https://gitee.com/aoe5188/foot-parent"
		articles[i] = article
	}
	news := message.NewNews(articles)
	return &message.Reply{MsgType: message.MsgTypeNews, MsgData: news}
}

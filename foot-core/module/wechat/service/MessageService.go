package service

import (
	"encoding/json"
	"fmt"
	"github.com/chanxuehong/wechat/mp/message/callback/response"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	"tesou.io/platform/foot-parent/foot-core/module/analy/service"
)

type MessageService struct {
	mysql.BaseService
	service.AnalyService
}

func (this *MessageService) Today() []response.Article {
	listData := this.AnalyService.ListDefaultData()
	articles := make([]response.Article,len(listData))
	for i, e := range listData {
		bytes, _ := json.Marshal(e)
		base.Log.Warn("比赛信息:" + string(bytes))
		matchDateStr := e.MatchDate.Format("01月02日15点04分")
		article := response.Article{}
		article.Title = fmt.Sprintf("%v", matchDateStr)
		article.Description = fmt.Sprintf("%v %v vs %v", e.LeagueName, e.MainTeamId, e.GuestTeamId)
		article.PicURL = "http://mmbiz.qpic.cn/sz_mmbiz_jpg/BePaFicK2B5QZuw0bf1HsiarnqQXzuWxE9XYC25oe2mVLeguvo6Rd1j1D2ibRibfmpu8eDqs0lfXaEfXR2bhslrPKQ/0?wx_fmt=jpeg"
		article.URL = "https://gitee.com/aoe5188/foot-parent"
		articles[i] = article
	}
	return articles
}

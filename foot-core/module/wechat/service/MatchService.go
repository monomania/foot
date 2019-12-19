package service

import (
	"encoding/json"
	"fmt"
	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/mp/material"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	"tesou.io/platform/foot-parent/foot-core/module/analy/service"
	"time"
)

type MatchService struct {
	mysql.BaseService
	service.RecommendService
}

func (this *MatchService) Today(wcClient *core.Client) string {
	listData := this.RecommendService.ListData()
	articles := make([]material.Article, len(listData)+1)
	first := material.Article{}
	first.Title = "今日推荐"
	matchDateStr := time.Now().Format("01月02日")
	first.Digest = matchDateStr
	first.ThumbMediaId = "chP-LBQxy9SVbAFjwZ4QEo81I0bHaY3YDYRwVGmf7o8"
	first.ShowCoverPic = 1
	//图文消息的原文地址，即点击“阅读原文”后的URL
	first.ContentSourceURL = "https://gitee.com/aoe5188/poem-parent"
	var first_content string

	for i, e := range listData {
		bytes, _ := json.Marshal(e)
		base.Log.Warn("比赛信息:" + string(bytes))
		matchDateStr := e.MatchDate.Format("15点04分")
		article := material.Article{}
		article.Title = fmt.Sprintf("%v %v %v vs %v", matchDateStr, e.LeagueName, e.MainTeamId, e.GuestTeamId)
		article.Digest = article.Title
		article.ThumbMediaId = "chP-LBQxy9SVbAFjwZ4QEmEjQNhcRlNZCM2b6YR_qVc"
		article.ShowCoverPic = 1
		article.ContentSourceURL = ""
		article.Content = string(bytes)
		articles[i+1] = article

		first_content += fmt.Sprintf("%v %v %v vs %v 推荐:\r\n", matchDateStr, e.LeagueName, e.MainTeamId, e.GuestTeamId, e.PreResult)
	}
	first.Content = first_content
	articles[0] = first

	mediaId, err := material.AddNews(wcClient, &material.News{Articles: articles})
	if err != nil {
		base.Log.Error(err)
		return ""
	}
	return mediaId
}

func (this *MatchService) Week(wcClient *core.Client) string {
	listData := this.RecommendService.ListData()
	articles := make([]material.Article, 1)
	first := material.Article{}
	first.Title = "本周战绩"
	first.Digest = "20191216-20191219"
	first.ThumbMediaId = "chP-LBQxy9SVbAFjwZ4QEpXfn8ShAn52EzP4-TrWvrM"
	first.ShowCoverPic = 1
	//图文消息的原文地址，即点击“阅读原文”后的URL
	first.ContentSourceURL = "https://gitee.com/aoe5188/poem-parent"
	var first_content string

	for _, e := range listData {
		matchDateStr := e.MatchDate.Format("15点04分")
		first_content += fmt.Sprintf("%v %v %v vs %v 推荐:\r\n", matchDateStr, e.LeagueName, e.MainTeamId, e.GuestTeamId, e.PreResult)
	}
	first.Content = first_content
	articles[0] = first

	mediaId, err := material.AddNews(wcClient, &material.News{Articles: articles})
	if err != nil {
		base.Log.Error(err)
		return ""
	}
	return mediaId
}

func (this *MatchService) Month(wcClient *core.Client) string {
	listData := this.RecommendService.ListData()
	articles := make([]material.Article, 1)
	first := material.Article{}
	first.Title = "本月战绩"
	first.Digest = "20191201-20191231"
	first.ThumbMediaId = "chP-LBQxy9SVbAFjwZ4QEpXfn8ShAn52EzP4-TrWvrM"
	first.ShowCoverPic = 1
	//图文消息的原文地址，即点击“阅读原文”后的URL
	first.ContentSourceURL = "https://gitee.com/aoe5188/poem-parent"
	var first_content string

	for _, e := range listData {
		matchDateStr := e.MatchDate.Format("15点04分")
		first_content += fmt.Sprintf("%v %v %v vs %v 推荐:\r\n", matchDateStr, e.LeagueName, e.MainTeamId, e.GuestTeamId, e.PreResult)
	}
	first.Content = first_content
	articles[0] = first

	mediaId, err := material.AddNews(wcClient, &material.News{Articles: articles})
	if err != nil {
		base.Log.Error(err)
		return ""
	}
	return mediaId
}

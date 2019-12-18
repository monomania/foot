package controller

import (
	"encoding/json"
	"fmt"
	_ "github.com/astaxie/beego"
	material2 "github.com/silenceper/wechat/material"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-core/common/base/controller"
	"tesou.io/platform/foot-parent/foot-core/module/analy/service"
)

type MaterialController struct {
	controller.BaseController
	service.RecommendService
}

func (this *MaterialController) AddImages() {
	material := wc.GetMaterial()
	mediaId, url, err := material.AddMaterial(material2.MediaTypeImage, "")
	if err != nil {
		base.Log.Error(err)
		return
	}
	base.Log.Info(fmt.Sprintf("mediaId is : %v ,url is : %v", mediaId, url))
}

func (this *MaterialController) AddNews() {
	listData := this.RecommendService.ListData()
	articles := make([]*material2.Article, len(listData))
	for _, e := range listData {
		bytes, _ := json.Marshal(e)
		base.Log.Warn("比赛信息:" + string(bytes))
		matchDateStr := e.MatchDate.Format("01月02日15点04分")
		article := new(material2.Article)
		article.Title = fmt.Sprintf("%v", matchDateStr)
		article.Digest = fmt.Sprintf("%v %v vs %v", e.LeagueName, e.MainTeamId, e.GuestTeamId)
		//-----
		article.ThumbMediaID = ""
		//-----
		article.ShowCoverPic = 1
		//图文消息的原文地址，即点击“阅读原文”后的URL
		article.ContentSourceURL = ""
		article.Content = string(bytes)
		articles = append(articles, article)
	}

	material := wc.GetMaterial()
	mediaId, err := material.AddNews(articles)
	if err != nil {
		base.Log.Error(err)
		return
	}
	base.Log.Info("mediaId is : ", mediaId)
}

package controller

import (
	_ "github.com/astaxie/beego"
	material2 "github.com/silenceper/wechat/material"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-core/common/base/controller"
)

type MaterialController struct {
	controller.BaseController
}

func (this *MaterialController) AddNews() {
	material := wc.GetMaterial()
	articles := make([]*material2.Article, 0)
	for i := 0; i < 5; i++ {
		article := new(material2.Article)
		
		articles = append(articles, article)
	}
	mediaId, err := material.AddNews(articles)
	if err != nil {
		base.Log.Error(err)
		return
	}
	base.Log.Info("mediaId is : ", mediaId)
}

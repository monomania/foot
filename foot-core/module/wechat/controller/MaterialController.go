package controller

import (
	"fmt"
	_ "github.com/astaxie/beego"
	"github.com/chanxuehong/wechat/mp/material"
	"io/ioutil"
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-core/common/base/controller"
	"tesou.io/platform/foot-parent/foot-core/module/wechat/service"
)

type MaterialController struct {
	controller.BaseController
	service.MatchService
}

func (this *MaterialController) Images() {

	infos, err := ioutil.ReadDir("assets/img")
	if err != nil {
		base.Log.Error(err)
		return
	}
	result := []string{}
	for _, e := range infos {
		name := e.Name()
		if !strings.HasSuffix(name, ".jpg") {
			continue
		}
		fileName := "assets/" + name
		mediaId, url, err := material.UploadImage(wcClient, fileName)
		if err != nil {
			base.Log.Error(err)
			return
		}
		data := fmt.Sprintf("fileName is: %v,mediaId is : %v ,url is : %v", fileName, mediaId, url)
		base.Log.Info(data)
		result = append(result, data)
	}

	this.Data["json"] = result
	this.ServeJSON()
}

func (this *MaterialController) News() {
	result := []string{}
	//today
	mediaId := this.MatchService.Today(wcClient)
	data := fmt.Sprintf("today mediaId is : %v", mediaId)
	base.Log.Info(data)
	result = append(result, data)
	//week
	mediaId = this.MatchService.Week(wcClient)
	data = fmt.Sprintf("week mediaId is : %v", mediaId)
	base.Log.Info(data)
	result = append(result, data)
	//month
	mediaId = this.MatchService.Month(wcClient)
	data = fmt.Sprintf("month mediaId is : %v", mediaId)
	base.Log.Info(data)
	result = append(result, data)

	this.Data["json"] = result
	this.ServeJSON()
}

func (this *MaterialController) ModifyNews() {
	this.MatchService.ModifyWeek(wcClient)
	this.Data["json"] = "ok"
	this.ServeJSON()
}

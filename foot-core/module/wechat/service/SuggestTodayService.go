package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/mp/material"
	"html/template"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/suggest/vo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	"tesou.io/platform/foot-parent/foot-core/common/utils"
	"tesou.io/platform/foot-parent/foot-core/module/analy/service"
	service2 "tesou.io/platform/foot-parent/foot-core/module/suggest/service"
	"time"
)

type SuggestTodayService struct {
	mysql.BaseService
	service.AnalyService
	service2.SuggestService
}

func preResultStr(val int) string {
	if 3 == val {
		return "主"
	} else if 1 == val {
		return "平"
	} else if 0 == val {
		return "客"
	}
	return "-"
}

func alFlagStr(str string) string {
	if "Asia20191206Service" == str {
		return "A1"
	} else if "Euro20191206Service" == str {
		return "E1"
	} else if "Euro20191212Service" == str {
		return "E2"
	}
	return "XX"
}

func getFuncMap() map[string]interface{} {
	funcMap := template.FuncMap{
		"preResultStr": preResultStr,
		"alFlagStr": alFlagStr,
	}
	return funcMap
}

/**
今日推荐
 */
func (this *SuggestTodayService) Today(wcClient *core.Client) string {
	listData := this.AnalyService.ListDefaultData()
	articles := make([]material.Article, 1)
	first := material.Article{}
	matchDateStr := time.Now().Format("01月02日")
	first.Title = fmt.Sprintf("今日推荐 %v", matchDateStr)
	first.Digest = fmt.Sprintf("%d场赛事", len(listData))
	first.ThumbMediaId = "chP-LBQxy9SVbAFjwZ4QEuxc8rI6Dy-bm5n3yZbsuJA"
	first.ShowCoverPic = 0
	//图文消息的原文地址，即:击“阅读原文”后的URL
	first.ContentSourceURL = "https://gitee.com/aoe5188/poem-parent"
	var first_content string
	for _, e := range listData {
		bytes, _ := json.Marshal(e)
		first_content += string(bytes) + "<br/>"
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

/**
更新今天推荐
 */
func (this *SuggestTodayService) ModifyToday(wcClient *core.Client) {
	param := new(vo.SuggestVO)
	now := time.Now()
	h12, _ := time.ParseDuration("-12h")
	beginDate := now.Add(h12)
	param.BeginDateStr = beginDate.Format("2006-01-02 15:04:05")
	h12, _ = time.ParseDuration("24h")
	endDate := now.Add(h12)
	param.EndDateStr = endDate.Format("2006-01-02 15:04:05")
	//今日推荐
	tempList := this.SuggestService.Query(param)
	//更新推送
	first := material.Article{}
	first.Title = fmt.Sprintf("今日推荐 %v", time.Now().Format("01月02日"))
	first.Digest = fmt.Sprintf("%d场赛事", len(tempList))
	first.ThumbMediaId = "chP-LBQxy9SVbAFjwZ4QEuxc8rI6Dy-bm5n3yZbsuJA"
	first.ContentSourceURL = "https://gitee.com/aoe5188/poem-parent"
	first.Author = utils.GetVal("wechat", "author")

	temp := vo.TodayVO{}
	temp.BeginDateStr = param.BeginDateStr
	temp.EndDateStr = param.EndDateStr
	temp.DataDateStr = now.Format("2006-01-02 15:04:05")
	temp.DataList = make([]vo.SuggestVO, len(tempList))
	for i, e := range tempList {
		e.MatchDateStr = e.MatchDate.Format("02号15:04")
		temp.DataList[i] = *e
	}

	var buf bytes.Buffer
	tpl, err := template.New("today.html").Funcs(getFuncMap()).ParseFiles("assets/wechat/html/today.html")
	if err != nil {
		base.Log.Error(err)
	}
	if err := tpl.Execute(&buf, &temp); err != nil {
		base.Log.Fatal(err)
	}
	first.Content = buf.String()

	err = material.UpdateNews(wcClient, "chP-LBQxy9SVbAFjwZ4QEoZGbUZaNED2Mf9jJauKvGo", 0, &first)
	if err != nil {
		base.Log.Error(err)
	}
}

/**
今日赛事分析
 */
func (this *SuggestTodayService) ModifyTodayDetail(wcClient *core.Client) {
	param := new(vo.SuggestVO)
	now := time.Now()
	h12, _ := time.ParseDuration("-12h")
	beginDate := now.Add(h12)
	param.BeginDateStr = beginDate.Format("2006-01-02 15:04:05")
	h12, _ = time.ParseDuration("24h")
	endDate := now.Add(h12)
	param.EndDateStr = endDate.Format("2006-01-02 15:04:05")
	//今日推荐
	tempList := this.SuggestService.Query(param)
	//更新推送
	first := material.Article{}
	first.Title = fmt.Sprintf("赛事解析")
	first.Digest = fmt.Sprintf("赛事的模型算法解析")
	first.ThumbMediaId = "chP-LBQxy9SVbAFjwZ4QEgIU_dXnFnXHvYzocwCpkM4"
	first.ContentSourceURL = "https://gitee.com/aoe5188/poem-parent"
	first.Author = utils.GetVal("wechat", "author")

	temp := vo.TodayVO{}
	temp.BeginDateStr = param.BeginDateStr
	temp.EndDateStr = param.EndDateStr
	temp.DataDateStr = now.Format("2006-01-02 15:04:05")
	temp.DataList = make([]vo.SuggestVO, len(tempList))
	for i, e := range tempList {
		e.MatchDateStr = e.MatchDate.Format("02号15:04")
		temp.DataList[i] = *e
	}

	var buf bytes.Buffer
	tpl, err := template.New("today_detail.html").Funcs(getFuncMap()).ParseFiles("assets/wechat/html/today_detail.html")
	if err != nil {
		base.Log.Error(err)
	}
	if err := tpl.Execute(&buf, &temp); err != nil {
		base.Log.Fatal(err)
	}
	first.Content = buf.String()

	err = material.UpdateNews(wcClient, "chP-LBQxy9SVbAFjwZ4QEoZGbUZaNED2Mf9jJauKvGo", 1, &first)
	if err != nil {
		base.Log.Error(err)
	}
}

/**
今日待选池比赛
 */
func (this *SuggestTodayService) ModifyTodayTbs(wcClient *core.Client) {
	param := new(vo.SuggestVO)
	now := time.Now()
	h12, _ := time.ParseDuration("-12h")
	beginDate := now.Add(h12)
	param.BeginDateStr = beginDate.Format("2006-01-02 15:04:05")
	h12, _ = time.ParseDuration("24h")
	endDate := now.Add(h12)
	param.EndDateStr = endDate.Format("2006-01-02 15:04:05")
	//今日推荐
	tempList := this.SuggestService.QueryTbs(param)
	//更新推送
	first := material.Article{}
	first.Title = fmt.Sprintf("待选池比赛")
	first.Digest = fmt.Sprintf("%d场赛事", len(tempList))
	first.ThumbMediaId = "chP-LBQxy9SVbAFjwZ4QEpOPdIm42ibP0pbNFt6VtAI"
	first.ContentSourceURL = "https://gitee.com/aoe5188/poem-parent"
	first.Author = utils.GetVal("wechat", "author")

	temp := vo.TodayVO{}
	temp.BeginDateStr = param.BeginDateStr
	temp.EndDateStr = param.EndDateStr
	temp.DataDateStr = now.Format("2006-01-02 15:04:05")
	temp.DataList = make([]vo.SuggestVO, len(tempList))
	for i, e := range tempList {
		e.MatchDateStr = e.MatchDate.Format("02号15:04")
		temp.DataList[i] = *e
	}

	var buf bytes.Buffer
	tpl, err := template.New("today_tbs.html").Funcs(getFuncMap()).ParseFiles("assets/wechat/html/today_tbs.html")
	if err != nil {
		base.Log.Error(err)
	}
	if err := tpl.Execute(&buf, &temp); err != nil {
		base.Log.Fatal(err)
	}
	first.Content = buf.String()

	err = material.UpdateNews(wcClient, "chP-LBQxy9SVbAFjwZ4QEoZGbUZaNED2Mf9jJauKvGo", 2, &first)
	if err != nil {
		base.Log.Error(err)
	}
}

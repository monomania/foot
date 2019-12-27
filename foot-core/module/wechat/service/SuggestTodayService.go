package service

import (
	"bytes"
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
		return "主胜"
	} else if 1 == val {
		return "平"
	} else if 0 == val {
		return "客胜"
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
	}else if "Euro20191226Service" == str{
		return "Q1"
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

func getAlFlag() string{
	al_flag := utils.GetVal("wechat", "al_flag")
	return al_flag;
}

const contentSourceURL = "https://gitee.com/aoe5188/foot"
const today_thumbMediaId = "chP-LBQxy9SVbAFjwZ4QEuxc8rI6Dy-bm5n3yZbsuJA"
const today_detail_thumbMediaId = "chP-LBQxy9SVbAFjwZ4QEgIU_dXnFnXHvYzocwCpkM4"
const today_tbs_thumbMediaId = "chP-LBQxy9SVbAFjwZ4QEpOPdIm42ibP0pbNFt6VtAI"
const today_mediaId = "chP-LBQxy9SVbAFjwZ4QEoZGbUZaNED2Mf9jJauKvGo"
/**
今日推荐
 */
func (this *SuggestTodayService) Today(wcClient *core.Client) string {
	articles := make([]material.Article, 1)
	first := material.Article{}
	matchDateStr := time.Now().Format("01月02日")
	first.Title = fmt.Sprintf("今日推荐 %v", matchDateStr)
	first.Digest = fmt.Sprintf("%d场赛事", 0)
	first.ThumbMediaId = today_thumbMediaId
	first.ShowCoverPic = 0
	//图文消息的原文地址，即:击“阅读原文”后的URL
	first.ContentSourceURL = today_thumbMediaId
	first.Content = "first_content"
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
	param.AlFlag = getAlFlag()
	tempList := this.SuggestService.Query(param)
	//更新推送
	first := material.Article{}
	first.Title = fmt.Sprintf("今日推荐 %v", time.Now().Format("01月02日"))
	first.Digest = fmt.Sprintf("%d场赛事", len(tempList))
	first.ThumbMediaId = today_thumbMediaId
	first.ContentSourceURL = contentSourceURL
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

	err = material.UpdateNews(wcClient, today_mediaId, 0, &first)
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
	param.AlFlag = getAlFlag()
	tempList := this.SuggestService.Query(param)
	//更新推送
	first := material.Article{}
	first.Title = fmt.Sprintf("赛事解析")
	first.Digest = fmt.Sprintf("赛事的模型算法解析")
	first.ThumbMediaId = today_detail_thumbMediaId
	first.ContentSourceURL = contentSourceURL
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

	err = material.UpdateNews(wcClient, today_mediaId, 1, &first)
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
	h12, _ := time.ParseDuration("-48h")
	beginDate := now.Add(h12)
	param.BeginDateStr = beginDate.Format("2006-01-02 15:04:05")
	h12, _ = time.ParseDuration("24h")
	endDate := now.Add(h12)
	param.EndDateStr = endDate.Format("2006-01-02 15:04:05")
	//今日推荐
	param.AlFlag = getAlFlag()
	tempList := this.SuggestService.QueryTbs(param)
	//更新推送
	first := material.Article{}
	first.Title = fmt.Sprintf("待选池比赛")
	first.Digest = fmt.Sprintf("%d场赛事", len(tempList))
	first.ThumbMediaId = today_tbs_thumbMediaId
	first.ContentSourceURL = contentSourceURL
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

	err = material.UpdateNews(wcClient, today_mediaId, 2, &first)
	if err != nil {
		base.Log.Error(err)
	}
}

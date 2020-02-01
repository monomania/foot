package service

import (
	"bytes"
	"fmt"
	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/mp/material"
	"html/template"
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/suggest/vo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	"tesou.io/platform/foot-parent/foot-core/common/utils"
	constants2 "tesou.io/platform/foot-parent/foot-core/module/analy/constants"
	"tesou.io/platform/foot-parent/foot-core/module/analy/service"
	"tesou.io/platform/foot-parent/foot-core/module/spider/constants"
	service2 "tesou.io/platform/foot-parent/foot-core/module/suggest/service"
	"time"
)

type SuggestTodayService struct {
	mysql.BaseService
	service.AnalyService
	service2.SuggestService
}

func preResultStr(val int, al_flag string) string {
	var result string
	if "E2" == al_flag {
		if 3 == val {
			result = "胜平"
		} else if 1 == val {
			result = "平"
		} else if 0 == val {
			result = "负平"
		}
	} else {
		if 3 == val {
			result = "主"
		} else if 1 == val {
			result = "走"
		} else if 0 == val {
			result = "客"
		}
	}
	return result
}

func color(str string) string {
	if "A1" == str {
		return "orange"
	} else if "C1" == str {
		return "yellow"
	} else if "E1" == str {
		return "blue"
	} else if "E2" == str {
		return "darkblue"
	} else if "Q1" == str {
		return "olivedrab"
	}
	return "XX"
}

func resultColor(str string) string {
	if str == constants2.HIT || str == constants2.HIT_1 {
		return "red"
	} else if str == constants2.UNHIT {
		return "gray"
	} else if str == constants2.WALKING_PLATE {
		return "greenyellow"
	}
	return "XX"
}

func getFuncMap() map[string]interface{} {
	funcMap := template.FuncMap{
		"preResultStr": preResultStr,
		"color":        color,
		"resultColor":  resultColor,
	}
	return funcMap
}

func getAlFlag() string {
	al_flag := utils.GetVal("wechat", "al_flag")
	return al_flag;
}

func getMainAlFlag() string {
	al_flag := utils.GetVal("wechat", "main_al_flag")
	return al_flag;
}

const contentSourceURL = "https://github.com/monomania/foot"
const today_thumbMediaId = "chP-LBQxy9SVbAFjwZ4QEuxc8rI6Dy-bm5n3yZbsuJA"
const today_detail_thumbMediaId = "chP-LBQxy9SVbAFjwZ4QEgIU_dXnFnXHvYzocwCpkM4"
const today_c1_thumbMediaId = "chP-LBQxy9SVbAFjwZ4QEgIU_dXnFnXHvYzocwCpkM4"
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
	param := new(vo.SuggStubVO)
	now := time.Now()
	h12, _ := time.ParseDuration("-24h")
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

	temp := vo.TTodayVO{}
	temp.SpiderDateStr = constants.SpiderDateStr
	temp.FullSpiderDateStr = constants.FullSpiderDateStr
	temp.BeginDateStr = param.BeginDateStr
	temp.EndDateStr = param.EndDateStr
	temp.DataDateStr = now.Format("2006-01-02 15:04:05")
	temp.DataList = make([]vo.SuggStubVO, len(tempList))
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
	param := new(vo.SuggStubVO)
	now := time.Now()
	h12, _ := time.ParseDuration("-24h")
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
	first.Title = fmt.Sprintf("赛事解析-A1,C1,E2")
	first.Digest = fmt.Sprintf("赛事的模型算法解析")
	first.ThumbMediaId = today_detail_thumbMediaId
	first.ContentSourceURL = contentSourceURL
	first.Author = utils.GetVal("wechat", "author")

	temp := vo.TTodayVO{}
	temp.SpiderDateStr = constants.SpiderDateStr
	temp.FullSpiderDateStr = constants.FullSpiderDateStr
	temp.BeginDateStr = param.BeginDateStr
	temp.EndDateStr = param.EndDateStr
	temp.DataDateStr = now.Format("2006-01-02 15:04:05")
	temp.DataList = make([]vo.SuggStubVO, len(tempList))
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
今日C1待选池比赛
 */
func (this *SuggestTodayService) ModifyTodayC1(wcClient *core.Client) {
	param := new(vo.SuggStubVO)
	now := time.Now()
	h12, _ := time.ParseDuration("-96h")
	beginDate := now.Add(h12)
	param.BeginDateStr = beginDate.Format("2006-01-02 15:04:05")
	h12, _ = time.ParseDuration("24h")
	endDate := now.Add(h12)
	param.EndDateStr = endDate.Format("2006-01-02 15:04:05")
	param.AlFlags = []string{"C1"}
	tempList := this.SuggestService.QueryTbs(param)
	//更新推送
	first := material.Article{}
	first.Title = fmt.Sprintf("推荐场次-C1")
	first.Digest = fmt.Sprintf("%d场赛事", len(tempList))
	first.ThumbMediaId = today_c1_thumbMediaId
	first.ContentSourceURL = contentSourceURL
	first.Author = utils.GetVal("wechat", "author")

	temp := vo.TTodayVO{}
	temp.SpiderDateStr = constants.SpiderDateStr
	temp.FullSpiderDateStr = constants.FullSpiderDateStr
	temp.BeginDateStr = param.BeginDateStr
	temp.EndDateStr = param.EndDateStr
	temp.DataDateStr = now.Format("2006-01-02 15:04:05")
	temp.DataList = make([]vo.SuggStubVO, len(tempList))
	for i, e := range tempList {
		e.MatchDateStr = e.MatchDate.Format("02号15:04")
		temp.DataList[i] = *e
	}

	var buf bytes.Buffer
	tpl, err := template.New("today_c1.html").Funcs(getFuncMap()).ParseFiles("assets/wechat/html/today_c1.html")
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
今日赛事分析
 */
func (this *SuggestTodayService) ModifyTodayDetailNew(wcClient *core.Client) {
	param := new(vo.SuggStubDetailVO)
	now := time.Now()
	h12, _ := time.ParseDuration("-24h")
	beginDate := now.Add(h12)
	param.BeginDateStr = beginDate.Format("2006-01-02 15:04:05")
	h12, _ = time.ParseDuration("24h")
	endDate := now.Add(h12)
	param.EndDateStr = endDate.Format("2006-01-02 15:04:05")
	//今日推荐
	param.AlFlag = getAlFlag()
	tempList := this.SuggestService.QueryDetail(param)
	//更新推送
	first := material.Article{}
	first.Title = fmt.Sprintf("赛事解析-A1,C1,E2")
	first.Digest = fmt.Sprintf("赛事的模型算法解析")
	first.ThumbMediaId = today_detail_thumbMediaId
	first.ContentSourceURL = contentSourceURL
	first.Author = utils.GetVal("wechat", "author")

	temp := vo.TTodayDetailVO{}
	temp.SpiderDateStr = constants.SpiderDateStr
	temp.FullSpiderDateStr = constants.FullSpiderDateStr
	temp.BeginDateStr = param.BeginDateStr
	temp.EndDateStr = param.EndDateStr
	temp.DataDateStr = now.Format("2006-01-02 15:04:05")
	temp.DataList = make([]vo.SuggStubDetailVO, len(tempList))
	for i, e := range tempList {
		e.MatchDateStr = e.MatchDate.Format("02号15:04")
		temp.DataList[i] = *e
	}

	teplate_paths := []string{
		"assets/wechat/html/today_detail_new.html",
		"assets/common/template/analycontent/001.html",
		"assets/common/template/analycontent/002.html",
		"assets/common/template/analycontent/004.html",
		"assets/common/template/analycontent/005.html",
		"assets/common/template/analycontent/footer.html",
		"assets/common/template/analycontent/wechat_today_detail_new.html",
	}

	var buf bytes.Buffer
	tpl, err := template.New("today_detail_new.html").Funcs(getFuncMap()).ParseFiles(teplate_paths...)
	//tpl, err := template.New("wechat_today_detail_new.html").Funcs(getFuncMap()).ParseFiles("assets/common/template/wechat_today_detail_new.html")
	if err != nil {
		base.Log.Error(err)
	}
	if err := tpl.Execute(&buf, &temp); err != nil {
		base.Log.Fatal(err)
	}
	first.Content = buf.String()
	first.Content = strings.TrimSpace(first.Content)
	first.Content = strings.ReplaceAll(first.Content,"\r\n","")

	err = material.UpdateNews(wcClient, today_mediaId, 2, &first)
	if err != nil {
		base.Log.Error(err)
	}
}


/**
今日A1待选池比赛
 */
func (this *SuggestTodayService) ModifyTodayA1(wcClient *core.Client) {
	param := new(vo.SuggStubVO)
	now := time.Now()
	h12, _ := time.ParseDuration("-96h")
	beginDate := now.Add(h12)
	param.BeginDateStr = beginDate.Format("2006-01-02 15:04:05")
	h12, _ = time.ParseDuration("24h")
	endDate := now.Add(h12)
	param.EndDateStr = endDate.Format("2006-01-02 15:04:05")
	param.AlFlags = []string{"A1"}
	tempList := this.SuggestService.QueryTbs(param)
	//更新推送
	first := material.Article{}
	first.Title = fmt.Sprintf("待选场次-A1")
	first.Digest = fmt.Sprintf("%d场赛事", len(tempList))
	first.ThumbMediaId = today_tbs_thumbMediaId
	first.ContentSourceURL = contentSourceURL
	first.Author = utils.GetVal("wechat", "author")

	temp := vo.TTodayVO{}
	temp.SpiderDateStr = constants.SpiderDateStr
	temp.FullSpiderDateStr = constants.FullSpiderDateStr
	temp.BeginDateStr = param.BeginDateStr
	temp.EndDateStr = param.EndDateStr
	temp.DataDateStr = now.Format("2006-01-02 15:04:05")
	temp.DataList = make([]vo.SuggStubVO, len(tempList))
	for i, e := range tempList {
		e.MatchDateStr = e.MatchDate.Format("02号15:04")
		temp.DataList[i] = *e
	}

	var buf bytes.Buffer
	tpl, err := template.New("today_a1.html").Funcs(getFuncMap()).ParseFiles("assets/wechat/html/today_a1.html")
	if err != nil {
		base.Log.Error(err)
	}
	if err := tpl.Execute(&buf, &temp); err != nil {
		base.Log.Fatal(err)
	}
	first.Content = buf.String()

	err = material.UpdateNews(wcClient, today_mediaId, 3, &first)
	if err != nil {
		base.Log.Error(err)
	}
}

/**
今日待选池比赛
 */
func (this *SuggestTodayService) ModifyTodayTbs(wcClient *core.Client) {
	param := new(vo.SuggStubVO)
	now := time.Now()
	h12, _ := time.ParseDuration("-48h")
	beginDate := now.Add(h12)
	param.BeginDateStr = beginDate.Format("2006-01-02 15:04:05")
	h12, _ = time.ParseDuration("24h")
	endDate := now.Add(h12)
	param.EndDateStr = endDate.Format("2006-01-02 15:04:05")
	param.AlFlags = []string{"E1", "E2", "Q1"}
	tempList := this.SuggestService.QueryTbs(param)
	//更新推送
	first := material.Article{}
	first.Title = fmt.Sprintf("待选场次-E1,E2,Q1")
	first.Digest = fmt.Sprintf("%d场赛事", len(tempList))
	first.ThumbMediaId = today_tbs_thumbMediaId
	first.ContentSourceURL = contentSourceURL
	first.Author = utils.GetVal("wechat", "author")

	temp := vo.TTodayVO{}
	temp.SpiderDateStr = constants.SpiderDateStr
	temp.FullSpiderDateStr = constants.FullSpiderDateStr
	temp.BeginDateStr = param.BeginDateStr
	temp.EndDateStr = param.EndDateStr
	temp.DataDateStr = now.Format("2006-01-02 15:04:05")
	temp.DataList = make([]vo.SuggStubVO, len(tempList))
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

	err = material.UpdateNews(wcClient, today_mediaId, 4, &first)
	if err != nil {
		base.Log.Error(err)
	}
}


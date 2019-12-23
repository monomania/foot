package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/mp/material"
	"html/template"
	"strconv"
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/suggest/vo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	"tesou.io/platform/foot-parent/foot-core/common/utils"
	"tesou.io/platform/foot-parent/foot-core/module/analy/constants"
	"tesou.io/platform/foot-parent/foot-core/module/analy/service"
	service2 "tesou.io/platform/foot-parent/foot-core/module/suggest/service"
	"time"
)

type MatchService struct {
	mysql.BaseService
	service.AnalyService
	service2.SuggestService
}

func (this *MatchService) Today(wcClient *core.Client) string {
	listData := this.AnalyService.ListDefaultData()
	articles := make([]material.Article, 1)
	first := material.Article{}
	matchDateStr := time.Now().Format("01月02日")
	first.Title = fmt.Sprintf("今日推荐 %v", matchDateStr)
	first.Digest = fmt.Sprintf("%d场赛事", len(listData))
	first.ThumbMediaId = "chP-chP-222222-bm5n3yZbsuJA"
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

func (this *MatchService) ModifyToday(wcClient *core.Client) {
	param := new(vo.SuggestVO)
	now := time.Now()
	h12, _ := time.ParseDuration("-12h")
	beginDate := now.Add(h12)
	param.BeginDateStr = beginDate.Format("2006-01-02 15:04:05")
	h12, _ = time.ParseDuration("12h")
	endDate := now.Add(h12)
	param.EndDateStr = endDate.Format("2006-01-02 15:04:05")
	//今日推荐
	tempList := this.SuggestService.Query(param)
	//更新推送
	first := material.Article{}
	first.Title = fmt.Sprintf("今日推荐 %v", time.Now().Format("01月02日"))
	first.Digest = fmt.Sprintf("%d场赛事", len(tempList))
	first.ThumbMediaId = "chP-chP-222222-bm5n3yZbsuJA"
	first.ContentSourceURL = "https://gitee.com/aoe5188/poem-parent"
	first.Author = utils.GetVal("wechat","author")

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
	tpl, err := template.ParseFiles("assets/wechat/html/today.html")
	if err != nil {
		base.Log.Error(err)
	}
	if err := tpl.Execute(&buf, &temp); err != nil {
		base.Log.Fatal(err)
	}
	first.Content = buf.String()

	err = material.UpdateNews(wcClient, "chP-chP-222222", 0, &first)
	if err != nil {
		base.Log.Error(err)
	}
}

func (this *MatchService) ModifyTodayDetail(wcClient *core.Client) {
	param := new(vo.SuggestVO)
	now := time.Now()
	h12, _ := time.ParseDuration("-12h")
	beginDate := now.Add(h12)
	param.BeginDateStr = beginDate.Format("2006-01-02 15:04:05")
	h12, _ = time.ParseDuration("12h")
	endDate := now.Add(h12)
	param.EndDateStr = endDate.Format("2006-01-02 15:04:05")
	//今日推荐
	tempList := this.SuggestService.Query(param)
	//更新推送
	first := material.Article{}
	first.Title = fmt.Sprintf("赛事解析")
	first.Digest = fmt.Sprintf("赛事的模型算法解析")
	first.ThumbMediaId = "chP-chP-222222"
	first.ContentSourceURL = "https://gitee.com/aoe5188/poem-parent"
	first.Author = utils.GetVal("wechat","author")

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
	tpl, err := template.ParseFiles("assets/wechat/html/today_detail.html")
	if err != nil {
		base.Log.Error(err)
	}
	if err := tpl.Execute(&buf, &temp); err != nil {
		base.Log.Fatal(err)
	}
	first.Content = buf.String()

	err = material.UpdateNews(wcClient, "chP-chP-222222", 1, &first)
	if err != nil {
		base.Log.Error(err)
	}
}


func (this *MatchService) ModifyTodayTbs(wcClient *core.Client) {
	param := new(vo.SuggestVO)
	now := time.Now()
	h12, _ := time.ParseDuration("-12h")
	beginDate := now.Add(h12)
	param.BeginDateStr = beginDate.Format("2006-01-02 15:04:05")
	h12, _ = time.ParseDuration("12h")
	endDate := now.Add(h12)
	param.EndDateStr = endDate.Format("2006-01-02 15:04:05")
	param.MinHitCount = 1
	hit_count_str := utils.GetVal(constants.SECTION_NAME, "hit_count")
	hit_count, _ := strconv.Atoi(hit_count_str)
	param.MaxHitCount = hit_count - 1
	//今日推荐
	tempList := this.SuggestService.Query(param)
	//更新推送
	first := material.Article{}
	first.Title = fmt.Sprintf("待选池比赛")
	first.Digest = fmt.Sprintf("%d场赛事", len(tempList))
	first.ThumbMediaId = "chP-chP-222222"
	first.ContentSourceURL = "https://gitee.com/aoe5188/poem-parent"
	first.Author = utils.GetVal("wechat","author")

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
	tpl, err := template.ParseFiles("assets/wechat/html/today_tbs.html")
	if err != nil {
		base.Log.Error(err)
	}
	if err := tpl.Execute(&buf, &temp); err != nil {
		base.Log.Fatal(err)
	}
	first.Content = buf.String()

	err = material.UpdateNews(wcClient, "chP-chP-222222", 2, &first)
	if err != nil {
		base.Log.Error(err)
	}
}

func (this *MatchService) Week(wcClient *core.Client) string {
	listData := this.AnalyService.ListDefaultData()
	articles := make([]material.Article, 1)
	first := material.Article{}
	first.Title = "最近七天战绩"
	first.Digest = "20191216-20191219"
	first.ThumbMediaId = "chP-chP-222222"

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

func (this *MatchService) ModifyWeek(wcClient *core.Client) {
	param := new(vo.SuggestVO)
	now := time.Now()
	h2, _ := time.ParseDuration("-2h")
	endDate := now.Add(h2)
	param.EndDateStr = endDate.Format("2006-01-02 15:04:05")
	h168, _ := time.ParseDuration("-168h")
	beginDate := now.Add(h168)
	param.BeginDateStr = beginDate.Format("2006-01-02 15:04:05")
	param.IsDesc = true
	//今日推荐
	tempList := this.SuggestService.Query(param)
	//更新推送
	first := material.Article{}
	first.Title = "最近七天战绩"
	first.Digest = fmt.Sprintf("%v-%v", beginDate.Format("2006年01月02日"), now.Format("2006年01月02日"))
	first.ThumbMediaId = "chP-chP-222222"
	first.ShowCoverPic = 0
	first.ContentSourceURL = "https://gitee.com/aoe5188/poem-parent"
	first.Author = utils.GetVal("wechat","author")
	temp := vo.WeekVO{}
	temp.BeginDateStr = param.BeginDateStr
	temp.EndDateStr = param.EndDateStr
	temp.MatchCount = int64(len(tempList))

	var redCount, walkCount, blackCount, linkRedCount, tempLinkRedCount, linkBlackCount, tempLinkBlackCount int64
	temp.DataList = make([]vo.SuggestVO, len(tempList))
	for i, e := range tempList {
		if strings.EqualFold("正确", e.Result) {
			redCount++
			tempLinkRedCount++
			tempLinkBlackCount = 0
		} else if strings.EqualFold("错误", e.Result) {
			blackCount++
			tempLinkBlackCount++
			tempLinkRedCount = 0
		} else {
			walkCount++
		}

		if tempLinkRedCount > linkRedCount {
			linkRedCount = tempLinkRedCount
		}
		if tempLinkBlackCount > linkBlackCount {
			linkBlackCount = tempLinkBlackCount
		}

		e.MatchDateStr = e.MatchDate.Format("02号15:04")
		temp.DataList[i] = *e
	}
	temp.RedCount = redCount
	temp.WalkCount = walkCount
	temp.BlackCount = blackCount
	temp.LinkRedCount = linkRedCount
	temp.LinkBlackCount = linkBlackCount

	var buf bytes.Buffer
	tpl, err := template.ParseFiles("assets/wechat/html/week.html")
	if err != nil {
		base.Log.Error(err)
	}
	if err := tpl.Execute(&buf, &temp); err != nil {
		base.Log.Fatal(err)
	}
	first.Content = buf.String()

	err = material.UpdateNews(wcClient, "chP-chP-222222-Ac", 0, &first)
	if err != nil {
		base.Log.Error(err)
	}
}

func (this *MatchService) Month(wcClient *core.Client) string {
	listData := this.AnalyService.ListDefaultData()
	articles := make([]material.Article, 1)
	first := material.Article{}
	first.Title = "最近一月战绩"
	first.Digest = "20191201-20191231"
	first.ThumbMediaId = "chP-chP-222222"

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

func (this *MatchService) ModifyMonth(wcClient *core.Client) {
	param := new(vo.SuggestVO)
	now := time.Now()
	h2, _ := time.ParseDuration("-2h")
	endDate := now.Add(h2)
	param.EndDateStr = endDate.Format("2006-01-02 15:04:05")
	h168, _ := time.ParseDuration("-720h")
	beginDate := now.Add(h168)
	param.BeginDateStr = beginDate.Format("2006-01-02 15:04:05")
	param.IsDesc = true
	//今日推荐
	tempList := this.SuggestService.Query(param)
	//更新推送
	first := material.Article{}
	first.Title = "最近一月战绩"
	first.Digest = fmt.Sprintf("%v-%v", beginDate.Format("2006年01月02日"), now.Format("2006年01月02日"))
	first.ThumbMediaId = "chP-222222"
	first.ContentSourceURL = "https://gitee.com/aoe5188/poem-parent"
	first.Author = utils.GetVal("wechat","author")

	temp := vo.WeekVO{}
	temp.BeginDateStr = param.BeginDateStr
	temp.EndDateStr = param.EndDateStr
	temp.MatchCount = int64(len(tempList))

	var redCount, walkCount, blackCount, linkRedCount, tempLinkRedCount, linkBlackCount, tempLinkBlackCount int64
	temp.DataList = make([]vo.SuggestVO, len(tempList))
	for i, e := range tempList {
		if strings.EqualFold("正确", e.Result) {
			redCount++
			tempLinkRedCount++
			tempLinkBlackCount = 0
		} else if strings.EqualFold("错误", e.Result) {
			blackCount++
			tempLinkBlackCount++
			tempLinkRedCount = 0
		} else {
			walkCount++
		}

		if tempLinkRedCount > linkRedCount {
			linkRedCount = tempLinkRedCount
		}
		if tempLinkBlackCount > linkBlackCount {
			linkBlackCount = tempLinkBlackCount
		}

		e.MatchDateStr = e.MatchDate.Format("02号15:04")
		temp.DataList[i] = *e
	}
	temp.RedCount = redCount
	temp.WalkCount = walkCount
	temp.BlackCount = blackCount
	temp.LinkRedCount = linkRedCount
	temp.LinkBlackCount = linkBlackCount

	var buf bytes.Buffer
	tpl, err := template.ParseFiles("assets/wechat/html/month.html")
	if err != nil {
		base.Log.Error(err)
	}
	if err := tpl.Execute(&buf, &temp); err != nil {
		base.Log.Fatal(err)
	}
	first.Content = buf.String()

	err = material.UpdateNews(wcClient, "chP-chP-222222-tpO49nJOac9ex8Q", 0, &first)
	if err != nil {
		base.Log.Error(err)
	}
}

package service

import (
	"bytes"
	"fmt"
	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/mp/material"
	"html/template"
	"strconv"
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	"tesou.io/platform/foot-parent/foot-api/module/suggest/vo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	"tesou.io/platform/foot-parent/foot-core/common/utils"
	constants2 "tesou.io/platform/foot-parent/foot-core/module/analy/constants"
	"tesou.io/platform/foot-parent/foot-core/module/analy/service"
	"tesou.io/platform/foot-parent/foot-core/module/spider/constants"
	service2 "tesou.io/platform/foot-parent/foot-core/module/suggest/service"
	"time"
)

type SuggestWeekService struct {
	mysql.BaseService
	service.AnalyService
	service2.SuggestService
}

const week_thumbMediaId = "chP-LBQxy9SVbAFjwZ4QEo81I0bHaY3YDYRwVGmf7o8"
const week_guts_thumbMediaId = "chP-LBQxy9SVbAFjwZ4QEthz2PGUtpVZdZccay_qCi0"
const week_mediaId = "chP-LBQxy9SVbAFjwZ4QEgcfOu5CZ67hiBgn5qnZ-Ac"

func (this *SuggestWeekService) Week(wcClient *core.Client) string {
	articles := make([]material.Article, 1)
	first := material.Article{}
	first.Title = "最近七天战绩"
	first.Digest = "20191216-20191219"
	first.ThumbMediaId = week_thumbMediaId

	first.Content = "first_content"
	articles[0] = first

	mediaId, err := material.AddNews(wcClient, &material.News{Articles: articles})
	if err != nil {
		base.Log.Error(err)
		return ""
	}
	return mediaId
}

func (this *SuggestWeekService) ModifyWeek(wcClient *core.Client) {
	param := new(vo.SuggStubVO)
	now := time.Now()
	h2, _ := time.ParseDuration("-2h")
	endDate := now.Add(h2)
	param.EndDateStr = endDate.Format("2006-01-02 15:04:05")
	h168, _ := time.ParseDuration("-168h")
	beginDate := now.Add(h168)
	param.BeginDateStr = beginDate.Format("2006-01-02 15:04:05")
	param.IsDesc = true
	//今日推荐
	param.AlFlag = getStatAlFlag()
	tempList := this.SuggestService.Query(param)
	//更新推送
	first := material.Article{}
	first.Title = "最近七天战绩"
	first.Digest = fmt.Sprintf("%v-%v", beginDate.Format("2006年01月02日"), now.Format("2006年01月02日"))
	first.ThumbMediaId = week_thumbMediaId
	first.ContentSourceURL = contentSourceURL
	first.ShowCoverPic = 0
	first.Author = utils.GetVal("wechat", "author")
	temp := vo.TWeekVO{}
	temp.SpiderDateStr = constants.SpiderDateStr
	temp.FullSpiderDateStr = constants.FullSpiderDateStr
	temp.BeginDateStr = param.BeginDateStr
	temp.EndDateStr = param.EndDateStr
	temp.MatchCount = int64(len(tempList))

	var redCount, walkCount, blackCount, linkRedCount, tempLinkRedCount, linkBlackCount, tempLinkBlackCount int64
	temp.DataList = make([]vo.SuggStubVO, len(tempList))
	for i, e := range tempList {
		if strings.EqualFold(constants2.HIT, e.Result) || strings.EqualFold(constants2.HIT_1, e.Result) {
			redCount++
			tempLinkRedCount++
			tempLinkBlackCount = 0
		} else if strings.EqualFold(constants2.UNHIT, e.Result) {
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
	val := float64(redCount) / (float64(redCount) + float64(blackCount)) * 100
	val, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", val), 64)
	temp.Val = strconv.FormatFloat(val, 'f', -1, 64) + "%"

	//计算主要模型胜率
	temp_main_alflag := getMainAlFlag()
	//计算单方向胜率
	var mainRedCount, mainBlackCount int64
	for _, e := range tempList {
		if !strings.Contains(temp_main_alflag,e.AlFlag){
			continue
		}
		last := new(pojo.MatchLast)
		last.Id = e.MatchId
		last.LeagueId = e.LeagueId
		last.MatchDate = e.MatchDate
		last.MainTeamId = e.MainTeam
		last.MainTeamGoals, _ = strconv.Atoi(e.MainTeamGoal)
		last.GuestTeamId = e.GuestTeam
		last.GuestTeamGoals, _ = strconv.Atoi(e.GuestTeamGoal)
		option := this.AnalyService.IsRight(last, &e.AnalyResult)
		if option == constants2.HIT ||  option == constants2.HIT_1{
			mainRedCount++
		}
		if option == constants2.UNHIT {
			mainBlackCount++
		}
	}
	temp.MainAlflag = temp_main_alflag
	temp.MainRedCount = mainRedCount
	temp.MainBlackCount = mainBlackCount
	mainVal := float64(mainRedCount) / (float64(mainRedCount) + float64(mainBlackCount)) * 100
	mainVal, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", mainVal), 64)
	temp.MainVal = strconv.FormatFloat(mainVal, 'f', -1, 64) + "%"

	var buf bytes.Buffer
	tpl, err := template.New("week.html").Funcs(getFuncMap()).ParseFiles("assets/wechat/html/week.html")
	if err != nil {
		base.Log.Error(err)
	}
	if err := tpl.Execute(&buf, &temp); err != nil {
		base.Log.Fatal(err)
	}
	first.Content = buf.String()

	err = material.UpdateNews(wcClient, week_mediaId, 0, &first)
	if err != nil {
		base.Log.Error(err)
	}
}


func (this *SuggestMonthService) ModifyGutsWeek(wcClient *core.Client) {
	param := new(vo.SuggStubVO)
	now := time.Now()
	h2, _ := time.ParseDuration("100h")
	endDate := now.Add(h2)
	param.EndDateStr = endDate.Format("2006-01-02 15:04:05")
	h168, _ := time.ParseDuration("-168h")
	beginDate := now.Add(h168)
	param.BeginDateStr = beginDate.Format("2006-01-02 15:04:05")
	param.IsDesc = true
	tempList := this.SuggestService.QueryGutsC2E2(param)
	//更新推送
	first := material.Article{}
	first.Title = "稳胆七天战绩"
	first.Digest = fmt.Sprintf("%v-%v", beginDate.Format("2006年01月02日"), now.Format("2006年01月02日"))
	first.ThumbMediaId = week_guts_thumbMediaId
	first.ContentSourceURL = contentSourceURL
	first.Author = utils.GetVal("wechat", "author")

	temp := vo.TMonthVO{}
	temp.SpiderDateStr = constants.SpiderDateStr
	temp.FullSpiderDateStr = constants.FullSpiderDateStr
	temp.BeginDateStr = param.BeginDateStr
	temp.EndDateStr = param.EndDateStr
	temp.MatchCount = int64(len(tempList))

	var redCount, walkCount, blackCount, linkRedCount, tempLinkRedCount, linkBlackCount, tempLinkBlackCount int64
	temp.DataList = make([]vo.SuggStubVO, len(tempList))
	for i, e := range tempList {
		if strings.EqualFold(constants2.HIT, e.Result) || strings.EqualFold(constants2.HIT_1, e.Result) {
			redCount++
			tempLinkRedCount++
			tempLinkBlackCount = 0
		} else if strings.EqualFold(constants2.UNHIT, e.Result) {
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
	val := float64(redCount) / (float64(redCount) + float64(blackCount)) * 100
	val, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", val), 64)
	temp.Val = strconv.FormatFloat(val, 'f', -1, 64) + "%"

	//计算单方向胜率
	var mainRedCount, mainBlackCount int64
	for _, e := range tempList {
		last := new(pojo.MatchLast)
		last.Id = e.MatchId
		last.LeagueId = e.LeagueId
		last.MatchDate = e.MatchDate
		last.MainTeamId = e.MainTeam
		last.MainTeamGoals, _ = strconv.Atoi(e.MainTeamGoal)
		last.GuestTeamId = e.GuestTeam
		last.GuestTeamGoals, _ = strconv.Atoi(e.GuestTeamGoal)
		option := this.AnalyService.IsRight(last, &e.AnalyResult)
		if option == constants2.HIT ||  option == constants2.HIT_1{
			mainRedCount++
		}
		if option == constants2.UNHIT {
			mainBlackCount++
		}
	}
	mainRedCount /= 2
	mainBlackCount /= 2
	temp.MainRedCount = mainRedCount
	temp.MainBlackCount = mainBlackCount
	value := float64(mainRedCount) / (float64(mainRedCount) + float64(mainBlackCount)) * 100
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	temp.MainVal = strconv.FormatFloat(value, 'f', -1, 64) + "%"

	var buf bytes.Buffer
	tpl, err := template.New("week_guts.html").Funcs(getFuncMap()).ParseFiles("assets/wechat/html/week_guts.html")
	if err != nil {
		base.Log.Error(err)
	}
	if err := tpl.Execute(&buf, &temp); err != nil {
		base.Log.Fatal(err)
	}
	first.Content = buf.String()

	err = material.UpdateNews(wcClient, week_mediaId, 1, &first)
	if err != nil {
		base.Log.Error(err)
	}
}
package proc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/pipeline"
	"github.com/hu17889/go_spider/core/spider"
	"regexp"
	"strconv"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-core/module/match/service"
	"tesou.io/platform/foot-parent/foot-spider/module/win007/down"
	"tesou.io/platform/foot-parent/foot-spider/module/win007/vo"
	"time"

	"strings"
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	"tesou.io/platform/foot-parent/foot-spider/module/win007"
)

type BaseFaceProcesser struct {
	service.BFScoreService
	service.BFBattleService
	service.BFFutureEventService

	MatchLastList      []*pojo.MatchLast
	Win007idMatchidMap map[string]string
}

func GetBaseFaceProcesser() *BaseFaceProcesser {
	return &BaseFaceProcesser{}
}

func (this *BaseFaceProcesser) Startup() {
	this.Win007idMatchidMap = map[string]string{}

	newSpider := spider.NewSpider(this, "BaseFaceProcesser")

	for _, v := range this.MatchLastList {
		i := v.Ext[win007.MODULE_FLAG]
		bytes, _ := json.Marshal(i)
		matchLastExt := new(pojo.MatchExt)
		json.Unmarshal(bytes, matchLastExt)

		win007_id := matchLastExt.Sid

		this.Win007idMatchidMap[win007_id] = v.Id

		url := strings.Replace(win007.WIN007_BASE_FACE_URL_PATTERN, "${matchId}", win007_id, 1)
		newSpider = newSpider.AddUrl(url, "html")
	}
	newSpider.SetDownloader(down.NewMWin007Downloader())
	newSpider = newSpider.AddPipeline(pipeline.NewPipelineConsole())
	newSpider.SetThreadnum(1).Run()
}

func (this *BaseFaceProcesser) Process(p *page.Page) {
	request := p.GetRequest()
	if !p.IsSucc() {
		base.Log.Info("URL:,", request.Url, p.Errormsg())
		return
	}

	var regex_temp = regexp.MustCompile(`(\d+).htm`)
	win007Id := strings.Split(regex_temp.FindString(request.Url), ".")[0]
	matchId := this.Win007idMatchidMap[win007Id]

	scoreList := this.score_process(matchId, p)
	battleList := this.battle_process(matchId, p)
	futureEventList := this.future_event_process(matchId, p)
	//保存数据
	this.BFScoreService.SaveList(scoreList)
	this.BFBattleService.SaveList(battleList)
	this.BFFutureEventService.SaveList(futureEventList)

}

//处理获取积分榜数据
func (this *BaseFaceProcesser) score_process(matchId string, p *page.Page) []interface{} {
	data_list_slice := make([]interface{}, 0)
	p.GetHtmlParser().Find(" table.mytable").EachWithBreak(func(i int, selection *goquery.Selection) bool {
		//只取前两个table
		if i > 1 {
			return false
		}

		ret, _ := selection.Html()
		fmt.Println(ret)

		selection.Find(" tr ").Each(func(i int, selection *goquery.Selection) {
			if i >= 1 {
				selection.Text()
				nodes := selection.Children().Nodes
				buf := bytes.Buffer{}
				buf.WriteString(nodes[0].Data)
				val0 := buf.String()

				buf.Reset()
				buf.WriteString(nodes[1].Data)
				val1 := buf.String()

				buf.Reset()
				buf.WriteString(nodes[2].Data)
				val2 := buf.String()

				buf.Reset()
				buf.WriteString(nodes[3].Data)
				val3 := buf.String()
				fmt.Println(val0, val1, val2, val3)
			}
		})
		return true
	})
	return data_list_slice
}

//处理对战数据获取
func (this *BaseFaceProcesser) battle_process(matchId string, p *page.Page) []interface{} {
	data_list_slice := make([]interface{}, 0)

	var hdata_str string
	p.GetHtmlParser().Find("script").Each(func(i int, selection *goquery.Selection) {
		text := selection.Text()
		if hdata_str == "" && strings.Contains(text, "var vsTeamInfo") {
			hdata_str = text
		} else {
			return
		}
	})
	if hdata_str == "" {
		return data_list_slice
	}

	// 获取script脚本中的，博彩公司信息
	hdata_str = strings.Replace(hdata_str, ";", "", 1)
	hdata_str = strings.Replace(hdata_str, "var vsTeamInfo = ", "", 1)
	base.Log.Info(hdata_str)

	var hdata_list = make([]*vo.BattleData, 0)
	json.Unmarshal(([]byte)(hdata_str), &hdata_list)

	//入库中
	for _, v := range hdata_list {
		battle := new(pojo.BFBattle)

		battle.MatchId = matchId
		battleMatchDate, _ := time.ParseInLocation("2006-01-02", v.Year+"-"+v.Date, time.Local)
		battle.BattleMatchDate = battleMatchDate
		battle.BattleLeagueId = v.SclassID
		battle.BattleMainTeamId = v.Home
		battle.BattleGuestTeamId = v.Guest

		half_goals := strings.Split(v.HT, "-")
		full_goals := strings.Split(v.FT, "-")
		battle.BattleMainTeamHalfGoals, _ = strconv.Atoi(half_goals[0])
		battle.BattleGuestTeamHalfGoals, _ = strconv.Atoi(half_goals[1])
		battle.BattleMainTeamGoals, _ = strconv.Atoi(full_goals[0])
		battle.BattleGuestTeamGoals, _ = strconv.Atoi(full_goals[1])

	}

	return data_list_slice
}

//处理获取示来对战数据
func (this *BaseFaceProcesser) future_event_process(matchId string, p *page.Page) []interface{} {
	data_list_slice := make([]interface{}, 0)
	p.GetHtmlParser().Find(" table.mytable3 tr").Each(func(i int, selection *goquery.Selection) {

	})

	return data_list_slice
}

func (this *BaseFaceProcesser) Finish() {
	base.Log.Info("基本面分析抓取解析完成 \r\n")

}

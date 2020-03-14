package proc

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/pipeline"
	"github.com/hu17889/go_spider/core/spider"
	"regexp"
	"strconv"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	service2 "tesou.io/platform/foot-parent/foot-core/module/elem/service"
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
	service.BFJinService
	service.BFFutureEventService
	service2.LeagueService
	//是否是单线程
	SingleThread bool
	MatchLastList      []*pojo.MatchLast
	Win007idMatchidMap map[string]string
}

func GetBaseFaceProcesser() *BaseFaceProcesser {
	processer := &BaseFaceProcesser{}
	processer.Init()
	return processer
}

func (this *BaseFaceProcesser) Init() {
	//初始化参数值
	this.Win007idMatchidMap = map[string]string{}
}

func (this *BaseFaceProcesser) Setup(temp *BaseFaceProcesser) {
	//设置参数值
}

func (this *BaseFaceProcesser) Startup() {

	var newSpider *spider.Spider
	processer := this
	newSpider = spider.NewSpider(processer, "BaseFaceProcesser")
	for i, v := range this.MatchLastList {

		if !this.SingleThread && i%10000 == 0 { //10000个比赛一个spider,一个赛季大概有30万场比赛,最多30spider
			processer = GetBaseFaceProcesser()
			processer.Setup(this)
			newSpider = spider.NewSpider(processer, "BaseFaceProcesser"+strconv.Itoa(i))
		}

		temp_flag := v.Ext[win007.MODULE_FLAG]
		bytes, _ := json.Marshal(temp_flag)
		matchLastExt := new(pojo.MatchExt)
		json.Unmarshal(bytes, matchLastExt)
		win007_id := matchLastExt.Sid

		processer.Win007idMatchidMap[win007_id] = v.Id

		url := strings.Replace(win007.WIN007_BASE_FACE_URL_PATTERN, "${matchId}", win007_id, 1)
		newSpider = newSpider.AddUrl(url, "html")
		if !this.SingleThread && i%10000 == 0 { //10000个比赛一个spider,一个赛季大概有30万场比赛,最多30spider
			newSpider.SetDownloader(down.NewMWin007Downloader())
			newSpider = newSpider.AddPipeline(pipeline.NewPipelineConsole())
			newSpider.SetSleepTime("rand", 100, 3000)
			newSpider.SetThreadnum(1).Run()
		}
	}

	newSpider.SetDownloader(down.NewMWin007Downloader())
	newSpider = newSpider.AddPipeline(pipeline.NewPipelineConsole())
	newSpider.SetSleepTime("rand", 100, 3000)
	newSpider.SetThreadnum(1).Run()

}

func (this *BaseFaceProcesser) Process(p *page.Page) {
	request := p.GetRequest()
	if !p.IsSucc() {
		base.Log.Error("URL:", request.Url, p.Errormsg())
		return
	}

	var regex_temp = regexp.MustCompile(`(\d+).htm`)
	win007Id := strings.Split(regex_temp.FindString(request.Url), ".")[0]
	matchId := this.Win007idMatchidMap[win007Id]

	//积分榜
	scoreSaveList := make([]interface{}, 0)
	scoreModifyList := make([]interface{}, 0)
	scoreList := this.score_process(matchId, p)
	for _, e := range scoreList {
		temp_id, exist := this.BFScoreService.Exist(e)
		if exist {
			e.Id = temp_id
			scoreModifyList = append(scoreModifyList, e)
		} else {
			scoreSaveList = append(scoreSaveList, e)
		}
	}
	this.BFScoreService.SaveList(scoreSaveList)
	this.BFScoreService.ModifyList(scoreModifyList)

	//对战历史
	battleSaveList := make([]interface{}, 0)
	battleModifyList := make([]interface{}, 0)
	battleList := this.battle_process(matchId, p)
	for _, e := range battleList {
		temp_id, exist := this.BFBattleService.Exist(e)
		if exist {
			e.Id = temp_id
			battleModifyList = append(battleModifyList, e)
		} else {
			battleSaveList = append(battleSaveList, e)
		}
	}

	this.BFBattleService.SaveList(battleSaveList)
	this.BFBattleService.ModifyList(battleModifyList)

	//近期对战
	jinSaveList := make([]interface{}, 0)
	jinModifyList := make([]interface{}, 0)
	jinList := this.jin_process(matchId, p)
	for _, e := range jinList {
		if len(string(e.ScheduleID)) <= 0 {
			continue
		}
		temp_id, exist := this.BFJinService.Exist(e)
		if exist {
			e.Id = temp_id
			jinModifyList = append(jinModifyList, e)
		} else {
			jinSaveList = append(jinSaveList, e)
		}
	}

	this.BFJinService.SaveList(jinSaveList)
	this.BFJinService.ModifyList(jinModifyList)

	//未来对战
	futureEventSaveList := make([]interface{}, 0)
	futureEventModifyList := make([]interface{}, 0)
	futureEventList := this.future_event_process(matchId, p)
	for _, e := range futureEventList {
		temp_id, exist := this.BFFutureEventService.Exist(e)
		if exist {
			e.Id = temp_id
			futureEventModifyList = append(futureEventModifyList, e)
		} else {
			futureEventSaveList = append(futureEventSaveList, e)
		}
	}
	this.BFFutureEventService.SaveList(futureEventSaveList)
	this.BFFutureEventService.ModifyList(futureEventModifyList)

}

//处理获取积分榜数据
func (this *BaseFaceProcesser) score_process(matchId string, p *page.Page) []*pojo.BFScore {
	data_list_slice := make([]*pojo.BFScore, 0)

	elem_table := p.GetHtmlParser().Find(" div.fenxiBar:contains('联赛积分排名')~table.mytable")
	elem_table.EachWithBreak(func(i int, selection *goquery.Selection) bool {
		//只取前两个table
		if i > 1 {
			return false
		}

		prev := selection.Prev()
		tempTeamId := strings.TrimSpace(prev.Text())

		selection.Find(" tr[align=center]  ").Each(func(i int, selection *goquery.Selection) {
			val_arr := make([]string, 0)
			selection.Children().Each(func(i int, selection *goquery.Selection) {
				val := selection.Text()
				val_arr = append(val_arr, strings.TrimSpace(val))
			})
			temp := new(pojo.BFScore)
			temp.MatchId = matchId
			temp.TeamId = tempTeamId
			temp.Type = val_arr[0]
			temp.MatchCount, _ = strconv.Atoi(val_arr[1])
			temp.WinCount, _ = strconv.Atoi(val_arr[2])
			temp.DrawCount, _ = strconv.Atoi(val_arr[3])
			temp.FailCount, _ = strconv.Atoi(val_arr[4])
			temp.GetGoal, _ = strconv.Atoi(val_arr[5])
			temp.LossGoal, _ = strconv.Atoi(val_arr[6])
			temp.DiffGoal, _ = strconv.Atoi(val_arr[7])
			temp.Score, _ = strconv.Atoi(val_arr[8])
			temp.Ranking, _ = strconv.Atoi(val_arr[9])
			temp_val := strings.Replace(val_arr[10], "%", "", 1)
			temp.WinRate, _ = strconv.ParseFloat(temp_val, 64)

			data_list_slice = append(data_list_slice, temp)
		})
		return true
	})
	return data_list_slice
}

//处理对战数据获取
func (this *BaseFaceProcesser) battle_process(matchId string, p *page.Page) []*pojo.BFBattle {
	request := p.GetRequest()
	data_list_slice := make([]*pojo.BFBattle, 0)

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
		base.Log.Error("hdata_str:为空,URL:", request.Url)
		return data_list_slice
	}

	base.Log.Info("hdata_str", hdata_str, "URL:", request.Url)
	// 获取script脚本中的，博彩公司信息
	temp_arr := strings.Split(hdata_str, "var vsTeamInfo = ")
	temp_arr = strings.Split(temp_arr[1], ";")
	hdata_str = strings.TrimSpace(temp_arr[0])
	if hdata_str == "" {
		base.Log.Info("hdata_str:解析失败,", hdata_str, "URL:", request.Url)
		return data_list_slice
	}
	var hdata_list = make([]*vo.BattleData, 0)
	json.Unmarshal(([]byte)(hdata_str), &hdata_list)

	//入库中
	for _, v := range hdata_list {
		temp := new(pojo.BFBattle)

		temp.MatchId = matchId
		battleMatchDate, _ := time.ParseInLocation("2006-01-02", v.Year+"-"+v.Date, time.Local)
		temp.BattleMatchDate = battleMatchDate
		temp.BattleLeagueId = v.SclassID
		temp.BattleMainTeamId = v.Home
		temp.BattleGuestTeamId = v.Guest

		half_goals := strings.Split(v.HT, "-")
		full_goals := strings.Split(v.FT, "-")
		temp.BattleMainTeamHalfGoals, _ = strconv.Atoi(half_goals[0])
		temp.BattleGuestTeamHalfGoals, _ = strconv.Atoi(half_goals[1])
		temp.BattleMainTeamGoals, _ = strconv.Atoi(full_goals[0])
		temp.BattleGuestTeamGoals, _ = strconv.Atoi(full_goals[1])

		data_list_slice = append(data_list_slice, temp)
	}

	return data_list_slice
}

//处理对战数据获取
func (this *BaseFaceProcesser) jin_process(matchId string, p *page.Page) []*pojo.BFJin {
	data_list_slice := make([]*pojo.BFJin, 0)

	var hdata_str string
	p.GetHtmlParser().Find("script").Each(func(i int, selection *goquery.Selection) {
		text := selection.Text()
		if hdata_str == "" && strings.Contains(text, "var nearInfo") {
			hdata_str = text
		} else {
			return
		}
	})
	if hdata_str == "" {
		return data_list_slice
	}

	// 获取script脚本中的，博彩公司信息
	temp_arr := strings.Split(hdata_str, "var nearInfo = ")
	temp_arr = strings.Split(temp_arr[1], ";")
	hdata_str = strings.TrimSpace(temp_arr[0])

	data := new(vo.JinData)
	json.Unmarshal(([]byte)(hdata_str), &data)

	for _, v := range data.HomeInfo {
		data_list_slice = append(data_list_slice, v)
	}
	for _, v := range data.GuestInfo {
		var exist bool
		for _, v2 := range data_list_slice {
			if v.ScheduleID == v2.ScheduleID {
				exist = true
				break
			}
		}
		if !exist {
			data_list_slice = append(data_list_slice, v)
		}
	}

	return data_list_slice
}

/**
将让球转换类型
*/
func (this *BaseFaceProcesser) ConvertLetball(letball string) float64 {
	var lb_sum float64
	slb_arr := strings.Split(letball, "/")
	slb_arr_0, _ := strconv.ParseFloat(slb_arr[0], 10)
	if len(slb_arr) > 1 {
		if strings.Index(slb_arr[0], "-") != -1 {
			lb_sum = slb_arr_0 - 0.25
		} else {
			lb_sum = slb_arr_0 + 0.25
		}
	} else {
		lb_sum = slb_arr_0
	}

	return lb_sum
}

//处理获取示来对战数据
func (this *BaseFaceProcesser) future_event_process(matchId string, p *page.Page) []*pojo.BFFutureEvent {
	data_list_slice := make([]*pojo.BFFutureEvent, 0)

	elem_table := p.GetHtmlParser().Find(" div.fenxiBar:contains('未来三场')~table.mytable")
	elem_table.Each(func(i int, selection *goquery.Selection) {
		if i > 1 {
			return
		}
		prev := selection.Prev()
		tempTeamId := strings.TrimSpace(prev.Text())

		selection.Find(" tr[align=center] ").Each(func(i int, selection *goquery.Selection) {
			val_arr := make([]string, 0)
			selection.Children().Each(func(i int, selection *goquery.Selection) {
				if i == 0 {
					selection.Find("div").Each(func(i int, selection *goquery.Selection) {
						val := selection.Text()
						val_arr = append(val_arr, strings.TrimSpace(val))
					})
				} else {
					val := selection.Text()
					val_arr = append(val_arr, strings.TrimSpace(val))
				}

			})
			temp := new(pojo.BFFutureEvent)
			temp.MatchId = matchId
			temp.TeamId = tempTeamId
			temp.EventMatchDate, _ = time.ParseInLocation("2006-01-02", val_arr[0], time.Local)
			temp.EventLeagueId = val_arr[1]
			temp.EventMainTeamId = val_arr[2]
			temp.EventGuestTeamId = val_arr[3]
			temp_val := strings.Replace(val_arr[4], "天", "", 1)
			temp.IntervalDay, _ = strconv.Atoi(temp_val)

			data_list_slice = append(data_list_slice, temp)
		})
	})
	return data_list_slice
}

func (this *BaseFaceProcesser) Finish() {
	base.Log.Info("基本面分析抓取解析完成 \r\n")

}

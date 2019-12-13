package proc

import (
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/pipeline"
	"github.com/hu17889/go_spider/core/spider"
	"math/rand"
	"strconv"
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	entity2 "tesou.io/platform/foot-parent/foot-api/module/elem/pojo"
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	service2 "tesou.io/platform/foot-parent/foot-core/module/elem/service"
	"tesou.io/platform/foot-parent/foot-core/module/match/service"
	"tesou.io/platform/foot-parent/foot-spider/module/win007"
	"tesou.io/platform/foot-parent/foot-spider/module/win007/down"
	"time"
)

type MatchPageProcesser struct {
	service.MatchLastService
	service2.LeagueService
	service2.CompService
	//抓取的url
	MatchlastUrl string
}

func GetMatchPageProcesser() *MatchPageProcesser {
	return &MatchPageProcesser{}
}

var (
	//联赛数据
	league_list           = make([]*entity2.League, 0)
	win007Id_leagueId_map = make(map[string]string)
	//比赛数据
	matchLast_list = make([]*pojo.MatchLast, 0)
)

func init() {

}

func (this *MatchPageProcesser) Startup() {
	if this.MatchlastUrl == "" {
		this.MatchlastUrl = "http://m.win007.com/phone/Schedule_0_0.txt"
	}
	this.MatchlastUrl = this.MatchlastUrl + "?flesh=" + strconv.FormatFloat(rand.Float64(), 'f', -1, 64)
	newSpider := spider.NewSpider(GetMatchPageProcesser(), "MatchPageProcesser")
	newSpider = newSpider.AddUrl(this.MatchlastUrl, "text")
	newSpider.SetDownloader(down.NewMWin007Downloader())
	newSpider = newSpider.AddPipeline(pipeline.NewPipelineConsole())
	newSpider.SetThreadnum(1).Run()
}

func (this *MatchPageProcesser) Process(p *page.Page) {
	request := p.GetRequest()
	if !p.IsSucc() {
		base.Log.Info("URL:,", request.Url, p.Errormsg())
		return
	}

	rawText := p.GetBodyStr()
	if rawText == "" {
		base.Log.Info("URL:,内容为空", request.Url)
		return
	}

	rawText_arr := strings.Split(rawText, "$$")
	if len(rawText_arr) < 2 {
		base.Log.Info("URL:,解析失败,rawTextArr长度为:,小于所必需要的长度3", request.Url, len(rawText_arr))
		return
	}

	flag := this.findParamVal(request.Url)
	var league_str string
	var match_str string
	if flag == "0" {
		league_str = rawText_arr[0]
		match_str = rawText_arr[1]
	} else {
		league_str = rawText_arr[1]
		match_str = rawText_arr[2]
	}

	base.Log.Info("联赛信息:", league_str)
	this.league_process(league_str)
	base.Log.Info("比赛信息:", match_str)
	this.match_process(match_str)
}

func (this *MatchPageProcesser) findParamVal(url string) string {
	paramUrl := strings.Split(url, "_")[2]
	paramArr := strings.Split(paramUrl, ".")
	return paramArr[0]
}

func (this *MatchPageProcesser) league_process(rawText string) {
	league_arr := strings.Split(rawText, "!")

	league_list = make([]*entity2.League, len(league_arr))
	var index int
	for _, v := range league_arr {
		league_info_arr := strings.Split(v, "^")
		if len(league_info_arr) < 3 {
			continue
		}
		name := league_info_arr[0]
		win007Id := league_info_arr[1]

		league := new(entity2.League)
		league.Id = win007Id
		league.Name = name
		//league.Ext = make(map[string]interface{})
		//league.Ext["win007Id"] = win007Id
		win007Id_leagueId_map[win007Id] = league.Id

		level_str := league_info_arr[2]
		level, _ := strconv.Atoi(level_str)
		league.Level = level

		league_list[index] = league
		index++
	}
}

/**
处理比赛信息
*/
func (this *MatchPageProcesser) match_process(rawText string) {
	match_arr := strings.Split(rawText, "!")
	match_len := len(match_arr)
	for i := 0; i < match_len; i++ {
		matchLast := new(pojo.MatchLast)
		//matchLast.Ext = make(map[string]interface{})
		//matchLast.Id = bson.NewObjectId().Hex()

		//match_arr[0] is
		//1503881^284^0^20180909170000^^町田泽维亚^水户蜀葵^0^0^0^0^0^0^0^0^0.5^181^^0^2^12^^^0^0^0^0
		match_info_arr := strings.Split(match_arr[i], "^")
		index := 0
		win007Id := match_info_arr[index]
		//matchLast.Ext["win007Id"] = win007Id
		matchLast.Id = win007Id
		index++
		matchLast.LeagueId = win007Id_leagueId_map[match_info_arr[index]]
		index++
		index++
		match_date_str := match_info_arr[index]
		matchLast.MatchDate, _ = time.ParseInLocation("20060102150405", match_date_str, time.Local)
		index++
		index++
		matchLast.MainTeamId = match_info_arr[index]
		/*		if regexp.MustCompile("^\\d*$").MatchString(dataDate_or_mainTeamName) {
					data_date_timestamp, _ := strconv.ParseInt(dataDate_or_mainTeamName, 10, 64)
					matchLast.DataDate = time.Unix(data_date_timestamp, 0).Format("2006-01-02 15:04:05")
				} else {
					matchLast.MainTeamId = dataDate_or_mainTeamName
				}*/
		index++
		matchLast.GuestTeamId = match_info_arr[index]
		index++
		mainTeamGoals_str := match_info_arr[index]
		mainTeamGoals, _ := strconv.Atoi(mainTeamGoals_str)
		matchLast.MainTeamGoals = mainTeamGoals
		index++
		guestTeamGoals_str := match_info_arr[index]
		guestTeamGoals, _ := strconv.Atoi(guestTeamGoals_str)
		matchLast.GuestTeamGoals = guestTeamGoals

		//最后加入数据中
		matchLast_list = append(matchLast_list, matchLast)
	}

}

func (this *MatchPageProcesser) Finish() {
	base.Log.Info("比赛抓取解析完成,执行入库 \r\n")

	league_list_slice := make([]interface{}, 0)
	for _, v := range league_list {
		if nil == v {
			continue
		}
		/*	bytes, _ := json.Marshal(v)
			base.Log.Info(string(bytes))*/
		exists := this.LeagueService.FindExistsById(v.Id)
		if exists {
			continue
		}
		league_list_slice = append(league_list_slice, v)
	}
	this.LeagueService.SaveList(league_list_slice)

	matchLast_list_slice := make([]interface{}, 0)
	for _, v := range matchLast_list {
		if nil == v {
			continue
		}

		exists := this.MatchLastService.FindExists(v)
		if exists {
			continue
		}
		//v.Id = v.Ext["win007Id"].(string);

		//处理比赛配置信息
		matchExt := new(pojo.MatchExt)
		/*		matchLast_elem := reflect.ValueOf(v).Elem()
				matchExt.MatchId = matchLast_elem.FieldByName("Id").String()*/
		//ext := matchLast_elem.FieldByName("Ext").Interface().(map[string]interface{})
		//matchExt.Sid = ext["win007Id"].(string)
		matchExt.Sid = v.Id
		v.Ext = make(map[string]interface{})
		v.Ext[win007.MODULE_FLAG] = matchExt
		matchLast_list_slice = append(matchLast_list_slice, v)
	}
	this.MatchLastService.SaveList(matchLast_list_slice)

}

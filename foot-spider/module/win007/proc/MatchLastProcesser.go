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

type MatchLastProcesser struct {
	service.MatchLastService
	service.MatchHisService
	service2.LeagueService
	service2.CompService
	//是否是单线程
	SingleThread bool
	//抓取的url
	MatchlastUrl string
	//联赛数据
	league_list           []*entity2.League
	win007Id_leagueId_map map[string]string
	//比赛数据
	matchLast_list []*pojo.MatchLast
	//比赛级别
	MatchLevel int
}

func GetMatchLastProcesser() *MatchLastProcesser {
	processer := &MatchLastProcesser{}
	processer.Init()
	return processer
}

func (this *MatchLastProcesser) Init() {
	//初始化参数值
	this.league_list = make([]*entity2.League, 0)
	this.win007Id_leagueId_map = make(map[string]string)
	this.matchLast_list = make([]*pojo.MatchLast, 0)
}

func (this *MatchLastProcesser) Startup() {
	if this.MatchlastUrl == "" {
		this.MatchlastUrl = "http://m.win007.com/phone/Schedule_0_0.txt"
	}
	this.MatchlastUrl = this.MatchlastUrl + "?flesh=" + strconv.FormatFloat(rand.Float64(), 'f', -1, 64)
	newSpider := spider.NewSpider(this, "MatchLastProcesser")
	newSpider = newSpider.AddUrl(this.MatchlastUrl, "text")
	newSpider.SetDownloader(down.NewMWin007Downloader())
	newSpider = newSpider.AddPipeline(pipeline.NewPipelineConsole())
	newSpider.SetSleepTime("rand", 1, 300)
	newSpider.SetThreadnum(1).Run()
}

func (this *MatchLastProcesser) Process(p *page.Page) {
	request := p.GetRequest()
	if !p.IsSucc() {
		base.Log.Error("URL:,", request.Url, p.Errormsg())
		return
	}

	rawText := p.GetBodyStr()
	if rawText == "" {
		base.Log.Error("rawText:为空.url:", request.Url)
		return
	}

	rawText_arr := strings.Split(rawText, "$$")
	if len(rawText_arr) < 2 {
		base.Log.Error("rawText:解析失败,rawTextArr长度小于所必需要的长度2,url:", request.Url, "内容:", rawText_arr)
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

	base.Log.Info("日期:TODAY", "联赛信息:", league_str)
	this.league_process(league_str)
	base.Log.Info("日期:TODAY", "比赛信息:", match_str)
	this.match_process(match_str)

	//now := time.Now()
	//获取明天赛程
	//h24h, _ := time.ParseDuration("24h")
	//t_1_date := now.Add(h24h).Format("2006-01-02")
	//this.futrueMatch(t_1_date)
	//获取后天赛程
	//h24h, _ = time.ParseDuration("48h")
	//t_1_date = now.Add(h24h).Format("2006-01-02")
	//this.futrueMatch(t_1_date)

}

type TomoReq struct {
	Date string `json:"date"`
}

func (this *MatchLastProcesser) futrueMatch(date string) {
	req := TomoReq{}
	req.Date = date

	url := "http://m.win007.com/ChangeDate.ashx?date=" + date
	rawText := GetText(url)
	//rawText := Post("http://m.win007.com/ChangeDate.ashx", &req)

	if rawText == "" {
		base.Log.Error("rawText:为空.url:", url)
		return
	}

	rawText_arr := strings.Split(rawText, "$")
	if len(rawText_arr) < 2 {
		base.Log.Error("rawText:解析失败,rawTextArr长度小于所必需要的长度2,url:", url, "内容:", rawText_arr)
		return
	}

	league_str := rawText_arr[0]
	match_str := rawText_arr[1]

	base.Log.Info("日期:", date, "联赛信息:", league_str)
	this.league_process(league_str)
	base.Log.Info("日期:", date, "比赛信息:", match_str)
	this.match_process(match_str)
}

func (this *MatchLastProcesser) findParamVal(url string) string {
	paramUrl := strings.Split(url, "_")[2]
	paramArr := strings.Split(paramUrl, ".")
	return paramArr[0]
}

func (this *MatchLastProcesser) league_process(rawText string) {
	league_arr := strings.Split(rawText, "!")

	for _, v := range league_arr {
		league_info_arr := strings.Split(v, "^")
		if len(league_info_arr) < 3 {
			continue
		}
		i := 0
		//联赛名称
		name := league_info_arr[i]
		//联赛ID
		i++
		win007Id := league_info_arr[i]
		league := new(entity2.League)
		league.Id = win007Id
		league.ShortName = name
		league.Name = name
		if this.MatchLevel > 0 {
			//设置联赛级别
			league.LevelAssist = 10000
		}
		//league.Ext = make(map[string]interface{})
		//league.Ext["win007Id"] = win007Id
		this.win007Id_leagueId_map[win007Id] = league.Id
		//联赛级别
		i++
		level_str := league_info_arr[i]
		level, _ := strconv.Atoi(level_str)
		league.Level = level
		//占位
		i++
		//赛事类别ID
		i++
		league.Sid = league_info_arr[i]
		//赛事类别名称
		i++
		league.SName = league_info_arr[i]

		//最后加入数据中
		this.league_list = append(this.league_list, league)
	}
}

/**
处理比赛信息
*/
func (this *MatchLastProcesser) match_process(rawText string) {
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
		//比赛ID
		matchLast.Id = win007Id
		//联赛ID
		index++
		matchLast.LeagueId = this.win007Id_leagueId_map[match_info_arr[index]]
		index++
		//比赛日期
		index++
		match_date_str := match_info_arr[index]
		matchLast.MatchDate, _ = time.ParseInLocation("20060102150405", match_date_str, time.Local)
		index++
		//主队客队名称
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
		//全场进球
		index++
		mainTeamGoals_str := match_info_arr[index]
		mainTeamGoals, _ := strconv.Atoi(mainTeamGoals_str)
		matchLast.MainTeamGoals = mainTeamGoals
		index++
		guestTeamGoals_str := match_info_arr[index]
		guestTeamGoals, _ := strconv.Atoi(guestTeamGoals_str)
		matchLast.GuestTeamGoals = guestTeamGoals
		//半场进球
		index++
		mainTeamHalfGoals_str := match_info_arr[index]
		mainTeamHalfGoals, _ := strconv.Atoi(mainTeamHalfGoals_str)
		matchLast.MainTeamHalfGoals = mainTeamHalfGoals
		index++
		guestTeamHalfGoals_str := match_info_arr[index]
		guestTeamHalfGoals, _ := strconv.Atoi(guestTeamHalfGoals_str)
		matchLast.GuestTeamHalfGoals = guestTeamHalfGoals
		//最后加入数据中
		this.matchLast_list = append(this.matchLast_list, matchLast)
	}

}

func (this *MatchLastProcesser) Finish() {
	base.Log.Info("比赛抓取解析完成,执行入库 \r\n")

	league_list_slice := make([]interface{}, 0)
	league_modify_list_slice := make([]interface{}, 0)
	for _, v := range this.league_list {
		if nil == v {
			continue
		}
		/*	bytes, _ := json.Marshal(v)
			base.Log.Info(string(bytes))*/
		exists := this.LeagueService.ExistById(v.Id)
		if exists {
			league_modify_list_slice = append(league_modify_list_slice, v)
			continue
		}
		league_list_slice = append(league_list_slice, v)
	}
	this.LeagueService.SaveList(league_list_slice)
	this.LeagueService.ModifyList(league_modify_list_slice)

	matchLast_list_slice := make([]interface{}, 0)
	matchLast_modify_list_slice := make([]interface{}, 0)
	matchHis_list_slice := make([]interface{}, 0)
	matchHis_modify_list_slice := make([]interface{}, 0)
	for _, v := range this.matchLast_list {
		if nil == v {
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
		exists := this.MatchLastService.Exist(v)
		if exists {
			matchLast_modify_list_slice = append(matchLast_modify_list_slice, v)
		} else {
			matchLast_list_slice = append(matchLast_list_slice, v)
		}

		his := new(pojo.MatchHis)
		his.MatchLast = *v
		his_exists := this.MatchHisService.Exist(his)
		if his_exists {
			matchHis_modify_list_slice = append(matchHis_modify_list_slice, his)
		} else {
			matchHis_list_slice = append(matchHis_list_slice, his)
		}
	}
	this.MatchLastService.SaveList(matchLast_list_slice)
	this.MatchLastService.ModifyList(matchLast_modify_list_slice)
	this.MatchHisService.SaveList(matchHis_list_slice)
	this.MatchHisService.ModifyList(matchHis_modify_list_slice)

}

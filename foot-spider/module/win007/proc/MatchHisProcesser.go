package proc

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/pipeline"
	"github.com/hu17889/go_spider/core/spider"
	"strconv"
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/elem/pojo"
	pojo2 "tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	service2 "tesou.io/platform/foot-parent/foot-core/module/elem/service"
	"tesou.io/platform/foot-parent/foot-core/module/match/service"
	"tesou.io/platform/foot-parent/foot-spider/module/win007"
	"tesou.io/platform/foot-parent/foot-spider/module/win007/down"
	"time"
)

type MatchHisProcesser struct {
	service.MatchHisService
	service2.LeagueService
	service2.LeagueSeasonService
	service2.LeagueSubService
	LeagueSeasonProcesser
	//是否是单线程
	SingleThread bool
	Season string
	//比赛数据
	MatchHis_list []*pojo2.MatchHis

	//联赛次级数据
	LeagueSeason_list []*pojo.LeagueSeason
	LeagueSub_list    []*pojo.LeagueSub
	SUrl_leagueId     map[string]string
	SUrl_Season       map[string]*pojo.LeagueSeason
	//----------------------
}

func GetMatchHisProcesser() *MatchHisProcesser {
	processer := &MatchHisProcesser{}
	processer.Init()
	return processer
}

func (this *MatchHisProcesser) Init() {
	//初始化参数值
	this.LeagueSeason_list = make([]*pojo.LeagueSeason, 0)
	this.LeagueSub_list = make([]*pojo.LeagueSub, 0)
	this.MatchHis_list = make([]*pojo2.MatchHis, 0)
	this.SUrl_leagueId = make(map[string]string)
	this.SUrl_Season = make(map[string]*pojo.LeagueSeason)
}

func (this *MatchHisProcesser) Setup(temp *MatchHisProcesser) {
	//设置参数值
}

func (this *MatchHisProcesser) Startup() {

	//1.获取所有的联赛赛季信息
	seasonList := this.LeagueSeasonService.FindBySeason(this.Season)
	//2.配置要抓取的路径
	var newSpider *spider.Spider
	processer := this
	newSpider = spider.NewSpider(processer, "MatchHisProcesser")
	for i, v := range seasonList {

		if i%10 == 0 { //10个联赛一个spider,总数1000多个联赛,最多100spider
			//先将前面的spider启动
			newSpider.SetDownloader(down.NewMWin007Downloader())
			newSpider = newSpider.AddPipeline(pipeline.NewPipelineConsole())
			newSpider.SetSleepTime("rand", win007.SLEEP_RAND_S, win007.SLEEP_RAND_E)
			newSpider.SetThreadnum(10).Run()

			processer = GetMatchHisProcesser()
			processer.Setup(this)
			newSpider = spider.NewSpider(processer, "MatchHisProcesser"+strconv.Itoa(i))
		}

		url := win007.WIN007_MATCH_HIS_PATTERN
		url = strings.Replace(url, "${season}", v.Season, 1)
		url = strings.Replace(url, "${leagueId}", v.LeagueId, 1)
		url = strings.Replace(url, "${subId}", "0", 1)

		index := 1
		for ; index <= v.Round; index++ {
			round_url := strings.Replace(url, "${round}", strconv.Itoa(index), 1)
			processer.SUrl_leagueId[round_url] = v.LeagueId
			processer.SUrl_Season[round_url] = v
			newSpider = newSpider.AddUrl(round_url, "html")
		}
	}

	newSpider.SetDownloader(down.NewMWin007Downloader())
	newSpider = newSpider.AddPipeline(pipeline.NewPipelineConsole())
	newSpider.SetSleepTime("rand", win007.SLEEP_RAND_S, win007.SLEEP_RAND_E)
	newSpider.SetThreadnum(1).Run()

}

func (this *MatchHisProcesser) Process(p *page.Page) {
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
	//1.处理season
	htmlParser := p.GetHtmlParser()
	leagueId := this.SUrl_leagueId[request.Url]
	this.LeagueSeasonProcesser.Init()
	this.LeagueSeasonProcesser.season_process(htmlParser, leagueId, request.Url)

	//1.处理比赛
	season := this.SUrl_Season[request.Url]
	htmlParser.Find("table[id='mainTable'] tr[onclick]").Each(func(i int, selection *goquery.Selection) {
		temp_id, exist := selection.Attr("onclick")
		if !exist {
			return
		}
		temp_id = strings.Replace(temp_id, "ToAnaly(", "", 1)
		temp_id = strings.Replace(temp_id, ",-1)", "", 1)
		temp_id = strings.TrimSpace(temp_id)
		base.Log.Info("比赛ID为:", temp_id, ",URL:"+request.Url)

		val_arr := make([]string, 0)
		selection.Find("td").Each(func(i int, selection *goquery.Selection) {
			val := selection.Text()
			val_arr = append(val_arr, strings.TrimSpace(val))
		})

		if len(val_arr) != 5 {
			return
		}

		his := new(pojo2.MatchHis)
		index := 0
		//比赛时间
		temp_matchDate := val_arr[index]
		seasonYear := season.Season
		if strings.Contains(season.Season, "-") {
			season_arr := strings.Split(season.Season, "-")
			month, _ := strconv.Atoi(temp_matchDate[:2])
			if month >= season.BeginMonth {
				seasonYear = season_arr[0]
			} else {
				seasonYear = season_arr[1]
			}
		}
		his.MatchDate, _ = time.ParseInLocation("200601-0215:04", seasonYear+temp_matchDate, time.Local)
		//比赛状态
		index++
		temp_status := val_arr[index]
		if !strings.EqualFold(temp_status, "完场") {
			return
		}
		//主队名称
		index++
		temp_mainTeam := val_arr[index]
		his.MainTeamId = temp_mainTeam
		//比分 全场半场
		index++
		temp_score := val_arr[index]
		temp_score = strings.Replace(temp_score, ")", "", 1)
		score_arr := strings.Split(temp_score, "(")
		full_arr := strings.Split(score_arr[0], ":")
		half_arr := strings.Split(score_arr[1], ":")
		his.MainTeamGoals, _ = strconv.Atoi(full_arr[0])
		his.MainTeamHalfGoals, _ = strconv.Atoi(half_arr[0])
		his.GuestTeamGoals, _ = strconv.Atoi(full_arr[1])
		his.GuestTeamHalfGoals, _ = strconv.Atoi(half_arr[1])
		//客队名称
		index++
		temp_guestTeam := val_arr[index]
		his.GuestTeamId = temp_guestTeam
		//设置id 联赛
		his.Id = temp_id
		his.LeagueId = season.LeagueId

		this.MatchHis_list = append(this.MatchHis_list, his)
	})

}

func (this *MatchHisProcesser) Finish() {
	base.Log.Info("历史比赛抓取解析完成,执行入库 \r\n")
	this.LeagueSeasonProcesser.Finish()

	matchHis_list_slice := make([]interface{}, 0)
	matchHis_modify_list_slice := make([]interface{}, 0)
	for _, v := range this.MatchHis_list {
		if nil == v {
			continue
		}

		//处理比赛配置信息
		matchExt := new(pojo2.MatchExt)
		matchExt.Sid = v.Id
		v.Ext = make(map[string]interface{})
		v.Ext[win007.MODULE_FLAG] = matchExt
		his_exists := this.MatchHisService.Exist(v)
		if his_exists {
			matchHis_modify_list_slice = append(matchHis_modify_list_slice, v)
		} else {
			matchHis_list_slice = append(matchHis_list_slice, v)
		}
	}
	this.MatchHisService.SaveList(matchHis_list_slice)
	this.MatchHisService.ModifyList(matchHis_modify_list_slice)
}

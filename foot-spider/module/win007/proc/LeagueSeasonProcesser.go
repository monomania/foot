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
	service2 "tesou.io/platform/foot-parent/foot-core/module/elem/service"
	"tesou.io/platform/foot-parent/foot-spider/module/win007"
	"tesou.io/platform/foot-parent/foot-spider/module/win007/down"
)

type LeagueSeasonProcesser struct {
	service2.LeagueService
	service2.LeagueSeasonService
	service2.LeagueSubService
	//是否是单线程
	SingleThread bool
	//联赛次级数据
	leagueSeason_list []*pojo.LeagueSeason
	leagueSub_list    []*pojo.LeagueSub
	sUrl_leagueId     map[string]string
}

func GetLeagueSeasonProcesser() *LeagueSeasonProcesser {
	processer := &LeagueSeasonProcesser{}
	processer.Init()
	return processer
}

func (this *LeagueSeasonProcesser) Init() {
	//初始化参数值
	this.leagueSeason_list = make([]*pojo.LeagueSeason, 0)
	this.leagueSub_list = make([]*pojo.LeagueSub, 0)
	this.sUrl_leagueId = make(map[string]string)
}

func (this *LeagueSeasonProcesser) Setup(temp *LeagueSeasonProcesser) {
	//设置参数值
}


func (this *LeagueSeasonProcesser) Startup() {
	//1.获取所有的联赛
	leaguesList := make([]*pojo.League, 0)
	this.LeagueService.FindAll(&leaguesList)
	//2.配置要抓取的路径
	var newSpider *spider.Spider
	processer := this
	newSpider = spider.NewSpider(processer, "LeagueSeasonProcesser")
	//index := 0
	for i, v := range leaguesList {
		//先不处理杯赛....
		if v.Cup {
			continue
		}
		//index++
		//if index > 10{
		//	break
		//}
		if i % 10 == 0 {//10个联赛一个spider,总数1000多个联赛,最多100spider
			processer = GetLeagueSeasonProcesser()
			processer.Setup(this)
			newSpider = spider.NewSpider(processer, "LeagueSeasonProcesser"+strconv.Itoa(i))
		}

		url := win007.WIN007_MATCH_HIS_PATTERN
		if v.SeasonCross {
			url = strings.Replace(url, "${season}", "2018-2019", 1)
		} else {
			url = strings.Replace(url, "${season}", "2019", 1)
		}
		url = strings.Replace(url, "${leagueId}", v.Id, 1)
		url = strings.Replace(url, "${subId}", "0", 1)
		url = strings.Replace(url, "${round}", "1", 1)

		processer.sUrl_leagueId[url] = v.Id
		newSpider = newSpider.AddUrl(url, "html")
		if i % 10 == 0 {//10个联赛一个spider,总数1000多个联赛,最多100spider
			newSpider.SetDownloader(down.NewMWin007Downloader())
			newSpider = newSpider.AddPipeline(pipeline.NewPipelineConsole())
			newSpider.SetSleepTime("rand", win007.SLEEP_RAND_S, win007.SLEEP_RAND_E)
			newSpider.SetThreadnum(10).Run()
		}
	}

	newSpider.SetDownloader(down.NewMWin007Downloader())
	newSpider = newSpider.AddPipeline(pipeline.NewPipelineConsole())
	newSpider.SetSleepTime("rand", win007.SLEEP_RAND_S, win007.SLEEP_RAND_E)
	newSpider.SetThreadnum(1).Run()

}

func (this *LeagueSeasonProcesser) Process(p *page.Page) {
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
	leagueId := this.sUrl_leagueId[request.Url]
	htmlParser := p.GetHtmlParser()
	this.season_process(htmlParser, leagueId, request.Url)

}

func (this *LeagueSeasonProcesser) season_process(htmlParser *goquery.Document, leagueId string, url string) {
	//获取赛季开始的月份
	var beginMonth int
	if strings.Contains(url, "-") {
		htmlParser.Find("table[id='mainTable'] tr[onclick]").Each(func(i int, selection *goquery.Selection) {
			temp_id, exist := selection.Attr("onclick")
			if !exist {
				return
			}
			temp_id = strings.Replace(temp_id, "ToAnaly(", "", 1)
			temp_id = strings.Replace(temp_id, ",-1)", "", 1)
			temp_id = strings.TrimSpace(temp_id)

			val_arr := make([]string, 0)
			selection.Find("td").Each(func(i int, selection *goquery.Selection) {
				val := selection.Text()
				val_arr = append(val_arr, strings.TrimSpace(val))
			})

			if len(val_arr) != 5 {
				return
			}

			index := 0
			//比赛时间
			temp_matchDate := val_arr[index]
			month, _ := strconv.Atoi(temp_matchDate[:2])
			beginMonth = month
		})
	}

	htmlParser.Find("select[id='selSeason'] option").Each(func(i int, selection *goquery.Selection) {
		temp_maxRound := 1
		temp_season := strings.TrimSpace(selection.Text())

		//2.处理sub
		htmlParser.Find("select[id='selSubSclass'] option").Each(func(i int, selection *goquery.Selection) {
			temp_subName := strings.TrimSpace(selection.Text())
			if len(temp_subName) <= 0 {
				return
			}
			val, exist := selection.Attr("value")
			if !exist {
				return
			}
			base.Log.Info("联赛Id:", leagueId, ",赛季:", temp_season, ",次级名称:", temp_subName, ",次级Id:", val)
			leagueSub := new(pojo.LeagueSub)
			leagueSub.LeagueId = leagueId
			leagueSub.Season = temp_season
			leagueSub.BeginMonth = beginMonth
			leagueSub.Round = temp_maxRound
			leagueSub.SubId = val
			leagueSub.SubName = temp_subName
			this.leagueSub_list = append(this.leagueSub_list, leagueSub)

		})
		//3.处理round
		htmlParser.Find("select[id='selRound'] option").Each(func(i int, selection *goquery.Selection) {
			temp_round_str, exist := selection.Attr("value")
			if !exist {
				return;
			}
			if len(temp_round_str) <= 0 {
				return
			}
			temp_round, _ := strconv.Atoi(temp_round_str)
			if temp_round > temp_maxRound {
				temp_maxRound = temp_round
			}
		})

		leagueSeason := new(pojo.LeagueSeason)
		leagueSeason.LeagueId = leagueId
		leagueSeason.Season = temp_season
		leagueSeason.BeginMonth = beginMonth
		leagueSeason.Round = temp_maxRound
		this.leagueSeason_list = append(this.leagueSeason_list, leagueSeason)
	})
}

func (this *LeagueSeasonProcesser) Finish() {
	base.Log.Info("联赛次级抓取解析完成,执行入库 \r\n")

	leagueSeason_list_slice := make([]interface{}, 0)
	leagueSeason_modify_list_slice := make([]interface{}, 0)
	for _, v := range this.leagueSeason_list {
		if nil == v {
			continue
		}
		id, exist := this.LeagueSeasonService.Exist(v)
		if exist {
			v.Id = id
			leagueSeason_modify_list_slice = append(leagueSeason_modify_list_slice, v)
			continue
		}
		leagueSeason_list_slice = append(leagueSeason_list_slice, v)
	}

	this.LeagueSeasonService.SaveList(leagueSeason_list_slice)
	this.LeagueSeasonService.ModifyList(leagueSeason_modify_list_slice)

	leagueSub_list_slice := make([]interface{}, 0)
	leagueSub_modify_list_slice := make([]interface{}, 0)
	for _, v := range this.leagueSub_list {
		if nil == v {
			continue
		}
		id, exist := this.LeagueSubService.Exist(v)
		if exist {
			v.Id = id
			leagueSub_modify_list_slice = append(leagueSub_modify_list_slice, v)
			continue
		}
		leagueSub_list_slice = append(leagueSub_list_slice, v)
	}

	this.LeagueSubService.SaveList(leagueSub_list_slice)
	this.LeagueSubService.ModifyList(leagueSub_modify_list_slice)

}

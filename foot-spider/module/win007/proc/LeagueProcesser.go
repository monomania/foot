package proc

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/pipeline"
	"github.com/hu17889/go_spider/core/spider"
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	entity2 "tesou.io/platform/foot-parent/foot-api/module/elem/pojo"
	service2 "tesou.io/platform/foot-parent/foot-core/module/elem/service"
	"tesou.io/platform/foot-parent/foot-spider/module/win007"
	"tesou.io/platform/foot-parent/foot-spider/module/win007/down"
)

type LeagueProcesser struct {
	service2.LeagueService
	service2.CompService
	//联赛数据
	league_list           []*entity2.League
	sUrl_Id map[string]string
	sUrl_Name map[string]string
}

func GetLeagueProcesser() *LeagueProcesser {
	return &LeagueProcesser{}
}

func (this *LeagueProcesser) Startup() {
	//初始化参数值
	this.league_list = make([]*entity2.League, 0)
	this.sUrl_Id = make(map[string]string)
	this.sUrl_Name = make(map[string]string)
	newSpider := spider.NewSpider(this, "LeagueProcesser")
	//sid 数据
	sid_stat_url := "http://m.win007.com/info.htm#section0";
	document, _ := GetDocument(sid_stat_url)
	document.Find("a[href*='sid']").Each(func(i int, selection *goquery.Selection) {
		sUrl, _ := selection.Attr("href")
		sId := strings.Split(sUrl, "sid=")[1]
		sName := strings.TrimSpace(selection.Text())
		base.Log.Info("sId:", sId, ",sName:", sName, ",sUrl:"+sUrl)
		this.sUrl_Id[win007.WIN007_BASE_URL+sUrl] = sId
		this.sUrl_Name[win007.WIN007_BASE_URL+sUrl] = sName
		newSpider = newSpider.AddUrl(win007.WIN007_BASE_URL+sUrl, "html")
	})
	newSpider.SetDownloader(down.NewMWin007Downloader())
	newSpider = newSpider.AddPipeline(pipeline.NewPipelineConsole())
	newSpider.SetSleepTime("rand",100,2000)
	newSpider.SetThreadnum(1).Run()
}

func (this *LeagueProcesser) Process(p *page.Page) {
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

	sUrl := request.Url
	sId := this.sUrl_Id[sUrl]
	sName := this.sUrl_Name[sUrl]

	p.GetHtmlParser().Find("a.gameItem[href*='info'][href*='htm']").Each(func(i int, selection *goquery.Selection) {
		lUrl, _ := selection.Attr("href")
		l_arr := strings.Split(lUrl, "/")
		lId_suffix := l_arr[len(l_arr)-1]
		lId := strings.ReplaceAll(lId_suffix,".htm","")
		lName := strings.TrimSpace(selection.Text())
		base.Log.Info("lId:", lId, ",lName:", lName, ",lUrl:"+lUrl)
		league := new(entity2.League)
		league.Id = lId
		league.Name = lName
		league.Sid = sId
		league.SName = sName
		league.ShortUrl = lUrl
		if strings.Contains(lUrl,"Cup"){
			league.Cup = true
		}
		if strings.Contains(lUrl,"-"){
			league.SeasonCross = true
		}

		this.league_list = append(this.league_list,league)
	})
}

func (this *LeagueProcesser) Finish() {
	base.Log.Info("联赛解析完成,执行入库 \r\n")

	league_list_slice := make([]interface{}, 0)
	league_modify_list_slice := make([]interface{}, 0)
	for _, v := range this.league_list {
		if nil == v {
			continue
		}
		exists := this.LeagueService.ExistById(v.Id)
		if exists {
			league_modify_list_slice = append(league_modify_list_slice,v)
			continue
		}
		league_list_slice = append(league_list_slice, v)
	}
	this.LeagueService.SaveList(league_list_slice)
	this.LeagueService.ModifyList(league_modify_list_slice)
}

package proc

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/pipeline"
	"github.com/hu17889/go_spider/core/spider"
	"strconv"
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	entity3 "tesou.io/platform/foot-parent/foot-api/module/odds/pojo"
	"tesou.io/platform/foot-parent/foot-core/module/elem/service"
	service2 "tesou.io/platform/foot-parent/foot-core/module/odds/service"
	"tesou.io/platform/foot-parent/foot-spider/module/win007"
	"tesou.io/platform/foot-parent/foot-spider/module/win007/down"
	"time"
)

type EuroTrackProcesser struct {
	service.CompService
	service2.EuroLastService
	service2.EuroHisService
	service2.EuroTrackService
	//入参
	//是否是单线程
	SingleThread bool
	MatchLastList []*pojo.MatchLast
	//博彩公司对应的win007id
	CompWin007Ids      []string
	Win007idMatchidMap map[string]string
}

func GetEuroTrackProcesser() *EuroTrackProcesser {
	processer := &EuroTrackProcesser{}
	processer.Init()
	return processer
}

func (this *EuroTrackProcesser) Init() {
	//初始化参数值
	this.Win007idMatchidMap = map[string]string{}
}

func (this *EuroTrackProcesser) Setup(temp *EuroTrackProcesser) {
	//设置参数值
	this.CompWin007Ids = temp.CompWin007Ids
}

func (this *EuroTrackProcesser) Startup() {

	var newSpider *spider.Spider
	processer := this
	newSpider = spider.NewSpider(processer, "EuroTrackProcesser")
	for i, v := range this.MatchLastList {

		if !this.SingleThread &&i%1000 == 0 { //10000个比赛一个spider,一个赛季大概有30万场比赛,最多30spider
			//先将前面的spider启动
			newSpider.SetDownloader(down.NewMWin007Downloader())
			newSpider = newSpider.AddPipeline(pipeline.NewPipelineConsole())
			newSpider.SetSleepTime("rand", 1, 300)
			newSpider.SetThreadnum(10).Run()

			processer = GetEuroTrackProcesser()
			processer.Setup(this)
			newSpider = spider.NewSpider(processer, "EuroTrackProcesser"+strconv.Itoa(i))
		}

		temp_flag := v.Ext[win007.MODULE_FLAG]
		bytes, _ := json.Marshal(temp_flag)
		matchExt := new(pojo.MatchExt)
		json.Unmarshal(bytes, matchExt)
		win007_id := matchExt.Sid
		processer.Win007idMatchidMap[win007_id] = v.Id

		base_url := strings.Replace(win007.WIN007_EUROODD_BET_URL_PATTERN, "${scheid}", win007_id, 1)
		for _, v := range processer.CompWin007Ids {
			url := strings.Replace(base_url, "${cId}", v, 1)
			newSpider = newSpider.AddUrl(url, "html")
		}
	}

	newSpider.SetDownloader(down.NewMWin007Downloader())
	newSpider = newSpider.AddPipeline(pipeline.NewPipelineConsole())
	newSpider.SetSleepTime("rand", 1, 300)
	newSpider.SetThreadnum(1).Run()

}

func (this *EuroTrackProcesser) findParamVal(url string, paramName string) string {
	paramUrl := strings.Split(url, "?")[1]
	paramArr := strings.Split(paramUrl, "&")
	for _, v := range paramArr {
		if strings.Contains(v, paramName) {
			return strings.Split(v, "=")[1]
		}
	}
	return ""
}

func (this *EuroTrackProcesser) Process(p *page.Page) {
	request := p.GetRequest()
	if !p.IsSucc() {
		base.Log.Error("URL:", request.Url, p.Errormsg())
		return
	}

	current_year := time.Now().Format("2006")

	win007_matchId := this.findParamVal(request.Url, "scheid")
	matchId := this.Win007idMatchidMap[win007_matchId]

	win007_betCompId := this.findParamVal(request.Url, "cId")

	var track_list = make([]*entity3.EuroTrack, 0)

	table_node := p.GetHtmlParser().Find(" table.mytable3 tr")
	table_node.Each(func(i int, selection *goquery.Selection) {
		if i < 2 {
			return
		}

		track := new(entity3.EuroTrack)
		track_list = append(track_list, track)
		track.MatchId = matchId
		track.CompId = win007_betCompId

		td_list_node := selection.Find(" td ")
		td_list_node.Each(func(ii int, selection *goquery.Selection) {
			val := strings.TrimSpace(selection.Text())
			if "" == val {
				return
			}

			switch ii {
			case 0:
				temp, _ := strconv.ParseFloat(val, 64)
				track.Sp3 = temp
			case 1:
				temp, _ := strconv.ParseFloat(val, 64)
				track.Sp1 = temp
			case 2:
				temp, _ := strconv.ParseFloat(val, 64)
				track.Sp0 = temp
			case 3:
				temp, _ := strconv.ParseFloat(val, 64)
				track.Payout = temp
			case 4:
				selection.Children().Each(func(iii int, selection *goquery.Selection) {
					val := selection.Text()
					switch iii {
					case 0:
						temp, _ := strconv.ParseFloat(val, 64)
						track.Kelly3 = temp
					case 1:
						temp, _ := strconv.ParseFloat(val, 64)
						track.Kelly1 = temp
					case 2:
						temp, _ := strconv.ParseFloat(val, 64)
						track.Kelly0 = temp
					}
				})
			case 5:
				var month_day string
				var hour_minute string
				selection.Children().Each(func(iii int, selection *goquery.Selection) {
					val := selection.Text()
					switch iii {
					case 0:
						month_day = val
					case 1:
						hour_minute = val
					}
				})
				track.OddDate = current_year + "-" + month_day + " " + hour_minute + ":00"
			}
		})
	})

	this.track_process(track_list)
}

func (this *EuroTrackProcesser) track_process(track_list []*entity3.EuroTrack) {
	track_lsit_len := len(track_list)
	if track_lsit_len < 1 {
		return
	}

	//将历史欧赔入库前，生成最后欧赔数据
	track_last := track_list[0]
	track_head := track_list[(track_lsit_len - 1)]

	last := new(entity3.EuroLast)
	last.MatchId = track_last.MatchId
	last.CompId = track_last.CompId
	last_temp_id, last_exists := this.EuroLastService.Exist(last)
	last.Sp3 = track_head.Sp3
	last.Sp1 = track_head.Sp1
	last.Sp0 = track_head.Sp0
	last.Ep3 = track_last.Sp3
	last.Ep1 = track_last.Sp1
	last.Ep0 = track_last.Sp0

	if last_exists {
		last.Id = last_temp_id
		this.EuroLastService.Modify(last)
	} else {
		this.EuroLastService.Save(last)
	}

	his := new(entity3.EuroHis)
	his.EuroLast = *last
	his_temp_id, his_exists := this.EuroHisService.Exist(his)
	if his_exists {
		his.Id = his_temp_id
		this.EuroHisService.Modify(his)
	} else {
		this.EuroHisService.Save(his)
	}

	//将历史赔率入库
	track_list_slice := make([]interface{}, 0)
	for _, v := range track_list {
		_, exists := this.EuroTrackService.Exist(v)
		if !exists {
			track_list_slice = append(track_list_slice, v)
		}
	}
	this.EuroTrackService.SaveList(track_list_slice)
}

func (this *EuroTrackProcesser) Finish() {
	base.Log.Info("欧赔历史抓取解析完成 \r\n")

}

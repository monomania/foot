package proc

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/pipeline"
	"github.com/hu17889/go_spider/core/spider"
	"log"
	"strconv"
	"strings"
	entity2 "tesou.io/platform/foot-parent/foot-api/module/elem/entity"
	"tesou.io/platform/foot-parent/foot-api/module/match/entity"
	entity3 "tesou.io/platform/foot-parent/foot-api/module/odds/entity"
	"tesou.io/platform/foot-parent/foot-core/module/elem/service"
	service2 "tesou.io/platform/foot-parent/foot-core/module/odds/service"
	"tesou.io/platform/foot-parent/foot-spider/module/win007"
	"time"
)

type EuroHisProcesser struct {
	service.CompService
	service.CompConfigService
	service2.EuroLastService
	service2.EuroHisService
	//博彩公司对应的win007id
	BetCompWin007Ids     []string
	MatchLastConfig_list []*entity.MatchLastConfig

	Win007Id_matchId_map   map[string]string
	Win007Id_betCompId_map map[string]string
}

func GetEuroHisProcesser() *EuroHisProcesser {
	return &EuroHisProcesser{}
}

func (this *EuroHisProcesser) Startup() {
	this.Win007Id_matchId_map = map[string]string{}
	this.Win007Id_betCompId_map = map[string]string{}

	newSpider := spider.NewSpider(this, "EuroHisProcesser")

	for _, v := range this.MatchLastConfig_list {

		match_win007_id := v.Sid

		this.Win007Id_matchId_map[match_win007_id] = v.MatchId

		url := strings.Replace(win007.WIN007_EUROODD_BET_URL_PATTERN, "${scheid}", match_win007_id, 1)
		for _, v := range this.BetCompWin007Ids {
			betCompConfig := new(entity2.CompConfig)
			betCompConfig.Sid = v
			existx := this.CompConfigService.FindBySId(betCompConfig)
			if !existx {
				comp := new(entity2.Comp)
				//comp.Id = bson.NewObjectId().Hex()
				comp.Id = v
				comp.Name = win007.MODULE_FLAG + "-" + v

				betCompConfig.CompId = comp.Id
				betCompConfig.S = win007.MODULE_FLAG
				this.CompService.Save(comp)
				this.CompConfigService.Save(betCompConfig)
			}
			this.Win007Id_betCompId_map[v] = betCompConfig.CompId

			url = strings.Replace(url, "${cId}", v, 1)
			newSpider = newSpider.AddUrl(url, "html")
		}
	}

	newSpider = newSpider.AddPipeline(pipeline.NewPipelineConsole())
	newSpider.SetThreadnum(1).Run()
}

func (this *EuroHisProcesser) findParamVal(url string, paramName string) string {
	paramUrl := strings.Split(url, "?")[1]
	paramArr := strings.Split(paramUrl, "&")
	for _, v := range paramArr {
		if strings.Contains(v, paramName) {
			return strings.Split(v, "=")[1]
		}
	}
	return ""
}

func (this *EuroHisProcesser) Process(p *page.Page) {
	request := p.GetRequest()
	if !p.IsSucc() {
		log.Println("URL:,", request.Url, p.Errormsg())
		return
	}

	current_year := time.Now().Format("2006")

	win007_matchId := this.findParamVal(request.Url, "scheid")
	matchId := this.Win007Id_matchId_map[win007_matchId]

	win007_betCompId := this.findParamVal(request.Url, "cId")
	compId := this.Win007Id_betCompId_map[win007_betCompId]

	var euroHis_list = make([]*entity3.EuroHis, 0)

	table_node := p.GetHtmlParser().Find(" table.mytable3 tr")
	table_node.Each(func(i int, selection *goquery.Selection) {
		if i < 2 {
			return
		}

		euroHis := new(entity3.EuroHis)
		euroHis_list = append(euroHis_list, euroHis)
		euroHis.MatchId = matchId
		euroHis.CompId = compId

		td_list_node := selection.Find(" td ")
		td_list_node.Each(func(ii int, selection *goquery.Selection) {
			val := strings.TrimSpace(selection.Text())
			if "" == val {
				return
			}

			switch ii {
			case 0:
				temp, _ := strconv.ParseFloat(val, 64)
				euroHis.Sp3 = temp
			case 1:
				temp, _ := strconv.ParseFloat(val, 64)
				euroHis.Sp1 = temp
			case 2:
				temp, _ := strconv.ParseFloat(val, 64)
				euroHis.Sp0 = temp
			case 3:
				temp, _ := strconv.ParseFloat(val, 64)
				euroHis.Payout = temp
			case 4:
				selection.Children().Each(func(iii int, selection *goquery.Selection) {
					val := selection.Text()
					switch iii {
					case 0:
						temp, _ := strconv.ParseFloat(val, 64)
						euroHis.Kelly3 = temp
					case 1:
						temp, _ := strconv.ParseFloat(val, 64)
						euroHis.Kelly1 = temp
					case 2:
						temp, _ := strconv.ParseFloat(val, 64)
						euroHis.Kelly0 = temp
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
				euroHis.OddDate = current_year + "-" + month_day + " " + hour_minute + ":00"
			}
		})
	})

	this.euroHis_process(euroHis_list)
}

func (this *EuroHisProcesser) euroHis_process(euroHis_lsit []*entity3.EuroHis) {
	euroHis_lsit_len := len(euroHis_lsit)
	if euroHis_lsit_len < 1 {
		return
	}

	//将历史欧赔入库前，生成最后欧赔表
	//暂时只处理616 -- 888Sport的数据
	euro_last := euroHis_lsit[0]
	if strings.EqualFold(euro_last.CompId, "616") {
		euro_head := euroHis_lsit[(euroHis_lsit_len - 1)]
		euro := new(entity3.EuroLast)
		euro.MatchId = euro_last.MatchId
		euro.CompId = euro_last.CompId
		euro_exists := this.EuroLastService.FindExists(euro)

		euro.Sp3 = euro_head.Sp3
		euro.Sp1 = euro_head.Sp1
		euro.Sp0 = euro_head.Sp0
		euro.Ep3 = euro_last.Sp3
		euro.Ep1 = euro_last.Sp1
		euro.Ep0 = euro_last.Sp0
		if euro_exists {
			this.EuroLastService.Modify(euro)
		} else {
			this.EuroLastService.Save(euro)
		}
	}

	//将历史赔率入库
	euroHis_list_slice := make([]interface{}, 0)
	for _, v := range euroHis_lsit {
		exists := this.EuroHisService.FindExists(v)
		if !exists {
			euroHis_list_slice = append(euroHis_list_slice, v)
		}
	}
	this.EuroHisService.SaveList(euroHis_list_slice)
}

func (this *EuroHisProcesser) Finish() {
	log.Println("欧赔历史抓取解析完成 \r\n")

}

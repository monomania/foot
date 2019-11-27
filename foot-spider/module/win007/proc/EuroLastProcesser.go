package proc

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/pipeline"
	"github.com/hu17889/go_spider/core/spider"
	"log"
	"regexp"
	"strconv"
	"strings"
	entity2 "tesou.io/platform/foot-parent/foot-api/module/elem/entity"
	"tesou.io/platform/foot-parent/foot-api/module/match/entity"
	entity3 "tesou.io/platform/foot-parent/foot-api/module/odds/entity"
	"tesou.io/platform/foot-parent/foot-core/module/elem/service"
	service2 "tesou.io/platform/foot-parent/foot-core/module/odds/service"
	"tesou.io/platform/foot-parent/foot-spider/module/win007"
	"tesou.io/platform/foot-parent/foot-spider/module/win007/vo"
)

type EuroLastProcesser struct {
	service.CompService
	service.CompConfigService
	service2.EuroLastService
	service2.EuroHisService
	//博彩公司对应的win007id
	BetCompWin007Ids     []string
	MatchLastConfig_list []*entity.MatchLastConfig
	Win007Id_matchId_map map[string]string
}

func GetEuroLastProcesser() *EuroLastProcesser {
	return &EuroLastProcesser{}
}

func (this *EuroLastProcesser) Startup() {
	this.Win007Id_matchId_map = map[string]string{}

	newSpider := spider.NewSpider(this, "EuroLastProcesser")

	for _, v := range this.MatchLastConfig_list {
		bytes, _ := json.Marshal(v)
		matchLastConfig := new(entity.MatchLastConfig)
		json.Unmarshal(bytes, matchLastConfig)

		win007_id := matchLastConfig.Sid

		this.Win007Id_matchId_map[win007_id] = matchLastConfig.MatchId

		url := strings.Replace(win007.WIN007_EUROODD_URL_PATTERN, "${matchId}", win007_id, 1)
		newSpider = newSpider.AddUrl(url, "html")
	}
	newSpider = newSpider.AddPipeline(pipeline.NewPipelineConsole())
	newSpider.SetThreadnum(1).Run()
}

func (this *EuroLastProcesser) Process(p *page.Page) {
	request := p.GetRequest()
	if !p.IsSucc() {
		log.Println("URL:,", request.Url, p.Errormsg())
		return
	}

	var hdata_str string
	p.GetHtmlParser().Find("script").Each(func(i int, selection *goquery.Selection) {
		text := selection.Text()
		if hdata_str == "" && strings.Contains(text, "var hData") {
			hdata_str = text
		} else {
			return
		}
	})
	if hdata_str == "" {
		return
	}

	// 获取script脚本中的，博彩公司信息
	hdata_str = strings.Replace(hdata_str, ";", "", 1)
	hdata_str = strings.Replace(hdata_str, "var hData = ", "", 1)
	log.Println(hdata_str)

	this.hdata_process(request.Url, hdata_str)
}

func (this *EuroLastProcesser) hdata_process(url string, hdata_str string) {

	var hdata_list = make([]*vo.HData, 0)
	json.Unmarshal(([]byte)(hdata_str), &hdata_list)
	var regex_temp = regexp.MustCompile(`(\d+).htm`)
	win007Id := strings.Split(regex_temp.FindString(url), ".")[0]
	matchId := this.Win007Id_matchId_map[win007Id]

	//入库中
	comp_list_slice := make([]interface{}, 0)
	compConfig_list_slice := make([]interface{}, 0)
	euro_list_slice := make([]interface{}, 0)
	euro_list_update_slice := make([]interface{}, 0)
	for _, v := range hdata_list {
		comp := new(entity2.Comp)
		comp.Name = v.Cn
		comp_exists := this.CompService.FindExistsByName(comp)
		if !comp_exists {
			//comp.Id = bson.NewObjectId().Hex()
			comp.Id = strconv.Itoa(v.CId)
			compConfig := new(entity2.CompConfig)
			compConfig.Id = comp.Id
			compConfig.CompId = comp.Id
			compConfig.Sid = strconv.Itoa(v.CId)
			compConfig.S = win007.MODULE_FLAG

			comp_list_slice = append(comp_list_slice, comp)
			compConfig_list_slice = append(compConfig_list_slice, compConfig)
		}

		//判断是否在配置的波菜抓取队列中
		if len(this.BetCompWin007Ids) > 0 {
			var equal bool
			for _, id := range this.BetCompWin007Ids {
				if strings.EqualFold(id, strconv.Itoa(v.CId)) {
					equal = true
					break
				}
			}
			if !equal {
				continue
			}
		}

		euro := new(entity3.EuroLast)
		euro.MatchId = matchId
		euro.CompId = comp.Id

		euro.Sp3 = v.Hw
		euro.Sp1 = v.So
		euro.Sp0 = v.Gw
		euro.Ep3 = v.Rh
		euro.Ep1 = v.Rs
		euro.Ep0 = v.Rg

		euro_exists := this.EuroLastService.FindExists(euro)
		if !euro_exists {
			euro_list_slice = append(euro_list_slice, euro)
		} else {
			euro_list_update_slice = append(euro_list_update_slice, euro)
		}

	}
	this.CompService.SaveList(comp_list_slice)
	this.CompService.SaveList(compConfig_list_slice)
	this.EuroLastService.SaveList(euro_list_slice)
	this.EuroLastService.ModifyList(euro_list_update_slice)
}

func (this *EuroLastProcesser) Finish() {
	log.Println("欧赔抓取解析完成 \r\n")

}

package proc

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"opensource.io/go_spider/core/common/page"
	"opensource.io/go_spider/core/pipeline"
	"opensource.io/go_spider/core/spider"
	"log"
	"tesou.io/platform/foot-parent/foot-core/module/match/entity"
	"tesou.io/platform/foot-parent/foot-spider/module/win007"
	"tesou.io/platform/foot-parent/foot-spider/module/win007/vo"
	"strings"
)

type MatchAnalyProcesser struct {
	//博彩公司对应的win007id
	MatchLastConfig_list []*entity.MatchLastConfig
	Win007Id_matchId_map map[string]string
}

func GetMatchAnalyProcesser() *MatchAnalyProcesser {
	return &MatchAnalyProcesser{}
}

func (this *MatchAnalyProcesser) Startup() {
	this.Win007Id_matchId_map = map[string]string{}

	newSpider := spider.NewSpider(this, "MatchAnalyProcesser")

 	for _, v := range this.MatchLastConfig_list {
		bytes, _ := json.Marshal(v)
		matchLastConfig := new(entity.MatchLastConfig)
		json.Unmarshal(bytes, matchLastConfig)

		win007_id := matchLastConfig.Sid

		this.Win007Id_matchId_map[win007_id] = matchLastConfig.MatchId

		url := strings.Replace(win007.WIN007_MATCH_ANALY_URL_PATTERN, "${matchId}", win007_id, 1)
		newSpider = newSpider.AddUrl(url, "html")
	}

	newSpider = newSpider.AddPipeline(pipeline.NewPipelineConsole())
	newSpider.SetThreadnum(1).Run()
}

func (this *MatchAnalyProcesser) Process(p *page.Page) {
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

func (this *MatchAnalyProcesser) hdata_process(url string, hdata_str string) {

	var hdata_list = make([]*vo.HData, 0)
	json.Unmarshal(([]byte)(hdata_str), &hdata_list)
	//var regex_temp = regexp.MustCompile(`(\d+).htm`)
	//win007Id := strings.Split(regex_temp.FindString(url), ".")[0]
	//matchId := this.Win007Id_matchId_map[win007Id]

	//入库中

}

func (this *MatchAnalyProcesser) Finish() {
	log.Println("比赛分析抓取解析完成 \r\n")

}

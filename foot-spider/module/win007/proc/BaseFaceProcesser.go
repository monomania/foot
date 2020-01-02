package proc

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/pipeline"
	"github.com/hu17889/go_spider/core/spider"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-spider/module/win007/down"

	"strings"
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	"tesou.io/platform/foot-parent/foot-spider/module/win007"
	"tesou.io/platform/foot-parent/foot-spider/module/win007/vo"
)

type BaseFaceProcesser struct {
	//博彩公司对应的win007id
	MatchLastList      []*pojo.MatchLast
	Win007idMatchidMap map[string]string
}

func GetBaseFaceProcesser() *BaseFaceProcesser {
	return &BaseFaceProcesser{}
}

func (this *BaseFaceProcesser) Startup() {
	this.Win007idMatchidMap = map[string]string{}

	newSpider := spider.NewSpider(this, "BaseFaceProcesser")

	for _, v := range this.MatchLastList {
		i := v.Ext[win007.MODULE_FLAG]
		bytes, _ := json.Marshal(i)
		matchLastExt := new(pojo.MatchExt)
		json.Unmarshal(bytes, matchLastExt)

		win007_id := matchLastExt.Sid

		this.Win007idMatchidMap[win007_id] = v.Id

		url := strings.Replace(win007.WIN007_BASE_FACE_URL_PATTERN, "${matchId}", win007_id, 1)
		newSpider = newSpider.AddUrl(url, "html")
	}
	newSpider.SetDownloader(down.NewMWin007Downloader())
	newSpider = newSpider.AddPipeline(pipeline.NewPipelineConsole())
	newSpider.SetThreadnum(1).Run()
}

func (this *BaseFaceProcesser) Process(p *page.Page) {
	request := p.GetRequest()
	if !p.IsSucc() {
		base.Log.Info("URL:,", request.Url, p.Errormsg())
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
	base.Log.Info(hdata_str)

	this.hdata_process(request.Url, hdata_str)
}

func (this *BaseFaceProcesser) hdata_process(url string, hdata_str string) {

	var hdata_list = make([]*vo.HData, 0)
	json.Unmarshal(([]byte)(hdata_str), &hdata_list)
	//var regex_temp = regexp.MustCompile(`(\d+).htm`)
	//win007Id := strings.Split(regex_temp.FindString(url), ".")[0]
	//matchId := this.Win007Id_matchId_map[win007Id]

	//入库中

}

func (this *BaseFaceProcesser) Finish() {
	base.Log.Info("比赛分析抓取解析完成 \r\n")

}

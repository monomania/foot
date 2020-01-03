package proc

import (
	"encoding/json"
	"fmt"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/pipeline"
	"github.com/hu17889/go_spider/core/spider"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-spider/module/win007/down"

	"strings"
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	"tesou.io/platform/foot-parent/foot-spider/module/win007"
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

	ret, _ := p.GetHtmlParser().Html()
	fmt.Println(ret)

}

func (this *BaseFaceProcesser) hdata_process(url string, hdata_str string) {



}

func (this *BaseFaceProcesser) Finish() {
	base.Log.Info("比赛分析抓取解析完成 \r\n")

}

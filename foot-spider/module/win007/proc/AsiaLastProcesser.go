package proc

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"opensource.io/go_spider/core/common/page"
	"opensource.io/go_spider/core/pipeline"
	"opensource.io/go_spider/core/spider"
	"log"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	"tesou.io/platform/foot-parent/foot-core/module/match/entity"
	entity3 "tesou.io/platform/foot-parent/foot-core/module/odds/entity"
	"tesou.io/platform/foot-parent/foot-spider/module/win007"
	"regexp"
	"strconv"
	"strings"
)

type AsiaLastProcesser struct {
	MatchLastConfig_list []*entity.MatchLastConfig
	Win007Id_matchId_map map[string]string
}

func GetAsiaLastProcesser() *AsiaLastProcesser {
	return &AsiaLastProcesser{}
}

func (this *AsiaLastProcesser) Startup() {
	this.Win007Id_matchId_map = map[string]string{}

	newSpider := spider.NewSpider(this, "AsiaLastProcesser")

	for _, v := range this.MatchLastConfig_list {
		bytes, _ := json.Marshal(v)
		matchLastConfig := new(entity.MatchLastConfig)
		json.Unmarshal(bytes, matchLastConfig)

		win007_id := matchLastConfig.Sid

		this.Win007Id_matchId_map[win007_id] = matchLastConfig.MatchId

		url := strings.Replace(win007.WIN007_ASIAODD_URL_PATTERN, "${matchId}", win007_id, 1)
		newSpider = newSpider.AddUrl(url, "html")
	}
	newSpider = newSpider.AddPipeline(pipeline.NewPipelineConsole())
	newSpider.SetThreadnum(1).Run()
}

func (this *AsiaLastProcesser) Process(p *page.Page) {
	request := p.GetRequest()
	if !p.IsSucc() {
		log.Println("URL:,", request.Url, p.Errormsg())
		return
	}

	var regex_temp = regexp.MustCompile(`(\d+).htm`)
	win007Id := strings.Split(regex_temp.FindString(request.Url), ".")[0]
	matchId := this.Win007Id_matchId_map[win007Id]
	asia_list_slice := make([]interface{}, 0)
	asia_list_update_slice := make([]interface{}, 0)
	p.GetHtmlParser().Find(" table.mytable3 tr").Each(func(i int, selection *goquery.Selection) {
		if i == 0 {
			return
		}

		asia := new(entity3.AsiaLast)
		asia.MatchId = matchId


		selection.Find("td").Each(func(td_index int, selection *goquery.Selection) {
			if td_index == 0 {
				//波菜公司名称
				asia.CompId = selection.Text()
			} else {
				selection.Children().Each(func(i int, selection *goquery.Selection) {
					if td_index == 1 {
						switch i {
						case 0:
							asia.Sp3, _ = strconv.ParseFloat(selection.Text(), 64)
							break
						case 1:
							asia.SLetBall = selection.Text()
							break
						case 2:
							asia.Sp0, _ = strconv.ParseFloat(selection.Text(), 64)
							break
						}
					} else if td_index == 2 {
						switch i {
						case 0:
							asia.Ep3, _ = strconv.ParseFloat(selection.Text(), 64)
							break
						case 1:
							asia.ELetBall = selection.Text()
							break
						case 2:
							asia.Ep0, _ = strconv.ParseFloat(selection.Text(), 64)
							break
						}
					}
				})
			}
		})

		asia_exists := asia.FindExists()
		if !asia_exists {
			asia_list_slice = append(asia_list_slice, asia)
		} else {
			asia_list_update_slice = append(asia_list_update_slice, asia)
		}

	})
	//执行入库
	mysql.SaveList(asia_list_slice)
	mysql.ModifyList(asia_list_update_slice)

}

func (this *AsiaLastProcesser) Finish() {
	log.Println("亚赔抓取解析完成 \r\n")

}

package proc

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/pipeline"
	"github.com/hu17889/go_spider/core/spider"
	"regexp"
	"strconv"
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	entity2 "tesou.io/platform/foot-parent/foot-api/module/odds/pojo"
	"tesou.io/platform/foot-parent/foot-core/module/odds/service"
	"tesou.io/platform/foot-parent/foot-spider/module/win007"
	"tesou.io/platform/foot-parent/foot-spider/module/win007/down"
)

type AsiaLastProcesser struct {
	service.AsiaLastService

	MatchLastList      []*pojo.MatchLast
	Win007idMatchidMap map[string]string
}

func GetAsiaLastProcesser() *AsiaLastProcesser {
	return &AsiaLastProcesser{}
}

func (this *AsiaLastProcesser) Startup() {
	this.Win007idMatchidMap = map[string]string{}

	newSpider := spider.NewSpider(this, "AsiaLastProcesser")

	for _, v := range this.MatchLastList {
		i := v.Ext[win007.MODULE_FLAG]
		bytes, _ := json.Marshal(i)
		matchExt := new(pojo.MatchExt)
		json.Unmarshal(bytes, matchExt)

		win007_id := matchExt.Sid

		this.Win007idMatchidMap[win007_id] = v.Id

		url := strings.Replace(win007.WIN007_ASIAODD_URL_PATTERN, "${matchId}", win007_id, 1)
		newSpider = newSpider.AddUrl(url, "html")
	}
	newSpider.SetDownloader(down.NewMWin007Downloader())
	newSpider = newSpider.AddPipeline(pipeline.NewPipelineConsole())
	newSpider.SetThreadnum(1).Run()
}

func (this *AsiaLastProcesser) Process(p *page.Page) {
	request := p.GetRequest()
	if !p.IsSucc() {
		base.Log.Info("URL:,", request.Url, p.Errormsg())
		return
	}

	var regex_temp = regexp.MustCompile(`(\d+).htm`)
	win007Id := strings.Split(regex_temp.FindString(request.Url), ".")[0]
	matchId := this.Win007idMatchidMap[win007Id]
	asia_list_slice := make([]interface{}, 0)
	asia_list_update_slice := make([]interface{}, 0)
	p.GetHtmlParser().Find(" table.mytable3 tr").Each(func(i int, selection *goquery.Selection) {
		if i == 0 {
			return
		}

		asia := new(entity2.AsiaLast)
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
							asia.SLetBall = ConvertLetball(selection.Text())
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
							asia.ELetBall = ConvertLetball(selection.Text())
							break
						case 2:
							asia.Ep0, _ = strconv.ParseFloat(selection.Text(), 64)
							break
						}
					}
				})
			}
		})

		asia_exists := this.AsiaLastService.FindExists(asia)
		if !asia_exists {
			asia_list_slice = append(asia_list_slice, asia)
		} else {
			asia_list_update_slice = append(asia_list_update_slice, asia)
		}

	})
	//执行入库
	this.AsiaLastService.SaveList(asia_list_slice)
	this.AsiaLastService.ModifyList(asia_list_update_slice)

}

/**
将让球转换类型
*/
func ConvertLetball(letball string) float64 {
	var lb_sum float64
	slb_arr := strings.Split(letball, "/")
	slb_arr_0, _ := strconv.ParseFloat(slb_arr[0], 10)
	if len(slb_arr) > 1 {
		if strings.Index(slb_arr[0], "-") != -1 {
			lb_sum = slb_arr_0 - 0.25
		} else {
			lb_sum = slb_arr_0 + 0.25
		}
	} else {
		lb_sum = slb_arr_0
	}

	return lb_sum
}

func (this *AsiaLastProcesser) Finish() {
	base.Log.Info("亚赔抓取解析完成 \r\n")

}

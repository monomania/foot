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
	service.AsiaHisService

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
	newSpider.SetSleepTime("rand",100,2000)
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
	asia_his_slice := make([]interface{}, 0)
	asia_his_update_slice := make([]interface{}, 0)
	p.GetHtmlParser().Find(" table.mytable3 tr").Each(func(i int, selection *goquery.Selection) {
		if i == 0 {
			return
		}

		last := new(entity2.AsiaLast)
		last.MatchId = matchId

		selection.Find("td").Each(func(td_index int, selection *goquery.Selection) {
			if td_index == 0 {
				//波菜公司名称
				last.CompId = selection.Text()
			} else {
				selection.Children().Each(func(i int, selection *goquery.Selection) {
					if td_index == 1 {
						switch i {
						case 0:
							last.Sp3, _ = strconv.ParseFloat(selection.Text(), 64)
							break
						case 1:
							last.SLetBall = ConvertLetball(selection.Text())
							break
						case 2:
							last.Sp0, _ = strconv.ParseFloat(selection.Text(), 64)
							break
						}
					} else if td_index == 2 {
						switch i {
						case 0:
							last.Ep3, _ = strconv.ParseFloat(selection.Text(), 64)
							break
						case 1:
							last.ELetBall = ConvertLetball(selection.Text())
							break
						case 2:
							last.Ep0, _ = strconv.ParseFloat(selection.Text(), 64)
							break
						}
					}
				})
			}
		})

		last_temp_id,last_exists := this.AsiaLastService.Exist(last)
		if !last_exists {
			asia_list_slice = append(asia_list_slice, last)
		} else {
			last.Id = last_temp_id
			asia_list_update_slice = append(asia_list_update_slice, last)
		}

		his := new(entity2.AsiaHis)
		his.CompId = last.CompId
		his.MatchId = last.MatchId
		his.OddDate = last.OddDate
		his.Sp0 = last.Sp0
		his.Sp3 = last.Sp3
		his.SLetBall = last.SLetBall
		his.Ep0 = last.Ep0
		his.Ep3 = last.Ep3
		his.ELetBall = last.ELetBall

		temp_id, his_exists := this.AsiaHisService.Exist(his)
		if !his_exists {
			asia_his_slice = append(asia_his_slice, his)
		} else {
			his.Id = temp_id
			asia_his_update_slice = append(asia_his_update_slice, his)
		}

	})
	//执行入库
	this.AsiaLastService.SaveList(asia_list_slice)
	this.AsiaLastService.ModifyList(asia_list_update_slice)
	this.AsiaHisService.SaveList(asia_his_slice)
	this.AsiaHisService.ModifyList(asia_his_update_slice)

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

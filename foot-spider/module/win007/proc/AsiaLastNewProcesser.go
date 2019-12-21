package proc

import (
	"encoding/json"
	"fmt"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/pipeline"
	"github.com/hu17889/go_spider/core/spider"
	"strconv"
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	entity2 "tesou.io/platform/foot-parent/foot-api/module/odds/pojo"
	"tesou.io/platform/foot-parent/foot-core/module/odds/service"
	"tesou.io/platform/foot-parent/foot-spider/module/win007"
	"tesou.io/platform/foot-parent/foot-spider/module/win007/down"
	"tesou.io/platform/foot-parent/foot-spider/module/win007/vo"
	"time"
)

type AsiaLastNewProcesser struct {
	service.AsiaLastService
	service.AsiaHisService

	MatchLastList      []*pojo.MatchLast
	Win007idMatchidMap map[string]string
}

func GetAsiaLastNewProcesser() *AsiaLastNewProcesser {
	return &AsiaLastNewProcesser{}
}

func (this *AsiaLastNewProcesser) Startup() {
	this.Win007idMatchidMap = map[string]string{}

	newSpider := spider.NewSpider(this, "AsiaLastNewProcesser")

	for _, v := range this.MatchLastList {
		i := v.Ext[win007.MODULE_FLAG]
		bytes, _ := json.Marshal(i)
		matchExt := new(pojo.MatchExt)
		json.Unmarshal(bytes, matchExt)

		win007_id := matchExt.Sid

		this.Win007idMatchidMap[win007_id] = v.Id

		url := strings.Replace(win007.WIN007_ASIAODD_NEW_URL_PATTERN, "${matchId}", win007_id, 1)
		newSpider = newSpider.AddUrl(url, "json")
	}
	newSpider.SetDownloader(down.NewMAsiaLastApiDownloader())
	newSpider = newSpider.AddPipeline(pipeline.NewPipelineConsole())
	newSpider.SetThreadnum(1).Run()
}

func (this *AsiaLastNewProcesser) Process(p *page.Page) {
	request := p.GetRequest()
	if !p.IsSucc() {
		base.Log.Info("URL:,", request.Url, p.Errormsg())
		return
	}


	asia_list_slice := make([]interface{}, 0)
	asia_list_update_slice := make([]interface{}, 0)
	asia_his_slice := make([]interface{}, 0)
	asia_his_update_slice := make([]interface{}, 0)
	str := p.GetBodyStr()
	fmt.Println(str)
	asiaData := &vo.AsiaData{}
	json.Unmarshal([]byte(str), asiaData)

	matchId := this.Win007idMatchidMap[strconv.Itoa(asiaData.ScheduleID)]
	//没有数据,则返回
	if nil == asiaData.Companies || len(asiaData.Companies) <= 0{
		return
	}
	for _, e := range asiaData.Companies {

		last := new(entity2.AsiaLast)
		last.MatchId = matchId
		last.CompId = e.NameCn

		odd := e.Details[0]
		last.Sp3= odd.FirstHomeOdds
		last.SLetBall = odd.FirstDrawOdds
		last.Sp0= odd.FirstAwayOdds
		last.Ep3= odd.HomeOdds
		last.ELetBall = odd.DrawOdds
		last.Ep0 = odd.AwayOdds
		if len(odd.ModifyTime) > 0 {
			tempMt, err := strconv.ParseInt(odd.ModifyTime,0,64)
			if nil != err{
				base.Log.Error(err.Error())
			}
			last.OddDate = time.Unix(tempMt, 0).Format("2006-01-02 15:04:05" )
		}
		last_exists := this.AsiaLastService.FindExists(last)
		if !last_exists {
			asia_list_slice = append(asia_list_slice, last)
		} else {
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

		his_exists := this.AsiaHisService.FindExists(his)
		if !his_exists {
			asia_his_slice = append(asia_his_slice, his)
		} else {
			asia_his_update_slice = append(asia_his_update_slice, his)
		}
	}

	//执行入库
	this.AsiaLastService.SaveList(asia_list_slice)
	this.AsiaLastService.ModifyList(asia_list_update_slice)
	this.AsiaHisService.SaveList(asia_his_slice)
	this.AsiaHisService.ModifyList(asia_his_update_slice)

}

func (this *AsiaLastNewProcesser) Finish() {
	base.Log.Info("亚赔抓取解析完成 \r\n")

}

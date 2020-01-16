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
	service.AsiaTrackService

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

	last_slice := make([]interface{}, 0)
	last_update_slice := make([]interface{}, 0)
	his_slice := make([]interface{}, 0)
	his_update_slice := make([]interface{}, 0)
	track_slice := make([]interface{}, 0)
	track_update_slice := make([]interface{}, 0)
	str := p.GetBodyStr()
	fmt.Println(str)
	asiaData := &vo.AsiaData{}
	json.Unmarshal([]byte(str), asiaData)

	matchId := this.Win007idMatchidMap[strconv.Itoa(asiaData.ScheduleID)]
	//没有数据,则返回
	if nil == asiaData.Companies || len(asiaData.Companies) <= 0 {
		return
	}
	for _, e := range asiaData.Companies {

		last := new(entity2.AsiaLast)
		last.MatchId = matchId
		last.CompId = e.NameCn

		odd := e.Details[0]
		last.Sp3 = odd.FirstHomeOdds
		last.SLetBall = odd.FirstDrawOdds
		last.Sp0 = odd.FirstAwayOdds
		last.Ep3 = odd.HomeOdds
		last.ELetBall = odd.DrawOdds
		last.Ep0 = odd.AwayOdds
		if len(odd.ModifyTime) > 0 {
			tempMt, err := strconv.ParseInt(odd.ModifyTime, 0, 64)
			if nil != err {
				base.Log.Error(err.Error())
			}
			last.OddDate = time.Unix(tempMt, 0).Format("2006-01-02 15:04:05")
		}
		last_exists := this.AsiaLastService.FindExists(last)
		if !last_exists {
			last_slice = append(last_slice, last)
		} else {
			last_update_slice = append(last_update_slice, last)
		}

		his := new(entity2.AsiaHis)
		his.AsiaLast = *last
		his_exists := this.AsiaHisService.FindExists(his)
		if !his_exists {
			his_slice = append(his_slice, his)
		} else {
			his_update_slice = append(his_update_slice, his)
		}

		track := new(entity2.AsiaTrack)
		track.CompId = last.CompId
		track.MatchId = last.MatchId
		track.OddDate = last.OddDate
		track.Sp0 = last.Sp0
		track.Sp3 = last.Sp3
		track.SLetBall = last.SLetBall
		track.Ep0 = last.Ep0
		track.Ep3 = last.Ep3
		track.ELetBall = last.ELetBall

		track_exists := this.AsiaTrackService.FindExists(track)
		if !track_exists {
			track_slice = append(track_slice, track)
		} else {
			track_update_slice = append(track_update_slice, track)
		}
	}

	//执行入库
	this.AsiaLastService.SaveList(last_slice)
	this.AsiaLastService.ModifyList(last_update_slice)

	this.AsiaHisService.SaveList(his_slice)
	this.AsiaHisService.ModifyList(his_update_slice)

	this.AsiaTrackService.SaveList(track_slice)
	this.AsiaTrackService.ModifyList(track_update_slice)

}

func (this *AsiaLastNewProcesser) Finish() {
	base.Log.Info("亚赔抓取解析完成 \r\n")

}

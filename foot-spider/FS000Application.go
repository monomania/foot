package main

import (
	"encoding/json"
	"regexp"
	"strings"
	"tesou.io/platform/foot-parent/foot-api/module/match/entity"
	"tesou.io/platform/foot-parent/foot-core/module/match/service"
	"tesou.io/platform/foot-parent/foot-spider/launch"
	"tesou.io/platform/foot-parent/foot-spider/module/win007"
)

func main() {
	Spider()
}

func Spider() {
	launch.Before_spider_match()
	launch.Before_spider_asiaLast()
	launch.Before_spider_euroLast()
	launch.Before_spider_euroHis()
	//执行抓取比赛数据
	//执行抓取比赛欧赔数据
	//执行抓取亚赔数据
	//执行抓取欧赔历史
	launch.Spider_match(4)
	launch.Spider_asiaLast()
	launch.Spider_euroLast()
	launch.Spider_euroHis()
}

func test() {
	lastConfigService := new(service.MatchLastConfigService)
	config := &entity.MatchLastConfig{}
	config.S = win007.MODULE_FLAG
	config.EOSpider = false
	matchLastConfig_list := lastConfigService.Query(config)

	for _, v := range matchLastConfig_list {
		bytes, _ := json.Marshal(v)

		println(string(bytes))
	}

	aaa := []string{"111", "222", "333"}
	println(strings.Join(aaa, ","))

	bbb := "http://m.win007.com/Compensate/1617919.htm"
	var digitsRegexp = regexp.MustCompile(`(\d+).htm`)
	println(digitsRegexp.FindString(bbb))

	map1 := map[string]string{}
	map1["11"] = "1111"
	println(map1["11"])
}

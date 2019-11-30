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
	launch.Spider(0)
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

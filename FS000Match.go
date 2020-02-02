package main

import (
	"strconv"
	"tesou.io/platform/foot-parent/foot-core/common/utils"
	"tesou.io/platform/foot-parent/foot-spider/launch"
)

func main() {
	matchLevelStr := utils.GetVal("spider", "match_level")
	if len(matchLevelStr) <= 0 {
		matchLevelStr = "4"
	}
	matchLevel, _ := strconv.Atoi(matchLevelStr)
	launch.Spider_match(matchLevel)
}


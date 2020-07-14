package main

import (
	"os"
	"strings"
	launch2 "tesou.io/platform/foot-parent/foot-core/launch"
	service4 "tesou.io/platform/foot-parent/foot-core/module/analy/service"
	"tesou.io/platform/foot-parent/foot-spider/launch"
)

func init() {

}

func main() {
	var input string
	if len(os.Args) > 1 {
		input = strings.ToLower(os.Args[1])
	} else {
		input = ""
	}

	switch input {
	case "exit\n", "exit", "quit\n", "quit":
		break;
	case "\n", "":
	case "init\n", "init":
		launch2.GenTable()
		launch2.TruncateTable()
	case "spider\n", "spider":
		launch.Spider()
	case "analy\n", "analy":
		launch2.Analy(false)
	case "alall\n", "alall":
		launch2.Analy(true)
	case "mr\n", "mr":
		//更新结果
		analyService := new(service4.AnalyService)
		analyService.ModifyAllResult()
	default:

	}

}

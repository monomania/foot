package main

import (
	"os"
	"strings"
	launch2 "tesou.io/platform/foot-parent/foot-core/launch"
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
	default:

	}

}

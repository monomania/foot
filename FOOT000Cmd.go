package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	launch2 "tesou.io/platform/foot-parent/foot-core/launch"
	"tesou.io/platform/foot-parent/foot-spider/launch"
)

func main() {
HEAD:
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please enter:")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}

	input = strings.ToLower(input)
	switch input {
	case "exit\n", "exit\r\n", "quit\n", "quit\r\n":
		break;
	case "\n", "\r\n":
		goto HEAD
	case "init\n", "init\r\n":
		launch2.GenTable()
		launch2.TruncateTable()
		goto HEAD
	case "spider\n", "spider\r\n":
		launch.Spider()
		goto HEAD
	case "analy\n", "analy\r\n":
		launch2.Analy(false)
		goto HEAD
	default:
		goto HEAD
	}

}

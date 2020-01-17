package main

import (
	"bufio"
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"io"
	"log"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-core/common/utils"

	"os"
	"strings"
	"tesou.io/platform/foot-parent/foot-core/launch"
	launch2 "tesou.io/platform/foot-parent/foot-spider/launch"
	"time"
)

func init() {

}

func main() {
	//buildWinForm()
	test()
}

func buildWinForm() {
	var inTE *walk.TextEdit
	window := MainWindow{
		Title:   "FOOT000GUI",
		MinSize: Size{400, 300},
		Layout:  VBox{},
		Children: []Widget{
			TextEdit{AssignTo: &inTE, ReadOnly: true, HScroll: false, VScroll: true},
			HSplitter{
				Children: []Widget{
					PushButton{
						Text: "Spider数据",
						OnClicked: func() {
							inTE.SetText("Spider数据...\r\n")
							go launch2.Spider()
							go showConsole(inTE)
						},
					},
					PushButton{
						Text: "分析数据",
						OnClicked: func() {
							inTE.SetText("分析数据...\r\n")
							go launch.Analy()
							go showConsole(inTE)
						},
					},
					PushButton{
						Text: "清理数据库",
						OnClicked: func() {
							inTE.SetText("清理数据库...\r\n")
							go showConsole(inTE)
						},
					},
					PushButton{
						Text: "清空日志",
						OnClicked: func() {
							inTE.SetText("清空日志...\r\n")
							logFile, err := os.OpenFile(base.output_Path, os.O_WRONLY|os.O_TRUNC, 0777)
							if err != nil {
								log.Fatal(err)
							}
							logFile.WriteString("")
							defer logFile.Close()

						},
					},
				},
			},
		},
	}
	window.Run()
}

func showConsole(edit *walk.TextEdit) {
	utils.FileMonitoring(base.output_Path, func(bytes []byte) {
		str := string(bytes)
		if strings.TrimSpace(str) == "" {
			return
		}
		str = str + "\r\n"
		edit.AppendText(str)
	})
}

func test() {
	logFile, err := os.OpenFile(base.output_Path, os.O_RDONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(logFile)
	for {
		var i int
		i += 1
		var str string
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			time.Sleep(200)
		} else if nil != err {
			str = "发生错误:" + err.Error()
		} else {
			str = string(line)
		}

		if strings.TrimSpace(str) == "" {
			continue
		}
		fmt.Println(str)
	}
}

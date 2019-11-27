package main

import (
	"encoding/json"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"io"
	"log"
	"os"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
	"tesou.io/platform/foot-parent/foot-core/common/utils"
	service2 "tesou.io/platform/foot-parent/foot-core/module/analy/service"
	"tesou.io/platform/foot-parent/foot-spider/launch"
)

func main() {
	build_win()
}

func build_win() {
	var inTE *walk.TextEdit
	window := MainWindow{
		Title:   "Foot",
		MinSize: Size{400, 300},
		Layout:  VBox{},
		Children: []Widget{
			TextEdit{AssignTo: &inTE, ReadOnly: true, HScroll: true},
			HSplitter{
				Children: []Widget{
					PushButton{
						Text: "Clean数据",
						OnClicked: func() {
							go GenTable()
							go TruncateTable()
							inTE.SetText("同步清理数据...")
						},
					},
					PushButton{
						Text: "Spider数据",
						OnClicked: func() {
							inTE.SetText("抓取数据...")
							go Spider()
						},
					},
					PushButton{
						Text: "Analy数据",
						OnClicked: func() {
							inTE.SetText("分析数据...")
							go Analy(func(i []interface{}) {
								var displayStr string
								for _, v := range i {
									bytes, _ := json.Marshal(v)
									displayStr += string(bytes)
									displayStr += "\r\n"
								}
								inTE.SetText(displayStr)
							})
						},
					},
				},
			},
		},
	}
	window.Run()
}

func ShowLog() {
	var logStr string
	for {
		utils.FileMonitoring("E:/home/logs/foot-gui.log", func(bytes []byte) {
			logStr += string(bytes)
		})
	}
}

func init() {
	//创建日志文件
	f, err := os.OpenFile("E:/home/logs/foot-gui.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	//完成后，延迟关闭
	defer f.Close()
	writers := []io.Writer{
		f,
		os.Stdout}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	//logger := log.New(fileAndStdoutWriter, "", log.Ldate|log.Ltime|log.Lshortfile)
	log.SetOutput(fileAndStdoutWriter)
	log.Println("初始化...")
}

func TruncateTable() {
	opsService := new(mysql.DBOpsService)
	//指定需要清空的数据表
	opsService.TruncateTable([]string{"t_match_last", "t_match_last_config", "t_euro_last", "t_euro_his", "t_asia_last"})
}

func GenTable() {
	generateService := new(mysql.DBOpsService)
	generateService.SyncTableStruct()
}

func Spider() {
	//执行抓取比赛数据
	launch.Before_spider_match()
	launch.Spider_match(0)
	//执行抓取比赛欧赔数据
	launch.Before_spider_euroLast()
	launch.Spider_euroLast()
	//执行抓取亚赔数据
	launch.Before_spider_asiaLast()
	launch.Spider_asiaLast()
	//执行抓取欧赔历史
	launch.Before_spider_euroHis()
	launch.Spider_euroHis()
}

func Analy(hookfn func([]interface{})){
	analysisService := new(service2.AnalyService)
	analysisService.MaxLetBall = 0.75
	analysisService.PrintOddData = false
	log.Println("-----------------------------------------------")
	log.Println("----------------计算欧86之差-------------------")
	log.Println("-----------------------------------------------")
	r1 := analysisService.Euro_Calc()
	log.Println("-----------------------------------------------")
	log.Println("---------------计算亚欧之差--------------------")
	log.Println("-----------------------------------------------")
	r2 := analysisService.Euro_Asia_Diff()

	i := append(r1, r2)
	hookfn(i)
}

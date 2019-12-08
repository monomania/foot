package main

import (
	"fmt"
	"os"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	launch2 "tesou.io/platform/foot-parent/foot-core/launch"
	"tesou.io/platform/foot-parent/foot-spider/launch"
	"tesou.io/platform/foot-parent/foot-web/module/leisu/service"
	"time"
)

func init() {

}

func main() {
	var input string
	if len(os.Args) > 1 {
		input = os.Args[1]
	} else {
		input = ""
	}

	switch input {
	case "init":
		launch2.GenTable()
		launch2.TruncateTable()
	case "spider":
		launch.Spider(4)
	case "analy":
		launch2.Analy()
	case "pub":
		pubService := new(service.PubService)
		pubService.PubBJDC()
	case "auto":
		for {
			base.Log.Info("--------程序开始运行--------")
			//1.安装数据库
			//2.配置好数据库连接,打包程序发布
			//3.程序执行流程,周期定制定为一天三次
			//3.1 FS000Application 爬取数据
			launch.Spider(4)
			//3.2 FC002AnalyApplication 分析得出推荐列表
			launch2.Analy()
			//3.3 FW001PubApplication 执行发布到雷速
			pubService := new(service.PubService)
			pubService.PubBJDC()
			base.Log.Info("--------程序周期结束--------")
			time.Sleep(7 * time.Hour)
		}
	default:
		fmt.Println("usage: init|spider|analy|pub|auto")
	}

}

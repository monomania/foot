package main

import (
	"fmt"
	"os"
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	launch2 "tesou.io/platform/foot-parent/foot-core/launch"
	service2 "tesou.io/platform/foot-parent/foot-core/module/core/service"
	"tesou.io/platform/foot-parent/foot-spider/launch"
	"time"
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
	case "init":
		launch2.GenTable()
		launch2.TruncateTable()
	case "spider":
		launch.Spider()
	case "analy":
		launch2.Analy()
	case "autospider":
		for {
			base.Log.Info("--------数据更新开始运行--------")
			//1.安装数据库
			//2.配置好数据库连接,打包程序发布
			//3.程序执行流程,周期定制定为一天三次
			//3.1 FS000Application 爬取数据
			launch.Spider()
			//3.2 FC002AnalyApplication 分析得出推荐列表
			launch2.Analy()
			configService := new(service2.ConfService)
			base.Log.Info("--------数据更新周期结束--------")
			time.Sleep(time.Duration(configService.GetSpiderCycleTime()) * time.Minute)
		}
	default:
		fmt.Println("usage: init|spider|analy|pub|auto")
	}

}

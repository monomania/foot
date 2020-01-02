package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	launch2 "tesou.io/platform/foot-parent/foot-core/launch"
	service2 "tesou.io/platform/foot-parent/foot-core/module/core/service"
	"tesou.io/platform/foot-parent/foot-spider/launch"
	"time"
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
		launch2.Analy()
		goto HEAD
	case "autospider\n", "autospider\r\n":
		go func() {
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
		}()
		goto HEAD
	case "auto\n", "auto\r\n":
		go func() {
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
		}()
		goto HEAD
	default:
		goto HEAD
	}

}

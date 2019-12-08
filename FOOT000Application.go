package main

import (
	"bufio"
	"fmt"
	"os"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	launch2 "tesou.io/platform/foot-parent/foot-core/launch"
	"tesou.io/platform/foot-parent/foot-spider/launch"
	"tesou.io/platform/foot-parent/foot-web/module/leisu/service"
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
	switch input {
	case "\n","\r\n":
		goto HEAD
	case "init\n","init\r\n":
		launch2.GenTable()
		launch2.TruncateTable()
		goto HEAD
	case "spider\n","spider\r\n":
		launch.Spider(4)
		goto HEAD
	case "analy\n","analy\r\n":
		launch2.Analy()
		goto HEAD
	case "pub\n","pub\r\n":
		pubService := new(service.PubService)
		pubService.PubBJDC()
		goto HEAD
	case "auto\n","auto\r\n":
		for {
			go func() {
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
			}()
			time.Sleep(8 * time.Hour)
		}
	default:
		goto HEAD
		fmt.Println("You are not welcome here! Goodbye!")
	}

}
